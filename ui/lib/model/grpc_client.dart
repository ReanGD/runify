import 'dart:io';
import 'package:grpc/grpc.dart';
import 'package:runify/model/settings.dart';
import 'package:runify/pb/runify.pbgrpc.dart';

RunifyClient newGrpcClient(Settings settings) {
  final channel = ClientChannel(
    InternetAddress(
      settings.grpcAddress,
      type: InternetAddressType.unix,
    ),
    options: const ChannelOptions(
      credentials: ChannelCredentials.insecure(),
    ),
  );

  return RunifyClient(channel);
}
