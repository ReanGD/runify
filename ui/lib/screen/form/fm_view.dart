import 'package:flutter/material.dart';

import 'package:reactive_forms/reactive_forms.dart';

import 'package:runify/style.dart';
import 'package:runify/screen/form/fm_controller.dart';

class _MarkupParser {
  var _isFirst = true;

  bool get autofocus {
    if (_isFirst) {
      _isFirst = false;
      return true;
    }
    return false;
  }

  Text parseTextWidget(Map<String, dynamic> markup) {
    return Text(markup["data"] as String);
  }

  ReactiveTextField parseTextFieldWidget(Map<String, dynamic> markup) {
    return ReactiveTextField(
      autofocus: autofocus,
      formControlName: markup["bind"] as String,
      obscureText: markup["obscureText"] ?? false,
      readOnly: markup["readOnly"] ?? false,
      maxLines: markup["maxLines"] ?? 1,
      textInputAction: TextInputAction.next,
    );
  }

  Column parseColumnWidget(Map<String, dynamic> markup) {
    final children = markup["children"] as List;
    return Column(
      crossAxisAlignment: CrossAxisAlignment.start,
      children: children.map((e) => parse(e)).toList(),
    );
  }

  Widget parse(Map<String, dynamic> markup) {
    final widgetType = markup["type"] as String;
    switch (widgetType) {
      case "Column":
        return parseColumnWidget(markup);
      case "Text":
        return parseTextWidget(markup);
      case "TextField":
        return parseTextFieldWidget(markup);
    }

    return Text("Unknown widget type: $widgetType");
  }

  static Widget build(Map<String, dynamic> markup) {
    return _MarkupParser().parse(markup);
  }
}

class FMView extends StatelessWidget {
  final FMController controller;

  const FMView(this.controller, {super.key});

  @override
  Widget build(BuildContext context) {
    final theme = Theme.of(context);
    final dialogTheme = theme.dialogTheme;

    final windowPadding = EdgeInsets.symmetric(
        horizontal: dialogTheme.horizontalOffset,
        vertical: dialogTheme.verticalOffset);

    return ReactiveForm(
      formGroup: controller.model,
      child: Padding(
        padding: windowPadding,
        child: _MarkupParser.build(controller.markup),
      ),
    );
  }
}
