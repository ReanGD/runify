import 'package:runify/system/logger.dart';
import 'package:runify/screen/router.dart';
import 'package:runify/pb/runify.pbgrpc.dart';

class ScreenRouterService {
  final Logger logger;
  final RunifyClient grpcClient;

  ScreenRouterService(this.logger, this.grpcClient);

  Future<void> waitShowWindow(ScreenRouter router) async {
    try {
      final stream = grpcClient.waitShowWindow(Empty());
      await for (var _ in stream) {
        router.showWindow();
      }
    } catch (e) {
      // TODO: add policy for reconnect
      logger.log('gRPC WaitShowWindow method ended with error: $e');
    }
  }
}
