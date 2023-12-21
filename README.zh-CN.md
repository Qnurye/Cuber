# Cuber

## 概述

该项目结合了 Arduino（PlatformIO）和 Go 代码，用于创建一个控制和与机器人设备通信的系统。Arduino 代码控制硬件，而 Go 代码管理与 Arduino 的通信，发送命令和处理响应。

## 组件

### [Arduino（PlatformIO）](Arduino/README.zh-CN.md)

#### 开发环境

项目的 Arduino 部分使用 PlatformIO 进行开发，这是一个针对嵌入式系统和物联网应用程序的强大平台。PlatformIO 简化了开发流程，提供了一个统一的环境，支持多种不同的嵌入式平台和硬件架构。

#### 开发板

该项目使用 Arduino Uno R3（ATMEGA328P）开发板。

#### 库

Arduino 代码包括与特定模块交互的库：
- [EESX670](Arduino/lib/EESX670): 光电传感器
- [MEGA996R](Arduino/lib/MEGA996R): 舵机
- [TB660](Arduino/lib/TB660): 步进电机

### [公式解析器及命令发送 (Go)](Go/README.zh-CN.md)

#### 开发环境

Go 代码设计用于在计算机上运行，并通过串口与 Arduino 设备通信。

#### 主要逻辑

Go 代码的主要逻辑涉及通过串口通信向 Arduino 发送命令。这些命令控制机器人设备的各种操作，如抓取、旋转等。

## 目录结构

```plaintext
├── Arduino
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
├── Go
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

## 入门

1. 克隆仓库。
2. 使用 PlatformIO 设置 [Arduino 环境](Arduino/README.zh-CN.md)。
3. 安装所需的 [Go 依赖](Go/README.zh-CN.md)。
4. 构建并运行 Go 代码。

## 使用

系统从 `formula.json` 文件中读取魔方公式，并通过串口通信向 Arduino 设备发送相应的命令。

## 贡献

欢迎贡献！如果您想为项目做贡献，请遵循标准的 GitHub 分叉和拉取请求工作流程。

## 许可证

本项目采用 [MIT 许可证](LICENSE)。
