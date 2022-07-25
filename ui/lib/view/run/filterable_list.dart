import 'dart:math';

import 'package:flutter/material.dart';
import 'package:provider/provider.dart';
import 'package:runify/model/cmd_storage.dart';
import 'package:runify/view/run/command_row.dart';
import 'package:runify/view/run/indexed_row.dart';
import 'package:scrollable_positioned_list/scrollable_positioned_list.dart';

class CustomSimulation extends Simulation {
  final double initPosition;
  final double velocity;

  CustomSimulation({required this.initPosition, required this.velocity});

  @override
  double x(double time) {
    return max(min(initPosition, 0.0), initPosition + velocity * time);
  }

  @override
  double dx(double time) {
    return velocity;
  }

  @override
  bool isDone(double time) {
    return false;
  }
}

class CustomScrollPhysics extends ScrollPhysics {
  @override
  ScrollPhysics applyTo(ScrollPhysics? ancestor) {
    return CustomScrollPhysics();
  }

  @override
  Simulation createBallisticSimulation(
      ScrollMetrics position, double velocity) {
    return CustomSimulation(initPosition: position.pixels, velocity: velocity);
  }
}

class FilterableList extends StatelessWidget {
  final ItemScrollController itemScrollController;
  final ItemPositionsListener itemPositionsListener;

  const FilterableList(
      {Key? key,
      required this.itemScrollController,
      required this.itemPositionsListener})
      : super(key: key);

  @override
  Widget build(BuildContext context) {
    final items = context.watch<CommandStorage>();
    // final selectIndex = context.read<SelectIndex>();
    // final index = max(min(selectIndex.index, items.length - 1), 0);
    // selectIndex.update(index);

    final w = ScrollablePositionedList.builder(
      physics: CustomScrollPhysics(),
      itemCount: items.length,
      itemBuilder: (context, index) {
        return IndexedRow(index: index, child: CommandRow(data: items[index]));
      },
      itemScrollController: itemScrollController,
      itemPositionsListener: itemPositionsListener,
      initialAlignment: 0.0,
    );

    return w;
  }
}
