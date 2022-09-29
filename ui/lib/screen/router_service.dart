import 'package:runify/system/logger.dart';
import 'package:runify/screen/router.dart';
import 'package:runify/pb/runify.pbgrpc.dart';

class ScreenRouterService {
  final Logger logger;
  final RunifyClient grpcClient;
  final ScreenRouter router;

  ScreenRouterService(this.logger, this.grpcClient, this.router);

  Future<void> waitShowWindow(ScreenRouter router) async {
    try {
      final stream = grpcClient.waitShowWindow(Empty());
      await for (var _ in stream) {
        router.showWindow();
      }
    } catch (e) {
      logger
          .log('gRPC WaitShowWindow method ended with error: $e. Stop runify.');
      await router.closeWindow();
    }
  }
}
