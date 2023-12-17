//
// Created by qnurye on 12/17/23.
//

#ifndef CUBER_MEGA996R_H
#define CUBER_MEGA996R_H

#include "Arduino.h"

class MEGA996R {
public:
    // 构造函数，传入舵机控制引脚
    explicit MEGA996R(int servoPin) {
        pwmPin = servoPin;
    };

    // 初始化舵机
    void initialize() const {
        pinMode(pwmPin, OUTPUT);
    };

    // 控制舵机角度
    void setAngle(int angle) const {
        int pulseWidth = static_cast<int>(map(angle, 0, 180, 500, 2500));

        // 计算脉冲周期
        int pulsePeriod = 20000;  // 20ms

        // 计算高电平部分的时间
        int pulseHigh = pulseWidth;

        // 计算低电平部分的时间
        int pulseLow = pulsePeriod - pulseHigh;

        // 发送脉冲信号
        digitalWrite(pwmPin, HIGH);
        delay(pulseHigh);
        digitalWrite(pwmPin, LOW);
        delay(pulseLow);
    };

private:
    int pwmPin;
};

#endif //CUBER_MEGA996R_H
