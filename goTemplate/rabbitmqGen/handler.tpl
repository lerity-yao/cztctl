package {{.PkgName}}

import (
    "context"
	{{.ImportPackages}}
)

func {{.HandlerName}}(ctx context.Context, c rabbitmq.RabbitListenerConf, svcCtx *svc.ServiceContext) service.Service {
	return rabbitmq.MustNewListener(ctx, c, {{.PkgName}}.New{{.LogicName}}(ctx, svcCtx))
}