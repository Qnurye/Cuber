#include <Arduino.h>
#include "SensorShieldPins.h"
#include "TB660.h"
#include "MEGA996R.h"
#include "EESX670.h"

EE_SX670 sensorL(L_LOC_SENSOR_PIN);
EE_SX670 sensorR(R_LOC_SENSOR_PIN);
TB660 stepDriverL(L_STEP_DIR_PIN, L_STEP_PUL_PIN, DEFAULT_DUTY_CYCLE);
TB660 stepDriverR(R_STEP_DIR_PIN, R_STEP_PUL_PIN, DEFAULT_DUTY_CYCLE);
MEGA996R servoL(L_SERVO_PIN);
MEGA996R servoR(R_SERVO_PIN);

void setup() {
    sensorL.initialize();
    sensorR.initialize();
    stepDriverL.initialize();
    stepDriverR.initialize();
    servoL.initialize();
    servoR.initialize();
    Serial.begin(9600);

    pinMode(LED_BUILTIN, OUTPUT);
    digitalWrite(LED_BUILTIN, HIGH);
    delay(2000);
    digitalWrite(LED_BUILTIN, LOW);
    delay(1000);

    if (!sensorL.getState()) {
        // 矫正左边
        stepDriverL.turn(Clockwise);
        while (!sensorL.getState()) {}
        stepDriverL.turn(Stop);
    }
    if (!sensorR.getState()) {
        // 矫正右边
        stepDriverR.turn(Clockwise);
        while (!sensorR.getState()) {}
        stepDriverR.turn(Stop);
    }
}

void loop() {
}
