//
// Created by qnurye on 12/18/23.
//

#include <Arduino.h>
#include "EESX670.h"

/**
 * EE_SX670 光电传感器类的构造函数
 * @param outputPin 输出引脚
 */
EE_SX670::EE_SX670(pin outputPin) {
    outPin = outputPin;
}

/**
 * 初始化光电传感器对象
 * 将输出引脚设置为上拉输入
 */
void EE_SX670::initialize() const {
    pinMode(outPin, INPUT_PULLUP);
}

/**
 * 获取光电传感器的当前状态
 * @return 如果传感器状态为高电平，返回 true；否则返回 false
 */
bool EE_SX670::getState() const {
    return digitalRead(outPin) == HIGH;
}
