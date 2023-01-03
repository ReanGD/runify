import 'package:flutter/material.dart';

import 'package:reactive_forms/reactive_forms.dart';

import 'package:runify/rpc/rpc_form_service.dart';
import 'package:runify/screen/form/fm_screen.dart';

class FMController {
  final FMService _service;

  FMController(this._service);

  Widget build() {
    return FMScreen(this);
  }

  FormGroup get model => _service.model;
  get markup => _service.markup;

  void onMenuActivate() {
    // TODO: implement
  }

  void onSubmit() {
    if (model.valid) {
      _service.submit(model.value);
    }
  }
}
