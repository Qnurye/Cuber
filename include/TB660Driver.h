//
// Created by qnurye on 12/16/23.
//

#ifndef CUBER_TB660DRIVER_H
#define CUBER_TB660DRIVER_H

#ifndef TB660_DRIVER_H
#define TB660_DRIVER_H

#include "SensorShieldPins.h"

enum dir {
    Clockwise,
    Counterclockwise
};

class TB660Driver {
public:
    TB660Driver(pin dirPin, pin pulPin);
    ~TB660Driver();

    // 初始化驱动器
    bool initialize();

    // 设置步进电机的脉冲和方向
    void setDirection(dir direction);

    // 发送步进脉冲
    void sendPulse();

private:
    pin dirPin;
    pin pulPin;
};

#endif // TB660_DRIVER_H


#endif //CUBER_TB660DRIVER_H
