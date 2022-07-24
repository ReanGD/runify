import 'package:flutter/material.dart';
import 'package:runify/model/command.dart';
import 'package:runify/view/run/style.dart';

class CommandRow extends StatelessWidget {
  final Command data;

  const CommandRow({Key? key, required this.data}) : super(key: key);

  @override
  Widget build(BuildContext context) {
    return Padding(
      padding: cmdRowPadding,
      child: Row(
        mainAxisAlignment: MainAxisAlignment.spaceBetween,
        children: <Widget>[
          Text(
            data.name(),
            style: cmdRowNameStyle,
          ),
          Text(
            data.category(),
            style: cmdRowCategoryStyle,
          ),
        ],
      ),
    );
  }
}
