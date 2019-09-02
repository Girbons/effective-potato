import 'package:flutter/material.dart';

class LightsPage extends StatefulWidget {
  static String tag = 'lights-page';

  @override
  _LightsPageState createState() => new _LightsPageState();
}

class _LightsPageState extends State<LightsPage> {
  static String tag = 'lights-page';

  List<Map> lights = [
    { 'key': 1, 'name': 'Room 1', 'status': false,},
    { 'key': 2, 'name': 'Room 2', 'status': false},
    { 'key': 3, 'name': 'Room 3', 'status': false},
    { 'key': 4, 'name': 'Room 4', 'status': false},
  ];

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(
        title: Text("Home"),
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
                    textAlign: TextAlign.justify,),
                ),
                Expanded(
                  child: Container(
                     alignment: Alignment.centerRight,
                     child: Switch(
                    value: lights[index]['status'], 
                    onChanged: (bool value) {
                      setState(() {
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
      )
    );
  }
}