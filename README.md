## grpc-gateway

### 安装

```bash
go install \
    github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway@latest \
    github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2@latest
```

### 生成代码

```bash
protoc -I ./proto --grpc-gateway_out=./pb --grpc-gateway_opt grpc_api_configuration=./proto/trip.yaml ./proto/trip.proto
```

### yaml 定义

```yaml
type: google.api.Service
config_version: 3

http:
  rules:
    - selector: com.github.qiyue365.coolcar.TripService.GetTrip
      get: /trip/{id}
```
