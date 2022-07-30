import 'package:flutter/material.dart';

// Prevent window from giving focus
// see: https://github.com/flutter/flutter/issues/86972
class DisableFocusTrapBehavior extends StatefulWidget {
  final Widget child;

  const DisableFocusTrapBehavior({Key? key, required this.child})
      : super(key: key);

  @override
  State<StatefulWidget> createState() => _DisableFocusTrapBehaviorState();
}

class _DisableFocusTrapBehaviorState extends State<DisableFocusTrapBehavior> {
  late FocusNode focusNode;

  @override
  void initState() {
    super.initState();
    focusNode =
        FocusManager.instance.primaryFocus ?? FocusManager.instance.rootScope;
    FocusManager.instance.addListener(_onFocusNodeChanged);
  }

  @override
  void dispose() {
    super.dispose();
    FocusManager.instance.removeListener(_onFocusNodeChanged);
  }

  @override
  Widget build(BuildContext context) {
    return FocusTrapArea(focusNode: focusNode, child: widget.child);
  }

  void _onFocusNodeChanged() {
    final newFocusNode =
        FocusManager.instance.primaryFocus ?? FocusManager.instance.rootScope;
    if (newFocusNode != focusNode) {
      setState(() {
        focusNode = newFocusNode;
      });
    }
  }
}
