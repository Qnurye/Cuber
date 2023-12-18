//
// Created by qnurye on 12/16/23.
//

#ifndef CUBER_EE_SX670_H
#define CUBER_EE_SX670_H
#include "Pin.h"

class EE_SX670 {
public:
    explicit EE_SX670(pin outputPin);

    // 初始化传感器
    void initialize() const;

    // 读取传感器输出
    bool getState() const;

private:
    pin outPin;
    volatile bool state;

//    static EE_SX670* instance;
//    static void RisingHandler();
//    static void FallingHandler();
//    static void ChangeHandler();
};

#endif //CUBER_EE_SX670_H
