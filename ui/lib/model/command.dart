import 'dart:ffi';

import 'package:fixnum/fixnum.dart';
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

class Command implements Matcher {
  Int64 id;
  String name;
  String category;
  String icon;

  Command(this.id, this.name, this.category, this.icon);

  @override
  bool match(String filter) {
    return name.toLowerCase().contains(filter);
  }
}

typedef CommandFilter = DataFilter<Command>;
