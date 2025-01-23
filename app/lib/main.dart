import "package:flutter/material.dart";
import "package:jawcon_app/App.dart";
import 'package:jawcon_app/theme.dart';

void main() {
  runApp(MaterialApp(
    home: App(),
    theme: AppTheme.colorscheme,
  ));
}
