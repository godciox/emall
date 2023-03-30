package main

import (
	"context"
	"database/sql"
	grpcc "github.com/go-micro/plugins/v4/client/grpc"
	_ "github.com/go-micro/plugins/v4/registry/kubernetes"
	grpcs "github.com/go-micro/plugins/v4/server/grpc"
	"github.com/go-micro/plugins/v4/wrapper/trace/opentelemetry"
	_ "github.com/go-sql-driver/mysql"
	"go-micro.dev/v4"
	"go-micro.dev/v4/logger"
	"go-micro.dev/v4/server"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/propagation"
	"strings"
	"time"
	"userservice/config"
	_const "userservice/const"
	db "userservice/db/sqlc"
	"userservice/handler"
	pb "userservice/proto"
	"userservice/tracing"
)

func main() {
	if err := config.Load(); err != nil {
		logger.Fatal(err)
		return
	}

	// Create service
	srv := micro.NewService(
		micro.Server(grpcs.NewServer()),
		micro.Client(grpcc.NewClient()),
		micro.Registry(config.NewRegistry()),
	)
	opts := []micro.Option{
		micro.Name(_const.Name),
		micro.Version(_const.Version),
		micro.Address(config.Address()),
	}

	//config.Register()
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
		traceOpts := []opentelemetry.Option{
			opentelemetry.WithHandleFilter(func(ctx context.Context, r server.Request) bool {
				if e := r.Endpoint(); strings.HasPrefix(e, "Health.") {
					return true
				}
				return false
			}),
		}
		opts = append(opts, micro.WrapHandler(opentelemetry.NewHandlerWrapper(traceOpts...)))
	}
	srv.Init(opts...)

	// Register handler
	if err := pb.RegisterUserServiceHandler(srv.Server(), new(handler.UserService)); err != nil {
		logger.Fatal(err)
	}
	if err := pb.RegisterHealthHandler(srv.Server(), new(handler.Health)); err != nil {
		logger.Fatal(err)
	}

	// open db
	conn, err := sql.Open(config.DBDriver(), config.DBSource())
	if err != nil {
		logger.Fatalf("db can not open, because of %s", err)
	}

	db.DB = db.NewStore(conn)

	// Run service
	if err := srv.Run(); err != nil {
		logger.Fatal(err)
	}

}
