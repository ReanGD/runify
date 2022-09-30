import 'package:runify/system/data_filter.dart';

class CommandAction implements Matcher {
  int id;
  String name;

  CommandAction(this.id, this.name);

  @override
  bool match(RegExp rexp) {
    return rexp.hasMatch(name);
  }
}

typedef CommandActions = List<CommandAction>;
typedef CommandActionFilter = DataFilter<CommandAction>;
