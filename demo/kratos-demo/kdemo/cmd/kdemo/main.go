package main

import (
	"flag"
	"fmt"
	"os"

	"kdemo/internal/conf"

	// consul "github.com/go-kratos/consul/registry"
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/config"
	"github.com/go-kratos/kratos/v2/config/file"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/go-kratos/kratos/v2/transport/http"
	"github.com/google/uuid"
	"github.com/hashicorp/consul/api"
)

// go build -ldflags "-X main.Version=x.y.z"
var (
	// Name is the name of the compiled software.
	Name string
	// Version is the version of the compiled software.
	Version string
	// flagconf is the config flag.
	flagconf string

	id, _ = os.Hostname()
)

func init() {
	flag.StringVar(&flagconf, "conf", "../../configs", "config path, eg: -conf config.yaml")
}

func newApp(logger log.Logger, hs *http.Server, gs *grpc.Server) *kratos.App {

	app := kratos.New(
		kratos.ID(id),
		kratos.Name(Name),
		kratos.Version(Version),
		kratos.Metadata(map[string]string{}),
		kratos.Logger(logger),
		kratos.Server(
			hs,
			gs,
		),
		// kratos.Registrar(consul.New(client)),
	)

	// client, err := api.NewClient(&api.Config{Address: "127.0.0.1:8500"})

	return app
}

func main() {

	flag.Parse()
	logger := log.With(log.NewStdLogger(os.Stdout),
		"ts", log.DefaultTimestamp,
		"caller", log.DefaultCaller,
		"service.id", id,
		"service.name", Name,
		"service.version", Version,
		"trace_id", log.TraceID(),
		"span_id", log.SpanID(),
	)
	c := config.New(
		config.WithSource(
			file.NewSource(flagconf),
		),
	)
	if err := c.Load(); err != nil {
		panic(err)
	}

	var bc conf.Bootstrap
	if err := c.Scan(&bc); err != nil {
		panic(err)
	}

	app, cleanup, err := initApp(bc.Server, bc.Data, logger)
	if err != nil {
		panic(err)
	}
	defer cleanup()

	// client, err := api.NewClient(api.DefaultConfig())
	client, err := api.NewClient(&api.Config{Address: "127.0.0.1:8500", Datacenter: "dc1"})

	serviceId := uuid.NewString()
	client.Agent().ServiceRegister(&api.AgentServiceRegistration{
		Address: "127.0.0.1",
		Port:    9000,
		Name:    "greetService1",
		ID:      serviceId,
	})

	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println("consul client created.")
	}

	defer func() {
		client.Agent().ServiceDeregister(serviceId)
	}()

	// start and wait for stop signal
	if err := app.Run(); err != nil {
		panic(err)
	}
}
