//
// Created by qnurye on 12/18/23.
//
#include <Arduino.h>
#include "EESX670.h"


//EE_SX670* EE_SX670::instance = nullptr;

EE_SX670::EE_SX670(pin outputPin) {
    outPin = outputPin;
    state = false;
//    instance = this;
}

void EE_SX670::initialize() const {
    pinMode(outPin, INPUT_PULLUP);
//    attachInterrupt(digitalPinToInterrupt(outPin), ChangeHandler, CHANGE);
//    attachInterrupt(digitalPinToInterrupt(outPin), RisingHandler, RISING);
//    attachInterrupt(digitalPinToInterrupt(outPin), FallingHandler, FALLING);
}

//void EE_SX670::ChangeHandler() {
//    if (instance != nullptr) {
//        instance->state = digitalRead(instance->outPin) == HIGH;
//    }
//}
//
//void EE_SX670::RisingHandler() {
//    if (instance != nullptr) {
//        instance->state = true;
//    }
//}
//
//void EE_SX670::FallingHandler() {
//    if (instance != nullptr) {
//        instance->state = false;
//    }
//}

bool EE_SX670::getState() const {
    return digitalRead(outPin) == HIGH;
}
