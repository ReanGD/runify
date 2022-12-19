import 'package:runify/global/data_filter.dart';

class ContextMenuRow implements Matcher<int> {
  final int id;
  final String name;

  ContextMenuRow(this.id, this.name);

  @override
  int get key => id;

  @override
  bool match(RegExp rexp) {
    return rexp.hasMatch(name);
  }
}

int contextMenuRowComparator(ContextMenuRow a, ContextMenuRow b) {
  return a.name.compareTo(b.name);
}

typedef ContextMenuRowFilter = DataFilter<int, ContextMenuRow>;

abstract class ContextMenuRpcClient {
  int get formID;
  ContextMenuRowFilter get filter;

  void setFilter(String value);
  void execute(int id);
  void formClosed();
}
