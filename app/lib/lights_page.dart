import 'dart:convert';

import 'package:app/add_light.dart';
import 'package:app/requests.dart';
import 'package:app/settings_page.dart';
import 'package:flutter/material.dart';

class Light {
  int id;
  String name;
  int pin;
  bool status;

  Light({this.id, this.name, this.pin, this.status});

  Light.fromJson(Map<String, dynamic> json) {
    id = json["ID"];
    name = json["name"];
    pin = json["pin"];
    status = json["status"];
  }

  Map<String, dynamic> toJson() {
    return {
      "id": id,
      "name": name,
      "pin": pin,
      "status": status
    };
  }
}

class LightsPage extends StatefulWidget {
  static String tag = 'lights-page';

  @override
  _LightsPageState createState() => new _LightsPageState();
}

class _LightsPageState extends State<LightsPage> {
  int _selectedIndex = 0;
  var lights = [];

  Future<List<Light>> _retrieveLights() async {
    var data = await fetchLights();
    List list = json.decode(data);
    var l = list.map((item) => Light.fromJson(item)).toList();
    return l;
  }

  void _onItemTapped(int index) {
    if (index == 1) {
      Navigator.of(context).pushNamed(SettingsPage.tag);
    }
  }

  @override
  void initState() {
    _retrieveLights().then((result) {
      setState(() {
        lights = result;
      });
    });
    super.initState();
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
                    lights[index].name, 
                    style: TextStyle(fontWeight: FontWeight.bold),
                    textAlign: TextAlign.justify),
                ),
                Expanded(
                  child: Container(
                     alignment: Alignment.centerRight,
                     child: Switch(
                    value: lights[index].status, 
                    onChanged: (bool value) {
                      setState(() {
                        if (value) {
                          turnOn(lights[index].pin);
                        } else {
                          turnOff(lights[index].pin);
                        }
                        lights[index].status = value;
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
      floatingActionButton: FloatingActionButton(
        onPressed: () => Navigator.of(context).pushNamed(LightAddPage.tag),
        tooltip: 'Add Item',
        child: Icon(Icons.add)
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