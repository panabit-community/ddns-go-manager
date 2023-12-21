# Panabit 网关插件系列之 DDNS-GO 启动器

一个 Panabit 智能应用网关插件，可在不开启 SSH 的条件下为用户安装和维护 [jeessy2/ddns-go](https://github.com/jeessy2/ddns-go)。

本项目的主要工作：

+ 适配 Panabit 网关插件的生命周期命令和钩子，使得安装和管理可以在 Web 界面完成
+ 额外的进程管理，可以控制 DDNS-GO 的启停和开机自动启动
+ 反向代理和垫片，利用 Web 管理页隐藏上游 DDNS-GO 没有 TLS 加密的接口

## 安装

下载 Release 版本，请先校验文件哈希无误。打开 Panabit 网关的 Web 管理页，前往**系统概况**->**应用商店**，点按右上角**安装更新**上传并安装插件即可。

## 自行构建

本地需要有 Go 语言工具链和 Linux 下常用工具。后期会提供 Docker 内构建方式。

克隆代码后，执行 `make` 命令即可。

## 相关项目

[jeessy2/ddns-go](https://github.com/jeessy2/ddns-go) 本项目的上游，DDNS 任务由其完成

[panabit-ttyd](https://github.com/szu17dmy/panabit-ttyd) 封装了一个 Web 终端模拟器的插件

## 授权许可

本项目使用 MIT 许可证。

Panabit 是北京派网软件有限公司的品牌/商标，本项目中的引用仅为辅助说明，代码等内容均与官方无关。
