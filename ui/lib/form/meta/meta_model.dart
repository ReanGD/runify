import 'package:runify/model/command.dart';
import 'package:runify/model/data_filter.dart';

class MetaModel {
  final DataFilter<Command> filter;

  MetaModel(Future<List<Command>> records)
      : filter = DataFilter.future(records);

  List<int> get visibleItems => filter.visibleItems;

  Command getListViewItem(int id) => filter[id];
}

class MetaMenuModel {
  final DataFilter<CommandAction> filter;

  MetaMenuModel(Future<List<CommandAction>> records)
      : filter = DataFilter.future(records);

  List<int> get visibleItems => filter.visibleItems;

  CommandAction getListViewItem(int id) => filter[id];
}
