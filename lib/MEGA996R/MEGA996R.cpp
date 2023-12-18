//
// Created by qnurye on 12/18/23.
//
#include <Arduino.h>
#include "MEGA996R.h"

MEGA996R::MEGA996R(pin servoPin) {
    pwmPin = servoPin;
}

void MEGA996R::initialize() {
    servo.attach(pwmPin);
//        pinMode(pwmPin, OUTPUT);
}

void MEGA996R::setAngle(int angle) {
    servo.write(angle);
    // 引脚5和6的默认PWM频率是976.56Hz。
    // 这意味着一个完整的PWM周期大约为1024微秒（1秒/976.56次）
//    int pulseWidth = static_cast<int>(map(angle, 0, 180, 0, 255));
//    analogWrite(pwmPin, pulseWidth);

//        int pulseWidth = static_cast<int>(map(angle, 0, 180, 500, 2500));
//
//        // 计算脉冲周期
//        int pulsePeriod = 20000;  // 20ms
//
//        // 计算高电平部分的时间
//        int pulseHigh = pulseWidth;
//
//        // 计算低电平部分的时间
//        int pulseLow = pulsePeriod - pulseHigh;
//
//        // 发送脉冲信号
//        digitalWrite(pwmPin, HIGH);
//        delayMicroseconds(pulseHigh);
//        digitalWrite(pwmPin, LOW);
//        delayMicroseconds(pulseLow);
}
