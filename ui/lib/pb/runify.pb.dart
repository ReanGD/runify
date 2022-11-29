///
//  Generated code. Do not modify.
//  source: runify.proto
//
// @dart = 2.12
// ignore_for_file: annotate_overrides,camel_case_types,constant_identifier_names,directives_ordering,library_prefixes,non_constant_identifier_names,prefer_final_fields,return_of_invalid_type,unnecessary_const,unnecessary_import,unnecessary_this,unused_import,unused_shown_name

import 'dart:core' as $core;

import 'package:fixnum/fixnum.dart' as $fixnum;
import 'package:protobuf/protobuf.dart' as $pb;

import 'runify.pbenum.dart';

export 'runify.pbenum.dart';

class Empty extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'Empty', package: const $pb.PackageName(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'runify'), createEmptyInstance: create)
    ..hasRequiredFields = false
  ;

  Empty._() : super();
  factory Empty() => create();
  factory Empty.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory Empty.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.deepCopy] instead. '
  'Will be removed in next major version')
  Empty clone() => Empty()..mergeFromMessage(this);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.rebuild] instead. '
  'Will be removed in next major version')
  Empty copyWith(void Function(Empty) updates) => super.copyWith((message) => updates(message as Empty)) as Empty; // ignore: deprecated_member_use
  $pb.BuilderInfo get info_ => _i;
  @$core.pragma('dart2js:noInline')
  static Empty create() => Empty._();
  Empty createEmptyInstance() => create();
  static $pb.PbList<Empty> createRepeated() => $pb.PbList<Empty>();
  @$core.pragma('dart2js:noInline')
  static Empty getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<Empty>(create);
  static Empty? _defaultInstance;
}

class CardItem extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'CardItem', package: const $pb.PackageName(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'runify'), createEmptyInstance: create)
    ..a<$fixnum.Int64>(1, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'cardID', $pb.PbFieldType.OU6, protoName: 'cardID', defaultOrMaker: $fixnum.Int64.ZERO)
    ..aOS(2, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'name')
    ..aOS(3, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'icon')
    ..hasRequiredFields = false
  ;

  CardItem._() : super();
  factory CardItem({
    $fixnum.Int64? cardID,
    $core.String? name,
    $core.String? icon,
  }) {
    final _result = create();
    if (cardID != null) {
      _result.cardID = cardID;
    }
    if (name != null) {
      _result.name = name;
    }
    if (icon != null) {
      _result.icon = icon;
    }
    return _result;
  }
  factory CardItem.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory CardItem.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.deepCopy] instead. '
  'Will be removed in next major version')
  CardItem clone() => CardItem()..mergeFromMessage(this);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.rebuild] instead. '
  'Will be removed in next major version')
  CardItem copyWith(void Function(CardItem) updates) => super.copyWith((message) => updates(message as CardItem)) as CardItem; // ignore: deprecated_member_use
  $pb.BuilderInfo get info_ => _i;
  @$core.pragma('dart2js:noInline')
  static CardItem create() => CardItem._();
  CardItem createEmptyInstance() => create();
  static $pb.PbList<CardItem> createRepeated() => $pb.PbList<CardItem>();
  @$core.pragma('dart2js:noInline')
  static CardItem getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<CardItem>(create);
  static CardItem? _defaultInstance;

  @$pb.TagNumber(1)
  $fixnum.Int64 get cardID => $_getI64(0);
  @$pb.TagNumber(1)
  set cardID($fixnum.Int64 v) { $_setInt64(0, v); }
  @$pb.TagNumber(1)
  $core.bool hasCardID() => $_has(0);
  @$pb.TagNumber(1)
  void clearCardID() => clearField(1);

  @$pb.TagNumber(2)
  $core.String get name => $_getSZ(1);
  @$pb.TagNumber(2)
  set name($core.String v) { $_setString(1, v); }
  @$pb.TagNumber(2)
  $core.bool hasName() => $_has(1);
  @$pb.TagNumber(2)
  void clearName() => clearField(2);

  @$pb.TagNumber(3)
  $core.String get icon => $_getSZ(2);
  @$pb.TagNumber(3)
  set icon($core.String v) { $_setString(2, v); }
  @$pb.TagNumber(3)
  $core.bool hasIcon() => $_has(2);
  @$pb.TagNumber(3)
  void clearIcon() => clearField(3);
}

class SelectedCard extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'SelectedCard', package: const $pb.PackageName(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'runify'), createEmptyInstance: create)
    ..a<$fixnum.Int64>(1, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'cardID', $pb.PbFieldType.OU6, protoName: 'cardID', defaultOrMaker: $fixnum.Int64.ZERO)
    ..hasRequiredFields = false
  ;

  SelectedCard._() : super();
  factory SelectedCard({
    $fixnum.Int64? cardID,
  }) {
    final _result = create();
    if (cardID != null) {
      _result.cardID = cardID;
    }
    return _result;
  }
  factory SelectedCard.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory SelectedCard.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.deepCopy] instead. '
  'Will be removed in next major version')
  SelectedCard clone() => SelectedCard()..mergeFromMessage(this);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.rebuild] instead. '
  'Will be removed in next major version')
  SelectedCard copyWith(void Function(SelectedCard) updates) => super.copyWith((message) => updates(message as SelectedCard)) as SelectedCard; // ignore: deprecated_member_use
  $pb.BuilderInfo get info_ => _i;
  @$core.pragma('dart2js:noInline')
  static SelectedCard create() => SelectedCard._();
  SelectedCard createEmptyInstance() => create();
  static $pb.PbList<SelectedCard> createRepeated() => $pb.PbList<SelectedCard>();
  @$core.pragma('dart2js:noInline')
  static SelectedCard getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<SelectedCard>(create);
  static SelectedCard? _defaultInstance;

  @$pb.TagNumber(1)
  $fixnum.Int64 get cardID => $_getI64(0);
  @$pb.TagNumber(1)
  set cardID($fixnum.Int64 v) { $_setInt64(0, v); }
  @$pb.TagNumber(1)
  $core.bool hasCardID() => $_has(0);
  @$pb.TagNumber(1)
  void clearCardID() => clearField(1);
}

class ActionItem extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'ActionItem', package: const $pb.PackageName(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'runify'), createEmptyInstance: create)
    ..a<$core.int>(1, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'actionID', $pb.PbFieldType.OU3, protoName: 'actionID')
    ..aOS(2, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'name')
    ..hasRequiredFields = false
  ;

  ActionItem._() : super();
  factory ActionItem({
    $core.int? actionID,
    $core.String? name,
  }) {
    final _result = create();
    if (actionID != null) {
      _result.actionID = actionID;
    }
    if (name != null) {
      _result.name = name;
    }
    return _result;
  }
  factory ActionItem.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory ActionItem.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.deepCopy] instead. '
  'Will be removed in next major version')
  ActionItem clone() => ActionItem()..mergeFromMessage(this);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.rebuild] instead. '
  'Will be removed in next major version')
  ActionItem copyWith(void Function(ActionItem) updates) => super.copyWith((message) => updates(message as ActionItem)) as ActionItem; // ignore: deprecated_member_use
  $pb.BuilderInfo get info_ => _i;
  @$core.pragma('dart2js:noInline')
  static ActionItem create() => ActionItem._();
  ActionItem createEmptyInstance() => create();
  static $pb.PbList<ActionItem> createRepeated() => $pb.PbList<ActionItem>();
  @$core.pragma('dart2js:noInline')
  static ActionItem getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<ActionItem>(create);
  static ActionItem? _defaultInstance;

  @$pb.TagNumber(1)
  $core.int get actionID => $_getIZ(0);
  @$pb.TagNumber(1)
  set actionID($core.int v) { $_setUnsignedInt32(0, v); }
  @$pb.TagNumber(1)
  $core.bool hasActionID() => $_has(0);
  @$pb.TagNumber(1)
  void clearActionID() => clearField(1);

  @$pb.TagNumber(2)
  $core.String get name => $_getSZ(1);
  @$pb.TagNumber(2)
  set name($core.String v) { $_setString(1, v); }
  @$pb.TagNumber(2)
  $core.bool hasName() => $_has(1);
  @$pb.TagNumber(2)
  void clearName() => clearField(2);
}

class Actions extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'Actions', package: const $pb.PackageName(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'runify'), createEmptyInstance: create)
    ..pc<ActionItem>(1, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'items', $pb.PbFieldType.PM, subBuilder: ActionItem.create)
    ..hasRequiredFields = false
  ;

  Actions._() : super();
  factory Actions({
    $core.Iterable<ActionItem>? items,
  }) {
    final _result = create();
    if (items != null) {
      _result.items.addAll(items);
    }
    return _result;
  }
  factory Actions.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory Actions.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.deepCopy] instead. '
  'Will be removed in next major version')
  Actions clone() => Actions()..mergeFromMessage(this);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.rebuild] instead. '
  'Will be removed in next major version')
  Actions copyWith(void Function(Actions) updates) => super.copyWith((message) => updates(message as Actions)) as Actions; // ignore: deprecated_member_use
  $pb.BuilderInfo get info_ => _i;
  @$core.pragma('dart2js:noInline')
  static Actions create() => Actions._();
  Actions createEmptyInstance() => create();
  static $pb.PbList<Actions> createRepeated() => $pb.PbList<Actions>();
  @$core.pragma('dart2js:noInline')
  static Actions getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<Actions>(create);
  static Actions? _defaultInstance;

  @$pb.TagNumber(1)
  $core.List<ActionItem> get items => $_getList(0);
}

class SelectedAction extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'SelectedAction', package: const $pb.PackageName(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'runify'), createEmptyInstance: create)
    ..a<$core.int>(1, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'actionID', $pb.PbFieldType.OU3, protoName: 'actionID')
    ..a<$fixnum.Int64>(2, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'cardID', $pb.PbFieldType.OU6, protoName: 'cardID', defaultOrMaker: $fixnum.Int64.ZERO)
    ..hasRequiredFields = false
  ;

  SelectedAction._() : super();
  factory SelectedAction({
    $core.int? actionID,
    $fixnum.Int64? cardID,
  }) {
    final _result = create();
    if (actionID != null) {
      _result.actionID = actionID;
    }
    if (cardID != null) {
      _result.cardID = cardID;
    }
    return _result;
  }
  factory SelectedAction.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory SelectedAction.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.deepCopy] instead. '
  'Will be removed in next major version')
  SelectedAction clone() => SelectedAction()..mergeFromMessage(this);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.rebuild] instead. '
  'Will be removed in next major version')
  SelectedAction copyWith(void Function(SelectedAction) updates) => super.copyWith((message) => updates(message as SelectedAction)) as SelectedAction; // ignore: deprecated_member_use
  $pb.BuilderInfo get info_ => _i;
  @$core.pragma('dart2js:noInline')
  static SelectedAction create() => SelectedAction._();
  SelectedAction createEmptyInstance() => create();
  static $pb.PbList<SelectedAction> createRepeated() => $pb.PbList<SelectedAction>();
  @$core.pragma('dart2js:noInline')
  static SelectedAction getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<SelectedAction>(create);
  static SelectedAction? _defaultInstance;

  @$pb.TagNumber(1)
  $core.int get actionID => $_getIZ(0);
  @$pb.TagNumber(1)
  set actionID($core.int v) { $_setUnsignedInt32(0, v); }
  @$pb.TagNumber(1)
  $core.bool hasActionID() => $_has(0);
  @$pb.TagNumber(1)
  void clearActionID() => clearField(1);

  @$pb.TagNumber(2)
  $fixnum.Int64 get cardID => $_getI64(1);
  @$pb.TagNumber(2)
  set cardID($fixnum.Int64 v) { $_setInt64(1, v); }
  @$pb.TagNumber(2)
  $core.bool hasCardID() => $_has(1);
  @$pb.TagNumber(2)
  void clearCardID() => clearField(2);
}

class Form extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'Form', package: const $pb.PackageName(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'runify'), createEmptyInstance: create)
    ..aOS(1, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'layout')
    ..pc<CardItem>(2, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'cards', $pb.PbFieldType.PM, subBuilder: CardItem.create)
    ..hasRequiredFields = false
  ;

  Form._() : super();
  factory Form({
    $core.String? layout,
    $core.Iterable<CardItem>? cards,
  }) {
    final _result = create();
    if (layout != null) {
      _result.layout = layout;
    }
    if (cards != null) {
      _result.cards.addAll(cards);
    }
    return _result;
  }
  factory Form.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory Form.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.deepCopy] instead. '
  'Will be removed in next major version')
  Form clone() => Form()..mergeFromMessage(this);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.rebuild] instead. '
  'Will be removed in next major version')
  Form copyWith(void Function(Form) updates) => super.copyWith((message) => updates(message as Form)) as Form; // ignore: deprecated_member_use
  $pb.BuilderInfo get info_ => _i;
  @$core.pragma('dart2js:noInline')
  static Form create() => Form._();
  Form createEmptyInstance() => create();
  static $pb.PbList<Form> createRepeated() => $pb.PbList<Form>();
  @$core.pragma('dart2js:noInline')
  static Form getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<Form>(create);
  static Form? _defaultInstance;

  @$pb.TagNumber(1)
  $core.String get layout => $_getSZ(0);
  @$pb.TagNumber(1)
  set layout($core.String v) { $_setString(0, v); }
  @$pb.TagNumber(1)
  $core.bool hasLayout() => $_has(0);
  @$pb.TagNumber(1)
  void clearLayout() => clearField(1);

  @$pb.TagNumber(2)
  $core.List<CardItem> get cards => $_getList(1);
}

class HideWindow extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'HideWindow', package: const $pb.PackageName(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'runify'), createEmptyInstance: create)
    ..aOS(1, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'error')
    ..hasRequiredFields = false
  ;

  HideWindow._() : super();
  factory HideWindow({
    $core.String? error,
  }) {
    final _result = create();
    if (error != null) {
      _result.error = error;
    }
    return _result;
  }
  factory HideWindow.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory HideWindow.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.deepCopy] instead. '
  'Will be removed in next major version')
  HideWindow clone() => HideWindow()..mergeFromMessage(this);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.rebuild] instead. '
  'Will be removed in next major version')
  HideWindow copyWith(void Function(HideWindow) updates) => super.copyWith((message) => updates(message as HideWindow)) as HideWindow; // ignore: deprecated_member_use
  $pb.BuilderInfo get info_ => _i;
  @$core.pragma('dart2js:noInline')
  static HideWindow create() => HideWindow._();
  HideWindow createEmptyInstance() => create();
  static $pb.PbList<HideWindow> createRepeated() => $pb.PbList<HideWindow>();
  @$core.pragma('dart2js:noInline')
  static HideWindow getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<HideWindow>(create);
  static HideWindow? _defaultInstance;

  @$pb.TagNumber(1)
  $core.String get error => $_getSZ(0);
  @$pb.TagNumber(1)
  set error($core.String v) { $_setString(0, v); }
  @$pb.TagNumber(1)
  $core.bool hasError() => $_has(0);
  @$pb.TagNumber(1)
  void clearError() => clearField(1);
}

enum Result_Payload {
  form, 
  empty, 
  hide, 
  notSet
}

class Result extends $pb.GeneratedMessage {
  static const $core.Map<$core.int, Result_Payload> _Result_PayloadByTag = {
    1 : Result_Payload.form,
    2 : Result_Payload.empty,
    3 : Result_Payload.hide,
    0 : Result_Payload.notSet
  };
  static final $pb.BuilderInfo _i = $pb.BuilderInfo(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'Result', package: const $pb.PackageName(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'runify'), createEmptyInstance: create)
    ..oo(0, [1, 2, 3])
    ..aOM<Form>(1, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'form', subBuilder: Form.create)
    ..aOM<Empty>(2, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'empty', subBuilder: Empty.create)
    ..aOM<HideWindow>(3, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'hide', subBuilder: HideWindow.create)
    ..hasRequiredFields = false
  ;

  Result._() : super();
  factory Result({
    Form? form,
    Empty? empty,
    HideWindow? hide,
  }) {
    final _result = create();
    if (form != null) {
      _result.form = form;
    }
    if (empty != null) {
      _result.empty = empty;
    }
    if (hide != null) {
      _result.hide = hide;
    }
    return _result;
  }
  factory Result.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory Result.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.deepCopy] instead. '
  'Will be removed in next major version')
  Result clone() => Result()..mergeFromMessage(this);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.rebuild] instead. '
  'Will be removed in next major version')
  Result copyWith(void Function(Result) updates) => super.copyWith((message) => updates(message as Result)) as Result; // ignore: deprecated_member_use
  $pb.BuilderInfo get info_ => _i;
  @$core.pragma('dart2js:noInline')
  static Result create() => Result._();
  Result createEmptyInstance() => create();
  static $pb.PbList<Result> createRepeated() => $pb.PbList<Result>();
  @$core.pragma('dart2js:noInline')
  static Result getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<Result>(create);
  static Result? _defaultInstance;

  Result_Payload whichPayload() => _Result_PayloadByTag[$_whichOneof(0)]!;
  void clearPayload() => clearField($_whichOneof(0));

  @$pb.TagNumber(1)
  Form get form => $_getN(0);
  @$pb.TagNumber(1)
  set form(Form v) { setField(1, v); }
  @$pb.TagNumber(1)
  $core.bool hasForm() => $_has(0);
  @$pb.TagNumber(1)
  void clearForm() => clearField(1);
  @$pb.TagNumber(1)
  Form ensureForm() => $_ensure(0);

  @$pb.TagNumber(2)
  Empty get empty => $_getN(1);
  @$pb.TagNumber(2)
  set empty(Empty v) { setField(2, v); }
  @$pb.TagNumber(2)
  $core.bool hasEmpty() => $_has(1);
  @$pb.TagNumber(2)
  void clearEmpty() => clearField(2);
  @$pb.TagNumber(2)
  Empty ensureEmpty() => $_ensure(1);

  @$pb.TagNumber(3)
  HideWindow get hide => $_getN(2);
  @$pb.TagNumber(3)
  set hide(HideWindow v) { setField(3, v); }
  @$pb.TagNumber(3)
  $core.bool hasHide() => $_has(2);
  @$pb.TagNumber(3)
  void clearHide() => clearField(3);
  @$pb.TagNumber(3)
  HideWindow ensureHide() => $_ensure(2);
}

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

class SetFormState extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'SetFormState', package: const $pb.PackageName(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'runify'), createEmptyInstance: create)
    ..e<FormStateType>(1, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'state', $pb.PbFieldType.OE, defaultOrMaker: FormStateType.SHOW, valueOf: FormStateType.valueOf, enumValues: FormStateType.values)
    ..hasRequiredFields = false
  ;

  SetFormState._() : super();
  factory SetFormState({
    FormStateType? state,
  }) {
    final _result = create();
    if (state != null) {
      _result.state = state;
    }
    return _result;
  }
  factory SetFormState.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory SetFormState.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.deepCopy] instead. '
  'Will be removed in next major version')
  SetFormState clone() => SetFormState()..mergeFromMessage(this);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.rebuild] instead. '
  'Will be removed in next major version')
  SetFormState copyWith(void Function(SetFormState) updates) => super.copyWith((message) => updates(message as SetFormState)) as SetFormState; // ignore: deprecated_member_use
  $pb.BuilderInfo get info_ => _i;
  @$core.pragma('dart2js:noInline')
  static SetFormState create() => SetFormState._();
  SetFormState createEmptyInstance() => create();
  static $pb.PbList<SetFormState> createRepeated() => $pb.PbList<SetFormState>();
  @$core.pragma('dart2js:noInline')
  static SetFormState getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<SetFormState>(create);
  static SetFormState? _defaultInstance;

  @$pb.TagNumber(1)
  FormStateType get state => $_getN(0);
  @$pb.TagNumber(1)
  set state(FormStateType v) { setField(1, v); }
  @$pb.TagNumber(1)
  $core.bool hasState() => $_has(0);
  @$pb.TagNumber(1)
  void clearState() => clearField(1);
}

enum ServiceMsgUI_Payload {
  writeLog, 
  notSet
}

class ServiceMsgUI extends $pb.GeneratedMessage {
  static const $core.Map<$core.int, ServiceMsgUI_Payload> _ServiceMsgUI_PayloadByTag = {
    1 : ServiceMsgUI_Payload.writeLog,
    0 : ServiceMsgUI_Payload.notSet
  };
  static final $pb.BuilderInfo _i = $pb.BuilderInfo(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'ServiceMsgUI', package: const $pb.PackageName(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'runify'), createEmptyInstance: create)
    ..oo(0, [1])
    ..aOM<WriteLog>(1, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'writeLog', protoName: 'writeLog', subBuilder: WriteLog.create)
    ..hasRequiredFields = false
  ;

  ServiceMsgUI._() : super();
  factory ServiceMsgUI({
    WriteLog? writeLog,
  }) {
    final _result = create();
    if (writeLog != null) {
      _result.writeLog = writeLog;
    }
    return _result;
  }
  factory ServiceMsgUI.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory ServiceMsgUI.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.deepCopy] instead. '
  'Will be removed in next major version')
  ServiceMsgUI clone() => ServiceMsgUI()..mergeFromMessage(this);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.rebuild] instead. '
  'Will be removed in next major version')
  ServiceMsgUI copyWith(void Function(ServiceMsgUI) updates) => super.copyWith((message) => updates(message as ServiceMsgUI)) as ServiceMsgUI; // ignore: deprecated_member_use
  $pb.BuilderInfo get info_ => _i;
  @$core.pragma('dart2js:noInline')
  static ServiceMsgUI create() => ServiceMsgUI._();
  ServiceMsgUI createEmptyInstance() => create();
  static $pb.PbList<ServiceMsgUI> createRepeated() => $pb.PbList<ServiceMsgUI>();
  @$core.pragma('dart2js:noInline')
  static ServiceMsgUI getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<ServiceMsgUI>(create);
  static ServiceMsgUI? _defaultInstance;

  ServiceMsgUI_Payload whichPayload() => _ServiceMsgUI_PayloadByTag[$_whichOneof(0)]!;
  void clearPayload() => clearField($_whichOneof(0));

  @$pb.TagNumber(1)
  WriteLog get writeLog => $_getN(0);
  @$pb.TagNumber(1)
  set writeLog(WriteLog v) { setField(1, v); }
  @$pb.TagNumber(1)
  $core.bool hasWriteLog() => $_has(0);
  @$pb.TagNumber(1)
  void clearWriteLog() => clearField(1);
  @$pb.TagNumber(1)
  WriteLog ensureWriteLog() => $_ensure(0);
}

enum ServiceMsgSrv_Payload {
  setFormState, 
  notSet
}

class ServiceMsgSrv extends $pb.GeneratedMessage {
  static const $core.Map<$core.int, ServiceMsgSrv_Payload> _ServiceMsgSrv_PayloadByTag = {
    1 : ServiceMsgSrv_Payload.setFormState,
    0 : ServiceMsgSrv_Payload.notSet
  };
  static final $pb.BuilderInfo _i = $pb.BuilderInfo(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'ServiceMsgSrv', package: const $pb.PackageName(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'runify'), createEmptyInstance: create)
    ..oo(0, [1])
    ..aOM<SetFormState>(1, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'setFormState', protoName: 'setFormState', subBuilder: SetFormState.create)
    ..hasRequiredFields = false
  ;

  ServiceMsgSrv._() : super();
  factory ServiceMsgSrv({
    SetFormState? setFormState,
  }) {
    final _result = create();
    if (setFormState != null) {
      _result.setFormState = setFormState;
    }
    return _result;
  }
  factory ServiceMsgSrv.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory ServiceMsgSrv.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.deepCopy] instead. '
  'Will be removed in next major version')
  ServiceMsgSrv clone() => ServiceMsgSrv()..mergeFromMessage(this);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.rebuild] instead. '
  'Will be removed in next major version')
  ServiceMsgSrv copyWith(void Function(ServiceMsgSrv) updates) => super.copyWith((message) => updates(message as ServiceMsgSrv)) as ServiceMsgSrv; // ignore: deprecated_member_use
  $pb.BuilderInfo get info_ => _i;
  @$core.pragma('dart2js:noInline')
  static ServiceMsgSrv create() => ServiceMsgSrv._();
  ServiceMsgSrv createEmptyInstance() => create();
  static $pb.PbList<ServiceMsgSrv> createRepeated() => $pb.PbList<ServiceMsgSrv>();
  @$core.pragma('dart2js:noInline')
  static ServiceMsgSrv getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<ServiceMsgSrv>(create);
  static ServiceMsgSrv? _defaultInstance;

  ServiceMsgSrv_Payload whichPayload() => _ServiceMsgSrv_PayloadByTag[$_whichOneof(0)]!;
  void clearPayload() => clearField($_whichOneof(0));

  @$pb.TagNumber(1)
  SetFormState get setFormState => $_getN(0);
  @$pb.TagNumber(1)
  set setFormState(SetFormState v) { setField(1, v); }
  @$pb.TagNumber(1)
  $core.bool hasSetFormState() => $_has(0);
  @$pb.TagNumber(1)
  void clearSetFormState() => clearField(1);
  @$pb.TagNumber(1)
  SetFormState ensureSetFormState() => $_ensure(0);
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

class RootListRows extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'RootListRows', package: const $pb.PackageName(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'runify'), createEmptyInstance: create)
    ..pc<RootListRow>(1, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'rows', $pb.PbFieldType.PM, subBuilder: RootListRow.create)
    ..hasRequiredFields = false
  ;

  RootListRows._() : super();
  factory RootListRows({
    $core.Iterable<RootListRow>? rows,
  }) {
    final _result = create();
    if (rows != null) {
      _result.rows.addAll(rows);
    }
    return _result;
  }
  factory RootListRows.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory RootListRows.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.deepCopy] instead. '
  'Will be removed in next major version')
  RootListRows clone() => RootListRows()..mergeFromMessage(this);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.rebuild] instead. '
  'Will be removed in next major version')
  RootListRows copyWith(void Function(RootListRows) updates) => super.copyWith((message) => updates(message as RootListRows)) as RootListRows; // ignore: deprecated_member_use
  $pb.BuilderInfo get info_ => _i;
  @$core.pragma('dart2js:noInline')
  static RootListRows create() => RootListRows._();
  RootListRows createEmptyInstance() => create();
  static $pb.PbList<RootListRows> createRepeated() => $pb.PbList<RootListRows>();
  @$core.pragma('dart2js:noInline')
  static RootListRows getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<RootListRows>(create);
  static RootListRows? _defaultInstance;

  @$pb.TagNumber(1)
  $core.List<RootListRow> get rows => $_getList(0);
}

class RootListRowsUpdate extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'RootListRowsUpdate', package: const $pb.PackageName(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'runify'), createEmptyInstance: create)
    ..pc<RootListRows>(1, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'create', $pb.PbFieldType.PM, subBuilder: RootListRows.create)
    ..pc<RootListRows>(2, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'change', $pb.PbFieldType.PM, subBuilder: RootListRows.create)
    ..pc<RootListRowGlobalID>(3, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'Remove', $pb.PbFieldType.PM, protoName: 'Remove', subBuilder: RootListRowGlobalID.create)
    ..hasRequiredFields = false
  ;

  RootListRowsUpdate._() : super();
  factory RootListRowsUpdate({
    $core.Iterable<RootListRows>? create_1,
    $core.Iterable<RootListRows>? change,
    $core.Iterable<RootListRowGlobalID>? remove,
  }) {
    final _result = create();
    if (create_1 != null) {
      _result.create_1.addAll(create_1);
    }
    if (change != null) {
      _result.change.addAll(change);
    }
    if (remove != null) {
      _result.remove.addAll(remove);
    }
    return _result;
  }
  factory RootListRowsUpdate.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory RootListRowsUpdate.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.deepCopy] instead. '
  'Will be removed in next major version')
  RootListRowsUpdate clone() => RootListRowsUpdate()..mergeFromMessage(this);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.rebuild] instead. '
  'Will be removed in next major version')
  RootListRowsUpdate copyWith(void Function(RootListRowsUpdate) updates) => super.copyWith((message) => updates(message as RootListRowsUpdate)) as RootListRowsUpdate; // ignore: deprecated_member_use
  $pb.BuilderInfo get info_ => _i;
  @$core.pragma('dart2js:noInline')
  static RootListRowsUpdate create() => RootListRowsUpdate._();
  RootListRowsUpdate createEmptyInstance() => create();
  static $pb.PbList<RootListRowsUpdate> createRepeated() => $pb.PbList<RootListRowsUpdate>();
  @$core.pragma('dart2js:noInline')
  static RootListRowsUpdate getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<RootListRowsUpdate>(create);
  static RootListRowsUpdate? _defaultInstance;

  @$pb.TagNumber(1)
  $core.List<RootListRows> get create_1 => $_getList(0);

  @$pb.TagNumber(2)
  $core.List<RootListRows> get change => $_getList(1);

  @$pb.TagNumber(3)
  $core.List<RootListRowGlobalID> get remove => $_getList(2);
}

class RootListForm extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'RootListForm', package: const $pb.PackageName(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'runify'), createEmptyInstance: create)
    ..aOM<RootListRows>(1, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'rows', subBuilder: RootListRows.create)
    ..hasRequiredFields = false
  ;

  RootListForm._() : super();
  factory RootListForm({
    RootListRows? rows,
  }) {
    final _result = create();
    if (rows != null) {
      _result.rows = rows;
    }
    return _result;
  }
  factory RootListForm.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory RootListForm.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.deepCopy] instead. '
  'Will be removed in next major version')
  RootListForm clone() => RootListForm()..mergeFromMessage(this);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.rebuild] instead. '
  'Will be removed in next major version')
  RootListForm copyWith(void Function(RootListForm) updates) => super.copyWith((message) => updates(message as RootListForm)) as RootListForm; // ignore: deprecated_member_use
  $pb.BuilderInfo get info_ => _i;
  @$core.pragma('dart2js:noInline')
  static RootListForm create() => RootListForm._();
  RootListForm createEmptyInstance() => create();
  static $pb.PbList<RootListForm> createRepeated() => $pb.PbList<RootListForm>();
  @$core.pragma('dart2js:noInline')
  static RootListForm getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<RootListForm>(create);
  static RootListForm? _defaultInstance;

  @$pb.TagNumber(1)
  RootListRows get rows => $_getN(0);
  @$pb.TagNumber(1)
  set rows(RootListRows v) { setField(1, v); }
  @$pb.TagNumber(1)
  $core.bool hasRows() => $_has(0);
  @$pb.TagNumber(1)
  void clearRows() => clearField(1);
  @$pb.TagNumber(1)
  RootListRows ensureRows() => $_ensure(0);
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

class ContextMenuRows extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'ContextMenuRows', package: const $pb.PackageName(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'runify'), createEmptyInstance: create)
    ..pc<ContextMenuRow>(1, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'rows', $pb.PbFieldType.PM, subBuilder: ContextMenuRow.create)
    ..hasRequiredFields = false
  ;

  ContextMenuRows._() : super();
  factory ContextMenuRows({
    $core.Iterable<ContextMenuRow>? rows,
  }) {
    final _result = create();
    if (rows != null) {
      _result.rows.addAll(rows);
    }
    return _result;
  }
  factory ContextMenuRows.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory ContextMenuRows.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.deepCopy] instead. '
  'Will be removed in next major version')
  ContextMenuRows clone() => ContextMenuRows()..mergeFromMessage(this);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.rebuild] instead. '
  'Will be removed in next major version')
  ContextMenuRows copyWith(void Function(ContextMenuRows) updates) => super.copyWith((message) => updates(message as ContextMenuRows)) as ContextMenuRows; // ignore: deprecated_member_use
  $pb.BuilderInfo get info_ => _i;
  @$core.pragma('dart2js:noInline')
  static ContextMenuRows create() => ContextMenuRows._();
  ContextMenuRows createEmptyInstance() => create();
  static $pb.PbList<ContextMenuRows> createRepeated() => $pb.PbList<ContextMenuRows>();
  @$core.pragma('dart2js:noInline')
  static ContextMenuRows getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<ContextMenuRows>(create);
  static ContextMenuRows? _defaultInstance;

  @$pb.TagNumber(1)
  $core.List<ContextMenuRow> get rows => $_getList(0);
}

class ContextMenuForm extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'ContextMenuForm', package: const $pb.PackageName(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'runify'), createEmptyInstance: create)
    ..aOM<ContextMenuRows>(1, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'rows', subBuilder: ContextMenuRows.create)
    ..hasRequiredFields = false
  ;

  ContextMenuForm._() : super();
  factory ContextMenuForm({
    ContextMenuRows? rows,
  }) {
    final _result = create();
    if (rows != null) {
      _result.rows = rows;
    }
    return _result;
  }
  factory ContextMenuForm.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory ContextMenuForm.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.deepCopy] instead. '
  'Will be removed in next major version')
  ContextMenuForm clone() => ContextMenuForm()..mergeFromMessage(this);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.rebuild] instead. '
  'Will be removed in next major version')
  ContextMenuForm copyWith(void Function(ContextMenuForm) updates) => super.copyWith((message) => updates(message as ContextMenuForm)) as ContextMenuForm; // ignore: deprecated_member_use
  $pb.BuilderInfo get info_ => _i;
  @$core.pragma('dart2js:noInline')
  static ContextMenuForm create() => ContextMenuForm._();
  ContextMenuForm createEmptyInstance() => create();
  static $pb.PbList<ContextMenuForm> createRepeated() => $pb.PbList<ContextMenuForm>();
  @$core.pragma('dart2js:noInline')
  static ContextMenuForm getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<ContextMenuForm>(create);
  static ContextMenuForm? _defaultInstance;

  @$pb.TagNumber(1)
  ContextMenuRows get rows => $_getN(0);
  @$pb.TagNumber(1)
  set rows(ContextMenuRows v) { setField(1, v); }
  @$pb.TagNumber(1)
  $core.bool hasRows() => $_has(0);
  @$pb.TagNumber(1)
  void clearRows() => clearField(1);
  @$pb.TagNumber(1)
  ContextMenuRows ensureRows() => $_ensure(0);
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

enum FormDataMsgUI_Payload {
  filterChanged, 
  rootListRowActivated, 
  rootListMenuActivated, 
  contextMenuRowActivated, 
  notSet
}

class FormDataMsgUI extends $pb.GeneratedMessage {
  static const $core.Map<$core.int, FormDataMsgUI_Payload> _FormDataMsgUI_PayloadByTag = {
    2 : FormDataMsgUI_Payload.filterChanged,
    3 : FormDataMsgUI_Payload.rootListRowActivated,
    4 : FormDataMsgUI_Payload.rootListMenuActivated,
    5 : FormDataMsgUI_Payload.contextMenuRowActivated,
    0 : FormDataMsgUI_Payload.notSet
  };
  static final $pb.BuilderInfo _i = $pb.BuilderInfo(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'FormDataMsgUI', package: const $pb.PackageName(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'runify'), createEmptyInstance: create)
    ..oo(0, [2, 3, 4, 5])
    ..a<$core.int>(1, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'formID', $pb.PbFieldType.OU3, protoName: 'formID')
    ..aOM<FilterData>(2, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'filterChanged', protoName: 'filterChanged', subBuilder: FilterData.create)
    ..aOM<RootListRowGlobalID>(3, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'rootListRowActivated', protoName: 'rootListRowActivated', subBuilder: RootListRowGlobalID.create)
    ..aOM<RootListRowGlobalID>(4, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'rootListMenuActivated', protoName: 'rootListMenuActivated', subBuilder: RootListRowGlobalID.create)
    ..aOM<ContextMenuRowID>(5, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'contextMenuRowActivated', protoName: 'contextMenuRowActivated', subBuilder: ContextMenuRowID.create)
    ..hasRequiredFields = false
  ;

  FormDataMsgUI._() : super();
  factory FormDataMsgUI({
    $core.int? formID,
    FilterData? filterChanged,
    RootListRowGlobalID? rootListRowActivated,
    RootListRowGlobalID? rootListMenuActivated,
    ContextMenuRowID? contextMenuRowActivated,
  }) {
    final _result = create();
    if (formID != null) {
      _result.formID = formID;
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
    return _result;
  }
  factory FormDataMsgUI.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory FormDataMsgUI.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.deepCopy] instead. '
  'Will be removed in next major version')
  FormDataMsgUI clone() => FormDataMsgUI()..mergeFromMessage(this);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.rebuild] instead. '
  'Will be removed in next major version')
  FormDataMsgUI copyWith(void Function(FormDataMsgUI) updates) => super.copyWith((message) => updates(message as FormDataMsgUI)) as FormDataMsgUI; // ignore: deprecated_member_use
  $pb.BuilderInfo get info_ => _i;
  @$core.pragma('dart2js:noInline')
  static FormDataMsgUI create() => FormDataMsgUI._();
  FormDataMsgUI createEmptyInstance() => create();
  static $pb.PbList<FormDataMsgUI> createRepeated() => $pb.PbList<FormDataMsgUI>();
  @$core.pragma('dart2js:noInline')
  static FormDataMsgUI getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<FormDataMsgUI>(create);
  static FormDataMsgUI? _defaultInstance;

  FormDataMsgUI_Payload whichPayload() => _FormDataMsgUI_PayloadByTag[$_whichOneof(0)]!;
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
  FilterData get filterChanged => $_getN(1);
  @$pb.TagNumber(2)
  set filterChanged(FilterData v) { setField(2, v); }
  @$pb.TagNumber(2)
  $core.bool hasFilterChanged() => $_has(1);
  @$pb.TagNumber(2)
  void clearFilterChanged() => clearField(2);
  @$pb.TagNumber(2)
  FilterData ensureFilterChanged() => $_ensure(1);

  @$pb.TagNumber(3)
  RootListRowGlobalID get rootListRowActivated => $_getN(2);
  @$pb.TagNumber(3)
  set rootListRowActivated(RootListRowGlobalID v) { setField(3, v); }
  @$pb.TagNumber(3)
  $core.bool hasRootListRowActivated() => $_has(2);
  @$pb.TagNumber(3)
  void clearRootListRowActivated() => clearField(3);
  @$pb.TagNumber(3)
  RootListRowGlobalID ensureRootListRowActivated() => $_ensure(2);

  @$pb.TagNumber(4)
  RootListRowGlobalID get rootListMenuActivated => $_getN(3);
  @$pb.TagNumber(4)
  set rootListMenuActivated(RootListRowGlobalID v) { setField(4, v); }
  @$pb.TagNumber(4)
  $core.bool hasRootListMenuActivated() => $_has(3);
  @$pb.TagNumber(4)
  void clearRootListMenuActivated() => clearField(4);
  @$pb.TagNumber(4)
  RootListRowGlobalID ensureRootListMenuActivated() => $_ensure(3);

  @$pb.TagNumber(5)
  ContextMenuRowID get contextMenuRowActivated => $_getN(4);
  @$pb.TagNumber(5)
  set contextMenuRowActivated(ContextMenuRowID v) { setField(5, v); }
  @$pb.TagNumber(5)
  $core.bool hasContextMenuRowActivated() => $_has(4);
  @$pb.TagNumber(5)
  void clearContextMenuRowActivated() => clearField(5);
  @$pb.TagNumber(5)
  ContextMenuRowID ensureContextMenuRowActivated() => $_ensure(4);
}

enum FormDataMsgSrv_Payload {
  openRootListForm, 
  updateRootListForm, 
  openContextMenuForm, 
  closeForm, 
  notSet
}

class FormDataMsgSrv extends $pb.GeneratedMessage {
  static const $core.Map<$core.int, FormDataMsgSrv_Payload> _FormDataMsgSrv_PayloadByTag = {
    2 : FormDataMsgSrv_Payload.openRootListForm,
    3 : FormDataMsgSrv_Payload.updateRootListForm,
    4 : FormDataMsgSrv_Payload.openContextMenuForm,
    5 : FormDataMsgSrv_Payload.closeForm,
    0 : FormDataMsgSrv_Payload.notSet
  };
  static final $pb.BuilderInfo _i = $pb.BuilderInfo(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'FormDataMsgSrv', package: const $pb.PackageName(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'runify'), createEmptyInstance: create)
    ..oo(0, [2, 3, 4, 5])
    ..a<$core.int>(1, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'formID', $pb.PbFieldType.OU3, protoName: 'formID')
    ..aOM<RootListForm>(2, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'openRootListForm', protoName: 'openRootListForm', subBuilder: RootListForm.create)
    ..aOM<RootListRowsUpdate>(3, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'updateRootListForm', protoName: 'updateRootListForm', subBuilder: RootListRowsUpdate.create)
    ..aOM<ContextMenuForm>(4, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'openContextMenuForm', protoName: 'openContextMenuForm', subBuilder: ContextMenuForm.create)
    ..aOM<CloseForm>(5, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'closeForm', protoName: 'closeForm', subBuilder: CloseForm.create)
    ..hasRequiredFields = false
  ;

  FormDataMsgSrv._() : super();
  factory FormDataMsgSrv({
    $core.int? formID,
    RootListForm? openRootListForm,
    RootListRowsUpdate? updateRootListForm,
    ContextMenuForm? openContextMenuForm,
    CloseForm? closeForm,
  }) {
    final _result = create();
    if (formID != null) {
      _result.formID = formID;
    }
    if (openRootListForm != null) {
      _result.openRootListForm = openRootListForm;
    }
    if (updateRootListForm != null) {
      _result.updateRootListForm = updateRootListForm;
    }
    if (openContextMenuForm != null) {
      _result.openContextMenuForm = openContextMenuForm;
    }
    if (closeForm != null) {
      _result.closeForm = closeForm;
    }
    return _result;
  }
  factory FormDataMsgSrv.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory FormDataMsgSrv.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.deepCopy] instead. '
  'Will be removed in next major version')
  FormDataMsgSrv clone() => FormDataMsgSrv()..mergeFromMessage(this);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.rebuild] instead. '
  'Will be removed in next major version')
  FormDataMsgSrv copyWith(void Function(FormDataMsgSrv) updates) => super.copyWith((message) => updates(message as FormDataMsgSrv)) as FormDataMsgSrv; // ignore: deprecated_member_use
  $pb.BuilderInfo get info_ => _i;
  @$core.pragma('dart2js:noInline')
  static FormDataMsgSrv create() => FormDataMsgSrv._();
  FormDataMsgSrv createEmptyInstance() => create();
  static $pb.PbList<FormDataMsgSrv> createRepeated() => $pb.PbList<FormDataMsgSrv>();
  @$core.pragma('dart2js:noInline')
  static FormDataMsgSrv getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<FormDataMsgSrv>(create);
  static FormDataMsgSrv? _defaultInstance;

  FormDataMsgSrv_Payload whichPayload() => _FormDataMsgSrv_PayloadByTag[$_whichOneof(0)]!;
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
  RootListForm get openRootListForm => $_getN(1);
  @$pb.TagNumber(2)
  set openRootListForm(RootListForm v) { setField(2, v); }
  @$pb.TagNumber(2)
  $core.bool hasOpenRootListForm() => $_has(1);
  @$pb.TagNumber(2)
  void clearOpenRootListForm() => clearField(2);
  @$pb.TagNumber(2)
  RootListForm ensureOpenRootListForm() => $_ensure(1);

  @$pb.TagNumber(3)
  RootListRowsUpdate get updateRootListForm => $_getN(2);
  @$pb.TagNumber(3)
  set updateRootListForm(RootListRowsUpdate v) { setField(3, v); }
  @$pb.TagNumber(3)
  $core.bool hasUpdateRootListForm() => $_has(2);
  @$pb.TagNumber(3)
  void clearUpdateRootListForm() => clearField(3);
  @$pb.TagNumber(3)
  RootListRowsUpdate ensureUpdateRootListForm() => $_ensure(2);

  @$pb.TagNumber(4)
  ContextMenuForm get openContextMenuForm => $_getN(3);
  @$pb.TagNumber(4)
  set openContextMenuForm(ContextMenuForm v) { setField(4, v); }
  @$pb.TagNumber(4)
  $core.bool hasOpenContextMenuForm() => $_has(3);
  @$pb.TagNumber(4)
  void clearOpenContextMenuForm() => clearField(4);
  @$pb.TagNumber(4)
  ContextMenuForm ensureOpenContextMenuForm() => $_ensure(3);

  @$pb.TagNumber(5)
  CloseForm get closeForm => $_getN(4);
  @$pb.TagNumber(5)
  set closeForm(CloseForm v) { setField(5, v); }
  @$pb.TagNumber(5)
  $core.bool hasCloseForm() => $_has(4);
  @$pb.TagNumber(5)
  void clearCloseForm() => clearField(5);
  @$pb.TagNumber(5)
  CloseForm ensureCloseForm() => $_ensure(4);
}

