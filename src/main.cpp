#include <Arduino.h>
#include "SensorShieldPins.h"
#include "EESX670.h"
#include "TB660Driver.h"
#include "MEGA996R.h"

EE_SX670 sensorR(R_LOC_SENSOR_PIN);
EE_SX670 sensorL(L_LOC_SENSOR_PIN);
TB660Driver stepDriverL(L_STEP_DIR_PIN, L_STEP_PUL_PIN, DEFAULT_DUTY_CYCLE);
TB660Driver stepDriverR(R_STEP_DIR_PIN, R_STEP_PUL_PIN, DEFAULT_DUTY_CYCLE);
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

    while (!sensorL.readSensor()) {
        // 矫正左边
        if (sensorL.readSensor()) {
            Serial.print("Left is OK !");
        }
        stepDriverL.turn(Clockwise);
    }

    while (!sensorR.readSensor()) {
        // 矫正右边
        if (sensorR.readSensor()) {
            Serial.print("Right is OK !");
        }
        stepDriverR.turn(Clockwise);
    }
}

void loop() {

    if (sensorL.readSensor()) {
        Serial.println("Left is OK !");
    }
    if (sensorR.readSensor()) {
        Serial.println("Right is OK !");
    }
    delay(50);
//    servoL.setAngle(90);
//    delay(1000);
//    servoR.setAngle(90);
//    delay(1000);

//    if(sensorL.readSensor() || sensorR.readSensor()) {
//        Serial.println("Oh Light !");
//    }

//    delay(1000);  // 延时1秒
//
//    stepDriverL.turn(Clockwise);
//    stepDriverR.turn(Clockwise);
//
//    delay(500);
//    stepDriverL.turn(Counterclockwise);
//    stepDriverR.turn(Counterclockwise);
//
//    delay(1000);
//
//    stepDriverL.turn(Clockwise);
//    stepDriverR.turn(Clockwise);
//    delay(500);
//
//    stepDriverL.turn(Stop);
//    stepDriverR.turn(Stop);
}