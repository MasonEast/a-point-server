## a-point-server

a-point 埋点系统的服务端，基于`go-zero`搭建单体服务

### 命令

启动服务

```bash
# 创建API服务
goctl api new core
# 启动服务
go run core.go -f etc/core-api.yaml
# 使用api文件生成代码
goctl api go -api core.api -dir . -style go_zero
```
