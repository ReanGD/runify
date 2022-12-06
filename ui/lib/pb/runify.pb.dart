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

enum FormDataMsgUI_Payload {
  filterChanged, 
  rootListRowActivated, 
  rootListMenuActivated, 
  contextMenuRowActivated, 
  formClosed, 
  notSet
}

class FormDataMsgUI extends $pb.GeneratedMessage {
  static const $core.Map<$core.int, FormDataMsgUI_Payload> _FormDataMsgUI_PayloadByTag = {
    2 : FormDataMsgUI_Payload.filterChanged,
    3 : FormDataMsgUI_Payload.rootListRowActivated,
    4 : FormDataMsgUI_Payload.rootListMenuActivated,
    5 : FormDataMsgUI_Payload.contextMenuRowActivated,
    6 : FormDataMsgUI_Payload.formClosed,
    0 : FormDataMsgUI_Payload.notSet
  };
  static final $pb.BuilderInfo _i = $pb.BuilderInfo(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'FormDataMsgUI', package: const $pb.PackageName(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'runify'), createEmptyInstance: create)
    ..oo(0, [2, 3, 4, 5, 6])
    ..a<$core.int>(1, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'formID', $pb.PbFieldType.OU3, protoName: 'formID')
    ..aOM<FilterData>(2, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'filterChanged', protoName: 'filterChanged', subBuilder: FilterData.create)
    ..aOM<RootListRowGlobalID>(3, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'rootListRowActivated', protoName: 'rootListRowActivated', subBuilder: RootListRowGlobalID.create)
    ..aOM<RootListRowGlobalID>(4, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'rootListMenuActivated', protoName: 'rootListMenuActivated', subBuilder: RootListRowGlobalID.create)
    ..aOM<ContextMenuRowID>(5, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'contextMenuRowActivated', protoName: 'contextMenuRowActivated', subBuilder: ContextMenuRowID.create)
    ..aOM<FormClosed>(6, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'formClosed', protoName: 'formClosed', subBuilder: FormClosed.create)
    ..hasRequiredFields = false
  ;

  FormDataMsgUI._() : super();
  factory FormDataMsgUI({
    $core.int? formID,
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

  @$pb.TagNumber(6)
  FormClosed get formClosed => $_getN(5);
  @$pb.TagNumber(6)
  set formClosed(FormClosed v) { setField(6, v); }
  @$pb.TagNumber(6)
  $core.bool hasFormClosed() => $_has(5);
  @$pb.TagNumber(6)
  void clearFormClosed() => clearField(6);
  @$pb.TagNumber(6)
  FormClosed ensureFormClosed() => $_ensure(5);
}

enum FormDataMsgSrv_Payload {
  rootListOpen, 
  rootListAddRows, 
  rootListChangeRows, 
  rootListRemoveRows, 
  contextMenuOpen, 
  formAction, 
  notSet
}

class FormDataMsgSrv extends $pb.GeneratedMessage {
  static const $core.Map<$core.int, FormDataMsgSrv_Payload> _FormDataMsgSrv_PayloadByTag = {
    2 : FormDataMsgSrv_Payload.rootListOpen,
    3 : FormDataMsgSrv_Payload.rootListAddRows,
    4 : FormDataMsgSrv_Payload.rootListChangeRows,
    5 : FormDataMsgSrv_Payload.rootListRemoveRows,
    6 : FormDataMsgSrv_Payload.contextMenuOpen,
    7 : FormDataMsgSrv_Payload.formAction,
    0 : FormDataMsgSrv_Payload.notSet
  };
  static final $pb.BuilderInfo _i = $pb.BuilderInfo(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'FormDataMsgSrv', package: const $pb.PackageName(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'runify'), createEmptyInstance: create)
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

  FormDataMsgSrv._() : super();
  factory FormDataMsgSrv({
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

