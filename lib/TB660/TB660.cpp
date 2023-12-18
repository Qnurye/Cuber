//
// Created by qnurye on 12/18/23.
//
#include "TB660.h"
#include "Arduino.h"

TB660::TB660(pin directionPin, pin pulsePin, percent dutyCycle) {
    dirPin = directionPin;
    pulPin = pulsePin;
    pulseWidth = 0;
    setPulseWidth(dutyCycle);
}

void TB660::initialize() const {
    pinMode(dirPin, OUTPUT);
    pinMode(pulPin, OUTPUT);
}

void TB660::setPulseWidth(percent dutyCycle) {
    pulseWidth = map(dutyCycle, 0, 100, 0, 255);
}

void TB660::turn(Dir dir) const {
    if (dir == Stop) {
        digitalWrite(dirPin, LOW);

        disableMotor();
    } else {
        digitalWrite(dirPin, dir == Clockwise ? HIGH : LOW);

        enableMotor();
    }
}

void TB660::enableMotor() const {
    analogWrite(pulPin, static_cast<int>(pulseWidth));
}

void TB660::disableMotor() const {
    analogWrite(pulPin, 0);
}
