
# Cuber

## Development Environment

本项目由 [PlatformIO](https://platformio.org/) 强力驱动 🚀

> [!quote] ChatGPT 🤖
>
> PlatformIO 是一个开发平台，专为嵌入式系统和物联网（IoT）应用程序的开发而设计。它提供了一个统一的开发环境，支持多种不同的嵌入式平台和硬件架构。PlatformIO 的目标是简化嵌入式系统的开发流程，使开发人员能够更轻松地构建、调试和部署嵌入式应用程序。

笨人开发环境如下：

1. Arch Linux 6.6.7-arch1-1 🐧
2. Clion 2023.3 Build # CL-233.11799.238
3. OpenJDK 64-Bit Server VM by JetBrains s.r.o.
4. Desktop Environment: i3

使用的 motherboard 是 Arduino Uno R3 (ATMEGA328P)。

## 引脚

据 [华控工作室. M-CubeRobot 说明书 [Z].](M-CubeRobot说明书.pdf)，Arduino Uno 与别的模块经由 Sensor Shield 间接连接。而 Arduino Uno 与 Sensor Shield 引脚实际一致。

如下为 Sensor Shield 的脚位连接注释：

| Sensor Shield 接口 |       连接至       |
|:----------------:|:---------------:|
|        0         |     通讯板 TX      |
|        1         |     通讯板 RX      |
|        2         |     R 定位传感器     |
|        3         |     L 定位传感器     |
|        5         |      R 舵机       |
|        6         |      L 舵机       |
|        8         | R 步进驱动器方向 (DIR) |
|        9         | L 步进驱动器方向 (DIR) |
|        10        | R 步进驱动器脉冲 (PUL) |
|        11        | R 步进驱动器脉冲 (PUL) | 

## 串口与指令
