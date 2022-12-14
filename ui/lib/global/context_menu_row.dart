import 'package:runify/global/data_filter.dart';

class ContextMenuRow implements Matcher<int> {
  final int id;
  final String name;

  ContextMenuRow(this.id, this.name);

  @override
  bool match(RegExp rexp) {
    return rexp.hasMatch(name);
  }

  @override
  bool equal(Iterable<int> keys) {
    return keys.any((int key) => key == id);
  }
}

int contextMenuRowComparator(ContextMenuRow a, ContextMenuRow b) {
  return a.name.compareTo(b.name);
}

typedef ContextMenuRowFilter = DataFilter<int, ContextMenuRow>;

abstract class ContextMenuRpcClient {
  ContextMenuRowFilter get filter;
  Future<void> execute(int id);
}
