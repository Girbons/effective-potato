import 'package:app/lights_page.dart';
import 'package:app/requests.dart';
import 'package:flutter/material.dart';

class LightAddPage extends StatefulWidget {
  static String tag = 'add-light-page';

  @override
  _LightAddPageState createState() => new _LightAddPageState();
}

class _LightAddPageState extends State<LightAddPage> {
  final _nameController = TextEditingController();
  final _pinController = TextEditingController();

  @override
  void dispose() {
    _nameController.dispose();
    _pinController.dispose();

    super.dispose();
  }

  void _add() {
    if (_nameController.text.length != 0 && _pinController.text.length != 0){
      addLight(_nameController.text, _pinController.text);
      Navigator.push(context, MaterialPageRoute(builder: (context) => LightsPage()));
    }
  }

  @override
  Widget build(BuildContext context) {
    final name = TextFormField(
      controller: _nameController,
      keyboardType: TextInputType.text,
      autofocus: false,
      decoration: InputDecoration(
        hintText: "Room 1",
        contentPadding: EdgeInsets.fromLTRB(20.0, 10.0, 20.0, 10.0),
        border: OutlineInputBorder(borderRadius: BorderRadius.circular(32.0)),
      )
    );

    final pin = TextFormField(
      controller: _pinController,
      keyboardType: TextInputType.number,
      autofocus: false,
      decoration: InputDecoration(
        hintText: "10",
        contentPadding: EdgeInsets.fromLTRB(20.0, 10.0, 20.0, 10.0),
        border: OutlineInputBorder(borderRadius: BorderRadius.circular(32.0)),
      )
    );

    final saveButton = Padding(
      padding: EdgeInsets.symmetric(vertical: 16.0),
      child: RaisedButton(
        shape: RoundedRectangleBorder(
          borderRadius: BorderRadius.circular(24),
        ),
        onPressed: _add,
        padding: EdgeInsets.all(12),
        color: Colors.lightBlueAccent,
        child: Text("Save", style: TextStyle(color: Colors.white)),
      ),
    );

    return Scaffold(
      appBar: AppBar(
        title: Text("Add Light"),
        backgroundColor: Colors.white,
      ),
      backgroundColor: Colors.white,
      body: Center(
        child: ListView(
          shrinkWrap: true,
          padding: EdgeInsets.only(left: 24.0, right: 24.0),
          children: <Widget>[
            name,
            SizedBox(height: 8.0),
            pin,
            SizedBox(height: 24.0),
            saveButton,
          ],
        ),
      ),
    );
  }
}