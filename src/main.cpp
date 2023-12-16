#include <Arduino.h>
#include "EESX670.h"
#include "SensorShieldPins.h"

EE_SX670 sensorR(R_LOC_SENSOR_PIN);
EE_SX670 sensorL(L_LOC_SENSOR_PIN);

void setup() {
    sensorR.begin();
    sensorL.begin();
    Serial.begin(9600);
}

void loop() {
    bool sensorValue = sensorL.readSensor();
    Serial.println(sensorValue ? "1\n" : "0\n");

    delay(10);  // 延时1秒
}