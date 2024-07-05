import 'package:runify/global/data_filter.dart';

class RootListRowID {
  final String providerID;
  final String rowID;

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
  command, // filter enabled
  link, // filter enabled
  unknown, // filter enabled
}

class RootListRow implements Matcher<RootListRowID> {
  final RootListRowID id;
  final RootListRowType rowType;
  final int priority;
  final String displayName;
  final String searchNames;
  final String icon;

  RootListRow(
    this.id,
    this.rowType,
    this.priority,
    this.displayName,
    this.searchNames,
    this.icon,
  );

  @override
  RootListRowID get key => id;

  get typeName {
    switch (rowType) {
      case RootListRowType.calculator:
        return "";
      case RootListRowType.application:
        return 'App';
      case RootListRowType.command:
        return 'Command';
      case RootListRowType.link:
        return 'Link';
      default:
        return "";
    }
  }

  @override
  bool match(RegExp rexp) {
    if (rowType == RootListRowType.calculator) {
      return true;
    }

    return rexp.hasMatch(searchNames);
  }
}

int rootListRowComparator(RootListRow a, RootListRow b) {
  if (a.priority != b.priority) {
    return b.priority.compareTo(a.priority);
  }

  return a.displayName.compareTo(b.displayName);
}

typedef RootListRowFilter = DataFilter<RootListRowID, RootListRow>;
