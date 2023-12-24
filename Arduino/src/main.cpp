#include <Arduino.h>
#include "SensorShieldPins.h"
#include "SerialCommands.h"
#include "TB660.h"
#include "MEGA996R.h"
#include "EESX670.h"
#include "Controllers.h"

EE_SX670 sensorL(R_STEP_PUL_PIN);
EE_SX670 sensorR(R_LOC_SENSOR_PIN); // #10 脚位用不了 pwm, 换了个口
TB660 stepDriverL(L_STEP_DIR_PIN, L_STEP_PUL_PIN, DEFAULT_DUTY_CYCLE);
TB660 stepDriverR(R_STEP_DIR_PIN, L_LOC_SENSOR_PIN, DEFAULT_DUTY_CYCLE);
MEGA996R servoL(L_SERVO_PIN);
MEGA996R servoR(R_SERVO_PIN);

unsigned char cmd = '\0';

void setup() {
    sensorR.initialize();
    sensorL.initialize();
    stepDriverR.initialize();
    stepDriverL.initialize();
    servoR.initialize();
    servoL.initialize();
    Serial.begin(9600);

    pinMode(LED_BUILTIN, OUTPUT);
    digitalWrite(LED_BUILTIN, HIGH);
    delay(1000);
    digitalWrite(LED_BUILTIN, LOW);
    delay(500);

    if (!sensorR.getState()) {
        // 矫正右边
        stepDriverR.turn(Clockwise);
        while (!sensorR.getState()) {}
        stepDriverR.turn(Stop);
    }
    if (!sensorL.getState()) {
        // 矫正左边
        stepDriverL.turn(Clockwise);
        while (!sensorL.getState()) {}
        stepDriverL.turn(Stop);
    }

    MEGA996RGrip(&servoL, Close);
    MEGA996RGrip(&servoR, Close);
}

void loop() {
    Serial.readBytes(&cmd, 1);
    if (cmd != '\0') {
        switch (cmd) {
            case CMD_L_ROTATE_CW_90:
                TB660Rotate(&stepDriverL, &sensorL, Clockwise, 90);
                break;
            case CMD_L_ROTATE_CW_180:
                TB660Rotate(&stepDriverL, &sensorL, Clockwise, 180);
                break;
            case CMD_L_ROTATE_CCW_90:
                TB660Rotate(&stepDriverL, &sensorL, Counterclockwise, 90);
                break;
            case CMD_L_ROTATE_CCW_180 | 'B':
                TB660Rotate(&stepDriverL, &sensorL, Counterclockwise, 180);
                break;
            case CMD_R_ROTATE_CW_90:
                TB660Rotate(&stepDriverR, &sensorR, Clockwise, 90);
                break;
            case CMD_R_ROTATE_CW_180:
                TB660Rotate(&stepDriverR, &sensorR, Clockwise, 180);
                break;
            case CMD_R_ROTATE_CCW_90:
                TB660Rotate(&stepDriverR, &sensorR, Counterclockwise, 90);
                break;
            case CMD_R_ROTATE_CCW_180 | 'C':
                TB660Rotate(&stepDriverR, &sensorR, Counterclockwise, 180);
                break;
            case CMD_L_GRIP_OPEN:
                MEGA996RGrip(&servoL, Open);
                break;
            case CMD_L_GRIP_CLOSE:
                MEGA996RGrip(&servoL, Close);
                break;
            case CMD_R_GRIP_OPEN:
                MEGA996RGrip(&servoR, Open);
                break;
            case CMD_R_GRIP_CLOSE:
                MEGA996RGrip(&servoR, Close);
                break;
            default:
                break;
        }
        cmd = '\0';
    }
}
