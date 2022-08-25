///
//  Generated code. Do not modify.
//  source: runify.proto
//
// @dart = 2.12
// ignore_for_file: annotate_overrides,camel_case_types,constant_identifier_names,deprecated_member_use_from_same_package,directives_ordering,library_prefixes,non_constant_identifier_names,prefer_final_fields,return_of_invalid_type,unnecessary_const,unnecessary_import,unnecessary_this,unused_import,unused_shown_name

import 'dart:core' as $core;
import 'dart:convert' as $convert;
import 'dart:typed_data' as $typed_data;
@$core.Deprecated('Use windowStateDescriptor instead')
const WindowState$json = const {
  '1': 'WindowState',
  '2': const [
    const {'1': 'SHOW', '2': 0},
    const {'1': 'HIDE', '2': 1},
  ],
};

/// Descriptor for `WindowState`. Decode as a `google.protobuf.EnumDescriptorProto`.
final $typed_data.Uint8List windowStateDescriptor = $convert.base64Decode('CgtXaW5kb3dTdGF0ZRIICgRTSE9XEAASCAoESElERRAB');
@$core.Deprecated('Use emptyDescriptor instead')
const Empty$json = const {
  '1': 'Empty',
};

/// Descriptor for `Empty`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List emptyDescriptor = $convert.base64Decode('CgVFbXB0eQ==');
@$core.Deprecated('Use selectedActionDescriptor instead')
const SelectedAction$json = const {
  '1': 'SelectedAction',
  '2': const [
    const {'1': 'actionID', '3': 1, '4': 1, '5': 13, '10': 'actionID'},
    const {'1': 'commandID', '3': 2, '4': 1, '5': 4, '10': 'commandID'},
  ],
};

/// Descriptor for `SelectedAction`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List selectedActionDescriptor = $convert.base64Decode('Cg5TZWxlY3RlZEFjdGlvbhIaCghhY3Rpb25JRBgBIAEoDVIIYWN0aW9uSUQSHAoJY29tbWFuZElEGAIgASgEUgljb21tYW5kSUQ=');
@$core.Deprecated('Use selectedCommandDescriptor instead')
const SelectedCommand$json = const {
  '1': 'SelectedCommand',
  '2': const [
    const {'1': 'commandID', '3': 1, '4': 1, '5': 4, '10': 'commandID'},
  ],
};

/// Descriptor for `SelectedCommand`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List selectedCommandDescriptor = $convert.base64Decode('Cg9TZWxlY3RlZENvbW1hbmQSHAoJY29tbWFuZElEGAEgASgEUgljb21tYW5kSUQ=');
@$core.Deprecated('Use actionDescriptor instead')
const Action$json = const {
  '1': 'Action',
  '2': const [
    const {'1': 'id', '3': 1, '4': 1, '5': 13, '10': 'id'},
    const {'1': 'name', '3': 2, '4': 1, '5': 9, '10': 'name'},
  ],
};

/// Descriptor for `Action`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List actionDescriptor = $convert.base64Decode('CgZBY3Rpb24SDgoCaWQYASABKA1SAmlkEhIKBG5hbWUYAiABKAlSBG5hbWU=');
@$core.Deprecated('Use actionsDescriptor instead')
const Actions$json = const {
  '1': 'Actions',
  '2': const [
    const {'1': 'data', '3': 1, '4': 3, '5': 11, '6': '.runify.Action', '10': 'data'},
  ],
};

/// Descriptor for `Actions`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List actionsDescriptor = $convert.base64Decode('CgdBY3Rpb25zEiIKBGRhdGEYASADKAsyDi5ydW5pZnkuQWN0aW9uUgRkYXRh');
@$core.Deprecated('Use commandDescriptor instead')
const Command$json = const {
  '1': 'Command',
  '2': const [
    const {'1': 'id', '3': 1, '4': 1, '5': 4, '10': 'id'},
    const {'1': 'name', '3': 2, '4': 1, '5': 9, '10': 'name'},
    const {'1': 'icon', '3': 3, '4': 1, '5': 9, '10': 'icon'},
  ],
};

/// Descriptor for `Command`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List commandDescriptor = $convert.base64Decode('CgdDb21tYW5kEg4KAmlkGAEgASgEUgJpZBISCgRuYW1lGAIgASgJUgRuYW1lEhIKBGljb24YAyABKAlSBGljb24=');
@$core.Deprecated('Use commandsDescriptor instead')
const Commands$json = const {
  '1': 'Commands',
  '2': const [
    const {'1': 'data', '3': 1, '4': 3, '5': 11, '6': '.runify.Command', '10': 'data'},
  ],
};

/// Descriptor for `Commands`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List commandsDescriptor = $convert.base64Decode('CghDb21tYW5kcxIjCgRkYXRhGAEgAygLMg8ucnVuaWZ5LkNvbW1hbmRSBGRhdGE=');
@$core.Deprecated('Use formDescriptor instead')
const Form$json = const {
  '1': 'Form',
};

/// Descriptor for `Form`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List formDescriptor = $convert.base64Decode('CgRGb3Jt');
@$core.Deprecated('Use messageDescriptor instead')
const Message$json = const {
  '1': 'Message',
  '2': const [
    const {'1': 'text', '3': 1, '4': 1, '5': 9, '10': 'text'},
  ],
};

/// Descriptor for `Message`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List messageDescriptor = $convert.base64Decode('CgdNZXNzYWdlEhIKBHRleHQYASABKAlSBHRleHQ=');
@$core.Deprecated('Use resultDescriptor instead')
const Result$json = const {
  '1': 'Result',
  '2': const [
    const {'1': 'commands', '3': 1, '4': 1, '5': 11, '6': '.runify.Commands', '9': 0, '10': 'commands'},
    const {'1': 'form', '3': 2, '4': 1, '5': 11, '6': '.runify.Form', '9': 0, '10': 'form'},
    const {'1': 'message', '3': 3, '4': 1, '5': 11, '6': '.runify.Message', '9': 0, '10': 'message'},
    const {'1': 'winState', '3': 4, '4': 1, '5': 14, '6': '.runify.WindowState', '9': 0, '10': 'winState'},
  ],
  '8': const [
    const {'1': 'payload'},
  ],
};

/// Descriptor for `Result`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List resultDescriptor = $convert.base64Decode('CgZSZXN1bHQSLgoIY29tbWFuZHMYASABKAsyEC5ydW5pZnkuQ29tbWFuZHNIAFIIY29tbWFuZHMSIgoEZm9ybRgCIAEoCzIMLnJ1bmlmeS5Gb3JtSABSBGZvcm0SKwoHbWVzc2FnZRgDIAEoCzIPLnJ1bmlmeS5NZXNzYWdlSABSB21lc3NhZ2USMQoId2luU3RhdGUYBCABKA4yEy5ydW5pZnkuV2luZG93U3RhdGVIAFIId2luU3RhdGVCCQoHcGF5bG9hZA==');
