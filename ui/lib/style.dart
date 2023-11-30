import 'package:flutter/material.dart';
import 'package:flex_color_scheme/flex_color_scheme.dart';

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
  double get actionsWidth => 350;
  double get actionsHeight => 300;
  double get verticalOffset => 10.0;
  double get horizontalOffset => 10.0;
}

enum TextStyleType {
  bodyRegular_100,
  bodyRegular_50,
  bodyLight_100,
  bodyLight_50,
  labelRegular_100,
  labelRegular_80,
  labelRegular_60,
  labelRegular_50,
}

extension RunifyTextTheme on TextTheme {
  static Map<TextStyleType, TextStyle?>? _textStyles;

  static TextStyle? _makeTextStyle(
      TextTheme theme, double fontSize, FontWeight fontWeight, double opacity) {
    return theme.bodyMedium?.copyWith(
      fontSize: fontSize,
      fontWeight: fontWeight,
      fontStyle: FontStyle.normal,
      letterSpacing: -0.1,
      wordSpacing: 0,
      color: theme.bodyMedium?.color?.withOpacity(opacity),
    );
  }

  static Map<TextStyleType, TextStyle?> _makeTextStyles(TextTheme theme) {
    // fonst size
    const bodyFS = 18.0;
    const labelFS = 16.0;
    // font weight
    const bodyRegularFW = FontWeight.w600;
    const bodyLightFW = FontWeight.w500;
    const labelRegularFW = FontWeight.w600;

    return <TextStyleType, TextStyle?>{
      TextStyleType.bodyRegular_100:
          _makeTextStyle(theme, bodyFS, bodyRegularFW, 1.0),
      TextStyleType.bodyRegular_50:
          _makeTextStyle(theme, bodyFS, bodyRegularFW, 0.5),
      TextStyleType.bodyLight_100:
          _makeTextStyle(theme, bodyFS, bodyLightFW, 1.0),
      TextStyleType.bodyLight_50:
          _makeTextStyle(theme, bodyFS, bodyLightFW, 0.5),
      TextStyleType.labelRegular_100:
          _makeTextStyle(theme, labelFS, labelRegularFW, 1.0),
      TextStyleType.labelRegular_80:
          _makeTextStyle(theme, labelFS, labelRegularFW, 0.8),
      TextStyleType.labelRegular_60:
          _makeTextStyle(theme, labelFS, labelRegularFW, 0.6),
      TextStyleType.labelRegular_50:
          _makeTextStyle(theme, labelFS, labelRegularFW, 0.5),
    };
  }

  static TextStyle? _getTextStyle(TextTheme theme, TextStyleType type) {
    _textStyles ??= _makeTextStyles(theme);
    return _textStyles?[type];
  }

  TextStyle? get bodyRegular =>
      _getTextStyle(this, TextStyleType.bodyRegular_100);
  TextStyle? get bodyRegularActive =>
      _getTextStyle(this, TextStyleType.bodyRegular_100);
  TextStyle? get bodyRegular_100 =>
      _getTextStyle(this, TextStyleType.bodyRegular_100);
  TextStyle? get bodyRegularInactive =>
      _getTextStyle(this, TextStyleType.bodyRegular_50);
  TextStyle? get bodyRegular_50 =>
      _getTextStyle(this, TextStyleType.bodyRegular_50);

  TextStyle? get bodyLight => _getTextStyle(this, TextStyleType.bodyLight_100);
  TextStyle? get bodyLightActive =>
      _getTextStyle(this, TextStyleType.bodyLight_100);
  TextStyle? get bodyLight_100 =>
      _getTextStyle(this, TextStyleType.bodyLight_100);
  TextStyle? get bodyLightInactive =>
      _getTextStyle(this, TextStyleType.bodyLight_50);
  TextStyle? get bodyLight_50 =>
      _getTextStyle(this, TextStyleType.bodyLight_50);

  TextStyle? get labelRegular =>
      _getTextStyle(this, TextStyleType.labelRegular_100);
  TextStyle? get labelRegular_100 =>
      _getTextStyle(this, TextStyleType.labelRegular_100);
  TextStyle? get labelRegular_80 =>
      _getTextStyle(this, TextStyleType.labelRegular_80);
  TextStyle? get labelRegular_60 =>
      _getTextStyle(this, TextStyleType.labelRegular_60);
  TextStyle? get labelRegular_50 =>
      _getTextStyle(this, TextStyleType.labelRegular_50);
}

extension RunifyCardTheme on CardTheme {
  EdgeInsets get commandPadding =>
      const EdgeInsets.symmetric(vertical: 10, horizontal: 10);
}
