import 'package:runify/global/data_filter.dart';

class ContextMenuRow implements Matcher<int> {
  final int id;
  final String value;

  ContextMenuRow(this.id, this.value);

  @override
  int get key => id;

  @override
  bool match(RegExp rexp) {
    return rexp.hasMatch(value);
  }
}

int contextMenuRowComparator(ContextMenuRow a, ContextMenuRow b) {
  return a.value.compareTo(b.value);
}

typedef ContextMenuRowFilter = DataFilter<int, ContextMenuRow>;
