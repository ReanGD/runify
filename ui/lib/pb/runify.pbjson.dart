///
//  Generated code. Do not modify.
//  source: runify.proto
//
// @dart = 2.12
// ignore_for_file: annotate_overrides,camel_case_types,constant_identifier_names,deprecated_member_use_from_same_package,directives_ordering,library_prefixes,non_constant_identifier_names,prefer_final_fields,return_of_invalid_type,unnecessary_const,unnecessary_import,unnecessary_this,unused_import,unused_shown_name

import 'dart:core' as $core;
import 'dart:convert' as $convert;
import 'dart:typed_data' as $typed_data;
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
