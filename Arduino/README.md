# Cuber

## Development Environment 💻

This project is powered by [PlatformIO](https://platformio.org/) 🚀

> [!quote] ChatGPT 🤖
>
> PlatformIO is a development platform designed for embedded systems and Internet of Things (IoT) applications. It provides a unified development environment that supports various embedded platforms and hardware architectures. PlatformIO aims to simplify the development process for embedded systems, making it easier for developers to build, debug, and deploy embedded applications.

The development environment details are as follows:

1. Arch Linux 6.6.7-arch1-1 🐧
2. Clion 2023.3 Build # CL-233.11799.238
3. OpenJDK 64-Bit Server VM by JetBrains s.r.o.
4. Desktop Environment: i3

The board used is Arduino Uno R3 (ATMEGA328P).

## Pins

According to [华控工作室. M-CubeRobot Manual [Z].](M-CubeRobot说明书.pdf), Arduino Uno indirectly connects with other modules via the Sensor Shield. The Arduino Uno and Sensor Shield pins are identical.

The following table provides annotations for the Sensor Shield pin connections:

| Sensor Shield Interface |            Connection            |
|:-----------------------:|:--------------------------------:|
|            0            |         Communication TX         |
|            1            |         Communication RX         |
|            2            |        R Position Sensor         |
|            3            |        L Position Sensor         |
|            5            |             R Servo              |
|            6            |             L Servo              |
|            8            | R Stepper Driver Direction (DIR) |
|            9            | L Stepper Driver Direction (DIR) |
|           10            |   R Stepper Driver Pulse (PUL)   |
|           11            |   R Stepper Driver Pulse (PUL)   |

## Serial Communication and Commands

The mechanical gripper responds to the following commands via Serial:

| Command | Action                           |
|---------|----------------------------------|
| 1       | R Gripper Close                  |
| 2       | R Gripper Open                   |
| 3       | L Gripper Close                  |
| 4       | L Gripper Open                   |
| 5       | R Gripper Rotate CW 90 degrees   |
| 6       | R Gripper Rotate CCW 90 degrees  |
| 7       | R Gripper Rotate CW 180 degrees  |
| 8       | L Gripper Rotate CW 90 degrees   |
| 9       | L Gripper Rotate CCW 90 degrees  |
| 0       | L Gripper Rotate CW 180 degrees  |
| B       | L Gripper Rotate CCW 180 degrees |
| C       | R Gripper Rotate CCW 180 degrees |

## Modules 🧩

### Photocell Sensor

[EESX670.h](lib/EESX670/EESX670.h)

> Status: `available ✅`

### Stepper Motor

[TB660.h](lib/TB660/TB660.h)

> Status: `available ✅`

Using Pulse Width Modulation (PWM) and `analogWrite()`, a rectangular pulse waveform is output to the stepper motor. The higher the frequency, the faster the rotation speed.

### Servo Motor

[MEGA996R.h](lib/MEGA996R/MEGA996R.h)

> Status: `available ✅`

The servo motor depends on the Arduino official library `Servo.h`.

## Project Structure 🕸

- [include/](include) Contains some header files
  - [Controllers.h](include/Controllers.h) Some controller functions, including control encapsulation for servo and stepper motors
  - [SensorShieldPin.h](include/SensorShieldPins.h) Configures pin interfaces, see [Pins](#pins)
  - [SerialCommands.h](include/SerialCommands.h) Configures serial commands, see [Serial Communication and Commands](#serial-communication-and-commands)
- [lib/](lib) Contains some libraries
  - [EESX670](lib/EESX670) Photocell sensor
  - [MEGA996R](lib/MEGA996R) Servo motor
  - [TB660](lib/TB660) Stepper motor
  - ~~[Pin](lib/Pin) For those who prefer treating pin as a type~~
- [src/](src)
  - [main.cpp](src/main.cpp) Main logic code
