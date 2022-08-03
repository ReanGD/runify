import 'package:flex_color_scheme/flex_color_scheme.dart';
import 'package:flutter/material.dart';

const double defaultRadius = 9.0;

ThemeData getLightTheme() {
  return FlexThemeData.light(
    colors: const FlexSchemeColor(
      primary: Color(0xff065808),
      primaryContainer: Color(0xff9bbc9c),
      secondary: Color(0xff365b37),
      secondaryContainer: Color(0xffaebdaf),
      tertiary: Color(0xff2c7e2e),
      tertiaryContainer: Color(0xffb8e6b9),
      appBarColor: Color(0xffb8e6b9),
      error: Color(0xffb00020),
    ),
    surfaceMode: FlexSurfaceMode.highScaffoldLowSurface,
    blendLevel: 20,
    appBarOpacity: 0.95,
    subThemesData: const FlexSubThemesData(
      blendOnLevel: 20,
      blendOnColors: false,
      defaultRadius: defaultRadius,
    ),
    visualDensity: FlexColorScheme.comfortablePlatformDensity,
    useMaterial3: true,
    // To use the playground font, add GoogleFonts package and uncomment
    // fontFamily: GoogleFonts.notoSans().fontFamily,
  );
}

ThemeData getDarkTheme() {
  return FlexThemeData.dark(
    colors: const FlexSchemeColor(
      primary: Color(0xff629f80),
      primaryContainer: Color(0xff273f33),
      secondary: Color(0xff81b39a),
      secondaryContainer: Color(0xff4d6b5c),
      tertiary: Color(0xff88c5a6),
      tertiaryContainer: Color(0xff356c50),
      appBarColor: Color(0xff356c50),
      error: Color(0xffcf6679),
    ),
    surfaceMode: FlexSurfaceMode.highScaffoldLowSurface,
    blendLevel: 15,
    appBarStyle: FlexAppBarStyle.background,
    appBarOpacity: 0.90,
    subThemesData: const FlexSubThemesData(
      blendOnLevel: 30,
      defaultRadius: defaultRadius,
    ),
    visualDensity: FlexColorScheme.comfortablePlatformDensity,
    useMaterial3: true,
    // To use the playground font, add GoogleFonts package and uncomment
    // fontFamily: GoogleFonts.notoSans().fontFamily,
  );
}

extension RunifyDialogTheme on DialogTheme {
  double get verticalOffset => 10.0;
  double get horizontalOffset => 10.0;
}

extension RunifyTextTheme on TextTheme {
  TextStyle? get majorText => bodyMedium?.copyWith(
        fontWeight: FontWeight.w600,
        fontSize: 18,
        letterSpacing: -0.2,
        wordSpacing: 0,
      );

  TextStyle? get minorText => bodyMedium?.copyWith(
        fontWeight: FontWeight.w600,
        fontSize: 18,
        letterSpacing: -0.2,
        wordSpacing: 0,
        color: bodyMedium?.color?.withAlpha(130),
      );
}

extension RunifyCardTheme on CardTheme {
  EdgeInsets get commandPadding =>
      const EdgeInsets.symmetric(vertical: 10, horizontal: 10);
}
