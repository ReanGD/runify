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
    ..a<$core.int>(1, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'providerID', $pb.PbFieldType.OU3, protoName: 'providerID')
    ..a<$core.int>(2, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'rowID', $pb.PbFieldType.OU3, protoName: 'rowID')
    ..a<$core.int>(3, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'priority', $pb.PbFieldType.OU3)
    ..aOS(4, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'icon')
    ..aOS(5, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'value')
    ..hasRequiredFields = false
  ;

  RootListRow._() : super();
  factory RootListRow({
    $core.int? providerID,
    $core.int? rowID,
    $core.int? priority,
    $core.String? icon,
    $core.String? value,
  }) {
    final _result = create();
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
    if (value != null) {
      _result.value = value;
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

  @$pb.TagNumber(3)
  $core.int get priority => $_getIZ(2);
  @$pb.TagNumber(3)
  set priority($core.int v) { $_setUnsignedInt32(2, v); }
  @$pb.TagNumber(3)
  $core.bool hasPriority() => $_has(2);
  @$pb.TagNumber(3)
  void clearPriority() => clearField(3);

  @$pb.TagNumber(4)
  $core.String get icon => $_getSZ(3);
  @$pb.TagNumber(4)
  set icon($core.String v) { $_setString(3, v); }
  @$pb.TagNumber(4)
  $core.bool hasIcon() => $_has(3);
  @$pb.TagNumber(4)
  void clearIcon() => clearField(4);

  @$pb.TagNumber(5)
  $core.String get value => $_getSZ(4);
  @$pb.TagNumber(5)
  set value($core.String v) { $_setString(4, v); }
  @$pb.TagNumber(5)
  $core.bool hasValue() => $_has(4);
  @$pb.TagNumber(5)
  void clearValue() => clearField(5);
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
    ..aOS(2, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'value')
    ..hasRequiredFields = false
  ;

  ContextMenuRow._() : super();
  factory ContextMenuRow({
    $core.int? rowID,
    $core.String? value,
  }) {
    final _result = create();
    if (rowID != null) {
      _result.rowID = rowID;
    }
    if (value != null) {
      _result.value = value;
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
  $core.String get value => $_getSZ(1);
  @$pb.TagNumber(2)
  set value($core.String v) { $_setString(1, v); }
  @$pb.TagNumber(2)
  $core.bool hasValue() => $_has(1);
  @$pb.TagNumber(2)
  void clearValue() => clearField(2);
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

class FormAction extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'FormAction', package: const $pb.PackageName(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'runify'), createEmptyInstance: create)
    ..e<FormActionType>(1, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'actionType', $pb.PbFieldType.OE, protoName: 'actionType', defaultOrMaker: FormActionType.CLOSE_ALL, valueOf: FormActionType.valueOf, enumValues: FormActionType.values)
    ..aOM<UserMessage>(2, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'message', subBuilder: UserMessage.create)
    ..hasRequiredFields = false
  ;

  FormAction._() : super();
  factory FormAction({
    FormActionType? actionType,
    UserMessage? message,
  }) {
    final _result = create();
    if (actionType != null) {
      _result.actionType = actionType;
    }
    if (message != null) {
      _result.message = message;
    }
    return _result;
  }
  factory FormAction.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory FormAction.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.deepCopy] instead. '
  'Will be removed in next major version')
  FormAction clone() => FormAction()..mergeFromMessage(this);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.rebuild] instead. '
  'Will be removed in next major version')
  FormAction copyWith(void Function(FormAction) updates) => super.copyWith((message) => updates(message as FormAction)) as FormAction; // ignore: deprecated_member_use
  $pb.BuilderInfo get info_ => _i;
  @$core.pragma('dart2js:noInline')
  static FormAction create() => FormAction._();
  FormAction createEmptyInstance() => create();
  static $pb.PbList<FormAction> createRepeated() => $pb.PbList<FormAction>();
  @$core.pragma('dart2js:noInline')
  static FormAction getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<FormAction>(create);
  static FormAction? _defaultInstance;

  @$pb.TagNumber(1)
  FormActionType get actionType => $_getN(0);
  @$pb.TagNumber(1)
  set actionType(FormActionType v) { setField(1, v); }
  @$pb.TagNumber(1)
  $core.bool hasActionType() => $_has(0);
  @$pb.TagNumber(1)
  void clearActionType() => clearField(1);

  @$pb.TagNumber(2)
  UserMessage get message => $_getN(1);
  @$pb.TagNumber(2)
  set message(UserMessage v) { setField(2, v); }
  @$pb.TagNumber(2)
  $core.bool hasMessage() => $_has(1);
  @$pb.TagNumber(2)
  void clearMessage() => clearField(2);
  @$pb.TagNumber(2)
  UserMessage ensureMessage() => $_ensure(1);
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

enum UIMessage_Payload {
  writeLog, 
  filterChanged, 
  rootListRowActivated, 
  rootListMenuActivated, 
  contextMenuRowActivated, 
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
    7 : UIMessage_Payload.formClosed,
    0 : UIMessage_Payload.notSet
  };
  static final $pb.BuilderInfo _i = $pb.BuilderInfo(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'UIMessage', package: const $pb.PackageName(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'runify'), createEmptyInstance: create)
    ..oo(0, [2, 3, 4, 5, 6, 7])
    ..a<$core.int>(1, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'formID', $pb.PbFieldType.OU3, protoName: 'formID')
    ..aOM<WriteLog>(2, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'writeLog', protoName: 'writeLog', subBuilder: WriteLog.create)
    ..aOM<FilterData>(3, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'filterChanged', protoName: 'filterChanged', subBuilder: FilterData.create)
    ..aOM<RootListRowGlobalID>(4, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'rootListRowActivated', protoName: 'rootListRowActivated', subBuilder: RootListRowGlobalID.create)
    ..aOM<RootListRowGlobalID>(5, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'rootListMenuActivated', protoName: 'rootListMenuActivated', subBuilder: RootListRowGlobalID.create)
    ..aOM<ContextMenuRowID>(6, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'contextMenuRowActivated', protoName: 'contextMenuRowActivated', subBuilder: ContextMenuRowID.create)
    ..aOM<FormClosed>(7, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'formClosed', protoName: 'formClosed', subBuilder: FormClosed.create)
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
  FormClosed get formClosed => $_getN(6);
  @$pb.TagNumber(7)
  set formClosed(FormClosed v) { setField(7, v); }
  @$pb.TagNumber(7)
  $core.bool hasFormClosed() => $_has(6);
  @$pb.TagNumber(7)
  void clearFormClosed() => clearField(7);
  @$pb.TagNumber(7)
  FormClosed ensureFormClosed() => $_ensure(6);
}

enum SrvMessage_Payload {
  rootListOpen, 
  rootListAddRows, 
  rootListChangeRows, 
  rootListRemoveRows, 
  contextMenuOpen, 
  formAction, 
  notSet
}

class SrvMessage extends $pb.GeneratedMessage {
  static const $core.Map<$core.int, SrvMessage_Payload> _SrvMessage_PayloadByTag = {
    2 : SrvMessage_Payload.rootListOpen,
    3 : SrvMessage_Payload.rootListAddRows,
    4 : SrvMessage_Payload.rootListChangeRows,
    5 : SrvMessage_Payload.rootListRemoveRows,
    6 : SrvMessage_Payload.contextMenuOpen,
    7 : SrvMessage_Payload.formAction,
    0 : SrvMessage_Payload.notSet
  };
  static final $pb.BuilderInfo _i = $pb.BuilderInfo(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'SrvMessage', package: const $pb.PackageName(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'runify'), createEmptyInstance: create)
    ..oo(0, [2, 3, 4, 5, 6, 7])
    ..a<$core.int>(1, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'formID', $pb.PbFieldType.OU3, protoName: 'formID')
    ..aOM<RootListOpen>(2, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'rootListOpen', protoName: 'rootListOpen', subBuilder: RootListOpen.create)
    ..aOM<RootListAddRows>(3, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'rootListAddRows', protoName: 'rootListAddRows', subBuilder: RootListAddRows.create)
    ..aOM<RootListChangeRows>(4, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'rootListChangeRows', protoName: 'rootListChangeRows', subBuilder: RootListChangeRows.create)
    ..aOM<RootListRemoveRows>(5, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'rootListRemoveRows', protoName: 'rootListRemoveRows', subBuilder: RootListRemoveRows.create)
    ..aOM<ContextMenuOpen>(6, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'contextMenuOpen', protoName: 'contextMenuOpen', subBuilder: ContextMenuOpen.create)
    ..aOM<FormAction>(7, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'formAction', protoName: 'formAction', subBuilder: FormAction.create)
    ..hasRequiredFields = false
  ;

  SrvMessage._() : super();
  factory SrvMessage({
    $core.int? formID,
    RootListOpen? rootListOpen,
    RootListAddRows? rootListAddRows,
    RootListChangeRows? rootListChangeRows,
    RootListRemoveRows? rootListRemoveRows,
    ContextMenuOpen? contextMenuOpen,
    FormAction? formAction,
  }) {
    final _result = create();
    if (formID != null) {
      _result.formID = formID;
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
    if (contextMenuOpen != null) {
      _result.contextMenuOpen = contextMenuOpen;
    }
    if (formAction != null) {
      _result.formAction = formAction;
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
  RootListOpen get rootListOpen => $_getN(1);
  @$pb.TagNumber(2)
  set rootListOpen(RootListOpen v) { setField(2, v); }
  @$pb.TagNumber(2)
  $core.bool hasRootListOpen() => $_has(1);
  @$pb.TagNumber(2)
  void clearRootListOpen() => clearField(2);
  @$pb.TagNumber(2)
  RootListOpen ensureRootListOpen() => $_ensure(1);

  @$pb.TagNumber(3)
  RootListAddRows get rootListAddRows => $_getN(2);
  @$pb.TagNumber(3)
  set rootListAddRows(RootListAddRows v) { setField(3, v); }
  @$pb.TagNumber(3)
  $core.bool hasRootListAddRows() => $_has(2);
  @$pb.TagNumber(3)
  void clearRootListAddRows() => clearField(3);
  @$pb.TagNumber(3)
  RootListAddRows ensureRootListAddRows() => $_ensure(2);

  @$pb.TagNumber(4)
  RootListChangeRows get rootListChangeRows => $_getN(3);
  @$pb.TagNumber(4)
  set rootListChangeRows(RootListChangeRows v) { setField(4, v); }
  @$pb.TagNumber(4)
  $core.bool hasRootListChangeRows() => $_has(3);
  @$pb.TagNumber(4)
  void clearRootListChangeRows() => clearField(4);
  @$pb.TagNumber(4)
  RootListChangeRows ensureRootListChangeRows() => $_ensure(3);

  @$pb.TagNumber(5)
  RootListRemoveRows get rootListRemoveRows => $_getN(4);
  @$pb.TagNumber(5)
  set rootListRemoveRows(RootListRemoveRows v) { setField(5, v); }
  @$pb.TagNumber(5)
  $core.bool hasRootListRemoveRows() => $_has(4);
  @$pb.TagNumber(5)
  void clearRootListRemoveRows() => clearField(5);
  @$pb.TagNumber(5)
  RootListRemoveRows ensureRootListRemoveRows() => $_ensure(4);

  @$pb.TagNumber(6)
  ContextMenuOpen get contextMenuOpen => $_getN(5);
  @$pb.TagNumber(6)
  set contextMenuOpen(ContextMenuOpen v) { setField(6, v); }
  @$pb.TagNumber(6)
  $core.bool hasContextMenuOpen() => $_has(5);
  @$pb.TagNumber(6)
  void clearContextMenuOpen() => clearField(6);
  @$pb.TagNumber(6)
  ContextMenuOpen ensureContextMenuOpen() => $_ensure(5);

  @$pb.TagNumber(7)
  FormAction get formAction => $_getN(6);
  @$pb.TagNumber(7)
  set formAction(FormAction v) { setField(7, v); }
  @$pb.TagNumber(7)
  $core.bool hasFormAction() => $_has(6);
  @$pb.TagNumber(7)
  void clearFormAction() => clearField(7);
  @$pb.TagNumber(7)
  FormAction ensureFormAction() => $_ensure(6);
}

