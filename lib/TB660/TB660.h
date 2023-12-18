//
// Created by qnurye on 12/16/23.
//

#ifndef CUBER_TB660_H
#define CUBER_TB660_H
//#define DEFAULT_PULSE_INTERVAL 1000
#define DEFAULT_DUTY_CYCLE 60
#include "Pin.h"

typedef unsigned percent;
typedef unsigned pwm;

enum Dir {
    Clockwise,
    Stop,
    Counterclockwise
};

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
    pwm pulseWidth;

    void enableMotor() const;

    void disableMotor() const;
};

#endif //CUBER_TB660_H
