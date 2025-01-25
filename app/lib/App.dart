import 'package:flame/flame.dart';
import 'package:flutter/material.dart';
import 'package:flutter/services.dart';
import 'package:gamepads/gamepads.dart';
import 'dart:io';
import 'dart:convert';

class App extends StatefulWidget {
  
  String name = "";
  String socket_state = "no connection";

  @override
  createState() => _App();
}

class _App extends State<App> {

  final String host = "192.168.1.13";
  final int port = 8890;

  @override
  void initState() {
    super.initState();
    initMe();
  }

  void initMe() async {
    
    // `````````````Gamepads.events.listen((event) {
    //   RawDatagramSocket.bind(InternetAddress.anyIPv4, port).then((socket){
    //     socket.send([event.key,event.value, event.type == KeyType.button ? 1 : 0, "endl"].join(",").codeUnits, InternetAddress(host), port);
    //     socket.close();
    //   });
    //   setState(() {
    //     widget.name = event.toString();
    //   });
    // });`````````````
    var previousKey = {};
    RawDatagramSocket.bind(InternetAddress.anyIPv4, port).then((socket){
      // socket.send([event.key,event.value, event.type == KeyType.button ? 1 : 0, "endl"].join(",").codeUnits, InternetAddress(host), port);
      Gamepads.events.listen((event){
        if(previousKey.containsKey(event.key)){
          if(event.type == KeyType.button){
            if(previousKey[event.key] == event.value){
              previousKey[event.key] = event.value;
              return;
            }
            socket.send([event.key,event.value, event.type == KeyType.button ? 1 : 0, "endl"].join(",").codeUnits, InternetAddress(host), port);
            previousKey[event.key] = event.value;
          }else{
            socket.send([event.key,event.value, event.type == KeyType.button ? 1 : 0, "endl"].join(",").codeUnits, InternetAddress(host), port);
          }
        }else{
          socket.send([event.key,event.value, event.type == KeyType.button ? 1 : 0, "endl"].join(",").codeUnits, InternetAddress(host), port);
          previousKey[event.key] = event.value;
        }
      });
    }).catchError((err){
      setState(() {
        widget.socket_state = "Socket Error: $err"; 
      });
    });
  }

  Future<List<String>> getConnected() async {
    List<String> name = [];
    final gamepad = await Gamepads.list();
    print(gamepad);
    gamepad.forEach((controller){
      name.add(controller.name.split(",")[0]);
    });
   return name; 
  }

  @override
  Widget build(BuildContext context) {
    SystemChrome.setEnabledSystemUIMode(SystemUiMode.manual, overlays: [
      SystemUiOverlay.top,
      SystemUiOverlay.bottom,
    ]);
    return Scaffold(
      // appBar: AppBar(
      //   title: Text("JawCon Client"),
      // ),
      // body: SafeArea(
      //   child: SingleChildScrollView(
      //     child: Container(
      //       child: Column(
      //         spacing: 20,
      //         children: [
      //           FutureBuilder(
      //             future: getConnected(),
      //             builder: (_, snap){
      //               if(snap.hasData){
      //                 return Text(snap.data!.join(","));
      //               }
      //               return Text("no connection");
      //             }
      //           ),
      //           Text(widget.name),
      //           Text(widget.socket_state),
      //           TextFormField(
      //             decoration: InputDecoration(
      //               border: OutlineInputBorder(),
      //               hintText: "ip"
      //             ),
      //           ),
      //           TextFormField(
      //             decoration: InputDecoration(
      //               border: OutlineInputBorder(),
      //               hintText: "port"
      //             ),
      //           ),
      //           FloatingActionButton.extended(
      //             elevation: 0.0,
      //             onPressed: (){}, 
      //             label: Text("Connect")
      //           )
      //         ],
      //       ),
      //     ),
      //   )
      // ),
    );
  }
}
