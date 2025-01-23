/**
* Copyright (c) 2025 bluespada <pentingmain@gmail.com>
*/
package main

import (
    "log"
    "time"
    "net"
    "github.com/ThomasT75/uinput"
)

//INFO: General Global variable
var SERVER_PORT = 8890
var MAX_BUFF_SIZE = 1024
var DEFAULT_CONTROLLER = DeviceInfo{
    // TODO: you cann change pid (Product ID) and vid (Vendor ID) here and set the name
    //pid: 0x02d1,
    //vid: 0x045e,
    // name: "Xbox One Controller",
    pid: 0xDEAD,
    vid: 0xDAFF,
    name: "JawCon LoneWolf X2 Pro Controller",
}

type DeviceInfo struct {
    pid uint16
    vid uint16
    name string
}

var DEVICE_WHITELIST map[string]*net.UDPAddr


func main(){
    // This code is prototyping
    // anything in here will be hardcoded

    // TODO: Create a Gamepad

    store_func()
    // TODO: Create a UDP Server to listen from App
}

func store_func(){
    // NOTE: this function will create uinput to create a virtual gamepad.
    gamepad, err := uinput.CreateGamepad(
        "/dev/uinput",
        []byte(DEFAULT_CONTROLLER.name),
        DEFAULT_CONTROLLER.pid,
        DEFAULT_CONTROLLER.vid,
    )
    if err != nil {
        log.Fatalln("Error: ", err)
    }


    defer gamepad.Close()


    for {
        // err := keyboard.KeyDown(uinput.KeyR)
        // if err != nil {
        //    log.Fatalln("Error: ", err)
        //}
        // 03006d98d10200005e04000001000000,Xbox One Controller,a:b0,b:b1,x:b2,y:b3,back:b8,guide:b10,start:b9,leftstick:b11,rightstick:b12,leftshoulder:b4,rightshoulder:b5,dpup:b13,dpdown:b14,dpleft:b15,dpright:b16,leftx:a0,lefty:a1,rightx:a3,righty:a4,lefttrigger:a2,righttrigger:a5,crc:986d,platform:Linux,
        gamepad.ButtonPress(uinput.ButtonWest)
        time.Sleep(
            500 * time.Millisecond,
        )
        gamepad.ButtonPress(uinput.ButtonNorth)
        time.Sleep(
            500 * time.Millisecond,
        )
        gamepad.ButtonPress(uinput.ButtonEast)
        time.Sleep(
            500 * time.Millisecond,
        )
        gamepad.ButtonPress(uinput.ButtonSouth)
        time.Sleep(
            500 * time.Millisecond,
        )
    }
}
