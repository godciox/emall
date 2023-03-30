package main

import (
	"context"
	_const "frontend/const"
	"frontend/routers"
	"frontend/service"
	"frontend/utils/tracing"
	"github.com/beego/beego/v2/server/web"
	"net/http"
	"os"
	"time"

	mgrpc "github.com/go-micro/plugins/v4/client/grpc"
	_ "github.com/go-micro/plugins/v4/registry/kubernetes"
	mhttp "github.com/go-micro/plugins/v4/server/http"
	"github.com/sirupsen/logrus"
	"go-micro.dev/v4"
	"go-micro.dev/v4/logger"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/propagation"

	"frontend/config"
	pb "frontend/proto"
)

type ctxKeySessionID struct{}

func init() {
	routers.HttpServer = web.BeeApp
}

func main() {
	// Load conigurations
	if err := config.Load(); err != nil {
		logger.Fatal(err)
	}

	// Create service
	srv := micro.NewService(
		micro.Server(mhttp.NewServer()),
		micro.Client(mgrpc.NewClient()),
		micro.Registry(config.NewRegistry()),
	)
	opts := []micro.Option{
		micro.Name(_const.Name),
		micro.Version(_const.Version),
		micro.Address(config.Address()),
	}
	if cfg := config.Tracing(); cfg.Enable {
		tp, err := tracing.NewTracerProvider(_const.Name, srv.Server().Options().Id, cfg.Jaeger.URL)
		if err != nil {
			logger.Fatal(err)
		}
		defer func() {
			ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
			defer cancel()
			if err := tp.Shutdown(ctx); err != nil {
				logger.Fatal(err)
			}
		}()
		otel.SetTracerProvider(tp)
		otel.SetTextMapPropagator(propagation.NewCompositeTextMapPropagator(propagation.TraceContext{}, propagation.Baggage{}))
	}
	srv.Init(opts...)

	log := logrus.New()
	log.Level = logrus.DebugLevel
	log.Formatter = &logrus.JSONFormatter{
		FieldMap: logrus.FieldMap{
			logrus.FieldKeyTime:  "timestamp",
			logrus.FieldKeyLevel: "severity",
			logrus.FieldKeyMsg:   "message",
		},
		TimestampFormat: time.RFC3339Nano,
	}
	log.Out = os.Stdout

	cfg, client := config.Get(), srv.Client()
	service.Svc = &service.FrontendServer{
		EmailService:     pb.NewEmailService(cfg.EmailService, client),
		UserService:      pb.NewUserService(cfg.UserService, client),
		OrderService:     pb.NewOrderService(cfg.OrderService, client),
		InventoryService: pb.NewInventoryService(cfg.InventoryService, client),
		CartService:      pb.NewCartService(cfg.CartService, client),
		ProductService:   pb.NewProductCatalogService(cfg.ProductService, client),
	}

	// 路由建立
	routers.Route()

	var handler http.Handler = routers.HttpServer.Handlers

	if err := micro.RegisterHandler(srv.Server(), handler); err != nil {
		logger.Fatal(err)
	}

	logger.Infof("starting server on %s", config.Address())
	if err := srv.Run(); err != nil {
		logger.Fatal(err)
	}
}
