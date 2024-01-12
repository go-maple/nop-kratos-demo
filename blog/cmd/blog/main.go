package main

import (
	"flag"
	"os"

	nacosconfig "github.com/go-kratos/kratos/contrib/config/nacos/v2"
	"github.com/go-maple/nop-kratos-demo/blog/internal/conf"
	"github.com/nacos-group/nacos-sdk-go/clients"
	"github.com/nacos-group/nacos-sdk-go/clients/config_client"
	"github.com/nacos-group/nacos-sdk-go/common/constant"
	"github.com/nacos-group/nacos-sdk-go/vo"
	"github.com/spf13/cast"
	"gopkg.in/yaml.v2"

	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/config"
	"github.com/go-kratos/kratos/v2/config/file"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/tracing"
	"github.com/go-kratos/kratos/v2/registry"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/go-kratos/kratos/v2/transport/http"
)

// go build -ldflags "-X main.Version=x.y.z"
var (
	// Name is the name of the compiled software.
	Name string = "blog"
	// Version is the version of the compiled software.
	Version string
	// flagconf is the config flag.
	flagconf string

	id, _ = os.Hostname()

	nacosHost, nacosPort, nacosNamespace, nacosUser, nacosPassword string
)

func init() {
	flag.StringVar(&flagconf, "conf", "../../configs", "config path, eg: -conf config.yaml")
	nacosHost = os.Getenv("REGISTER_HOST")
	nacosPort = os.Getenv("REGISTER_PORT")
	nacosNamespace = os.Getenv("REGISTER_NAMESPACE")
	nacosUser = os.Getenv("REGISTER_USERNAME")
	nacosPassword = os.Getenv("REGISTER_PASSWORD")
}

func newApp(logger log.Logger, hs *http.Server, gs *grpc.Server, registry registry.Registrar) *kratos.App {
	return kratos.New(
		kratos.ID(id),
		kratos.Name(Name),
		kratos.Version(Version),
		kratos.Metadata(map[string]string{}),
		kratos.Logger(logger),
		kratos.Server(
			hs,
			gs,
		),
		kratos.Registrar(registry),
	)
}

// Set global trace provider
func setTracerProvider(url string) error {
	// Create the Jaeger exporter
	// exp, err := jaeger.New(jaeger.WithCollectorEndpoint(jaeger.WithEndpoint(url)))
	// if err != nil {
	// 	return err
	// }
	// tp := tracesdk.NewTracerProvider(
	// 	// Set the sampling rate based on the parent span to 100%
	// 	tracesdk.WithSampler(tracesdk.ParentBased(tracesdk.TraceIDRatioBased(1.0))),
	// 	// Always be sure to batch in production.
	// 	tracesdk.WithBatcher(exp),
	// 	// Record information about this application in an Resource.
	// 	tracesdk.WithResource(resource.NewSchemaless(
	// 		semconv.ServiceNameKey.String(Name),
	// 		attribute.String("env", "dev"),
	// 	)),
	// )
	// otel.SetTracerProvider(tp)
	return nil
}

func main() {

	logger := log.With(log.NewStdLogger(os.Stdout),
		"ts", log.DefaultTimestamp,
		"caller", log.DefaultCaller,
		"service.id", id,
		"service.name", Name,
		"service.version", Version,
		"trace.id", tracing.TraceID(),
		"span.id", tracing.SpanID(),
	)

	bc := loadConfig()

	app, cleanup, err := initApp(Name, bc.Server, bc.Data, bc.Registry, logger)
	if err != nil {
		panic(err)
	}
	defer cleanup()

	// start and wait for stop signal
	if err := app.Run(); err != nil {
		panic(err)
	}
}

func loadConfig() *conf.Bootstrap {
	flag.Parse()
	if flagconf != "" {
		c := config.New(
			config.WithSource(
				file.NewSource(flagconf),
			),
		)
		defer c.Close()

		if err := c.Load(); err != nil {
			panic(err)
		}

		var bc = &conf.Bootstrap{}
		if err := c.Scan(bc); err != nil {
			panic(err)
		}

		return bc
	}

	return loadNacosConfig()
}

func loadNacosConfig() *conf.Bootstrap {
	client, err := getNacosConfigClient()
	if err != nil {
		return nil
	}

	configSource := nacosconfig.NewConfigSource(client, nacosconfig.WithGroup("DEFAULT_GROUP"), nacosconfig.WithDataID(Name))

	c := config.New(
		config.WithSource(
			configSource,
		),
		config.WithDecoder(func(kv *config.KeyValue, v map[string]interface{}) error {
			return yaml.Unmarshal(kv.Value, v)
		}),
	)

	if err = c.Load(); err != nil {
		return nil
	}

	bc := &conf.Bootstrap{}
	if err = c.Scan(bc); err != nil {
		return nil
	}

	return bc
}

func getNacosServerAndClientConfig() ([]constant.ServerConfig, constant.ClientConfig) {
	return []constant.ServerConfig{
			*constant.NewServerConfig(nacosHost, cast.ToUint64(nacosPort)),
		},
		constant.ClientConfig{
			AppName:             Name,
			NamespaceId:         nacosNamespace, //namespace id
			TimeoutMs:           5000,
			NotLoadCacheAtStart: true,
			LogLevel:            "error",
			Username:            nacosUser,
			Password:            nacosPassword,
		}
}

func getNacosConfigClient() (config_client.IConfigClient, error) {
	sc, cc := getNacosServerAndClientConfig()
	// a more graceful way to create naming client
	client, err := clients.NewConfigClient(
		vo.NacosClientParam{
			ClientConfig:  &cc,
			ServerConfigs: sc,
		},
	)

	if err != nil {
		return nil, err
	}
	return client, nil
}
