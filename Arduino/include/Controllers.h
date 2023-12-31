//
// Created by qnurye on 12/18/23.
//

#ifndef CUBER_CONTROLLER_H
#define CUBER_CONTROLLER_H

#include "TB660.h"
#include "EESX670.h"

void TB660Rotate(TB660* tb660, EE_SX670* eeSx670, Dir dir, int degree) {
    switch (degree) {
        case 180:
            tb660->turn(dir);
            while (eeSx670->getState()) {} // 先等到sensor变false
            while (!eeSx670->getState()) {} // 然后转到true, 为一个90度
            tb660->turn(Stop);
            // 180度不break，直接再转一次
        case 90:
            tb660->turn(dir);
            while (eeSx670->getState()) {}
            while (!eeSx670->getState()) {}
            tb660->turn(Stop);
            break;
        default:
            break;
    }
    Serial.write('.');
}

enum GripStatus {
    Open,
    Close
};

void MEGA996RGrip(MEGA996R* mega996R, GripStatus status) {
    if (status == Close) {
        mega996R->setAngle(85);
    } else {
        mega996R->setAngle(180);
    }
    delay(400);
    Serial.write('.');
}

#endif //CUBER_CONTROLLER_H
