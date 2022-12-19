import 'package:runify/global/data_filter.dart';

class RootListRowID {
  final int providerID;
  final int rowID;

  RootListRowID(this.providerID, this.rowID);

  @override
  bool operator ==(other) =>
      other is RootListRowID &&
      other.providerID == providerID &&
      other.rowID == rowID;

  @override
  int get hashCode => Object.hash(providerID, rowID);
}

class RootListRow implements Matcher<RootListRowID> {
  final RootListRowID id;
  final int priority;
  final String name;
  final String category;
  final String icon;

  RootListRow(this.id, this.priority, this.name, this.category, this.icon);

  @override
  RootListRowID get key => id;

  @override
  bool match(RegExp rexp) {
    return rexp.hasMatch(name);
  }
}

int rootListRowComparator(RootListRow a, RootListRow b) {
  if (a.priority != b.priority) {
    return b.priority.compareTo(a.priority);
  }

  return a.name.compareTo(b.name);
}

typedef RootListRowFilter = DataFilter<RootListRowID, RootListRow>;

abstract class RootListRpcClient {
  int get formID;
  RootListRowFilter get filter;

  void setFilter(String value);
  void execute(RootListRowID id);
  void menuActivate(RootListRowID id);
  void formClosed();
}
