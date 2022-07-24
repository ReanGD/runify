import 'package:flutter/material.dart';
import 'package:provider/provider.dart';
import 'package:runify/view/run/select_index.dart';
import 'package:runify/view/run/style.dart';

class IndexedRow extends StatelessWidget {
  final int index;
  final Widget child;

  const IndexedRow({Key? key, required this.index, required this.child})
      : super(key: key);

  @override
  Widget build(BuildContext context) {
    final select = context.watch<SelectIndex>();
    if (select.index == index) {
      return ColoredBox(
        color: cmdRowSelectedColor,
        child: child,
      );
    }

    return child;
  }
}
