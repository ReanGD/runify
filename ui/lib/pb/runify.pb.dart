///
//  Generated code. Do not modify.
//  source: runify.proto
//
// @dart = 2.12
// ignore_for_file: annotate_overrides,camel_case_types,constant_identifier_names,directives_ordering,library_prefixes,non_constant_identifier_names,prefer_final_fields,return_of_invalid_type,unnecessary_const,unnecessary_import,unnecessary_this,unused_import,unused_shown_name

import 'dart:core' as $core;

import 'package:protobuf/protobuf.dart' as $pb;

import 'runify.pbenum.dart';

export 'runify.pbenum.dart';

class WriteLog extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'WriteLog', package: const $pb.PackageName(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'runify'), createEmptyInstance: create)
    ..e<LogLevel>(1, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'level', $pb.PbFieldType.OE, defaultOrMaker: LogLevel.DEBUG, valueOf: LogLevel.valueOf, enumValues: LogLevel.values)
    ..aOS(2, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'message')
    ..hasRequiredFields = false
  ;

  WriteLog._() : super();
  factory WriteLog({
    LogLevel? level,
    $core.String? message,
  }) {
    final _result = create();
    if (level != null) {
      _result.level = level;
    }
    if (message != null) {
      _result.message = message;
    }
    return _result;
  }
  factory WriteLog.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory WriteLog.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.deepCopy] instead. '
  'Will be removed in next major version')
  WriteLog clone() => WriteLog()..mergeFromMessage(this);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.rebuild] instead. '
  'Will be removed in next major version')
  WriteLog copyWith(void Function(WriteLog) updates) => super.copyWith((message) => updates(message as WriteLog)) as WriteLog; // ignore: deprecated_member_use
  $pb.BuilderInfo get info_ => _i;
  @$core.pragma('dart2js:noInline')
  static WriteLog create() => WriteLog._();
  WriteLog createEmptyInstance() => create();
  static $pb.PbList<WriteLog> createRepeated() => $pb.PbList<WriteLog>();
  @$core.pragma('dart2js:noInline')
  static WriteLog getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<WriteLog>(create);
  static WriteLog? _defaultInstance;

  @$pb.TagNumber(1)
  LogLevel get level => $_getN(0);
  @$pb.TagNumber(1)
  set level(LogLevel v) { setField(1, v); }
  @$pb.TagNumber(1)
  $core.bool hasLevel() => $_has(0);
  @$pb.TagNumber(1)
  void clearLevel() => clearField(1);

  @$pb.TagNumber(2)
  $core.String get message => $_getSZ(1);
  @$pb.TagNumber(2)
  set message($core.String v) { $_setString(1, v); }
  @$pb.TagNumber(2)
  $core.bool hasMessage() => $_has(1);
  @$pb.TagNumber(2)
  void clearMessage() => clearField(2);
}

class FormMarkup extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'FormMarkup', package: const $pb.PackageName(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'runify'), createEmptyInstance: create)
    ..aOS(1, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'json')
    ..hasRequiredFields = false
  ;

  FormMarkup._() : super();
  factory FormMarkup({
    $core.String? json,
  }) {
    final _result = create();
    if (json != null) {
      _result.json = json;
    }
    return _result;
  }
  factory FormMarkup.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory FormMarkup.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.deepCopy] instead. '
  'Will be removed in next major version')
  FormMarkup clone() => FormMarkup()..mergeFromMessage(this);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.rebuild] instead. '
  'Will be removed in next major version')
  FormMarkup copyWith(void Function(FormMarkup) updates) => super.copyWith((message) => updates(message as FormMarkup)) as FormMarkup; // ignore: deprecated_member_use
  $pb.BuilderInfo get info_ => _i;
  @$core.pragma('dart2js:noInline')
  static FormMarkup create() => FormMarkup._();
  FormMarkup createEmptyInstance() => create();
  static $pb.PbList<FormMarkup> createRepeated() => $pb.PbList<FormMarkup>();
  @$core.pragma('dart2js:noInline')
  static FormMarkup getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<FormMarkup>(create);
  static FormMarkup? _defaultInstance;

  @$pb.TagNumber(1)
  $core.String get json => $_getSZ(0);
  @$pb.TagNumber(1)
  set json($core.String v) { $_setString(0, v); }
  @$pb.TagNumber(1)
  $core.bool hasJson() => $_has(0);
  @$pb.TagNumber(1)
  void clearJson() => clearField(1);
}

class FormModel extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'FormModel', package: const $pb.PackageName(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'runify'), createEmptyInstance: create)
    ..aOS(1, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'json')
    ..hasRequiredFields = false
  ;

  FormModel._() : super();
  factory FormModel({
    $core.String? json,
  }) {
    final _result = create();
    if (json != null) {
      _result.json = json;
    }
    return _result;
  }
  factory FormModel.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory FormModel.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.deepCopy] instead. '
  'Will be removed in next major version')
  FormModel clone() => FormModel()..mergeFromMessage(this);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.rebuild] instead. '
  'Will be removed in next major version')
  FormModel copyWith(void Function(FormModel) updates) => super.copyWith((message) => updates(message as FormModel)) as FormModel; // ignore: deprecated_member_use
  $pb.BuilderInfo get info_ => _i;
  @$core.pragma('dart2js:noInline')
  static FormModel create() => FormModel._();
  FormModel createEmptyInstance() => create();
  static $pb.PbList<FormModel> createRepeated() => $pb.PbList<FormModel>();
  @$core.pragma('dart2js:noInline')
  static FormModel getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<FormModel>(create);
  static FormModel? _defaultInstance;

  @$pb.TagNumber(1)
  $core.String get json => $_getSZ(0);
  @$pb.TagNumber(1)
  set json($core.String v) { $_setString(0, v); }
  @$pb.TagNumber(1)
  $core.bool hasJson() => $_has(0);
  @$pb.TagNumber(1)
  void clearJson() => clearField(1);
}

class FormData extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'FormData', package: const $pb.PackageName(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'runify'), createEmptyInstance: create)
    ..aOS(1, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'json')
    ..hasRequiredFields = false
  ;

  FormData._() : super();
  factory FormData({
    $core.String? json,
  }) {
    final _result = create();
    if (json != null) {
      _result.json = json;
    }
    return _result;
  }
  factory FormData.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory FormData.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.deepCopy] instead. '
  'Will be removed in next major version')
  FormData clone() => FormData()..mergeFromMessage(this);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.rebuild] instead. '
  'Will be removed in next major version')
  FormData copyWith(void Function(FormData) updates) => super.copyWith((message) => updates(message as FormData)) as FormData; // ignore: deprecated_member_use
  $pb.BuilderInfo get info_ => _i;
  @$core.pragma('dart2js:noInline')
  static FormData create() => FormData._();
  FormData createEmptyInstance() => create();
  static $pb.PbList<FormData> createRepeated() => $pb.PbList<FormData>();
  @$core.pragma('dart2js:noInline')
  static FormData getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<FormData>(create);
  static FormData? _defaultInstance;

  @$pb.TagNumber(1)
  $core.String get json => $_getSZ(0);
  @$pb.TagNumber(1)
  set json($core.String v) { $_setString(0, v); }
  @$pb.TagNumber(1)
  $core.bool hasJson() => $_has(0);
  @$pb.TagNumber(1)
  void clearJson() => clearField(1);
}

class FilterData extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'FilterData', package: const $pb.PackageName(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'runify'), createEmptyInstance: create)
    ..aOS(1, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'value')
    ..hasRequiredFields = false
  ;

  FilterData._() : super();
  factory FilterData({
    $core.String? value,
  }) {
    final _result = create();
    if (value != null) {
      _result.value = value;
    }
    return _result;
  }
  factory FilterData.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory FilterData.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.deepCopy] instead. '
  'Will be removed in next major version')
  FilterData clone() => FilterData()..mergeFromMessage(this);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.rebuild] instead. '
  'Will be removed in next major version')
  FilterData copyWith(void Function(FilterData) updates) => super.copyWith((message) => updates(message as FilterData)) as FilterData; // ignore: deprecated_member_use
  $pb.BuilderInfo get info_ => _i;
  @$core.pragma('dart2js:noInline')
  static FilterData create() => FilterData._();
  FilterData createEmptyInstance() => create();
  static $pb.PbList<FilterData> createRepeated() => $pb.PbList<FilterData>();
  @$core.pragma('dart2js:noInline')
  static FilterData getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<FilterData>(create);
  static FilterData? _defaultInstance;

  @$pb.TagNumber(1)
  $core.String get value => $_getSZ(0);
  @$pb.TagNumber(1)
  set value($core.String v) { $_setString(0, v); }
  @$pb.TagNumber(1)
  $core.bool hasValue() => $_has(0);
  @$pb.TagNumber(1)
  void clearValue() => clearField(1);
}

class RootListRowGlobalID extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'RootListRowGlobalID', package: const $pb.PackageName(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'runify'), createEmptyInstance: create)
    ..a<$core.int>(1, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'providerID', $pb.PbFieldType.OU3, protoName: 'providerID')
    ..a<$core.int>(2, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'rowID', $pb.PbFieldType.OU3, protoName: 'rowID')
    ..hasRequiredFields = false
  ;

  RootListRowGlobalID._() : super();
  factory RootListRowGlobalID({
    $core.int? providerID,
    $core.int? rowID,
  }) {
    final _result = create();
    if (providerID != null) {
      _result.providerID = providerID;
    }
    if (rowID != null) {
      _result.rowID = rowID;
    }
    return _result;
  }
  factory RootListRowGlobalID.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory RootListRowGlobalID.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.deepCopy] instead. '
  'Will be removed in next major version')
  RootListRowGlobalID clone() => RootListRowGlobalID()..mergeFromMessage(this);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.rebuild] instead. '
  'Will be removed in next major version')
  RootListRowGlobalID copyWith(void Function(RootListRowGlobalID) updates) => super.copyWith((message) => updates(message as RootListRowGlobalID)) as RootListRowGlobalID; // ignore: deprecated_member_use
  $pb.BuilderInfo get info_ => _i;
  @$core.pragma('dart2js:noInline')
  static RootListRowGlobalID create() => RootListRowGlobalID._();
  RootListRowGlobalID createEmptyInstance() => create();
  static $pb.PbList<RootListRowGlobalID> createRepeated() => $pb.PbList<RootListRowGlobalID>();
  @$core.pragma('dart2js:noInline')
  static RootListRowGlobalID getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<RootListRowGlobalID>(create);
  static RootListRowGlobalID? _defaultInstance;

  @$pb.TagNumber(1)
  $core.int get providerID => $_getIZ(0);
  @$pb.TagNumber(1)
  set providerID($core.int v) { $_setUnsignedInt32(0, v); }
  @$pb.TagNumber(1)
  $core.bool hasProviderID() => $_has(0);
  @$pb.TagNumber(1)
  void clearProviderID() => clearField(1);

  @$pb.TagNumber(2)
  $core.int get rowID => $_getIZ(1);
  @$pb.TagNumber(2)
  set rowID($core.int v) { $_setUnsignedInt32(1, v); }
  @$pb.TagNumber(2)
  $core.bool hasRowID() => $_has(1);
  @$pb.TagNumber(2)
  void clearRowID() => clearField(2);
}

class RootListRow extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'RootListRow', package: const $pb.PackageName(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'runify'), createEmptyInstance: create)
    ..e<RootListRowType>(1, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'rowType', $pb.PbFieldType.OE, protoName: 'rowType', defaultOrMaker: RootListRowType.CALCULATOR, valueOf: RootListRowType.valueOf, enumValues: RootListRowType.values)
    ..a<$core.int>(2, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'providerID', $pb.PbFieldType.OU3, protoName: 'providerID')
    ..a<$core.int>(3, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'rowID', $pb.PbFieldType.OU3, protoName: 'rowID')
    ..a<$core.int>(4, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'priority', $pb.PbFieldType.OU3)
    ..aOS(5, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'icon')
    ..aOS(6, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'displayName', protoName: 'displayName')
    ..aOS(7, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'searchNames', protoName: 'searchNames')
    ..hasRequiredFields = false
  ;

  RootListRow._() : super();
  factory RootListRow({
    RootListRowType? rowType,
    $core.int? providerID,
    $core.int? rowID,
    $core.int? priority,
    $core.String? icon,
    $core.String? displayName,
    $core.String? searchNames,
  }) {
    final _result = create();
    if (rowType != null) {
      _result.rowType = rowType;
    }
    if (providerID != null) {
      _result.providerID = providerID;
    }
    if (rowID != null) {
      _result.rowID = rowID;
    }
    if (priority != null) {
      _result.priority = priority;
    }
    if (icon != null) {
      _result.icon = icon;
    }
    if (displayName != null) {
      _result.displayName = displayName;
    }
    if (searchNames != null) {
      _result.searchNames = searchNames;
    }
    return _result;
  }
  factory RootListRow.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory RootListRow.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.deepCopy] instead. '
  'Will be removed in next major version')
  RootListRow clone() => RootListRow()..mergeFromMessage(this);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.rebuild] instead. '
  'Will be removed in next major version')
  RootListRow copyWith(void Function(RootListRow) updates) => super.copyWith((message) => updates(message as RootListRow)) as RootListRow; // ignore: deprecated_member_use
  $pb.BuilderInfo get info_ => _i;
  @$core.pragma('dart2js:noInline')
  static RootListRow create() => RootListRow._();
  RootListRow createEmptyInstance() => create();
  static $pb.PbList<RootListRow> createRepeated() => $pb.PbList<RootListRow>();
  @$core.pragma('dart2js:noInline')
  static RootListRow getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<RootListRow>(create);
  static RootListRow? _defaultInstance;

  @$pb.TagNumber(1)
  RootListRowType get rowType => $_getN(0);
  @$pb.TagNumber(1)
  set rowType(RootListRowType v) { setField(1, v); }
  @$pb.TagNumber(1)
  $core.bool hasRowType() => $_has(0);
  @$pb.TagNumber(1)
  void clearRowType() => clearField(1);

  @$pb.TagNumber(2)
  $core.int get providerID => $_getIZ(1);
  @$pb.TagNumber(2)
  set providerID($core.int v) { $_setUnsignedInt32(1, v); }
  @$pb.TagNumber(2)
  $core.bool hasProviderID() => $_has(1);
  @$pb.TagNumber(2)
  void clearProviderID() => clearField(2);

  @$pb.TagNumber(3)
  $core.int get rowID => $_getIZ(2);
  @$pb.TagNumber(3)
  set rowID($core.int v) { $_setUnsignedInt32(2, v); }
  @$pb.TagNumber(3)
  $core.bool hasRowID() => $_has(2);
  @$pb.TagNumber(3)
  void clearRowID() => clearField(3);

  @$pb.TagNumber(4)
  $core.int get priority => $_getIZ(3);
  @$pb.TagNumber(4)
  set priority($core.int v) { $_setUnsignedInt32(3, v); }
  @$pb.TagNumber(4)
  $core.bool hasPriority() => $_has(3);
  @$pb.TagNumber(4)
  void clearPriority() => clearField(4);

  @$pb.TagNumber(5)
  $core.String get icon => $_getSZ(4);
  @$pb.TagNumber(5)
  set icon($core.String v) { $_setString(4, v); }
  @$pb.TagNumber(5)
  $core.bool hasIcon() => $_has(4);
  @$pb.TagNumber(5)
  void clearIcon() => clearField(5);

  @$pb.TagNumber(6)
  $core.String get displayName => $_getSZ(5);
  @$pb.TagNumber(6)
  set displayName($core.String v) { $_setString(5, v); }
  @$pb.TagNumber(6)
  $core.bool hasDisplayName() => $_has(5);
  @$pb.TagNumber(6)
  void clearDisplayName() => clearField(6);

  @$pb.TagNumber(7)
  $core.String get searchNames => $_getSZ(6);
  @$pb.TagNumber(7)
  set searchNames($core.String v) { $_setString(6, v); }
  @$pb.TagNumber(7)
  $core.bool hasSearchNames() => $_has(6);
  @$pb.TagNumber(7)
  void clearSearchNames() => clearField(7);
}

class FormOpen extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'FormOpen', package: const $pb.PackageName(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'runify'), createEmptyInstance: create)
    ..aOM<FormMarkup>(1, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'markup', subBuilder: FormMarkup.create)
    ..aOM<FormModel>(2, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'model', subBuilder: FormModel.create)
    ..hasRequiredFields = false
  ;

  FormOpen._() : super();
  factory FormOpen({
    FormMarkup? markup,
    FormModel? model,
  }) {
    final _result = create();
    if (markup != null) {
      _result.markup = markup;
    }
    if (model != null) {
      _result.model = model;
    }
    return _result;
  }
  factory FormOpen.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory FormOpen.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.deepCopy] instead. '
  'Will be removed in next major version')
  FormOpen clone() => FormOpen()..mergeFromMessage(this);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.rebuild] instead. '
  'Will be removed in next major version')
  FormOpen copyWith(void Function(FormOpen) updates) => super.copyWith((message) => updates(message as FormOpen)) as FormOpen; // ignore: deprecated_member_use
  $pb.BuilderInfo get info_ => _i;
  @$core.pragma('dart2js:noInline')
  static FormOpen create() => FormOpen._();
  FormOpen createEmptyInstance() => create();
  static $pb.PbList<FormOpen> createRepeated() => $pb.PbList<FormOpen>();
  @$core.pragma('dart2js:noInline')
  static FormOpen getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<FormOpen>(create);
  static FormOpen? _defaultInstance;

  @$pb.TagNumber(1)
  FormMarkup get markup => $_getN(0);
  @$pb.TagNumber(1)
  set markup(FormMarkup v) { setField(1, v); }
  @$pb.TagNumber(1)
  $core.bool hasMarkup() => $_has(0);
  @$pb.TagNumber(1)
  void clearMarkup() => clearField(1);
  @$pb.TagNumber(1)
  FormMarkup ensureMarkup() => $_ensure(0);

  @$pb.TagNumber(2)
  FormModel get model => $_getN(1);
  @$pb.TagNumber(2)
  set model(FormModel v) { setField(2, v); }
  @$pb.TagNumber(2)
  $core.bool hasModel() => $_has(1);
  @$pb.TagNumber(2)
  void clearModel() => clearField(2);
  @$pb.TagNumber(2)
  FormModel ensureModel() => $_ensure(1);
}

class FieldCheckRequest extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'FieldCheckRequest', package: const $pb.PackageName(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'runify'), createEmptyInstance: create)
    ..a<$core.int>(1, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'requestID', $pb.PbFieldType.OU3, protoName: 'requestID')
    ..aOS(2, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'fieldName', protoName: 'fieldName')
    ..aOM<FormData>(3, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'data', subBuilder: FormData.create)
    ..hasRequiredFields = false
  ;

  FieldCheckRequest._() : super();
  factory FieldCheckRequest({
    $core.int? requestID,
    $core.String? fieldName,
    FormData? data,
  }) {
    final _result = create();
    if (requestID != null) {
      _result.requestID = requestID;
    }
    if (fieldName != null) {
      _result.fieldName = fieldName;
    }
    if (data != null) {
      _result.data = data;
    }
    return _result;
  }
  factory FieldCheckRequest.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory FieldCheckRequest.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.deepCopy] instead. '
  'Will be removed in next major version')
  FieldCheckRequest clone() => FieldCheckRequest()..mergeFromMessage(this);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.rebuild] instead. '
  'Will be removed in next major version')
  FieldCheckRequest copyWith(void Function(FieldCheckRequest) updates) => super.copyWith((message) => updates(message as FieldCheckRequest)) as FieldCheckRequest; // ignore: deprecated_member_use
  $pb.BuilderInfo get info_ => _i;
  @$core.pragma('dart2js:noInline')
  static FieldCheckRequest create() => FieldCheckRequest._();
  FieldCheckRequest createEmptyInstance() => create();
  static $pb.PbList<FieldCheckRequest> createRepeated() => $pb.PbList<FieldCheckRequest>();
  @$core.pragma('dart2js:noInline')
  static FieldCheckRequest getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<FieldCheckRequest>(create);
  static FieldCheckRequest? _defaultInstance;

  @$pb.TagNumber(1)
  $core.int get requestID => $_getIZ(0);
  @$pb.TagNumber(1)
  set requestID($core.int v) { $_setUnsignedInt32(0, v); }
  @$pb.TagNumber(1)
  $core.bool hasRequestID() => $_has(0);
  @$pb.TagNumber(1)
  void clearRequestID() => clearField(1);

  @$pb.TagNumber(2)
  $core.String get fieldName => $_getSZ(1);
  @$pb.TagNumber(2)
  set fieldName($core.String v) { $_setString(1, v); }
  @$pb.TagNumber(2)
  $core.bool hasFieldName() => $_has(1);
  @$pb.TagNumber(2)
  void clearFieldName() => clearField(2);

  @$pb.TagNumber(3)
  FormData get data => $_getN(2);
  @$pb.TagNumber(3)
  set data(FormData v) { setField(3, v); }
  @$pb.TagNumber(3)
  $core.bool hasData() => $_has(2);
  @$pb.TagNumber(3)
  void clearData() => clearField(3);
  @$pb.TagNumber(3)
  FormData ensureData() => $_ensure(2);
}

class FieldCheckResponse extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'FieldCheckResponse', package: const $pb.PackageName(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'runify'), createEmptyInstance: create)
    ..a<$core.int>(1, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'requestID', $pb.PbFieldType.OU3, protoName: 'requestID')
    ..aOB(2, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'result')
    ..aOS(3, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'error')
    ..hasRequiredFields = false
  ;

  FieldCheckResponse._() : super();
  factory FieldCheckResponse({
    $core.int? requestID,
    $core.bool? result,
    $core.String? error,
  }) {
    final _result = create();
    if (requestID != null) {
      _result.requestID = requestID;
    }
    if (result != null) {
      _result.result = result;
    }
    if (error != null) {
      _result.error = error;
    }
    return _result;
  }
  factory FieldCheckResponse.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory FieldCheckResponse.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.deepCopy] instead. '
  'Will be removed in next major version')
  FieldCheckResponse clone() => FieldCheckResponse()..mergeFromMessage(this);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.rebuild] instead. '
  'Will be removed in next major version')
  FieldCheckResponse copyWith(void Function(FieldCheckResponse) updates) => super.copyWith((message) => updates(message as FieldCheckResponse)) as FieldCheckResponse; // ignore: deprecated_member_use
  $pb.BuilderInfo get info_ => _i;
  @$core.pragma('dart2js:noInline')
  static FieldCheckResponse create() => FieldCheckResponse._();
  FieldCheckResponse createEmptyInstance() => create();
  static $pb.PbList<FieldCheckResponse> createRepeated() => $pb.PbList<FieldCheckResponse>();
  @$core.pragma('dart2js:noInline')
  static FieldCheckResponse getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<FieldCheckResponse>(create);
  static FieldCheckResponse? _defaultInstance;

  @$pb.TagNumber(1)
  $core.int get requestID => $_getIZ(0);
  @$pb.TagNumber(1)
  set requestID($core.int v) { $_setUnsignedInt32(0, v); }
  @$pb.TagNumber(1)
  $core.bool hasRequestID() => $_has(0);
  @$pb.TagNumber(1)
  void clearRequestID() => clearField(1);

  @$pb.TagNumber(2)
  $core.bool get result => $_getBF(1);
  @$pb.TagNumber(2)
  set result($core.bool v) { $_setBool(1, v); }
  @$pb.TagNumber(2)
  $core.bool hasResult() => $_has(1);
  @$pb.TagNumber(2)
  void clearResult() => clearField(2);

  @$pb.TagNumber(3)
  $core.String get error => $_getSZ(2);
  @$pb.TagNumber(3)
  set error($core.String v) { $_setString(2, v); }
  @$pb.TagNumber(3)
  $core.bool hasError() => $_has(2);
  @$pb.TagNumber(3)
  void clearError() => clearField(3);
}

class FormSubmit extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'FormSubmit', package: const $pb.PackageName(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'runify'), createEmptyInstance: create)
    ..aOM<FormData>(1, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'data', subBuilder: FormData.create)
    ..hasRequiredFields = false
  ;

  FormSubmit._() : super();
  factory FormSubmit({
    FormData? data,
  }) {
    final _result = create();
    if (data != null) {
      _result.data = data;
    }
    return _result;
  }
  factory FormSubmit.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory FormSubmit.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.deepCopy] instead. '
  'Will be removed in next major version')
  FormSubmit clone() => FormSubmit()..mergeFromMessage(this);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.rebuild] instead. '
  'Will be removed in next major version')
  FormSubmit copyWith(void Function(FormSubmit) updates) => super.copyWith((message) => updates(message as FormSubmit)) as FormSubmit; // ignore: deprecated_member_use
  $pb.BuilderInfo get info_ => _i;
  @$core.pragma('dart2js:noInline')
  static FormSubmit create() => FormSubmit._();
  FormSubmit createEmptyInstance() => create();
  static $pb.PbList<FormSubmit> createRepeated() => $pb.PbList<FormSubmit>();
  @$core.pragma('dart2js:noInline')
  static FormSubmit getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<FormSubmit>(create);
  static FormSubmit? _defaultInstance;

  @$pb.TagNumber(1)
  FormData get data => $_getN(0);
  @$pb.TagNumber(1)
  set data(FormData v) { setField(1, v); }
  @$pb.TagNumber(1)
  $core.bool hasData() => $_has(0);
  @$pb.TagNumber(1)
  void clearData() => clearField(1);
  @$pb.TagNumber(1)
  FormData ensureData() => $_ensure(0);
}

class RootListOpen extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'RootListOpen', package: const $pb.PackageName(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'runify'), createEmptyInstance: create)
    ..pc<RootListRow>(1, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'rows', $pb.PbFieldType.PM, subBuilder: RootListRow.create)
    ..hasRequiredFields = false
  ;

  RootListOpen._() : super();
  factory RootListOpen({
    $core.Iterable<RootListRow>? rows,
  }) {
    final _result = create();
    if (rows != null) {
      _result.rows.addAll(rows);
    }
    return _result;
  }
  factory RootListOpen.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory RootListOpen.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.deepCopy] instead. '
  'Will be removed in next major version')
  RootListOpen clone() => RootListOpen()..mergeFromMessage(this);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.rebuild] instead. '
  'Will be removed in next major version')
  RootListOpen copyWith(void Function(RootListOpen) updates) => super.copyWith((message) => updates(message as RootListOpen)) as RootListOpen; // ignore: deprecated_member_use
  $pb.BuilderInfo get info_ => _i;
  @$core.pragma('dart2js:noInline')
  static RootListOpen create() => RootListOpen._();
  RootListOpen createEmptyInstance() => create();
  static $pb.PbList<RootListOpen> createRepeated() => $pb.PbList<RootListOpen>();
  @$core.pragma('dart2js:noInline')
  static RootListOpen getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<RootListOpen>(create);
  static RootListOpen? _defaultInstance;

  @$pb.TagNumber(1)
  $core.List<RootListRow> get rows => $_getList(0);
}

class RootListAddRows extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'RootListAddRows', package: const $pb.PackageName(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'runify'), createEmptyInstance: create)
    ..pc<RootListRow>(1, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'rows', $pb.PbFieldType.PM, subBuilder: RootListRow.create)
    ..hasRequiredFields = false
  ;

  RootListAddRows._() : super();
  factory RootListAddRows({
    $core.Iterable<RootListRow>? rows,
  }) {
    final _result = create();
    if (rows != null) {
      _result.rows.addAll(rows);
    }
    return _result;
  }
  factory RootListAddRows.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory RootListAddRows.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.deepCopy] instead. '
  'Will be removed in next major version')
  RootListAddRows clone() => RootListAddRows()..mergeFromMessage(this);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.rebuild] instead. '
  'Will be removed in next major version')
  RootListAddRows copyWith(void Function(RootListAddRows) updates) => super.copyWith((message) => updates(message as RootListAddRows)) as RootListAddRows; // ignore: deprecated_member_use
  $pb.BuilderInfo get info_ => _i;
  @$core.pragma('dart2js:noInline')
  static RootListAddRows create() => RootListAddRows._();
  RootListAddRows createEmptyInstance() => create();
  static $pb.PbList<RootListAddRows> createRepeated() => $pb.PbList<RootListAddRows>();
  @$core.pragma('dart2js:noInline')
  static RootListAddRows getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<RootListAddRows>(create);
  static RootListAddRows? _defaultInstance;

  @$pb.TagNumber(1)
  $core.List<RootListRow> get rows => $_getList(0);
}

class RootListChangeRows extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'RootListChangeRows', package: const $pb.PackageName(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'runify'), createEmptyInstance: create)
    ..pc<RootListRow>(1, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'rows', $pb.PbFieldType.PM, subBuilder: RootListRow.create)
    ..hasRequiredFields = false
  ;

  RootListChangeRows._() : super();
  factory RootListChangeRows({
    $core.Iterable<RootListRow>? rows,
  }) {
    final _result = create();
    if (rows != null) {
      _result.rows.addAll(rows);
    }
    return _result;
  }
  factory RootListChangeRows.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory RootListChangeRows.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.deepCopy] instead. '
  'Will be removed in next major version')
  RootListChangeRows clone() => RootListChangeRows()..mergeFromMessage(this);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.rebuild] instead. '
  'Will be removed in next major version')
  RootListChangeRows copyWith(void Function(RootListChangeRows) updates) => super.copyWith((message) => updates(message as RootListChangeRows)) as RootListChangeRows; // ignore: deprecated_member_use
  $pb.BuilderInfo get info_ => _i;
  @$core.pragma('dart2js:noInline')
  static RootListChangeRows create() => RootListChangeRows._();
  RootListChangeRows createEmptyInstance() => create();
  static $pb.PbList<RootListChangeRows> createRepeated() => $pb.PbList<RootListChangeRows>();
  @$core.pragma('dart2js:noInline')
  static RootListChangeRows getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<RootListChangeRows>(create);
  static RootListChangeRows? _defaultInstance;

  @$pb.TagNumber(1)
  $core.List<RootListRow> get rows => $_getList(0);
}

class RootListRemoveRows extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'RootListRemoveRows', package: const $pb.PackageName(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'runify'), createEmptyInstance: create)
    ..pc<RootListRowGlobalID>(1, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'rows', $pb.PbFieldType.PM, subBuilder: RootListRowGlobalID.create)
    ..hasRequiredFields = false
  ;

  RootListRemoveRows._() : super();
  factory RootListRemoveRows({
    $core.Iterable<RootListRowGlobalID>? rows,
  }) {
    final _result = create();
    if (rows != null) {
      _result.rows.addAll(rows);
    }
    return _result;
  }
  factory RootListRemoveRows.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory RootListRemoveRows.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.deepCopy] instead. '
  'Will be removed in next major version')
  RootListRemoveRows clone() => RootListRemoveRows()..mergeFromMessage(this);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.rebuild] instead. '
  'Will be removed in next major version')
  RootListRemoveRows copyWith(void Function(RootListRemoveRows) updates) => super.copyWith((message) => updates(message as RootListRemoveRows)) as RootListRemoveRows; // ignore: deprecated_member_use
  $pb.BuilderInfo get info_ => _i;
  @$core.pragma('dart2js:noInline')
  static RootListRemoveRows create() => RootListRemoveRows._();
  RootListRemoveRows createEmptyInstance() => create();
  static $pb.PbList<RootListRemoveRows> createRepeated() => $pb.PbList<RootListRemoveRows>();
  @$core.pragma('dart2js:noInline')
  static RootListRemoveRows getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<RootListRemoveRows>(create);
  static RootListRemoveRows? _defaultInstance;

  @$pb.TagNumber(1)
  $core.List<RootListRowGlobalID> get rows => $_getList(0);
}

class ContextMenuRowID extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'ContextMenuRowID', package: const $pb.PackageName(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'runify'), createEmptyInstance: create)
    ..a<$core.int>(1, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'rowID', $pb.PbFieldType.OU3, protoName: 'rowID')
    ..hasRequiredFields = false
  ;

  ContextMenuRowID._() : super();
  factory ContextMenuRowID({
    $core.int? rowID,
  }) {
    final _result = create();
    if (rowID != null) {
      _result.rowID = rowID;
    }
    return _result;
  }
  factory ContextMenuRowID.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory ContextMenuRowID.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.deepCopy] instead. '
  'Will be removed in next major version')
  ContextMenuRowID clone() => ContextMenuRowID()..mergeFromMessage(this);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.rebuild] instead. '
  'Will be removed in next major version')
  ContextMenuRowID copyWith(void Function(ContextMenuRowID) updates) => super.copyWith((message) => updates(message as ContextMenuRowID)) as ContextMenuRowID; // ignore: deprecated_member_use
  $pb.BuilderInfo get info_ => _i;
  @$core.pragma('dart2js:noInline')
  static ContextMenuRowID create() => ContextMenuRowID._();
  ContextMenuRowID createEmptyInstance() => create();
  static $pb.PbList<ContextMenuRowID> createRepeated() => $pb.PbList<ContextMenuRowID>();
  @$core.pragma('dart2js:noInline')
  static ContextMenuRowID getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<ContextMenuRowID>(create);
  static ContextMenuRowID? _defaultInstance;

  @$pb.TagNumber(1)
  $core.int get rowID => $_getIZ(0);
  @$pb.TagNumber(1)
  set rowID($core.int v) { $_setUnsignedInt32(0, v); }
  @$pb.TagNumber(1)
  $core.bool hasRowID() => $_has(0);
  @$pb.TagNumber(1)
  void clearRowID() => clearField(1);
}

class FormClosed extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'FormClosed', package: const $pb.PackageName(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'runify'), createEmptyInstance: create)
    ..hasRequiredFields = false
  ;

  FormClosed._() : super();
  factory FormClosed() => create();
  factory FormClosed.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory FormClosed.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.deepCopy] instead. '
  'Will be removed in next major version')
  FormClosed clone() => FormClosed()..mergeFromMessage(this);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.rebuild] instead. '
  'Will be removed in next major version')
  FormClosed copyWith(void Function(FormClosed) updates) => super.copyWith((message) => updates(message as FormClosed)) as FormClosed; // ignore: deprecated_member_use
  $pb.BuilderInfo get info_ => _i;
  @$core.pragma('dart2js:noInline')
  static FormClosed create() => FormClosed._();
  FormClosed createEmptyInstance() => create();
  static $pb.PbList<FormClosed> createRepeated() => $pb.PbList<FormClosed>();
  @$core.pragma('dart2js:noInline')
  static FormClosed getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<FormClosed>(create);
  static FormClosed? _defaultInstance;
}

class ContextMenuRow extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'ContextMenuRow', package: const $pb.PackageName(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'runify'), createEmptyInstance: create)
    ..a<$core.int>(1, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'rowID', $pb.PbFieldType.OU3, protoName: 'rowID')
    ..aOS(2, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'displayName', protoName: 'displayName')
    ..aOS(3, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'searchNames', protoName: 'searchNames')
    ..hasRequiredFields = false
  ;

  ContextMenuRow._() : super();
  factory ContextMenuRow({
    $core.int? rowID,
    $core.String? displayName,
    $core.String? searchNames,
  }) {
    final _result = create();
    if (rowID != null) {
      _result.rowID = rowID;
    }
    if (displayName != null) {
      _result.displayName = displayName;
    }
    if (searchNames != null) {
      _result.searchNames = searchNames;
    }
    return _result;
  }
  factory ContextMenuRow.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory ContextMenuRow.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.deepCopy] instead. '
  'Will be removed in next major version')
  ContextMenuRow clone() => ContextMenuRow()..mergeFromMessage(this);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.rebuild] instead. '
  'Will be removed in next major version')
  ContextMenuRow copyWith(void Function(ContextMenuRow) updates) => super.copyWith((message) => updates(message as ContextMenuRow)) as ContextMenuRow; // ignore: deprecated_member_use
  $pb.BuilderInfo get info_ => _i;
  @$core.pragma('dart2js:noInline')
  static ContextMenuRow create() => ContextMenuRow._();
  ContextMenuRow createEmptyInstance() => create();
  static $pb.PbList<ContextMenuRow> createRepeated() => $pb.PbList<ContextMenuRow>();
  @$core.pragma('dart2js:noInline')
  static ContextMenuRow getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<ContextMenuRow>(create);
  static ContextMenuRow? _defaultInstance;

  @$pb.TagNumber(1)
  $core.int get rowID => $_getIZ(0);
  @$pb.TagNumber(1)
  set rowID($core.int v) { $_setUnsignedInt32(0, v); }
  @$pb.TagNumber(1)
  $core.bool hasRowID() => $_has(0);
  @$pb.TagNumber(1)
  void clearRowID() => clearField(1);

  @$pb.TagNumber(2)
  $core.String get displayName => $_getSZ(1);
  @$pb.TagNumber(2)
  set displayName($core.String v) { $_setString(1, v); }
  @$pb.TagNumber(2)
  $core.bool hasDisplayName() => $_has(1);
  @$pb.TagNumber(2)
  void clearDisplayName() => clearField(2);

  @$pb.TagNumber(3)
  $core.String get searchNames => $_getSZ(2);
  @$pb.TagNumber(3)
  set searchNames($core.String v) { $_setString(2, v); }
  @$pb.TagNumber(3)
  $core.bool hasSearchNames() => $_has(2);
  @$pb.TagNumber(3)
  void clearSearchNames() => clearField(3);
}

class ContextMenuOpen extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'ContextMenuOpen', package: const $pb.PackageName(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'runify'), createEmptyInstance: create)
    ..pc<ContextMenuRow>(1, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'rows', $pb.PbFieldType.PM, subBuilder: ContextMenuRow.create)
    ..hasRequiredFields = false
  ;

  ContextMenuOpen._() : super();
  factory ContextMenuOpen({
    $core.Iterable<ContextMenuRow>? rows,
  }) {
    final _result = create();
    if (rows != null) {
      _result.rows.addAll(rows);
    }
    return _result;
  }
  factory ContextMenuOpen.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory ContextMenuOpen.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.deepCopy] instead. '
  'Will be removed in next major version')
  ContextMenuOpen clone() => ContextMenuOpen()..mergeFromMessage(this);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.rebuild] instead. '
  'Will be removed in next major version')
  ContextMenuOpen copyWith(void Function(ContextMenuOpen) updates) => super.copyWith((message) => updates(message as ContextMenuOpen)) as ContextMenuOpen; // ignore: deprecated_member_use
  $pb.BuilderInfo get info_ => _i;
  @$core.pragma('dart2js:noInline')
  static ContextMenuOpen create() => ContextMenuOpen._();
  ContextMenuOpen createEmptyInstance() => create();
  static $pb.PbList<ContextMenuOpen> createRepeated() => $pb.PbList<ContextMenuOpen>();
  @$core.pragma('dart2js:noInline')
  static ContextMenuOpen getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<ContextMenuOpen>(create);
  static ContextMenuOpen? _defaultInstance;

  @$pb.TagNumber(1)
  $core.List<ContextMenuRow> get rows => $_getList(0);
}

class UserMessage extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'UserMessage', package: const $pb.PackageName(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'runify'), createEmptyInstance: create)
    ..e<MessageType>(1, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'messageType', $pb.PbFieldType.OE, protoName: 'messageType', defaultOrMaker: MessageType.TYPE_ERROR, valueOf: MessageType.valueOf, enumValues: MessageType.values)
    ..aOS(2, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'message')
    ..hasRequiredFields = false
  ;

  UserMessage._() : super();
  factory UserMessage({
    MessageType? messageType,
    $core.String? message,
  }) {
    final _result = create();
    if (messageType != null) {
      _result.messageType = messageType;
    }
    if (message != null) {
      _result.message = message;
    }
    return _result;
  }
  factory UserMessage.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory UserMessage.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.deepCopy] instead. '
  'Will be removed in next major version')
  UserMessage clone() => UserMessage()..mergeFromMessage(this);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.rebuild] instead. '
  'Will be removed in next major version')
  UserMessage copyWith(void Function(UserMessage) updates) => super.copyWith((message) => updates(message as UserMessage)) as UserMessage; // ignore: deprecated_member_use
  $pb.BuilderInfo get info_ => _i;
  @$core.pragma('dart2js:noInline')
  static UserMessage create() => UserMessage._();
  UserMessage createEmptyInstance() => create();
  static $pb.PbList<UserMessage> createRepeated() => $pb.PbList<UserMessage>();
  @$core.pragma('dart2js:noInline')
  static UserMessage getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<UserMessage>(create);
  static UserMessage? _defaultInstance;

  @$pb.TagNumber(1)
  MessageType get messageType => $_getN(0);
  @$pb.TagNumber(1)
  set messageType(MessageType v) { setField(1, v); }
  @$pb.TagNumber(1)
  $core.bool hasMessageType() => $_has(0);
  @$pb.TagNumber(1)
  void clearMessageType() => clearField(1);

  @$pb.TagNumber(2)
  $core.String get message => $_getSZ(1);
  @$pb.TagNumber(2)
  set message($core.String v) { $_setString(1, v); }
  @$pb.TagNumber(2)
  $core.bool hasMessage() => $_has(1);
  @$pb.TagNumber(2)
  void clearMessage() => clearField(2);
}

class CloseForm extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'CloseForm', package: const $pb.PackageName(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'runify'), createEmptyInstance: create)
    ..hasRequiredFields = false
  ;

  CloseForm._() : super();
  factory CloseForm() => create();
  factory CloseForm.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory CloseForm.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.deepCopy] instead. '
  'Will be removed in next major version')
  CloseForm clone() => CloseForm()..mergeFromMessage(this);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.rebuild] instead. '
  'Will be removed in next major version')
  CloseForm copyWith(void Function(CloseForm) updates) => super.copyWith((message) => updates(message as CloseForm)) as CloseForm; // ignore: deprecated_member_use
  $pb.BuilderInfo get info_ => _i;
  @$core.pragma('dart2js:noInline')
  static CloseForm create() => CloseForm._();
  CloseForm createEmptyInstance() => create();
  static $pb.PbList<CloseForm> createRepeated() => $pb.PbList<CloseForm>();
  @$core.pragma('dart2js:noInline')
  static CloseForm getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<CloseForm>(create);
  static CloseForm? _defaultInstance;
}

class HideUI extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'HideUI', package: const $pb.PackageName(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'runify'), createEmptyInstance: create)
    ..aOM<UserMessage>(1, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'message', subBuilder: UserMessage.create)
    ..hasRequiredFields = false
  ;

  HideUI._() : super();
  factory HideUI({
    UserMessage? message,
  }) {
    final _result = create();
    if (message != null) {
      _result.message = message;
    }
    return _result;
  }
  factory HideUI.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory HideUI.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.deepCopy] instead. '
  'Will be removed in next major version')
  HideUI clone() => HideUI()..mergeFromMessage(this);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.rebuild] instead. '
  'Will be removed in next major version')
  HideUI copyWith(void Function(HideUI) updates) => super.copyWith((message) => updates(message as HideUI)) as HideUI; // ignore: deprecated_member_use
  $pb.BuilderInfo get info_ => _i;
  @$core.pragma('dart2js:noInline')
  static HideUI create() => HideUI._();
  HideUI createEmptyInstance() => create();
  static $pb.PbList<HideUI> createRepeated() => $pb.PbList<HideUI>();
  @$core.pragma('dart2js:noInline')
  static HideUI getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<HideUI>(create);
  static HideUI? _defaultInstance;

  @$pb.TagNumber(1)
  UserMessage get message => $_getN(0);
  @$pb.TagNumber(1)
  set message(UserMessage v) { setField(1, v); }
  @$pb.TagNumber(1)
  $core.bool hasMessage() => $_has(0);
  @$pb.TagNumber(1)
  void clearMessage() => clearField(1);
  @$pb.TagNumber(1)
  UserMessage ensureMessage() => $_ensure(0);
}

class CloseUI extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'CloseUI', package: const $pb.PackageName(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'runify'), createEmptyInstance: create)
    ..hasRequiredFields = false
  ;

  CloseUI._() : super();
  factory CloseUI() => create();
  factory CloseUI.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory CloseUI.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.deepCopy] instead. '
  'Will be removed in next major version')
  CloseUI clone() => CloseUI()..mergeFromMessage(this);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.rebuild] instead. '
  'Will be removed in next major version')
  CloseUI copyWith(void Function(CloseUI) updates) => super.copyWith((message) => updates(message as CloseUI)) as CloseUI; // ignore: deprecated_member_use
  $pb.BuilderInfo get info_ => _i;
  @$core.pragma('dart2js:noInline')
  static CloseUI create() => CloseUI._();
  CloseUI createEmptyInstance() => create();
  static $pb.PbList<CloseUI> createRepeated() => $pb.PbList<CloseUI>();
  @$core.pragma('dart2js:noInline')
  static CloseUI getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<CloseUI>(create);
  static CloseUI? _defaultInstance;
}

enum UIMessage_Payload {
  writeLog, 
  filterChanged, 
  rootListRowActivated, 
  rootListMenuActivated, 
  contextMenuRowActivated, 
  fieldCheckRequest, 
  formSubmit, 
  formClosed, 
  notSet
}

class UIMessage extends $pb.GeneratedMessage {
  static const $core.Map<$core.int, UIMessage_Payload> _UIMessage_PayloadByTag = {
    2 : UIMessage_Payload.writeLog,
    3 : UIMessage_Payload.filterChanged,
    4 : UIMessage_Payload.rootListRowActivated,
    5 : UIMessage_Payload.rootListMenuActivated,
    6 : UIMessage_Payload.contextMenuRowActivated,
    7 : UIMessage_Payload.fieldCheckRequest,
    8 : UIMessage_Payload.formSubmit,
    9 : UIMessage_Payload.formClosed,
    0 : UIMessage_Payload.notSet
  };
  static final $pb.BuilderInfo _i = $pb.BuilderInfo(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'UIMessage', package: const $pb.PackageName(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'runify'), createEmptyInstance: create)
    ..oo(0, [2, 3, 4, 5, 6, 7, 8, 9])
    ..a<$core.int>(1, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'formID', $pb.PbFieldType.OU3, protoName: 'formID')
    ..aOM<WriteLog>(2, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'writeLog', protoName: 'writeLog', subBuilder: WriteLog.create)
    ..aOM<FilterData>(3, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'filterChanged', protoName: 'filterChanged', subBuilder: FilterData.create)
    ..aOM<RootListRowGlobalID>(4, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'rootListRowActivated', protoName: 'rootListRowActivated', subBuilder: RootListRowGlobalID.create)
    ..aOM<RootListRowGlobalID>(5, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'rootListMenuActivated', protoName: 'rootListMenuActivated', subBuilder: RootListRowGlobalID.create)
    ..aOM<ContextMenuRowID>(6, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'contextMenuRowActivated', protoName: 'contextMenuRowActivated', subBuilder: ContextMenuRowID.create)
    ..aOM<FieldCheckRequest>(7, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'fieldCheckRequest', protoName: 'fieldCheckRequest', subBuilder: FieldCheckRequest.create)
    ..aOM<FormSubmit>(8, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'formSubmit', protoName: 'formSubmit', subBuilder: FormSubmit.create)
    ..aOM<FormClosed>(9, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'formClosed', protoName: 'formClosed', subBuilder: FormClosed.create)
    ..hasRequiredFields = false
  ;

  UIMessage._() : super();
  factory UIMessage({
    $core.int? formID,
    WriteLog? writeLog,
    FilterData? filterChanged,
    RootListRowGlobalID? rootListRowActivated,
    RootListRowGlobalID? rootListMenuActivated,
    ContextMenuRowID? contextMenuRowActivated,
    FieldCheckRequest? fieldCheckRequest,
    FormSubmit? formSubmit,
    FormClosed? formClosed,
  }) {
    final _result = create();
    if (formID != null) {
      _result.formID = formID;
    }
    if (writeLog != null) {
      _result.writeLog = writeLog;
    }
    if (filterChanged != null) {
      _result.filterChanged = filterChanged;
    }
    if (rootListRowActivated != null) {
      _result.rootListRowActivated = rootListRowActivated;
    }
    if (rootListMenuActivated != null) {
      _result.rootListMenuActivated = rootListMenuActivated;
    }
    if (contextMenuRowActivated != null) {
      _result.contextMenuRowActivated = contextMenuRowActivated;
    }
    if (fieldCheckRequest != null) {
      _result.fieldCheckRequest = fieldCheckRequest;
    }
    if (formSubmit != null) {
      _result.formSubmit = formSubmit;
    }
    if (formClosed != null) {
      _result.formClosed = formClosed;
    }
    return _result;
  }
  factory UIMessage.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory UIMessage.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.deepCopy] instead. '
  'Will be removed in next major version')
  UIMessage clone() => UIMessage()..mergeFromMessage(this);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.rebuild] instead. '
  'Will be removed in next major version')
  UIMessage copyWith(void Function(UIMessage) updates) => super.copyWith((message) => updates(message as UIMessage)) as UIMessage; // ignore: deprecated_member_use
  $pb.BuilderInfo get info_ => _i;
  @$core.pragma('dart2js:noInline')
  static UIMessage create() => UIMessage._();
  UIMessage createEmptyInstance() => create();
  static $pb.PbList<UIMessage> createRepeated() => $pb.PbList<UIMessage>();
  @$core.pragma('dart2js:noInline')
  static UIMessage getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<UIMessage>(create);
  static UIMessage? _defaultInstance;

  UIMessage_Payload whichPayload() => _UIMessage_PayloadByTag[$_whichOneof(0)]!;
  void clearPayload() => clearField($_whichOneof(0));

  @$pb.TagNumber(1)
  $core.int get formID => $_getIZ(0);
  @$pb.TagNumber(1)
  set formID($core.int v) { $_setUnsignedInt32(0, v); }
  @$pb.TagNumber(1)
  $core.bool hasFormID() => $_has(0);
  @$pb.TagNumber(1)
  void clearFormID() => clearField(1);

  @$pb.TagNumber(2)
  WriteLog get writeLog => $_getN(1);
  @$pb.TagNumber(2)
  set writeLog(WriteLog v) { setField(2, v); }
  @$pb.TagNumber(2)
  $core.bool hasWriteLog() => $_has(1);
  @$pb.TagNumber(2)
  void clearWriteLog() => clearField(2);
  @$pb.TagNumber(2)
  WriteLog ensureWriteLog() => $_ensure(1);

  @$pb.TagNumber(3)
  FilterData get filterChanged => $_getN(2);
  @$pb.TagNumber(3)
  set filterChanged(FilterData v) { setField(3, v); }
  @$pb.TagNumber(3)
  $core.bool hasFilterChanged() => $_has(2);
  @$pb.TagNumber(3)
  void clearFilterChanged() => clearField(3);
  @$pb.TagNumber(3)
  FilterData ensureFilterChanged() => $_ensure(2);

  @$pb.TagNumber(4)
  RootListRowGlobalID get rootListRowActivated => $_getN(3);
  @$pb.TagNumber(4)
  set rootListRowActivated(RootListRowGlobalID v) { setField(4, v); }
  @$pb.TagNumber(4)
  $core.bool hasRootListRowActivated() => $_has(3);
  @$pb.TagNumber(4)
  void clearRootListRowActivated() => clearField(4);
  @$pb.TagNumber(4)
  RootListRowGlobalID ensureRootListRowActivated() => $_ensure(3);

  @$pb.TagNumber(5)
  RootListRowGlobalID get rootListMenuActivated => $_getN(4);
  @$pb.TagNumber(5)
  set rootListMenuActivated(RootListRowGlobalID v) { setField(5, v); }
  @$pb.TagNumber(5)
  $core.bool hasRootListMenuActivated() => $_has(4);
  @$pb.TagNumber(5)
  void clearRootListMenuActivated() => clearField(5);
  @$pb.TagNumber(5)
  RootListRowGlobalID ensureRootListMenuActivated() => $_ensure(4);

  @$pb.TagNumber(6)
  ContextMenuRowID get contextMenuRowActivated => $_getN(5);
  @$pb.TagNumber(6)
  set contextMenuRowActivated(ContextMenuRowID v) { setField(6, v); }
  @$pb.TagNumber(6)
  $core.bool hasContextMenuRowActivated() => $_has(5);
  @$pb.TagNumber(6)
  void clearContextMenuRowActivated() => clearField(6);
  @$pb.TagNumber(6)
  ContextMenuRowID ensureContextMenuRowActivated() => $_ensure(5);

  @$pb.TagNumber(7)
  FieldCheckRequest get fieldCheckRequest => $_getN(6);
  @$pb.TagNumber(7)
  set fieldCheckRequest(FieldCheckRequest v) { setField(7, v); }
  @$pb.TagNumber(7)
  $core.bool hasFieldCheckRequest() => $_has(6);
  @$pb.TagNumber(7)
  void clearFieldCheckRequest() => clearField(7);
  @$pb.TagNumber(7)
  FieldCheckRequest ensureFieldCheckRequest() => $_ensure(6);

  @$pb.TagNumber(8)
  FormSubmit get formSubmit => $_getN(7);
  @$pb.TagNumber(8)
  set formSubmit(FormSubmit v) { setField(8, v); }
  @$pb.TagNumber(8)
  $core.bool hasFormSubmit() => $_has(7);
  @$pb.TagNumber(8)
  void clearFormSubmit() => clearField(8);
  @$pb.TagNumber(8)
  FormSubmit ensureFormSubmit() => $_ensure(7);

  @$pb.TagNumber(9)
  FormClosed get formClosed => $_getN(8);
  @$pb.TagNumber(9)
  set formClosed(FormClosed v) { setField(9, v); }
  @$pb.TagNumber(9)
  $core.bool hasFormClosed() => $_has(8);
  @$pb.TagNumber(9)
  void clearFormClosed() => clearField(9);
  @$pb.TagNumber(9)
  FormClosed ensureFormClosed() => $_ensure(8);
}

enum SrvMessage_Payload {
  formOpen, 
  rootListOpen, 
  rootListAddRows, 
  rootListChangeRows, 
  rootListRemoveRows, 
  fieldCheckResponse, 
  contextMenuOpen, 
  userMessage, 
  closeForm, 
  hideUI, 
  closeUI, 
  notSet
}

class SrvMessage extends $pb.GeneratedMessage {
  static const $core.Map<$core.int, SrvMessage_Payload> _SrvMessage_PayloadByTag = {
    2 : SrvMessage_Payload.formOpen,
    3 : SrvMessage_Payload.rootListOpen,
    4 : SrvMessage_Payload.rootListAddRows,
    5 : SrvMessage_Payload.rootListChangeRows,
    6 : SrvMessage_Payload.rootListRemoveRows,
    7 : SrvMessage_Payload.fieldCheckResponse,
    8 : SrvMessage_Payload.contextMenuOpen,
    9 : SrvMessage_Payload.userMessage,
    10 : SrvMessage_Payload.closeForm,
    11 : SrvMessage_Payload.hideUI,
    12 : SrvMessage_Payload.closeUI,
    0 : SrvMessage_Payload.notSet
  };
  static final $pb.BuilderInfo _i = $pb.BuilderInfo(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'SrvMessage', package: const $pb.PackageName(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'runify'), createEmptyInstance: create)
    ..oo(0, [2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12])
    ..a<$core.int>(1, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'formID', $pb.PbFieldType.OU3, protoName: 'formID')
    ..aOM<FormOpen>(2, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'formOpen', protoName: 'formOpen', subBuilder: FormOpen.create)
    ..aOM<RootListOpen>(3, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'rootListOpen', protoName: 'rootListOpen', subBuilder: RootListOpen.create)
    ..aOM<RootListAddRows>(4, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'rootListAddRows', protoName: 'rootListAddRows', subBuilder: RootListAddRows.create)
    ..aOM<RootListChangeRows>(5, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'rootListChangeRows', protoName: 'rootListChangeRows', subBuilder: RootListChangeRows.create)
    ..aOM<RootListRemoveRows>(6, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'rootListRemoveRows', protoName: 'rootListRemoveRows', subBuilder: RootListRemoveRows.create)
    ..aOM<FieldCheckResponse>(7, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'fieldCheckResponse', protoName: 'fieldCheckResponse', subBuilder: FieldCheckResponse.create)
    ..aOM<ContextMenuOpen>(8, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'contextMenuOpen', protoName: 'contextMenuOpen', subBuilder: ContextMenuOpen.create)
    ..aOM<UserMessage>(9, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'userMessage', protoName: 'userMessage', subBuilder: UserMessage.create)
    ..aOM<CloseForm>(10, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'closeForm', protoName: 'closeForm', subBuilder: CloseForm.create)
    ..aOM<HideUI>(11, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'hideUI', protoName: 'hideUI', subBuilder: HideUI.create)
    ..aOM<CloseUI>(12, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'closeUI', protoName: 'closeUI', subBuilder: CloseUI.create)
    ..hasRequiredFields = false
  ;

  SrvMessage._() : super();
  factory SrvMessage({
    $core.int? formID,
    FormOpen? formOpen,
    RootListOpen? rootListOpen,
    RootListAddRows? rootListAddRows,
    RootListChangeRows? rootListChangeRows,
    RootListRemoveRows? rootListRemoveRows,
    FieldCheckResponse? fieldCheckResponse,
    ContextMenuOpen? contextMenuOpen,
    UserMessage? userMessage,
    CloseForm? closeForm,
    HideUI? hideUI,
    CloseUI? closeUI,
  }) {
    final _result = create();
    if (formID != null) {
      _result.formID = formID;
    }
    if (formOpen != null) {
      _result.formOpen = formOpen;
    }
    if (rootListOpen != null) {
      _result.rootListOpen = rootListOpen;
    }
    if (rootListAddRows != null) {
      _result.rootListAddRows = rootListAddRows;
    }
    if (rootListChangeRows != null) {
      _result.rootListChangeRows = rootListChangeRows;
    }
    if (rootListRemoveRows != null) {
      _result.rootListRemoveRows = rootListRemoveRows;
    }
    if (fieldCheckResponse != null) {
      _result.fieldCheckResponse = fieldCheckResponse;
    }
    if (contextMenuOpen != null) {
      _result.contextMenuOpen = contextMenuOpen;
    }
    if (userMessage != null) {
      _result.userMessage = userMessage;
    }
    if (closeForm != null) {
      _result.closeForm = closeForm;
    }
    if (hideUI != null) {
      _result.hideUI = hideUI;
    }
    if (closeUI != null) {
      _result.closeUI = closeUI;
    }
    return _result;
  }
  factory SrvMessage.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory SrvMessage.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.deepCopy] instead. '
  'Will be removed in next major version')
  SrvMessage clone() => SrvMessage()..mergeFromMessage(this);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.rebuild] instead. '
  'Will be removed in next major version')
  SrvMessage copyWith(void Function(SrvMessage) updates) => super.copyWith((message) => updates(message as SrvMessage)) as SrvMessage; // ignore: deprecated_member_use
  $pb.BuilderInfo get info_ => _i;
  @$core.pragma('dart2js:noInline')
  static SrvMessage create() => SrvMessage._();
  SrvMessage createEmptyInstance() => create();
  static $pb.PbList<SrvMessage> createRepeated() => $pb.PbList<SrvMessage>();
  @$core.pragma('dart2js:noInline')
  static SrvMessage getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<SrvMessage>(create);
  static SrvMessage? _defaultInstance;

  SrvMessage_Payload whichPayload() => _SrvMessage_PayloadByTag[$_whichOneof(0)]!;
  void clearPayload() => clearField($_whichOneof(0));

  @$pb.TagNumber(1)
  $core.int get formID => $_getIZ(0);
  @$pb.TagNumber(1)
  set formID($core.int v) { $_setUnsignedInt32(0, v); }
  @$pb.TagNumber(1)
  $core.bool hasFormID() => $_has(0);
  @$pb.TagNumber(1)
  void clearFormID() => clearField(1);

  @$pb.TagNumber(2)
  FormOpen get formOpen => $_getN(1);
  @$pb.TagNumber(2)
  set formOpen(FormOpen v) { setField(2, v); }
  @$pb.TagNumber(2)
  $core.bool hasFormOpen() => $_has(1);
  @$pb.TagNumber(2)
  void clearFormOpen() => clearField(2);
  @$pb.TagNumber(2)
  FormOpen ensureFormOpen() => $_ensure(1);

  @$pb.TagNumber(3)
  RootListOpen get rootListOpen => $_getN(2);
  @$pb.TagNumber(3)
  set rootListOpen(RootListOpen v) { setField(3, v); }
  @$pb.TagNumber(3)
  $core.bool hasRootListOpen() => $_has(2);
  @$pb.TagNumber(3)
  void clearRootListOpen() => clearField(3);
  @$pb.TagNumber(3)
  RootListOpen ensureRootListOpen() => $_ensure(2);

  @$pb.TagNumber(4)
  RootListAddRows get rootListAddRows => $_getN(3);
  @$pb.TagNumber(4)
  set rootListAddRows(RootListAddRows v) { setField(4, v); }
  @$pb.TagNumber(4)
  $core.bool hasRootListAddRows() => $_has(3);
  @$pb.TagNumber(4)
  void clearRootListAddRows() => clearField(4);
  @$pb.TagNumber(4)
  RootListAddRows ensureRootListAddRows() => $_ensure(3);

  @$pb.TagNumber(5)
  RootListChangeRows get rootListChangeRows => $_getN(4);
  @$pb.TagNumber(5)
  set rootListChangeRows(RootListChangeRows v) { setField(5, v); }
  @$pb.TagNumber(5)
  $core.bool hasRootListChangeRows() => $_has(4);
  @$pb.TagNumber(5)
  void clearRootListChangeRows() => clearField(5);
  @$pb.TagNumber(5)
  RootListChangeRows ensureRootListChangeRows() => $_ensure(4);

  @$pb.TagNumber(6)
  RootListRemoveRows get rootListRemoveRows => $_getN(5);
  @$pb.TagNumber(6)
  set rootListRemoveRows(RootListRemoveRows v) { setField(6, v); }
  @$pb.TagNumber(6)
  $core.bool hasRootListRemoveRows() => $_has(5);
  @$pb.TagNumber(6)
  void clearRootListRemoveRows() => clearField(6);
  @$pb.TagNumber(6)
  RootListRemoveRows ensureRootListRemoveRows() => $_ensure(5);

  @$pb.TagNumber(7)
  FieldCheckResponse get fieldCheckResponse => $_getN(6);
  @$pb.TagNumber(7)
  set fieldCheckResponse(FieldCheckResponse v) { setField(7, v); }
  @$pb.TagNumber(7)
  $core.bool hasFieldCheckResponse() => $_has(6);
  @$pb.TagNumber(7)
  void clearFieldCheckResponse() => clearField(7);
  @$pb.TagNumber(7)
  FieldCheckResponse ensureFieldCheckResponse() => $_ensure(6);

  @$pb.TagNumber(8)
  ContextMenuOpen get contextMenuOpen => $_getN(7);
  @$pb.TagNumber(8)
  set contextMenuOpen(ContextMenuOpen v) { setField(8, v); }
  @$pb.TagNumber(8)
  $core.bool hasContextMenuOpen() => $_has(7);
  @$pb.TagNumber(8)
  void clearContextMenuOpen() => clearField(8);
  @$pb.TagNumber(8)
  ContextMenuOpen ensureContextMenuOpen() => $_ensure(7);

  @$pb.TagNumber(9)
  UserMessage get userMessage => $_getN(8);
  @$pb.TagNumber(9)
  set userMessage(UserMessage v) { setField(9, v); }
  @$pb.TagNumber(9)
  $core.bool hasUserMessage() => $_has(8);
  @$pb.TagNumber(9)
  void clearUserMessage() => clearField(9);
  @$pb.TagNumber(9)
  UserMessage ensureUserMessage() => $_ensure(8);

  @$pb.TagNumber(10)
  CloseForm get closeForm => $_getN(9);
  @$pb.TagNumber(10)
  set closeForm(CloseForm v) { setField(10, v); }
  @$pb.TagNumber(10)
  $core.bool hasCloseForm() => $_has(9);
  @$pb.TagNumber(10)
  void clearCloseForm() => clearField(10);
  @$pb.TagNumber(10)
  CloseForm ensureCloseForm() => $_ensure(9);

  @$pb.TagNumber(11)
  HideUI get hideUI => $_getN(10);
  @$pb.TagNumber(11)
  set hideUI(HideUI v) { setField(11, v); }
  @$pb.TagNumber(11)
  $core.bool hasHideUI() => $_has(10);
  @$pb.TagNumber(11)
  void clearHideUI() => clearField(11);
  @$pb.TagNumber(11)
  HideUI ensureHideUI() => $_ensure(10);

  @$pb.TagNumber(12)
  CloseUI get closeUI => $_getN(11);
  @$pb.TagNumber(12)
  set closeUI(CloseUI v) { setField(12, v); }
  @$pb.TagNumber(12)
  $core.bool hasCloseUI() => $_has(11);
  @$pb.TagNumber(12)
  void clearCloseUI() => clearField(12);
  @$pb.TagNumber(12)
  CloseUI ensureCloseUI() => $_ensure(11);
}

