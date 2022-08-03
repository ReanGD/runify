import 'package:flutter/material.dart';

class TextSizeCalculator {
  final Map<String, double> _cache = {};

  TextSizeCalculator._();

  static final TextSizeCalculator instance = TextSizeCalculator._();

  Size getSize(BuildContext context, String text, TextStyle? style) {
    return (TextPainter(
            text: TextSpan(text: text, style: style),
            maxLines: 1,
            textScaleFactor: MediaQuery.of(context).textScaleFactor,
            textDirection: TextDirection.ltr)
          ..layout())
        .size;
  }

  double getCachedHeight(
      BuildContext context, String cacheKey, TextStyle? style) {
    final result = _cache[cacheKey];
    if (result != null) {
      return result;
    }

    final height = getSize(context, "AaZzАаЯя", style).height;
    _cache[cacheKey] = height;
    return height;
  }
}
