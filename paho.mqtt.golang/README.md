# paho.mqtt.golang

MQTT（Message Queuing Telemetry Transport）是一种轻量级、基于发布-订阅模式的消息传输协议，适用于资源受限的设备和低带宽、高延迟或不稳定的网络环境。它在物联网应用中广受欢迎，能够实现传感器、执行器和其它设备之间的高效通信。

MQTT 是基于发布-订阅模式的通信协议，由 MQTT 客户端通过主题（Topic）发布或订阅消息，通过 MQTT Broker 集中管理消息路由，并依据预设的服务质量等级(QoS)确保端到端消息传递可靠性。

MQTT Broker 是负责处理客户端请求的关键组件，包括建立连接、断开连接、订阅和取消订阅等操作，同时还负责消息的转发。

一个高效强大的 MQTT Broker 能够轻松应对海量连接和百万级消息吞吐量，从而帮助物联网服务提供商专注于业务发展，快速构建可靠的 MQTT 应用。MQTT Broker的实现见参考5。

## EMQX

本仓库使用 EMQX 作为 MQTT Broker，并使用 [paho.mqtt.golang](https://github.com/eclipse/paho.mqtt.golang) 作为 MQTT 客户端库，实现了一个简单的 MQTT 客户端示例，用于发布和订阅消息。注意**该模块只支持MQTT 3.1/3.11协议**。而最新MQTT协议已经是6.0了。

EMQX 是目前物联网应用中最具扩展性的 MQTT Broker。它能够以亚毫秒级的延迟在一秒钟内处理百万级的 MQTT 消息，并支持在一个集群内连接高达 1 亿个客户端进行消息传输。EMQX 兼容 MQTT 5.0 和 3.x 版本。

它是分布式物联网网络的理想选择，可以在 Microsoft Azure、Amazon Web Services 和 Google Cloud 等云上运行。EMQX 支持 MQTT over TLS/SSL，并支持多种认证机制，如 PSK、JWT 和 X.5093。与 Mosquitto 不同，EMQX 支持通过 CLI、HTTP API 和 Dashboard 进行集群管理。

emqx使用docker安装：

```shell
docker run -d --name emqx \
    -e EMQX_LISTENERS_TCP_DEFAULT_BIND=1884 \
    -p 18083:18083 \
    -p 1884:1884 \
emqx/emqx:6.2.0-alpha.1
```

## paho.mqtt.golang 的参考资料

1. [如何在 Golang 中使用 MQTT.](https://www.emqx.com/zh/blog/how-to-use-mqtt-in-golang)
2. [开发者指南.](https://docs.emqx.com/zh/emqx/latest/connect-emqx/developer-guide.html)
3. [MQTT 发布/订阅模式介绍.](https://www.emqx.com/zh/blog/mqtt-5-introduction-to-publish-subscribe-model)
4. [MQTT 协议快速入门 2025：基础知识和实用教程.](https://www.emqx.com/zh/blog/the-easiest-guide-to-getting-started-with-mqtt)
5. [MQTT服务器（MQTT Broker）：工作原理与快速入门指南.](https://www.emqx.com/zh/blog/the-ultimate-guide-to-mqtt-broker-comparison)
6. [Eclipse Paho MQTT Go client go dev.](https://pkg.go.dev/github.com/eclipse/paho.mqtt.golang)
7. [云消息队列 MQTT 版配置 DDoS 高防.](https://help.aliyun.com/zh/apsaramq-for-mqtt/mqtt-upgraded/use-cases/configure-anti-ddos-high-for-mqtt?spm=a2c4g.11186623.help-menu-100973.d_3_0_0.4cdb50b9nOdvga)
