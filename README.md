# nestos-kubernetes-deployer

## 介绍
NKD（NestOS Kubernetes Deployer）是专为在NestOS上部署和维护Kubernetes集群而打造的解决方案。其旨在通过在集群外提供一系列服务，涵盖了基础设施和Kubernetes核心组件的部署、更新和配置管理等，从而简化了集群部署和升级的流程。NKD的设计目标在于提供更为便捷的集群操作体验，使得用户能够轻松完成复杂的管理任务，从而提高整体部署和维护的效率。

#### 支持平台
支持多种部署平台，NKD根据集群需求，连接基础设施提供商动态创建所需的IaaS资源，目前支持OpenStack和libvirt平台。

## 软件架构
详细内容请见[软件架构说明](docs/zh/overall_design.md)

## 安装部署
详细内容请见[用户操作手册](docs/zh/manual.md)

## 未来规划
NKD的最终愿景是通过长期驻留的服务形式，为运维提供支持，并同时支持多个集群的高效管理。它将提供持久的配置变更记录、证书管理、多种更新升级策略以及镜像源频道等丰富功能。未来，我们将持续精进NKD的功能和性能，并引入更多智能化特性，例如自动化故障处理和资源优化等。我们的目标是将NKD塑造成NestOS生态的核心组成部分，为云原生场景下的运维工作提供全方位支持，从而进一步推动云原生技术的发展和广泛应用。

## 参与贡献
非常欢迎对本项目感兴趣的伙伴加入我们，并参与贡献。

## License
Apache License 2.0