import 'package:runify/global/data_filter.dart';

class RootListRowID {
  final int providerID;
  final int rowID;

  RootListRowID(this.providerID, this.rowID);
}

class RootListRow implements Matcher<RootListRowID> {
  final RootListRowID id;
  final int priority;
  final String name;
  final String category;
  final String icon;

  RootListRow(this.id, this.priority, this.name, this.category, this.icon);

  @override
  bool match(RegExp rexp) {
    return rexp.hasMatch(name);
  }

  @override
  bool equal(Iterable<RootListRowID> keys) {
    return keys.any((RootListRowID key) => key == id);
  }
}

int rootListRowComparator(RootListRow a, RootListRow b) {
  if (a.priority != b.priority) {
    return a.priority.compareTo(b.priority);
  }

  return a.name.compareTo(b.name);
}

typedef RootListRowFilter = DataFilter<RootListRowID, RootListRow>;

abstract class RootListRpcClient {
  RootListRowFilter get filter;
  Future<void> execute(RootListRowID id);
  Future<void> menuActivate(RootListRowID id);
}
