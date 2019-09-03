import 'dart:convert';
import 'dart:io';
import 'package:http/http.dart' as http;
import 'package:app/storage.dart';


Future<String> getGatewayAddress() async {
  return await storage.read(key: 'gatewayIP');
}

Future<int> login(String username, password) async {
  var body = json.encode({
    "username": username,
    "password": password
  });

  var baseURL = await getGatewayAddress();

  var loginEndpoint = baseURL + "/login/";
  var response = await http.post(
    loginEndpoint, 
    headers: {"Content-Type": "application/json"},
    body: body
  );
  var parsedResponse = jsonDecode(response.body);
  print(response.request.url);
  storage.write(key: "token", value: parsedResponse["token"]);

  return response.statusCode;
}

void logout() {
  storage.deleteAll();
}

void turnOn(int pin) async {
  var baseURL = await getGatewayAddress();
  var token = await storage.read(key: "token");

  var turnOnEndpoint = baseURL + "/api/pin/on/$pin/";
  await http.get(turnOnEndpoint, headers: {
    HttpHeaders.authorizationHeader: "$token",
    "Content-Type": "application/json",
  });
}

Future<int> getUserProfile() async {
  var baseURL = await getGatewayAddress();
  var token = await storage.read(key: "token");

  var profileEndpoint = baseURL + "/user/profile/";
  var response = await http.get(profileEndpoint, headers: {
    HttpHeaders.authorizationHeader: "$token",
    "Content-Type": "application/json",
  });

  return response.statusCode;
}

void turnOff(int pin) async {
  var baseURL = await getGatewayAddress();
  var token = await storage.read(key: "token");
  
  var turnOffEndpoint = baseURL + "/api/pin/off/$pin/";
  await http.get(turnOffEndpoint, headers: {
    HttpHeaders.authorizationHeader: "$token",
    "Content-Type": "application/json",
  });
}
void saveGateway(String value) async {
  storage.write(key: "gatewayIP", value: value);
}