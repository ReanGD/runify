import 'package:runify/model/data_filter.dart';

abstract class CommandAction implements Matcher {
  String get name;
  void execute();
}

typedef CommandActions = List<CommandAction>;
typedef CommandActionFilter = DataFilter<CommandAction>;

abstract class Command implements Matcher {
  String get name;
  String get category;
  String? get iconPath;
  CommandActions get actions;
}

typedef CommandFilter = DataFilter<Command>;
