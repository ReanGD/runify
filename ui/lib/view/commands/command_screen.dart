import 'package:flutter/material.dart';
import 'package:provider/provider.dart';
import 'package:runify/model/cmd_filter.dart';
import 'package:runify/model/cmd_loader.dart';
import 'package:runify/widgets/data_list_view.dart';
import 'package:runify/view/commands/command_list.dart';
import 'package:runify/widgets/disable_focus_trap_behavior.dart';

class CommandListController extends DataListController {
  @override
  List<int> getVisibleItems(BuildContext context) {
    return context.watch<CommandFilter>().visibleItems;
  }
}

class CommandScreen extends StatelessWidget {
  final DataListController controller = CommandListController();

  CommandScreen({Key? key}) : super(key: key);

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      body: MultiProvider(
        providers: [
          ChangeNotifierProvider<CommandFilter>(
              create: (_) => CommandFilter(CommandLoader.instance.commands)),
        ],
        child: DisableFocusTrapBehavior(
          child: Shortcuts(
            shortcuts: controller.getShortcuts(),
            child: Actions(
              actions: controller.getActions(),
              child: CommandList(controller: controller),
            ),
          ),
        ),
      ),
    );
  }
}
