import 'package:app/add_light.dart';
import 'package:app/lights_page.dart';
import 'package:app/settings_page.dart';
import 'package:flutter/material.dart';
import 'login_page.dart';

void main() => runApp(MyApp());

class MyApp extends StatelessWidget {

  final routes = <String, WidgetBuilder>{
    LoginPage.tag: (context) => LoginPage(),
    LightsPage.tag: (context) => LightsPage(),
    LightAddPage.tag: (context) => LightAddPage(),
    SettingsPage.tag: (context) => SettingsPage(),
  };

  @override
  Widget build(BuildContext context) {
    return MaterialApp(
      title: 'Effective Potato',
      debugShowCheckedModeBanner: false,
      theme: ThemeData(
        primarySwatch: Colors.lightBlue,
        fontFamily: 'Nunito',
      ),
      home: LoginPage(),
      routes: routes,
    );
  }
}