import 'dart:async';

import 'package:runify/system/logger.dart';
import 'package:runify/screen/router.dart';
import 'package:runify/pb/runify.pbgrpc.dart';

class ScreenRouterService {
  final RunifyClient grpcClient;
  final ScreenRouter router;
  final _sender = StreamController<ServiceMsgUI>();
  late final Logger logger;

  ScreenRouterService(this.grpcClient, this.router) {
    logger = Logger(_sender);
  }

  Future<void> serviceChannel(ScreenRouter router) async {
    try {
      final stream = grpcClient.serviceChannel(_sender.stream);
      await for (var _ in stream) {
        router.showWindow();
      }
    } catch (e) {
      // ignore: avoid_print
      print("gRPC WaitShowWindow method ended with error: $e. Stop runify.");
      await router.closeWindow();
    }
  }
}
