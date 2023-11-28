import 'package:runify/global/data_filter.dart';

class ContextMenuRow implements Matcher<int> {
  final int id;
  final String displayName;
  final String searchNames;

  ContextMenuRow(this.id, this.displayName, this.searchNames);

  @override
  int get key => id;

  @override
  bool match(RegExp rexp) {
    return rexp.hasMatch(searchNames);
  }
}

int contextMenuRowComparator(ContextMenuRow a, ContextMenuRow b) {
  return a.displayName.compareTo(b.displayName);
}

typedef ContextMenuRowFilter = DataFilter<int, ContextMenuRow>;
