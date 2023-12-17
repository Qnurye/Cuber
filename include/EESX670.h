//
// Created by qnurye on 12/16/23.
//

#ifndef CUBER_EE_SX670_H
#define CUBER_EE_SX670_H

#include "SensorShieldPins.h"
#include "Arduino.h"

class EE_SX670 {
public:
    // 构造函数，传入传感器的引脚
    explicit EE_SX670(pin outputPin) {
        outPin = outputPin;
    };

    // 初始化传感器
    void initialize() const {
        pinMode(outPin, INPUT);
    };

    // 读取传感器输出
    bool readSensor() const {
        return digitalRead(outPin) == HIGH;  // NPN输出类型，高电平表示有光
    };

private:
    int outPin;  // 传感器输出引脚
};

#endif //CUBER_EE_SX670_H
