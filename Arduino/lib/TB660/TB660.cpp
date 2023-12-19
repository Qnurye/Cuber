//
// Created by qnurye on 12/18/23.
//

#include "TB660.h"
#include "Arduino.h"

/**
 * TB660 型号步进电机类的构造函数
 * @param directionPin 控制方向的引脚 
 * @param pulsePin 控制脉冲的引脚
 * @param dutyCycle 占空比，单位是 %
 */
TB660::TB660(pin directionPin, pin pulsePin, percent dutyCycle) {
    dirPin = directionPin;
    pulPin = pulsePin;
    pulseWidth = 0;
    setPulseWidth(dutyCycle);
}

/**
 * 初始化步进电机对象
 * 将 dir 和 pul 引脚设置为输出引脚
 */
void TB660::initialize() const {
    pinMode(dirPin, OUTPUT);
    pinMode(pulPin, OUTPUT);
}

/**
 * 设置占空比
 * @param dutyCycle 占空比
 */
void TB660::setPulseWidth(percent dutyCycle) {
    // 占空比 (dutyCycle) 取值于 [0, 100]
    // 将其映射至 pwm 输出的取值空间 [0,255]
    pulseWidth = static_cast<int>(map(dutyCycle, 0, 100, 0, 255));
}

/**
 * 控制步进电机转动或停止
 * @param dir 转动方向 
 */
void TB660::turn(Dir dir) const {
    if (dir == Stop) {
        // 停止转动时，调整所有输出为低电平
        digitalWrite(dirPin, LOW);
        disableMotor();
    } else {
        digitalWrite(dirPin, dir == Clockwise ? HIGH : LOW);
        enableMotor();
    }
}

/**
 * 启用步进电机，设置脉冲输出占空比
 */
void TB660::enableMotor() const {
    analogWrite(pulPin, pulseWidth);
}

/**
 * 停用步进电机，将脉冲输出设为零
 */
void TB660::disableMotor() const {
    analogWrite(pulPin, 0);
}
