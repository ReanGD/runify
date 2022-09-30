import 'package:fixnum/fixnum.dart';
import 'package:runify/system/data_filter.dart';

class Command implements Matcher {
  Int64 id;
  String name;
  String category;
  String icon;

  Command(this.id, this.name, this.category, this.icon);

  @override
  bool match(RegExp rexp) {
    return rexp.hasMatch(name);
  }
}

typedef CommandFilter = DataFilter<Command>;
