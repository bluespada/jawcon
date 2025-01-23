/**
* Copyright (c) 2025 bluespada <pentingmain@gmail.com>
 */
package main

import (
	"log"
	"net"
	"strconv"
	"strings"
	"time"

	"github.com/ThomasT75/uinput"
)

// INFO: General Global variable
var SERVER_PORT = 8890
var MAX_BUFF_SIZE = 512
var DEFAULT_CONTROLLER = DeviceInfo{
	// TODO: you cann change pid (Product ID) and vid (Vendor ID) here and set the name
	//pid: 0x02d1,
	//vid: 0x045e,
	// name: "Xbox One Controller",
	pid:  0xDEAD,
	vid:  0xDAFF,
	name: "JawCon LoneWolf X2 Pro Controller",
}

type DeviceInfo struct {
	pid  uint16
	vid  uint16
	name string
}

var DEVICE_WHITELIST map[string]*net.UDPAddr

var MAP_BUTTON_KEYCODE = map[string]int{
	"KEYCODE_BUTTON_A":      uinput.ButtonSouth,
	"KEYCODE_BUTTON_B":      uinput.ButtonEast,
	"KEYCODE_BUTTON_Y":      uinput.ButtonWest,
	"KEYCODE_BUTTON_X":      uinput.ButtonNorth,
	"KEYCODE_BUTTON_R1":     uinput.ButtonBumperRight,
	"KEYCODE_BUTTON_R2":     uinput.ButtonTriggerRight,
	"KEYCODE_BUTTON_L1":     uinput.ButtonBumperLeft,
	"KEYCODE_BUTTON_L2":     uinput.ButtonTriggerLeft,
	"KEYCODE_BUTTON_START":  uinput.ButtonStart,
	"KEYCODE_BUTTON_SELECT": uinput.ButtonSelect,
	"KEYCODE_BUTTON_MODE":   uinput.ButtonMode,
	"KEYCODE_BUTTON_THUMBL": uinput.ButtonThumbLeft,
	"KEYCODE_BUTTON_THUMBR": uinput.ButtonThumbRight,
}

func main() {
	// This code is prototyping
	// anything in here will be hardcoded

	// TODO: Create a Gamepad
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

	// store_func()
	// TODO: Create a UDP Server to listen from App
	udpAddr, err := net.ResolveUDPAddr("udp", "0.0.0.0:8890")

	if err != nil {
		log.Fatalln(err)
	}

	// Start listening for UDP packages on the given address
	conn, err := net.ListenUDP("udp", udpAddr)

	if err != nil {
		log.Fatalln(err)
	}

	// Read from UDP listener in endless loop
	for {
		var buf [1024]byte
		_, _, err := conn.ReadFromUDP(buf[0:])
		if err != nil {
			log.Fatalln(err)
			return
		}

		var input = strings.Split(string(buf[0:]), ",")
		gamepad_hanlder(gamepad, input)

	}
}

func gamepad_hanlder(gamepad uinput.Gamepad, data []string) {
	var KEYCODE = strings.TrimSpace(data[0])
	var KEYVAL, _ = strconv.ParseFloat(strings.TrimSpace(data[1]), 32)
	var KEYTYPE, _ = strconv.Atoi(strings.TrimSpace(data[2]))
	if KEYTYPE == 1 {
		if key, ok := MAP_BUTTON_KEYCODE[KEYCODE]; ok {
			if KEYVAL == 0.0 {
				gamepad.ButtonDown(key)
			}
			if KEYVAL == 1.0 {
				gamepad.ButtonUp(key)
			}
		}
	} else if KEYTYPE == 0 {
		switch KEYCODE {
		case "AXIS_HAT_Y":
			if KEYVAL == -1.0 {
				gamepad.ButtonDown(uinput.ButtonDpadDown)
			} else {
				gamepad.ButtonUp(uinput.ButtonDpadDown)
			}

			if KEYVAL == 1.0 {
				gamepad.ButtonDown(uinput.ButtonDpadUp)
			} else {
				gamepad.ButtonUp(uinput.ButtonDpadUp)
			}
		case "AXIS_HAT_X":
			if KEYVAL == -1.0 {
				gamepad.ButtonDown(uinput.ButtonDpadLeft)
			} else {
				gamepad.ButtonUp(uinput.ButtonDpadLeft)
			}

			if KEYVAL == 1.0 {
				gamepad.ButtonDown(uinput.ButtonDpadRight)
			} else {
				gamepad.ButtonUp(uinput.ButtonDpadRight)
			}
		case "AXIS_X":
			gamepad.LeftStickMoveX(float32(KEYVAL))
		case "AXIS_Y":
			gamepad.LeftStickMoveY(float32(KEYVAL) * -1)
		case "AXIS_RZ":
			gamepad.RightStickMoveY(float32(KEYVAL) * -1)
		case "AXIS_Z":
			gamepad.RightStickMoveX(float32(KEYVAL))
		case "AXIS_LTRIGGER":
			log.Println("VAL", KEYVAL)
			gamepad.LeftTriggerForce(float32(KEYVAL))
		case "AXIS_RTRIGGER":
			gamepad.RightTriggerForce(float32(KEYVAL))
		}
	}
}

func store_func() {
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
