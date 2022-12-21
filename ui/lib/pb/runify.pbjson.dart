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
@$core.Deprecated('Use rootListRowTypeDescriptor instead')
const RootListRowType$json = const {
  '1': 'RootListRowType',
  '2': const [
    const {'1': 'TYPE_CALC', '2': 0},
    const {'1': 'TYPE_APP', '2': 1},
    const {'1': 'TYPE_OTHER', '2': 2},
  ],
};

/// Descriptor for `RootListRowType`. Decode as a `google.protobuf.EnumDescriptorProto`.
final $typed_data.Uint8List rootListRowTypeDescriptor = $convert.base64Decode('Cg9Sb290TGlzdFJvd1R5cGUSDQoJVFlQRV9DQUxDEAASDAoIVFlQRV9BUFAQARIOCgpUWVBFX09USEVSEAI=');
@$core.Deprecated('Use messageTypeDescriptor instead')
const MessageType$json = const {
  '1': 'MessageType',
  '2': const [
    const {'1': 'TYPE_ERROR', '2': 0},
  ],
};

/// Descriptor for `MessageType`. Decode as a `google.protobuf.EnumDescriptorProto`.
final $typed_data.Uint8List messageTypeDescriptor = $convert.base64Decode('CgtNZXNzYWdlVHlwZRIOCgpUWVBFX0VSUk9SEAA=');
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
@$core.Deprecated('Use filterDataDescriptor instead')
const FilterData$json = const {
  '1': 'FilterData',
  '2': const [
    const {'1': 'value', '3': 1, '4': 1, '5': 9, '10': 'value'},
  ],
};

/// Descriptor for `FilterData`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List filterDataDescriptor = $convert.base64Decode('CgpGaWx0ZXJEYXRhEhQKBXZhbHVlGAEgASgJUgV2YWx1ZQ==');
@$core.Deprecated('Use rootListRowGlobalIDDescriptor instead')
const RootListRowGlobalID$json = const {
  '1': 'RootListRowGlobalID',
  '2': const [
    const {'1': 'providerID', '3': 1, '4': 1, '5': 13, '10': 'providerID'},
    const {'1': 'rowID', '3': 2, '4': 1, '5': 13, '10': 'rowID'},
  ],
};

/// Descriptor for `RootListRowGlobalID`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List rootListRowGlobalIDDescriptor = $convert.base64Decode('ChNSb290TGlzdFJvd0dsb2JhbElEEh4KCnByb3ZpZGVySUQYASABKA1SCnByb3ZpZGVySUQSFAoFcm93SUQYAiABKA1SBXJvd0lE');
@$core.Deprecated('Use rootListRowDescriptor instead')
const RootListRow$json = const {
  '1': 'RootListRow',
  '2': const [
    const {'1': 'rowType', '3': 1, '4': 1, '5': 14, '6': '.runify.RootListRowType', '10': 'rowType'},
    const {'1': 'providerID', '3': 2, '4': 1, '5': 13, '10': 'providerID'},
    const {'1': 'rowID', '3': 3, '4': 1, '5': 13, '10': 'rowID'},
    const {'1': 'priority', '3': 4, '4': 1, '5': 13, '10': 'priority'},
    const {'1': 'icon', '3': 5, '4': 1, '5': 9, '10': 'icon'},
    const {'1': 'value', '3': 6, '4': 1, '5': 9, '10': 'value'},
  ],
};

/// Descriptor for `RootListRow`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List rootListRowDescriptor = $convert.base64Decode('CgtSb290TGlzdFJvdxIxCgdyb3dUeXBlGAEgASgOMhcucnVuaWZ5LlJvb3RMaXN0Um93VHlwZVIHcm93VHlwZRIeCgpwcm92aWRlcklEGAIgASgNUgpwcm92aWRlcklEEhQKBXJvd0lEGAMgASgNUgVyb3dJRBIaCghwcmlvcml0eRgEIAEoDVIIcHJpb3JpdHkSEgoEaWNvbhgFIAEoCVIEaWNvbhIUCgV2YWx1ZRgGIAEoCVIFdmFsdWU=');
@$core.Deprecated('Use rootListOpenDescriptor instead')
const RootListOpen$json = const {
  '1': 'RootListOpen',
  '2': const [
    const {'1': 'rows', '3': 1, '4': 3, '5': 11, '6': '.runify.RootListRow', '10': 'rows'},
  ],
};

/// Descriptor for `RootListOpen`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List rootListOpenDescriptor = $convert.base64Decode('CgxSb290TGlzdE9wZW4SJwoEcm93cxgBIAMoCzITLnJ1bmlmeS5Sb290TGlzdFJvd1IEcm93cw==');
@$core.Deprecated('Use rootListAddRowsDescriptor instead')
const RootListAddRows$json = const {
  '1': 'RootListAddRows',
  '2': const [
    const {'1': 'rows', '3': 1, '4': 3, '5': 11, '6': '.runify.RootListRow', '10': 'rows'},
  ],
};

/// Descriptor for `RootListAddRows`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List rootListAddRowsDescriptor = $convert.base64Decode('Cg9Sb290TGlzdEFkZFJvd3MSJwoEcm93cxgBIAMoCzITLnJ1bmlmeS5Sb290TGlzdFJvd1IEcm93cw==');
@$core.Deprecated('Use rootListChangeRowsDescriptor instead')
const RootListChangeRows$json = const {
  '1': 'RootListChangeRows',
  '2': const [
    const {'1': 'rows', '3': 1, '4': 3, '5': 11, '6': '.runify.RootListRow', '10': 'rows'},
  ],
};

/// Descriptor for `RootListChangeRows`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List rootListChangeRowsDescriptor = $convert.base64Decode('ChJSb290TGlzdENoYW5nZVJvd3MSJwoEcm93cxgBIAMoCzITLnJ1bmlmeS5Sb290TGlzdFJvd1IEcm93cw==');
@$core.Deprecated('Use rootListRemoveRowsDescriptor instead')
const RootListRemoveRows$json = const {
  '1': 'RootListRemoveRows',
  '2': const [
    const {'1': 'rows', '3': 1, '4': 3, '5': 11, '6': '.runify.RootListRowGlobalID', '10': 'rows'},
  ],
};

/// Descriptor for `RootListRemoveRows`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List rootListRemoveRowsDescriptor = $convert.base64Decode('ChJSb290TGlzdFJlbW92ZVJvd3MSLwoEcm93cxgBIAMoCzIbLnJ1bmlmeS5Sb290TGlzdFJvd0dsb2JhbElEUgRyb3dz');
@$core.Deprecated('Use contextMenuRowIDDescriptor instead')
const ContextMenuRowID$json = const {
  '1': 'ContextMenuRowID',
  '2': const [
    const {'1': 'rowID', '3': 1, '4': 1, '5': 13, '10': 'rowID'},
  ],
};

/// Descriptor for `ContextMenuRowID`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List contextMenuRowIDDescriptor = $convert.base64Decode('ChBDb250ZXh0TWVudVJvd0lEEhQKBXJvd0lEGAEgASgNUgVyb3dJRA==');
@$core.Deprecated('Use formClosedDescriptor instead')
const FormClosed$json = const {
  '1': 'FormClosed',
};

/// Descriptor for `FormClosed`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List formClosedDescriptor = $convert.base64Decode('CgpGb3JtQ2xvc2Vk');
@$core.Deprecated('Use contextMenuRowDescriptor instead')
const ContextMenuRow$json = const {
  '1': 'ContextMenuRow',
  '2': const [
    const {'1': 'rowID', '3': 1, '4': 1, '5': 13, '10': 'rowID'},
    const {'1': 'value', '3': 2, '4': 1, '5': 9, '10': 'value'},
  ],
};

/// Descriptor for `ContextMenuRow`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List contextMenuRowDescriptor = $convert.base64Decode('Cg5Db250ZXh0TWVudVJvdxIUCgVyb3dJRBgBIAEoDVIFcm93SUQSFAoFdmFsdWUYAiABKAlSBXZhbHVl');
@$core.Deprecated('Use contextMenuOpenDescriptor instead')
const ContextMenuOpen$json = const {
  '1': 'ContextMenuOpen',
  '2': const [
    const {'1': 'rows', '3': 1, '4': 3, '5': 11, '6': '.runify.ContextMenuRow', '10': 'rows'},
  ],
};

/// Descriptor for `ContextMenuOpen`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List contextMenuOpenDescriptor = $convert.base64Decode('Cg9Db250ZXh0TWVudU9wZW4SKgoEcm93cxgBIAMoCzIWLnJ1bmlmeS5Db250ZXh0TWVudVJvd1IEcm93cw==');
@$core.Deprecated('Use userMessageDescriptor instead')
const UserMessage$json = const {
  '1': 'UserMessage',
  '2': const [
    const {'1': 'messageType', '3': 1, '4': 1, '5': 14, '6': '.runify.MessageType', '10': 'messageType'},
    const {'1': 'message', '3': 2, '4': 1, '5': 9, '10': 'message'},
  ],
};

/// Descriptor for `UserMessage`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List userMessageDescriptor = $convert.base64Decode('CgtVc2VyTWVzc2FnZRI1CgttZXNzYWdlVHlwZRgBIAEoDjITLnJ1bmlmeS5NZXNzYWdlVHlwZVILbWVzc2FnZVR5cGUSGAoHbWVzc2FnZRgCIAEoCVIHbWVzc2FnZQ==');
@$core.Deprecated('Use closeFormDescriptor instead')
const CloseForm$json = const {
  '1': 'CloseForm',
};

/// Descriptor for `CloseForm`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List closeFormDescriptor = $convert.base64Decode('CglDbG9zZUZvcm0=');
@$core.Deprecated('Use hideUIDescriptor instead')
const HideUI$json = const {
  '1': 'HideUI',
  '2': const [
    const {'1': 'message', '3': 1, '4': 1, '5': 11, '6': '.runify.UserMessage', '10': 'message'},
  ],
};

/// Descriptor for `HideUI`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List hideUIDescriptor = $convert.base64Decode('CgZIaWRlVUkSLQoHbWVzc2FnZRgBIAEoCzITLnJ1bmlmeS5Vc2VyTWVzc2FnZVIHbWVzc2FnZQ==');
@$core.Deprecated('Use closeUIDescriptor instead')
const CloseUI$json = const {
  '1': 'CloseUI',
};

/// Descriptor for `CloseUI`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List closeUIDescriptor = $convert.base64Decode('CgdDbG9zZVVJ');
@$core.Deprecated('Use uIMessageDescriptor instead')
const UIMessage$json = const {
  '1': 'UIMessage',
  '2': const [
    const {'1': 'formID', '3': 1, '4': 1, '5': 13, '10': 'formID'},
    const {'1': 'writeLog', '3': 2, '4': 1, '5': 11, '6': '.runify.WriteLog', '9': 0, '10': 'writeLog'},
    const {'1': 'filterChanged', '3': 3, '4': 1, '5': 11, '6': '.runify.FilterData', '9': 0, '10': 'filterChanged'},
    const {'1': 'rootListRowActivated', '3': 4, '4': 1, '5': 11, '6': '.runify.RootListRowGlobalID', '9': 0, '10': 'rootListRowActivated'},
    const {'1': 'rootListMenuActivated', '3': 5, '4': 1, '5': 11, '6': '.runify.RootListRowGlobalID', '9': 0, '10': 'rootListMenuActivated'},
    const {'1': 'contextMenuRowActivated', '3': 6, '4': 1, '5': 11, '6': '.runify.ContextMenuRowID', '9': 0, '10': 'contextMenuRowActivated'},
    const {'1': 'formClosed', '3': 7, '4': 1, '5': 11, '6': '.runify.FormClosed', '9': 0, '10': 'formClosed'},
  ],
  '8': const [
    const {'1': 'payload'},
  ],
};

/// Descriptor for `UIMessage`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List uIMessageDescriptor = $convert.base64Decode('CglVSU1lc3NhZ2USFgoGZm9ybUlEGAEgASgNUgZmb3JtSUQSLgoId3JpdGVMb2cYAiABKAsyEC5ydW5pZnkuV3JpdGVMb2dIAFIId3JpdGVMb2cSOgoNZmlsdGVyQ2hhbmdlZBgDIAEoCzISLnJ1bmlmeS5GaWx0ZXJEYXRhSABSDWZpbHRlckNoYW5nZWQSUQoUcm9vdExpc3RSb3dBY3RpdmF0ZWQYBCABKAsyGy5ydW5pZnkuUm9vdExpc3RSb3dHbG9iYWxJREgAUhRyb290TGlzdFJvd0FjdGl2YXRlZBJTChVyb290TGlzdE1lbnVBY3RpdmF0ZWQYBSABKAsyGy5ydW5pZnkuUm9vdExpc3RSb3dHbG9iYWxJREgAUhVyb290TGlzdE1lbnVBY3RpdmF0ZWQSVAoXY29udGV4dE1lbnVSb3dBY3RpdmF0ZWQYBiABKAsyGC5ydW5pZnkuQ29udGV4dE1lbnVSb3dJREgAUhdjb250ZXh0TWVudVJvd0FjdGl2YXRlZBI0Cgpmb3JtQ2xvc2VkGAcgASgLMhIucnVuaWZ5LkZvcm1DbG9zZWRIAFIKZm9ybUNsb3NlZEIJCgdwYXlsb2Fk');
@$core.Deprecated('Use srvMessageDescriptor instead')
const SrvMessage$json = const {
  '1': 'SrvMessage',
  '2': const [
    const {'1': 'formID', '3': 1, '4': 1, '5': 13, '10': 'formID'},
    const {'1': 'rootListOpen', '3': 2, '4': 1, '5': 11, '6': '.runify.RootListOpen', '9': 0, '10': 'rootListOpen'},
    const {'1': 'rootListAddRows', '3': 3, '4': 1, '5': 11, '6': '.runify.RootListAddRows', '9': 0, '10': 'rootListAddRows'},
    const {'1': 'rootListChangeRows', '3': 4, '4': 1, '5': 11, '6': '.runify.RootListChangeRows', '9': 0, '10': 'rootListChangeRows'},
    const {'1': 'rootListRemoveRows', '3': 5, '4': 1, '5': 11, '6': '.runify.RootListRemoveRows', '9': 0, '10': 'rootListRemoveRows'},
    const {'1': 'contextMenuOpen', '3': 6, '4': 1, '5': 11, '6': '.runify.ContextMenuOpen', '9': 0, '10': 'contextMenuOpen'},
    const {'1': 'userMessage', '3': 7, '4': 1, '5': 11, '6': '.runify.UserMessage', '9': 0, '10': 'userMessage'},
    const {'1': 'closeForm', '3': 8, '4': 1, '5': 11, '6': '.runify.CloseForm', '9': 0, '10': 'closeForm'},
    const {'1': 'hideUI', '3': 9, '4': 1, '5': 11, '6': '.runify.HideUI', '9': 0, '10': 'hideUI'},
    const {'1': 'closeUI', '3': 10, '4': 1, '5': 11, '6': '.runify.CloseUI', '9': 0, '10': 'closeUI'},
  ],
  '8': const [
    const {'1': 'payload'},
  ],
};

/// Descriptor for `SrvMessage`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List srvMessageDescriptor = $convert.base64Decode('CgpTcnZNZXNzYWdlEhYKBmZvcm1JRBgBIAEoDVIGZm9ybUlEEjoKDHJvb3RMaXN0T3BlbhgCIAEoCzIULnJ1bmlmeS5Sb290TGlzdE9wZW5IAFIMcm9vdExpc3RPcGVuEkMKD3Jvb3RMaXN0QWRkUm93cxgDIAEoCzIXLnJ1bmlmeS5Sb290TGlzdEFkZFJvd3NIAFIPcm9vdExpc3RBZGRSb3dzEkwKEnJvb3RMaXN0Q2hhbmdlUm93cxgEIAEoCzIaLnJ1bmlmeS5Sb290TGlzdENoYW5nZVJvd3NIAFIScm9vdExpc3RDaGFuZ2VSb3dzEkwKEnJvb3RMaXN0UmVtb3ZlUm93cxgFIAEoCzIaLnJ1bmlmeS5Sb290TGlzdFJlbW92ZVJvd3NIAFIScm9vdExpc3RSZW1vdmVSb3dzEkMKD2NvbnRleHRNZW51T3BlbhgGIAEoCzIXLnJ1bmlmeS5Db250ZXh0TWVudU9wZW5IAFIPY29udGV4dE1lbnVPcGVuEjcKC3VzZXJNZXNzYWdlGAcgASgLMhMucnVuaWZ5LlVzZXJNZXNzYWdlSABSC3VzZXJNZXNzYWdlEjEKCWNsb3NlRm9ybRgIIAEoCzIRLnJ1bmlmeS5DbG9zZUZvcm1IAFIJY2xvc2VGb3JtEigKBmhpZGVVSRgJIAEoCzIOLnJ1bmlmeS5IaWRlVUlIAFIGaGlkZVVJEisKB2Nsb3NlVUkYCiABKAsyDy5ydW5pZnkuQ2xvc2VVSUgAUgdjbG9zZVVJQgkKB3BheWxvYWQ=');
