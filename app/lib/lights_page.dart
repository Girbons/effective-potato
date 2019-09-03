import 'package:app/requests.dart';
import 'package:app/settings_page.dart';
import 'package:flutter/material.dart';

class LightsPage extends StatefulWidget {
  static String tag = 'lights-page';

  @override
  _LightsPageState createState() => new _LightsPageState();
}

class _LightsPageState extends State<LightsPage> {
  static String tag = 'lights-page';
  int _selectedIndex = 0;

  List<Map> lights = [
    { 'key': 1, 'name': 'Room 1', 'status': false, 'pin': 10},
    { 'key': 2, 'name': 'Room 2', 'status': false, 'pin': 14},
    { 'key': 3, 'name': 'Room 3', 'status': false, 'pin': 12},
  ];

  void _onItemTapped(int index) {
    if (index == 1) {
      Navigator.of(context).pushNamed(SettingsPage.tag);
    }
  }

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(
        title: Text("Lights"),
        backgroundColor: Colors.white,
        automaticallyImplyLeading: false,
      ),
      body: ListView.builder(
        padding: const EdgeInsets.all(8.0),
        itemCount: lights.length,
        physics: const NeverScrollableScrollPhysics(),
        itemBuilder: (BuildContext context, int index) {
          return Container(
            height: 50,
            child: Row(
              children: <Widget>[
                Expanded(
                  child: Text(
                    lights[index]['name'], 
                    style: TextStyle(fontWeight: FontWeight.bold),
                    textAlign: TextAlign.justify),
                ),
                Expanded(
                  child: Container(
                     alignment: Alignment.centerRight,
                     child: Switch(
                    value: lights[index]['status'], 
                    onChanged: (bool value) {
                      setState(() {
                        if (value) {
                          turnOn(lights[index]['pin']);
                        } else {
                          turnOff(lights[index]['pin']);
                        }

                        lights[index]['status'] = value;
                      });
                    }
                  ),
                 )
                ),
              ],
            )
          );
        }
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