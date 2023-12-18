//
// Created by qnurye on 12/17/23.
//

#ifndef CUBER_MEGA996R_H
#define CUBER_MEGA996R_H

#include "Pin.h"
#include "Servo.h"

class MEGA996R {
public:
    // 构造函数，传入舵机控制引脚
    explicit MEGA996R(pin servoPin);

    // 初始化舵机
    void initialize();

    // 控制舵机角度
    void setAngle(int angle);

private:
    pin pwmPin;
    Servo servo;
};

#endif //CUBER_MEGA996R_H
