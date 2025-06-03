# Sparkplug B 边缘采集框架（Go）

## 项目简介

本项目为通用的 Sparkplug B 边缘采集框架，支持多设备、插件式采集协议（如 Modbus、SNMP）、完整生命周期管理，适用于工业物联网场景。

## 主要特性
- MQTT 通信与重连、LWT
- Sparkplug 报文（NBIRTH/NDATA/NDEATH/DBIRTH/DDATA/DDEATH）
- 多设备建模与管理
- 插件式采集协议接口（Collector）
- 状态与生命周期管理
- 指令通道（NCMD/DCMD）
- 配置与日志模块

## 目录结构

```bash
sparkplug-edge/
├── main.go
├── config/
├── core/
├── mqtt/
├── proto/
├── data/
└── utils/
```

## 快速开始
```bash
go run main.go
```

## 依赖
- github.com/eclipse/paho.mqtt.golang
- github.com/eclipse/tahu
- github.com/spf13/viper
- go.uber.org/zap

## 扩展建议
- Collector 插件支持 Modbus/SNMP/模拟器
- 支持设备热插拔、远程命令、Prometheus 监控
