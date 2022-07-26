import 'package:flutter/material.dart';
import 'package:runify/view/run/style.dart';

class SearchWidget extends StatefulWidget {
  final String text;
  final ValueChanged<String> onChanged;
  final String hintText;

  const SearchWidget({
    Key? key,
    required this.text,
    required this.onChanged,
    required this.hintText,
  }) : super(key: key);

  @override
  State<SearchWidget> createState() => _SearchWidgetState();
}

class _SearchWidgetState extends State<SearchWidget> {
  final controller = TextEditingController();
  late FocusNode focusNode;

  @override
  void initState() {
    super.initState();

    focusNode = FocusNode();
    focusNode.addListener(() {
      if (!focusNode.hasFocus) {
        focusNode.requestFocus();
      }
    });
  }

  @override
  void dispose() {
    focusNode.dispose();

    super.dispose();
  }

  @override
  Widget build(BuildContext context) {
    const styleActive = TextStyle(color: Colors.black);
    const styleHint = TextStyle(color: Colors.black87);
    final style = widget.text.isEmpty ? styleHint : styleActive;

    return TextField(
      controller: controller,
      focusNode: focusNode,
      style: styleNormalText,
      autofocus: true,
      decoration: InputDecoration(
        icon: Icon(Icons.search, color: style.color),
        suffixIcon: widget.text.isNotEmpty
            ? GestureDetector(
                child: Icon(Icons.close, color: style.color),
                onTap: () {
                  controller.clear();
                  widget.onChanged('');
                  FocusScope.of(context).requestFocus(FocusNode());
                },
              )
            : null,
        hintText: widget.hintText,
        border: InputBorder.none,
      ),
      onChanged: widget.onChanged,
    );
  }
}
