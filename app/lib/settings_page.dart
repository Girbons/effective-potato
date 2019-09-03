import 'package:app/lights_page.dart';
import 'package:flutter/material.dart';
import 'package:app/requests.dart';

class SettingsPage extends StatefulWidget {
  static String tag = 'settings-page';

  @override
  _SettingsPageState createState() => new _SettingsPageState();
}

class _SettingsPageState extends State<SettingsPage> {
  final _gatewatyController = TextEditingController();
  int _selectedIndex = 1;

  void dispose() {
    _gatewatyController.dispose();
    super.dispose();
  }

  void setInitialValue() async {
    var baseURL = await getGatewayAddress();
    _gatewatyController.text = baseURL;
  }

  void initState() {
    setInitialValue();

    super.initState();
  }

  void _onItemTapped(int index) {
    if (index == 0) {
      Navigator.of(context).pop(LightsPage.tag);
    }

  }

  @override
  Widget build(BuildContext context) {
    final gatewayForm = TextFormField(
      controller: _gatewatyController,
      keyboardType: TextInputType.text,
      autofocus: false,
      decoration: InputDecoration(
        hintText: "http://192.168.1.10",
        contentPadding: EdgeInsets.fromLTRB(20.0, 10.0, 20.0, 10.0),
        border: OutlineInputBorder(borderRadius: BorderRadius.circular(32.0)),
      )
    );

    final saveButton = Padding(
      padding: EdgeInsets.symmetric(vertical: 16.0),
      child: RaisedButton(
        onPressed: () {
          if (_gatewatyController.text.length != 0) {
            saveGateway(_gatewatyController.text);
          }

          return null;
        },
        shape: RoundedRectangleBorder(
          borderRadius: BorderRadius.circular(24),
        ),
        padding: EdgeInsets.all(12),
        color: Colors.lightBlueAccent,
        child: Text('Save', style: TextStyle(color: Colors.white)),
      ),
    );

    final logoutButton = Padding(
      padding: EdgeInsets.symmetric(vertical: 16.0),
      child: RaisedButton(
        onPressed: () => logout(),
        shape: RoundedRectangleBorder(
          borderRadius: BorderRadius.circular(24),
        ),
        padding: EdgeInsets.all(12),
        color: Colors.lightBlueAccent,
        child: Text('Logout', style: TextStyle(color: Colors.white)),
      ),
    );

    return Scaffold(
      appBar: AppBar(
        title: Text("Settings"),
        backgroundColor: Colors.white,
        automaticallyImplyLeading: false,
      ),
      body: Center(
        child: ListView (
          shrinkWrap: true,
          padding: EdgeInsets.only(left: 24.0, right: 24.0),
          children: <Widget> [
            gatewayForm,
            SizedBox(height: 8.0),
            saveButton,
            SizedBox(height: 30.0),
            logoutButton,
          ]
        )
      ),
      bottomNavigationBar: BottomNavigationBar(
        items: const <BottomNavigationBarItem>[
          BottomNavigationBarItem(
            icon: Icon(Icons.home),
            title: Text('Home'),
          ),
          BottomNavigationBarItem(
            icon: Icon(Icons.settings),
            title: Text('Settings'),
          ),
        ],
        currentIndex: _selectedIndex,
        onTap: _onItemTapped,
      ),
    );
  }
}