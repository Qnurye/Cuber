//
// Created by qnurye on 12/16/23.
//

#ifndef CUBER_TB660_H
#define CUBER_TB660_H
#define DEFAULT_DUTY_CYCLE 60

#include "Pin.h"

// 强迫症喜欢把单位定义成类型
typedef unsigned percent;

// 强迫症喜欢使用枚举
enum Dir {
    Clockwise,
    Stop,
    Counterclockwise
};

/**
 * TB660 步进电机
 */
class TB660 {
public:
    explicit TB660(pin directionPin, pin pulsePin, percent dutyCycle);

    // 初始化驱动器
    void initialize() const;

    void setPulseWidth(percent dutyCycle);

    // 发送步进脉冲
    void turn(Dir dir) const;

private:
    pin dirPin;
    pin pulPin;
    int pulseWidth;

    void enableMotor() const;

    void disableMotor() const;
};

#endif //CUBER_TB660_H
