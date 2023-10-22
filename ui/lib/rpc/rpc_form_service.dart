import 'dart:async';
import 'dart:convert';

import 'package:reactive_forms/reactive_forms.dart';

import 'package:runify/system/logger.dart';
import 'package:runify/rpc/rpc_types.dart';
import 'package:runify/rpc/rpc_proto_client.dart';
import 'package:runify/pb/runify.pbgrpc.dart' as pb;
import 'package:runify/rpc/rpc_service_storage.dart';

class _AsyncValidator extends AsyncValidator<dynamic> {
  final AsyncValidatorFunction _func;

  _AsyncValidator(this._func);

  @override
  Future<Map<String, dynamic>?> validate(
          AbstractControl<dynamic> control) async =>
      _func(control);
}

class FMService implements Service {
  var nextRequestID = 0;
  final Logger _logger;
  final ProtoClient _pClient;
  final ServiceStorage _storage;
  final Map<String, dynamic> _markup;
  final fieldCheckChan = StreamController<pb.FieldCheckResponse>.broadcast();
  late final FormGroup _model;

  FMService(this._storage, this._pClient, this._logger, pb.FormMarkup markup,
      pb.FormModel model)
      : _markup = json.decode(markup.json) {
    _model = _parseModel(model.json);
  }

  get model => _model;
  get markup => _markup;
  get formID => _pClient.formID;

  Future<Map<String, dynamic>> _asyncValidator(
      String name, AbstractControl<dynamic> control) async {
    final requestID = nextRequestID;
    nextRequestID++;
    _pClient.fieldCheckRequest(requestID, name, json.encode(_model.value));

    final respose = await fieldCheckChan.stream
        .firstWhere((item) => item.requestID == requestID)
        .timeout(const Duration(seconds: 10),
            onTimeout: () => pb.FieldCheckResponse(
                requestID: requestID, result: false, error: "Timeout"));

    if (respose.result) {
      return {};
    }

    return {respose.error: true};
  }

  FormGroup _parseModel(String data) {
    final Map<String, AbstractControl<dynamic>> controls = {};
    for (var it in json.decode(data)) {
      final item = it as Map<String, dynamic>;

      final name = item["name"];
      final validatorsJson = item["validators"];
      final List<Validator<dynamic>> validators = [];
      final List<AsyncValidator<dynamic>> asyncValidators = [];
      if (validatorsJson != null) {
        final items = validatorsJson as Map<String, dynamic>;
        for (var item in items.entries) {
          switch (item.key) {
            case "minLength":
              validators.add(Validators.minLength(item.value));
              break;
            case "maxLength":
              validators.add(Validators.maxLength(item.value));
              break;
            case "required":
              validators.add(Validators.required);
              break;
            case "serverSide":
              asyncValidators.add(
                  _AsyncValidator((value) => _asyncValidator(name, value)));
              break;
          }
        }
      }

      final value = item["value"];
      if (value is String) {
        controls[name] = FormControl<String>(
            value: value,
            validators: validators,
            asyncValidators: asyncValidators);
      } else if (value is int) {
        controls[name] = FormControl<int>(
            value: value,
            validators: validators,
            asyncValidators: asyncValidators);
      } else if (value is bool) {
        controls[name] = FormControl<bool>(
            value: value,
            validators: validators,
            asyncValidators: asyncValidators);
      } else if (value is double) {
        controls[name] = FormControl<double>(
            value: value,
            validators: validators,
            asyncValidators: asyncValidators);
      } else {
        _logger.error("Unexpected value type: ${value.runtimeType}");
      }
    }

    return FormGroup(controls);
  }

  @override
  void onRootListAddRows(List<pb.RootListRow> rows) {
    _logger.error("Unexpected grpc message 'RootListAddRows' for form handler");
  }

  @override
  void onRootListChangeRows(List<pb.RootListRow> rows) {
    _logger
        .error("Unexpected grpc message 'RootListChangeRows' for form handler");
  }

  @override
  void onRootListRemoveRows(List<pb.RootListRowGlobalID> rows) {
    _logger
        .error("Unexpected grpc message 'RootListRemoveRows' for form handler");
  }

  @override
  void onFieldCheckResponse(pb.FieldCheckResponse msg) {
    fieldCheckChan.add(msg);
  }

  void submit() {
    _pClient.formSubmit(json.encode(_model.value));
  }

  void formClosed() {
    if (_storage.remove(formID)) {
      _pClient.formClosed();
    }
  }
}
