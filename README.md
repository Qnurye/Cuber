[简体中文](README.zh-CN.md)

# Cuber

## Overview

This project combines Arduino (PlatformIO) and Go code to create a system for controlling and communicating with a robotic device. The Arduino code controls the hardware, while the Go code manages the communication with the Arduino, sending commands and handling responses.

## Components

### [Arduino (PlatformIO)](Arduino/README.md)

#### Development Environment

The Arduino part of the project is developed using PlatformIO, a powerful platform for embedded systems and IoT applications. PlatformIO simplifies the development process, providing a unified environment supporting various embedded platforms and hardware architectures.

#### Board

The project utilizes the Arduino Uno R3 (ATMEGA328P) board.

#### Libraries

The Arduino code includes libraries for interfacing with specific modules:
- [EESX670](Arduino/lib/EESX670): Photocell sensor
- [MEGA996R](Arduino/lib/MEGA996R): Servo motor
- [TB660](Arduino/lib/TB660): Stepper motor

### [Formula Parser and Command sender (Go)](Go/README.md)

#### Development Environment

The Go code is designed to run on a computer and communicate with the Arduino device over a serial connection.

#### Main Logic

The main logic of the Go code involves sending commands to the Arduino via serial communication. The commands control various actions of a robotic device, such as gripping, rotating, and other operations.

## Directory Structure

```plaintext
├── [Arduino](Arduino/README.md)
│   ├── lib
│   │   ├── EESX670
│   │   ├── MEGA996R
│   │   └── TB660
│   ├── include
│   │   ├── Controllers.h
│   │   ├── SensorShieldPins.h
│   │   └── SerialCommands.h
│   ├── src
│   │   └── main.cpp
│   ├── .gitignore
│   ├── platformio.ini
│   └── M-CubeRobot 说明书.pdf
├── [Go](Go/README.md)
│   ├── build
│   ├── config
│   │   ├── cmd.json
│   │   ├── delay.json
│   │   └── formula.json
│   ├── go.mod
│   ├── go.sum
│   ├── main.go
│   ├── pkg
│   │   ├── config
│   │   │   ├── cmd.go
│   │   │   ├── delay.go
│   │   │   └── formula.go
│   │   └── rubiksCube
│   │       ├── parser.go
│   │       └── README.md
│   ├── README.md
│   └── test
│       ├── cmd_test.go
│       ├── delay_test.go
│       ├── formula_test.go
│       └── parser_test.go
└── .gitignore
```

## Getting Started

1. Clone the repository.
2. Set up the [Arduino environment](Arduino/README.md) using PlatformIO.
3. Install the required [Go dependencies](Go/README.md).
4. Build and run the Go code.

## Usage

The system reads a Rubik's Cube formula from the `formula.json` file and sends corresponding commands to the Arduino device via serial communication.

## Contributions

Contributions are welcome! If you'd like to contribute to the project, please follow the standard GitHub fork and pull request workflow.

## License

This project is licensed under the [MIT License](LICENSE).
