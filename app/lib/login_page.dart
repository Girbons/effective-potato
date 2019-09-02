import 'package:app/lights_page.dart';
import 'package:app/requests.dart';
import 'package:app/storage.dart';
import 'package:flutter/material.dart';


class LoginPage extends StatefulWidget {
  static String tag = 'login-page';

  @override
  _LoginPageState createState() => new _LoginPageState();
}

class _LoginPageState extends State<LoginPage> {
  final _usernameController = TextEditingController();
  final _passwordController = TextEditingController();

  void checkToken() async {
  }

  void initState() {
    super.initState();
    storage.deleteAll();

    ()async {
      var token = await storage.read(key: 'token');
      if (token.length != 0) {
        Navigator.push(context, MaterialPageRoute(builder: (context) => LightsPage()));
      }
    }();
  }

  void dispose() {
    _usernameController.dispose();
    _passwordController.dispose();
  }

  bool _buttonDisabled() {
    if (_usernameController.text == '') {
      return true;
    }

    if (_passwordController.text == ''){
       return true;
    }

    return false;
  }

  @override
  Widget build(BuildContext context) {
    final textLogo = Text(
      "Potato System",
      textAlign: TextAlign.center,
    );

    final username = TextFormField(
      controller: _usernameController,
      keyboardType: TextInputType.text,
      autofocus: false,
      decoration: InputDecoration(
        hintText: 'username',
        contentPadding: EdgeInsets.fromLTRB(20.0, 10.0, 20.0, 10.0),
        border: OutlineInputBorder(borderRadius: BorderRadius.circular(32.0)),
      )
    );

    final password = TextFormField(
      controller: _passwordController,
      autofocus: false,
      obscureText: true,
      decoration: InputDecoration(
        hintText: 'password',
        contentPadding: EdgeInsets.fromLTRB(20.0, 10.0, 20.0, 10.0),
        border: OutlineInputBorder(borderRadius: BorderRadius.circular(32.0)),
      )
    );

    final loginButton = Padding(
      padding: EdgeInsets.symmetric(vertical: 16.0),
      child: RaisedButton(
        shape: RoundedRectangleBorder(
          borderRadius: BorderRadius.circular(24),
        ),
        onPressed: () {
          if (_buttonDisabled()) {
            return null;
          }
          Future<int> result = login(_usernameController.text, _passwordController.text);
          result.then((value) => {
            if (value == 200) {
              Navigator.of(context).pushNamed(LightsPage.tag)
            }
          })
          .catchError((error) => print(error));
        },
        padding: EdgeInsets.all(12),
        color: Colors.lightBlueAccent,
        child: Text('Log In', style: TextStyle(color: Colors.white)),
      ),
    );

    return Scaffold(
      backgroundColor: Colors.white,
      body: Center(
        child: ListView(
          shrinkWrap: true,
          padding: EdgeInsets.only(left: 24.0, right: 24.0),
          children: <Widget>[
            textLogo,
            SizedBox(height: 50.0),
            username,
            SizedBox(height: 8.0),
            password,
            SizedBox(height: 24.0),
            loginButton,
          ],
        ),
      ),
    );
  }
}
