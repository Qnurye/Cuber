//
// Created by qnurye on 12/16/23.
//

#ifndef CUBER_TB660DRIVER_H
#define CUBER_TB660DRIVER_H

#include "SensorShieldPins.h"
#define DEFAULT_PULSE_INTERVAL 1000
#define DEFAULT_DUTY_CYCLE 50

typedef unsigned long ms;
typedef unsigned percent;
typedef unsigned pwm;

enum Dir {
    Clockwise,
    Stop,
    Counterclockwise
};

class TB660Driver {
public:
    explicit TB660Driver(pin directionPin, pin pulsePin, percent dutyCycle) {
        dirPin = directionPin;
        pulPin = pulsePin;
        pulseWidth = 0;
        setPulseWidth(dutyCycle);
    };

    // 初始化驱动器
    void initialize() const {
        pinMode(dirPin, OUTPUT);
        pinMode(pulPin, OUTPUT);
    };

    void setPulseWidth(percent dutyCycle) {
        pulseWidth = map(dutyCycle, 0, 100, 0, 255);
    }

    // 发送步进脉冲
    void turn(Dir dir) const {
        if (dir == Stop) {
            digitalWrite(dirPin, LOW);

            disableMotor();
        } else {
            digitalWrite(dirPin, dir == Clockwise ? HIGH : LOW);

            enableMotor();
        }
    };

private:
    pin dirPin;
    pin pulPin;
    pwm pulseWidth;

    void enableMotor() const {
        analogWrite(pulPin, static_cast<int>(pulseWidth));
    };

    void disableMotor() const {
        analogWrite(pulPin, 0);
    };
};

#endif //CUBER_TB660DRIVER_H
