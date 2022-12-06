///
//  Generated code. Do not modify.
//  source: runify.proto
//
// @dart = 2.12
// ignore_for_file: annotate_overrides,camel_case_types,constant_identifier_names,directives_ordering,library_prefixes,non_constant_identifier_names,prefer_final_fields,return_of_invalid_type,unnecessary_const,unnecessary_import,unnecessary_this,unused_import,unused_shown_name

// ignore_for_file: UNDEFINED_SHOWN_NAME
import 'dart:core' as $core;
import 'package:protobuf/protobuf.dart' as $pb;

class LogLevel extends $pb.ProtobufEnum {
  static const LogLevel DEBUG = LogLevel._(0, const $core.bool.fromEnvironment('protobuf.omit_enum_names') ? '' : 'DEBUG');
  static const LogLevel INFO = LogLevel._(1, const $core.bool.fromEnvironment('protobuf.omit_enum_names') ? '' : 'INFO');
  static const LogLevel WARNING = LogLevel._(2, const $core.bool.fromEnvironment('protobuf.omit_enum_names') ? '' : 'WARNING');
  static const LogLevel ERROR = LogLevel._(3, const $core.bool.fromEnvironment('protobuf.omit_enum_names') ? '' : 'ERROR');

  static const $core.List<LogLevel> values = <LogLevel> [
    DEBUG,
    INFO,
    WARNING,
    ERROR,
  ];

  static final $core.Map<$core.int, LogLevel> _byValue = $pb.ProtobufEnum.initByValue(values);
  static LogLevel? valueOf($core.int value) => _byValue[value];

  const LogLevel._($core.int v, $core.String n) : super(v, n);
}

class FormStateType extends $pb.ProtobufEnum {
  static const FormStateType SHOW = FormStateType._(0, const $core.bool.fromEnvironment('protobuf.omit_enum_names') ? '' : 'SHOW');
  static const FormStateType HIDE = FormStateType._(1, const $core.bool.fromEnvironment('protobuf.omit_enum_names') ? '' : 'HIDE');

  static const $core.List<FormStateType> values = <FormStateType> [
    SHOW,
    HIDE,
  ];

  static final $core.Map<$core.int, FormStateType> _byValue = $pb.ProtobufEnum.initByValue(values);
  static FormStateType? valueOf($core.int value) => _byValue[value];

  const FormStateType._($core.int v, $core.String n) : super(v, n);
}

class MessageType extends $pb.ProtobufEnum {
  static const MessageType TYPE_ERROR = MessageType._(0, const $core.bool.fromEnvironment('protobuf.omit_enum_names') ? '' : 'TYPE_ERROR');

  static const $core.List<MessageType> values = <MessageType> [
    TYPE_ERROR,
  ];

  static final $core.Map<$core.int, MessageType> _byValue = $pb.ProtobufEnum.initByValue(values);
  static MessageType? valueOf($core.int value) => _byValue[value];

  const MessageType._($core.int v, $core.String n) : super(v, n);
}

class FormActionType extends $pb.ProtobufEnum {
  static const FormActionType CLOSE_ALL = FormActionType._(0, const $core.bool.fromEnvironment('protobuf.omit_enum_names') ? '' : 'CLOSE_ALL');
  static const FormActionType CLOSE_ONE = FormActionType._(1, const $core.bool.fromEnvironment('protobuf.omit_enum_names') ? '' : 'CLOSE_ONE');
  static const FormActionType SHOW_MESSAGE = FormActionType._(2, const $core.bool.fromEnvironment('protobuf.omit_enum_names') ? '' : 'SHOW_MESSAGE');

  static const $core.List<FormActionType> values = <FormActionType> [
    CLOSE_ALL,
    CLOSE_ONE,
    SHOW_MESSAGE,
  ];

  static final $core.Map<$core.int, FormActionType> _byValue = $pb.ProtobufEnum.initByValue(values);
  static FormActionType? valueOf($core.int value) => _byValue[value];

  const FormActionType._($core.int v, $core.String n) : super(v, n);
}

