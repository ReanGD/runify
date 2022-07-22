import 'dart:io';

import 'package:flutter/material.dart';
import 'package:runify/view/run/run_screen.dart';

void main() {
  runApp(
    MaterialApp(
      title: "Runify",
      theme: ThemeData(
        brightness: Brightness.dark,
        primarySwatch: Colors.blue,
      ),
      home: const RunScreen(),
      debugShowCheckedModeBanner: false,
    ),
  );
}
