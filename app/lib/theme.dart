import 'package:flutter/material.dart';
import 'package:flex_color_scheme/flex_color_scheme.dart';

class AppTheme {
  static ThemeData colorscheme = FlexColorScheme.dark(
    primary: Color(0xFF00BFFF),
    secondary: Color(0xFF0071C2)
  ).toTheme;
}
