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
@$core.Deprecated('Use messageTypeDescriptor instead')
const MessageType$json = const {
  '1': 'MessageType',
  '2': const [
    const {'1': 'TYPE_ERROR', '2': 0},
  ],
};

/// Descriptor for `MessageType`. Decode as a `google.protobuf.EnumDescriptorProto`.
final $typed_data.Uint8List messageTypeDescriptor = $convert.base64Decode('CgtNZXNzYWdlVHlwZRIOCgpUWVBFX0VSUk9SEAA=');
@$core.Deprecated('Use formActionTypeDescriptor instead')
const FormActionType$json = const {
  '1': 'FormActionType',
  '2': const [
    const {'1': 'CLOSE_ALL', '2': 0},
    const {'1': 'CLOSE_ONE', '2': 1},
    const {'1': 'SHOW_MESSAGE', '2': 2},
  ],
};

/// Descriptor for `FormActionType`. Decode as a `google.protobuf.EnumDescriptorProto`.
final $typed_data.Uint8List formActionTypeDescriptor = $convert.base64Decode('Cg5Gb3JtQWN0aW9uVHlwZRINCglDTE9TRV9BTEwQABINCglDTE9TRV9PTkUQARIQCgxTSE9XX01FU1NBR0UQAg==');
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
    const {'1': 'providerID', '3': 1, '4': 1, '5': 13, '10': 'providerID'},
    const {'1': 'rowID', '3': 2, '4': 1, '5': 13, '10': 'rowID'},
    const {'1': 'priority', '3': 3, '4': 1, '5': 13, '10': 'priority'},
    const {'1': 'icon', '3': 4, '4': 1, '5': 9, '10': 'icon'},
    const {'1': 'value', '3': 5, '4': 1, '5': 9, '10': 'value'},
  ],
};

/// Descriptor for `RootListRow`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List rootListRowDescriptor = $convert.base64Decode('CgtSb290TGlzdFJvdxIeCgpwcm92aWRlcklEGAEgASgNUgpwcm92aWRlcklEEhQKBXJvd0lEGAIgASgNUgVyb3dJRBIaCghwcmlvcml0eRgDIAEoDVIIcHJpb3JpdHkSEgoEaWNvbhgEIAEoCVIEaWNvbhIUCgV2YWx1ZRgFIAEoCVIFdmFsdWU=');
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
@$core.Deprecated('Use formActionDescriptor instead')
const FormAction$json = const {
  '1': 'FormAction',
  '2': const [
    const {'1': 'actionType', '3': 1, '4': 1, '5': 14, '6': '.runify.FormActionType', '10': 'actionType'},
    const {'1': 'message', '3': 2, '4': 1, '5': 11, '6': '.runify.UserMessage', '10': 'message'},
  ],
};

/// Descriptor for `FormAction`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List formActionDescriptor = $convert.base64Decode('CgpGb3JtQWN0aW9uEjYKCmFjdGlvblR5cGUYASABKA4yFi5ydW5pZnkuRm9ybUFjdGlvblR5cGVSCmFjdGlvblR5cGUSLQoHbWVzc2FnZRgCIAEoCzITLnJ1bmlmeS5Vc2VyTWVzc2FnZVIHbWVzc2FnZQ==');
@$core.Deprecated('Use filterDataDescriptor instead')
const FilterData$json = const {
  '1': 'FilterData',
  '2': const [
    const {'1': 'value', '3': 1, '4': 1, '5': 9, '10': 'value'},
  ],
};

/// Descriptor for `FilterData`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List filterDataDescriptor = $convert.base64Decode('CgpGaWx0ZXJEYXRhEhQKBXZhbHVlGAEgASgJUgV2YWx1ZQ==');
@$core.Deprecated('Use formDataMsgUIDescriptor instead')
const FormDataMsgUI$json = const {
  '1': 'FormDataMsgUI',
  '2': const [
    const {'1': 'formID', '3': 1, '4': 1, '5': 13, '10': 'formID'},
    const {'1': 'filterChanged', '3': 2, '4': 1, '5': 11, '6': '.runify.FilterData', '9': 0, '10': 'filterChanged'},
    const {'1': 'rootListRowActivated', '3': 3, '4': 1, '5': 11, '6': '.runify.RootListRowGlobalID', '9': 0, '10': 'rootListRowActivated'},
    const {'1': 'rootListMenuActivated', '3': 4, '4': 1, '5': 11, '6': '.runify.RootListRowGlobalID', '9': 0, '10': 'rootListMenuActivated'},
    const {'1': 'contextMenuRowActivated', '3': 5, '4': 1, '5': 11, '6': '.runify.ContextMenuRowID', '9': 0, '10': 'contextMenuRowActivated'},
    const {'1': 'formClosed', '3': 6, '4': 1, '5': 11, '6': '.runify.FormClosed', '9': 0, '10': 'formClosed'},
  ],
  '8': const [
    const {'1': 'payload'},
  ],
};

/// Descriptor for `FormDataMsgUI`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List formDataMsgUIDescriptor = $convert.base64Decode('Cg1Gb3JtRGF0YU1zZ1VJEhYKBmZvcm1JRBgBIAEoDVIGZm9ybUlEEjoKDWZpbHRlckNoYW5nZWQYAiABKAsyEi5ydW5pZnkuRmlsdGVyRGF0YUgAUg1maWx0ZXJDaGFuZ2VkElEKFHJvb3RMaXN0Um93QWN0aXZhdGVkGAMgASgLMhsucnVuaWZ5LlJvb3RMaXN0Um93R2xvYmFsSURIAFIUcm9vdExpc3RSb3dBY3RpdmF0ZWQSUwoVcm9vdExpc3RNZW51QWN0aXZhdGVkGAQgASgLMhsucnVuaWZ5LlJvb3RMaXN0Um93R2xvYmFsSURIAFIVcm9vdExpc3RNZW51QWN0aXZhdGVkElQKF2NvbnRleHRNZW51Um93QWN0aXZhdGVkGAUgASgLMhgucnVuaWZ5LkNvbnRleHRNZW51Um93SURIAFIXY29udGV4dE1lbnVSb3dBY3RpdmF0ZWQSNAoKZm9ybUNsb3NlZBgGIAEoCzISLnJ1bmlmeS5Gb3JtQ2xvc2VkSABSCmZvcm1DbG9zZWRCCQoHcGF5bG9hZA==');
@$core.Deprecated('Use formDataMsgSrvDescriptor instead')
const FormDataMsgSrv$json = const {
  '1': 'FormDataMsgSrv',
  '2': const [
    const {'1': 'formID', '3': 1, '4': 1, '5': 13, '10': 'formID'},
    const {'1': 'rootListOpen', '3': 2, '4': 1, '5': 11, '6': '.runify.RootListOpen', '9': 0, '10': 'rootListOpen'},
    const {'1': 'rootListAddRows', '3': 3, '4': 1, '5': 11, '6': '.runify.RootListAddRows', '9': 0, '10': 'rootListAddRows'},
    const {'1': 'rootListChangeRows', '3': 4, '4': 1, '5': 11, '6': '.runify.RootListChangeRows', '9': 0, '10': 'rootListChangeRows'},
    const {'1': 'rootListRemoveRows', '3': 5, '4': 1, '5': 11, '6': '.runify.RootListRemoveRows', '9': 0, '10': 'rootListRemoveRows'},
    const {'1': 'contextMenuOpen', '3': 6, '4': 1, '5': 11, '6': '.runify.ContextMenuOpen', '9': 0, '10': 'contextMenuOpen'},
    const {'1': 'formAction', '3': 7, '4': 1, '5': 11, '6': '.runify.FormAction', '9': 0, '10': 'formAction'},
  ],
  '8': const [
    const {'1': 'payload'},
  ],
};

/// Descriptor for `FormDataMsgSrv`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List formDataMsgSrvDescriptor = $convert.base64Decode('Cg5Gb3JtRGF0YU1zZ1NydhIWCgZmb3JtSUQYASABKA1SBmZvcm1JRBI6Cgxyb290TGlzdE9wZW4YAiABKAsyFC5ydW5pZnkuUm9vdExpc3RPcGVuSABSDHJvb3RMaXN0T3BlbhJDCg9yb290TGlzdEFkZFJvd3MYAyABKAsyFy5ydW5pZnkuUm9vdExpc3RBZGRSb3dzSABSD3Jvb3RMaXN0QWRkUm93cxJMChJyb290TGlzdENoYW5nZVJvd3MYBCABKAsyGi5ydW5pZnkuUm9vdExpc3RDaGFuZ2VSb3dzSABSEnJvb3RMaXN0Q2hhbmdlUm93cxJMChJyb290TGlzdFJlbW92ZVJvd3MYBSABKAsyGi5ydW5pZnkuUm9vdExpc3RSZW1vdmVSb3dzSABSEnJvb3RMaXN0UmVtb3ZlUm93cxJDCg9jb250ZXh0TWVudU9wZW4YBiABKAsyFy5ydW5pZnkuQ29udGV4dE1lbnVPcGVuSABSD2NvbnRleHRNZW51T3BlbhI0Cgpmb3JtQWN0aW9uGAcgASgLMhIucnVuaWZ5LkZvcm1BY3Rpb25IAFIKZm9ybUFjdGlvbkIJCgdwYXlsb2Fk');
