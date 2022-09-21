import 'package:runify/model/data_filter.dart';

class CommandAction implements Matcher {
  int id;
  String name;

  CommandAction(this.id, this.name);

  @override
  bool match(String filter) {
    return name.toLowerCase().contains(filter);
  }
}

typedef CommandActions = List<CommandAction>;
typedef CommandActionFilter = DataFilter<CommandAction>;
