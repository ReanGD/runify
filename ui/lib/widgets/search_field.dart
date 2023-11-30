import 'package:runify/style.dart';
import 'package:flutter/material.dart';

class SearchField extends StatelessWidget {
  final String hintText;
  final EdgeInsetsGeometry? padding;
  final ValueChanged<String> onChanged;

  const SearchField({
    super.key,
    required this.hintText,
    this.padding,
    required this.onChanged,
  });

  @override
  Widget build(BuildContext context) {
    final theme = Theme.of(context);
    const border = UnderlineInputBorder(
      borderSide: BorderSide.none,
    );

    final child = TextField(
      autofocus: true,
      style: theme.textTheme.bodyRegular,
      textAlignVertical: TextAlignVertical.center,
      decoration: InputDecoration(
        hintText: hintText,
        hintStyle: theme.textTheme.bodyLightInactive,
        filled: false,
        isDense: true,
        errorBorder: border,
        focusedBorder: border,
        focusedErrorBorder: border,
        disabledBorder: border,
        enabledBorder: border,
        border: border,
        contentPadding: const EdgeInsets.all(0),
      ),
      onChanged: onChanged,
    );

    if (padding != null) {
      return Padding(
        padding: padding!,
        child: child,
      );
    }

    return child;
  }
}
