package main

import (
	"flag"
	"fmt"
	"os"
	"time"

	"bubblelist/internal/conf"

	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/config"
	"github.com/go-kratos/kratos/v2/config/env"
	"github.com/go-kratos/kratos/v2/config/file"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/tracing"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/go-kratos/kratos/v2/transport/http"

	kratoszap "github.com/go-kratos/kratos/contrib/log/zap/v2"
	_ "go.uber.org/automaxprocs"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
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

func newApp(logger log.Logger, gs *grpc.Server, hs *http.Server) *kratos.App {
	return kratos.New(
		kratos.ID(id),
		kratos.Name(Name),
		kratos.Version(Version),
		kratos.Metadata(map[string]string{}),
		kratos.Logger(logger),
		kratos.Server(
			gs,
			hs,
		),
	)
}

func main() {
	flag.Parse()
	// 1.将日志写入文件里
	f, _ := os.OpenFile("test.log", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0666)

	// 2.zap 日志库
	// Setting up a global logger with zap log

	f, err := os.OpenFile("test.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		return
	}
	writeSyncer := zapcore.AddSync(f)

	encoder := zapcore.NewJSONEncoder(zap.NewProductionEncoderConfig())
	core := zapcore.NewCore(encoder, writeSyncer, zapcore.DebugLevel)
	z := zap.New(core)

	// Print log
	log.Info("info")
	log.Debug("debug")

	log.NewStdLogger(f)
	// logger := log.With(log.NewStdLogger(os.Stdout),
	// logger := log.With(log.NewStdLogger(f),
	logger := log.With(kratoszap.NewLogger(z),
		"ts", log.DefaultTimestamp,
		"caller", log.DefaultCaller,
		"service.id", id,
		"service.name", Name,
		"service.version", Version,
		"trace.id", tracing.TraceID(),
		"span.id", tracing.SpanID(),
	)

	logger.Log(log.LevelDebug, "msg", "logger init success") //2.KRATOS不推荐使用Logger,不够方便
	// 3.KRATOS推荐使用Helper
	helper := log.NewHelper(log.NewFilter(logger,
		log.FilterKey("password")))
	helper.Debug("logger init success")
	helper.Infof("today is %v", time.Now())
	helper.Warnw("name", "qimi", "age", 18, "password", "1234567") //password 将被过滤
	// 4.设置全局字段
	/*
			logger := log.With(log.NewStdLogger(f),
			"ts", log.DefaultTimestamp,
			"caller", log.DefaultCaller,
			"service.id", id,
			"service.name", Name,
			"service.version", Version,
			"trace.id", tracing.TraceID(),
			"span.id", tracing.SpanID(),
		)
	*/

	// 5.日志过滤
	/*
		FileterLevel:等级过滤
		FilterKey:敏感字段过滤
		FilterValue:按照value过滤
		FilterFunc:自定义过滤规则
	*/

	// 创建配置对象
	c := config.New(
		config.WithSource(
			//指定配置的来源
			env.NewSource("BUBBLE_"),
			file.NewSource(flagconf),
		),
	)
	defer c.Close()

	// 加载配置(从配置文件/配置中心/环境变量加载配置)
	if err := c.Load(); err != nil {
		panic(err)
	}

	// 创建配置结构体变量bc
	var bc conf.Bootstrap

	// 将配置数据扫描到结构体变量bc中
	if err := c.Scan(&bc); err != nil {
		panic(err)
	}
	fmt.Println(bc.Server.Http)
	fmt.Println(bc.Server)
	fmt.Println(bc.Data)
	fmt.Println(bc.Mode)
	app, cleanup, err := wireApp(bc.Server, bc.Data, logger)
	if err != nil {
		panic(err)
	}
	defer cleanup()

	// start and wait for stop signal
	if err := app.Run(); err != nil {
		panic(err)
	}
}
