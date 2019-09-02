import 'dart:convert';
import 'package:http/http.dart' as http;
import 'package:app/storage.dart';

var baseURL = 'http://192.168.247.122:8080';
var loginEndpoint = baseURL + '/login/';
var logoutEndpoint = baseURL + '/logout/';

Future<int> login(String username, password) async {
  var body = json.encode({
    "username": username,
    "password": password
  });

  var response = await http.post(
    loginEndpoint, 
    headers: {"Content-Type": "application/json"},
    body: body
  );
  var parsedResponse = jsonDecode(response.body);
  storage.write(key: 'token', value: parsedResponse['token']);

  return response.statusCode;
}

void turnOn(int pin) async {
  var turnOnEndpoint = baseURL + '/api/pin/on/{$pin}';
  await http.get(turnOnEndpoint);
}

void turnOff(int pin) async {
  var turnOffEndpoint = baseURL + '/api/pin/off/${pin}';
  await http.get(turnOffEndpoint);
}