class Metrics {
  final bool enabled;

  Metrics(this.enabled);

  grpcCall(int dtMcs, String name) {
    if (enabled) {
      // TODO: send to server
      // ignore: avoid_print
      print("grpc method $name = $dtMcs mcs");
    }
  }
}
