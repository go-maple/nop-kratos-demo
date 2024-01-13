package data

import (
	"context"
	"errors"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/middleware/tracing"
	"github.com/go-kratos/kratos/v2/registry"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/go-maple/nop-kratos-demo/blog/internal/biz"
	"github.com/go-maple/nop-kratos-demo/blog/internal/conf"
	vite "github.com/go-maple/nop-kratos-demo/server/api/server/v1"

	grpcx "google.golang.org/grpc"
)

type viteRepo struct {
	viteCli vite.GreeterClient
}

// sayHello implements biz.ViteRepo.
func (*viteRepo) sayHello(ctx context.Context, content string) (*biz.HelloMessage, error) {
	panic("unimplemented")
}

func NewViteRepo(conf *conf.Server, data *conf.Data, r registry.Discovery, logger log.Logger) biz.ViteRepo {
	conn, clean, err := getGrpcConn(conf, data, r, logger)
	if err != nil {
		log.Error(err)
		clean()
		return nil
	}
	return &viteRepo{
		viteCli: conn,
	}
}

// service join schema
func joinSchemaGRPC(serviceName string) string {
	return serviceName + ".grpc"
}

// newRpcClient new rpc client
func getGrpcConn(conf *conf.Server, data *conf.Data, r registry.Discovery, logger log.Logger) (vite.GreeterClient, func(), error) {

	log := log.NewHelper(logger)

	viteEndpoint := joinSchemaGRPC(data.Facade.ViteServer)

	log.Infof("grpc Timeout: %d", conf.Grpc.Timeout.AsDuration())

	vc, err := grpc.DialInsecure(
		context.Background(),
		grpc.WithEndpoint(viteEndpoint),
		grpc.WithDiscovery(r),
		grpc.WithMiddleware(
			recovery.Recovery(),
			tracing.Client(),
		),
		grpc.WithTimeout(conf.Grpc.Timeout.AsDuration()),
		grpc.WithOptions(grpcx.WithStatsHandler(&tracing.ClientHandler{})),
	)

	if err != nil {
		return nil, nil, err
	}

	log.Infof("viteEndpoint: %s", viteEndpoint)

	return vite.NewGreeterClient(vc), func() {
		_ = vc.Close()
	}, nil
}

// sayHello implements biz.ViteRepo.
func (v *viteRepo) SayHello(ctx context.Context, content string) (*biz.HelloMessage, error) {
	vite, err := v.viteCli.SayHello(ctx, &vite.HelloRequest{
		Name: content,
	})

	if err != nil {
		return nil, errors.New("request failed")
	}

	return &biz.HelloMessage{
		Message: vite.Message,
	}, err
}

func (v *viteRepo) SayHelloError(ctx context.Context, content string) (*biz.HelloMessage, error) {
	_, err := v.viteCli.SayHelloError(ctx, &vite.HelloRequest{
		Name: content,
	})
	return nil, err
}
