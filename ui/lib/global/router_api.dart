import 'package:flutter/material.dart';

abstract class Controller {
  int get formID;

  Widget build();
  void onFormClosed();
}
