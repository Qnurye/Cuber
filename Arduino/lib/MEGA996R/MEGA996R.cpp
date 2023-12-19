//
// Created by qnurye on 12/18/23.
//

#include "MEGA996R.h"

/**
 * MEGA996R 伺服电机类的构造函数
 * @param servoPin 伺服电机控制引脚
 */
MEGA996R::MEGA996R(pin servoPin) {
    pwmPin = servoPin;
}

/**
 * 初始化伺服电机对象
 * 设置伺服电机控制引脚
 */
void MEGA996R::initialize() {
    servo.attach(pwmPin);
}

/**
 * 设置伺服电机的旋转角度
 * @param angle 旋转角度，取值范围为 [0, 180]
 */
void MEGA996R::setAngle(int angle) {
    servo.write(angle);
}
