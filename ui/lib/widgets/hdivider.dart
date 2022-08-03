import 'package:flutter/material.dart';

class HDivider extends StatelessWidget {
  final EdgeInsetsGeometry? padding;

  const HDivider({super.key, this.padding});

  @override
  Widget build(BuildContext context) {
    final Color color =
        DividerTheme.of(context).color ?? Theme.of(context).dividerColor;

    final child = DecoratedBox(
      decoration: BoxDecoration(
        border: Border(
          bottom: BorderSide(
            color: color,
          ),
        ),
      ),
      child: LimitedBox(
        maxWidth: 0.0,
        maxHeight: 0.0,
        child: ConstrainedBox(constraints: const BoxConstraints.expand()),
      ),
    );

    if (padding == null) {
      return child;
    }

    return Padding(
      padding: padding!,
      child: child,
    );
  }
}
