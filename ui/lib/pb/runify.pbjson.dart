///
//  Generated code. Do not modify.
//  source: runify.proto
//
// @dart = 2.12
// ignore_for_file: annotate_overrides,camel_case_types,constant_identifier_names,deprecated_member_use_from_same_package,directives_ordering,library_prefixes,non_constant_identifier_names,prefer_final_fields,return_of_invalid_type,unnecessary_const,unnecessary_import,unnecessary_this,unused_import,unused_shown_name

import 'dart:core' as $core;
import 'dart:convert' as $convert;
import 'dart:typed_data' as $typed_data;
@$core.Deprecated('Use logLevelDescriptor instead')
const LogLevel$json = const {
  '1': 'LogLevel',
  '2': const [
    const {'1': 'DEBUG', '2': 0},
    const {'1': 'INFO', '2': 1},
    const {'1': 'WARNING', '2': 2},
    const {'1': 'ERROR', '2': 3},
  ],
};

/// Descriptor for `LogLevel`. Decode as a `google.protobuf.EnumDescriptorProto`.
final $typed_data.Uint8List logLevelDescriptor = $convert.base64Decode('CghMb2dMZXZlbBIJCgVERUJVRxAAEggKBElORk8QARILCgdXQVJOSU5HEAISCQoFRVJST1IQAw==');
@$core.Deprecated('Use formStateTypeDescriptor instead')
const FormStateType$json = const {
  '1': 'FormStateType',
  '2': const [
    const {'1': 'SHOW', '2': 0},
    const {'1': 'HIDE', '2': 1},
  ],
};

/// Descriptor for `FormStateType`. Decode as a `google.protobuf.EnumDescriptorProto`.
final $typed_data.Uint8List formStateTypeDescriptor = $convert.base64Decode('Cg1Gb3JtU3RhdGVUeXBlEggKBFNIT1cQABIICgRISURFEAE=');
@$core.Deprecated('Use emptyDescriptor instead')
const Empty$json = const {
  '1': 'Empty',
};

/// Descriptor for `Empty`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List emptyDescriptor = $convert.base64Decode('CgVFbXB0eQ==');
@$core.Deprecated('Use cardItemDescriptor instead')
const CardItem$json = const {
  '1': 'CardItem',
  '2': const [
    const {'1': 'cardID', '3': 1, '4': 1, '5': 4, '10': 'cardID'},
    const {'1': 'name', '3': 2, '4': 1, '5': 9, '10': 'name'},
    const {'1': 'icon', '3': 3, '4': 1, '5': 9, '10': 'icon'},
  ],
};

/// Descriptor for `CardItem`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List cardItemDescriptor = $convert.base64Decode('CghDYXJkSXRlbRIWCgZjYXJkSUQYASABKARSBmNhcmRJRBISCgRuYW1lGAIgASgJUgRuYW1lEhIKBGljb24YAyABKAlSBGljb24=');
@$core.Deprecated('Use selectedCardDescriptor instead')
const SelectedCard$json = const {
  '1': 'SelectedCard',
  '2': const [
    const {'1': 'cardID', '3': 1, '4': 1, '5': 4, '10': 'cardID'},
  ],
};

/// Descriptor for `SelectedCard`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List selectedCardDescriptor = $convert.base64Decode('CgxTZWxlY3RlZENhcmQSFgoGY2FyZElEGAEgASgEUgZjYXJkSUQ=');
@$core.Deprecated('Use actionItemDescriptor instead')
const ActionItem$json = const {
  '1': 'ActionItem',
  '2': const [
    const {'1': 'actionID', '3': 1, '4': 1, '5': 13, '10': 'actionID'},
    const {'1': 'name', '3': 2, '4': 1, '5': 9, '10': 'name'},
  ],
};

/// Descriptor for `ActionItem`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List actionItemDescriptor = $convert.base64Decode('CgpBY3Rpb25JdGVtEhoKCGFjdGlvbklEGAEgASgNUghhY3Rpb25JRBISCgRuYW1lGAIgASgJUgRuYW1l');
@$core.Deprecated('Use actionsDescriptor instead')
const Actions$json = const {
  '1': 'Actions',
  '2': const [
    const {'1': 'items', '3': 1, '4': 3, '5': 11, '6': '.runify.ActionItem', '10': 'items'},
  ],
};

/// Descriptor for `Actions`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List actionsDescriptor = $convert.base64Decode('CgdBY3Rpb25zEigKBWl0ZW1zGAEgAygLMhIucnVuaWZ5LkFjdGlvbkl0ZW1SBWl0ZW1z');
@$core.Deprecated('Use selectedActionDescriptor instead')
const SelectedAction$json = const {
  '1': 'SelectedAction',
  '2': const [
    const {'1': 'actionID', '3': 1, '4': 1, '5': 13, '10': 'actionID'},
    const {'1': 'cardID', '3': 2, '4': 1, '5': 4, '10': 'cardID'},
  ],
};

/// Descriptor for `SelectedAction`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List selectedActionDescriptor = $convert.base64Decode('Cg5TZWxlY3RlZEFjdGlvbhIaCghhY3Rpb25JRBgBIAEoDVIIYWN0aW9uSUQSFgoGY2FyZElEGAIgASgEUgZjYXJkSUQ=');
@$core.Deprecated('Use formDescriptor instead')
const Form$json = const {
  '1': 'Form',
  '2': const [
    const {'1': 'layout', '3': 1, '4': 1, '5': 9, '10': 'layout'},
    const {'1': 'cards', '3': 2, '4': 3, '5': 11, '6': '.runify.CardItem', '10': 'cards'},
  ],
};

/// Descriptor for `Form`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List formDescriptor = $convert.base64Decode('CgRGb3JtEhYKBmxheW91dBgBIAEoCVIGbGF5b3V0EiYKBWNhcmRzGAIgAygLMhAucnVuaWZ5LkNhcmRJdGVtUgVjYXJkcw==');
@$core.Deprecated('Use showWindowDescriptor instead')
const ShowWindow$json = const {
  '1': 'ShowWindow',
};

/// Descriptor for `ShowWindow`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List showWindowDescriptor = $convert.base64Decode('CgpTaG93V2luZG93');
@$core.Deprecated('Use hideWindowDescriptor instead')
const HideWindow$json = const {
  '1': 'HideWindow',
  '2': const [
    const {'1': 'error', '3': 1, '4': 1, '5': 9, '10': 'error'},
  ],
};

/// Descriptor for `HideWindow`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List hideWindowDescriptor = $convert.base64Decode('CgpIaWRlV2luZG93EhQKBWVycm9yGAEgASgJUgVlcnJvcg==');
@$core.Deprecated('Use resultDescriptor instead')
const Result$json = const {
  '1': 'Result',
  '2': const [
    const {'1': 'form', '3': 1, '4': 1, '5': 11, '6': '.runify.Form', '9': 0, '10': 'form'},
    const {'1': 'empty', '3': 2, '4': 1, '5': 11, '6': '.runify.Empty', '9': 0, '10': 'empty'},
    const {'1': 'hide', '3': 3, '4': 1, '5': 11, '6': '.runify.HideWindow', '9': 0, '10': 'hide'},
  ],
  '8': const [
    const {'1': 'payload'},
  ],
};

/// Descriptor for `Result`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List resultDescriptor = $convert.base64Decode('CgZSZXN1bHQSIgoEZm9ybRgBIAEoCzIMLnJ1bmlmeS5Gb3JtSABSBGZvcm0SJQoFZW1wdHkYAiABKAsyDS5ydW5pZnkuRW1wdHlIAFIFZW1wdHkSKAoEaGlkZRgDIAEoCzISLnJ1bmlmeS5IaWRlV2luZG93SABSBGhpZGVCCQoHcGF5bG9hZA==');
@$core.Deprecated('Use writeLogDescriptor instead')
const WriteLog$json = const {
  '1': 'WriteLog',
  '2': const [
    const {'1': 'level', '3': 1, '4': 1, '5': 14, '6': '.runify.LogLevel', '10': 'level'},
    const {'1': 'message', '3': 2, '4': 1, '5': 9, '10': 'message'},
  ],
};

/// Descriptor for `WriteLog`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List writeLogDescriptor = $convert.base64Decode('CghXcml0ZUxvZxImCgVsZXZlbBgBIAEoDjIQLnJ1bmlmeS5Mb2dMZXZlbFIFbGV2ZWwSGAoHbWVzc2FnZRgCIAEoCVIHbWVzc2FnZQ==');
@$core.Deprecated('Use setFormStateDescriptor instead')
const SetFormState$json = const {
  '1': 'SetFormState',
  '2': const [
    const {'1': 'state', '3': 1, '4': 1, '5': 14, '6': '.runify.FormStateType', '10': 'state'},
  ],
};

/// Descriptor for `SetFormState`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List setFormStateDescriptor = $convert.base64Decode('CgxTZXRGb3JtU3RhdGUSKwoFc3RhdGUYASABKA4yFS5ydW5pZnkuRm9ybVN0YXRlVHlwZVIFc3RhdGU=');
@$core.Deprecated('Use serviceMsgUIDescriptor instead')
const ServiceMsgUI$json = const {
  '1': 'ServiceMsgUI',
  '2': const [
    const {'1': 'writeLog', '3': 1, '4': 1, '5': 11, '6': '.runify.WriteLog', '9': 0, '10': 'writeLog'},
  ],
  '8': const [
    const {'1': 'payload'},
  ],
};

/// Descriptor for `ServiceMsgUI`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List serviceMsgUIDescriptor = $convert.base64Decode('CgxTZXJ2aWNlTXNnVUkSLgoId3JpdGVMb2cYASABKAsyEC5ydW5pZnkuV3JpdGVMb2dIAFIId3JpdGVMb2dCCQoHcGF5bG9hZA==');
@$core.Deprecated('Use serviceMsgSrvDescriptor instead')
const ServiceMsgSrv$json = const {
  '1': 'ServiceMsgSrv',
  '2': const [
    const {'1': 'setFormState', '3': 1, '4': 1, '5': 11, '6': '.runify.SetFormState', '9': 0, '10': 'setFormState'},
  ],
  '8': const [
    const {'1': 'payload'},
  ],
};

/// Descriptor for `ServiceMsgSrv`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List serviceMsgSrvDescriptor = $convert.base64Decode('Cg1TZXJ2aWNlTXNnU3J2EjoKDHNldEZvcm1TdGF0ZRgBIAEoCzIULnJ1bmlmeS5TZXRGb3JtU3RhdGVIAFIMc2V0Rm9ybVN0YXRlQgkKB3BheWxvYWQ=');
