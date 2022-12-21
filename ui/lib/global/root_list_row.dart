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

enum RootListRowType {
  calculator, // filter disabled
  application, // filter enabled
  unknown, // filter enabled
}

class RootListRow implements Matcher<RootListRowID> {
  final RootListRowID id;
  final RootListRowType rowType;
  final int priority;
  final String value;
  final String icon;

  RootListRow(this.id, this.rowType, this.priority, this.value, this.icon);

  @override
  RootListRowID get key => id;

  get typeName {
    switch (rowType) {
      case RootListRowType.calculator:
        return "";
      case RootListRowType.application:
        return 'App';
      default:
        return "";
    }
  }

  @override
  bool match(RegExp rexp) {
    if (rowType == RootListRowType.calculator) {
      return true;
    }

    return rexp.hasMatch(value);
  }
}

int rootListRowComparator(RootListRow a, RootListRow b) {
  if (a.priority != b.priority) {
    return b.priority.compareTo(a.priority);
  }

  return a.value.compareTo(b.value);
}

typedef RootListRowFilter = DataFilter<RootListRowID, RootListRow>;
