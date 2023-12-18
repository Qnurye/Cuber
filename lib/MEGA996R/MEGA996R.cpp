//
// Created by qnurye on 12/18/23.
//
#include <Arduino.h>
#include "MEGA996R.h"
#include "Pin.h"

MEGA996R::MEGA996R(pin servoPin) {
        pwmPin = servoPin;
}

void MEGA996R::initialize() const {
        pinMode(pwmPin, OUTPUT);
}

void MEGA996R::setAngle(int angle) const {
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
}
