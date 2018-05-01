package main

import (
	"flag"
	"github.com/avegao/gocondi"
	"github.com/sirupsen/logrus"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"github.com/heroku/rollrus"
	"net"
	"google.golang.org/grpc/reflection"
	"google.golang.org/grpc"
	pb "github.com/avegao/iot-fronius/resource/grpc"
	"github.com/avegao/iot-fronius/service"
	_ "github.com/lib/pq"
)

const (
	version = "1.0.1"
)

var (
	debug      = flag.Bool("debug", false, "Print debug logs")
	grcpPort   = flag.Int("port", 50000, "gRPC Server port. Default = 50000")
	buildDate  string
	commitHash string
	container  *gocondi.Container
	parameters map[string]interface{}
	server     *grpc.Server
)

func initContainer() {
	flag.Parse()

	parameters = map[string]interface{}{
		"build_date":  buildDate,
		"debug":       *debug,
		"commit_hash": commitHash,
		"version":     version,
	}

	logger := initLogger()
	gocondi.Initialize(logger)
	container = gocondi.GetContainer()

	for name, value := range parameters {
		container.SetParameter(name, value)
	}
}

func initLogger() *logrus.Logger {
	logLevel := logrus.InfoLevel
	environment := "release"
	log := logrus.New()

	if *debug {
		logLevel = logrus.DebugLevel
		environment = "debug"
	} else {
		hook := rollrus.NewHook(fmt.Sprintf("%v", parameters["rollbar_token"]), environment)
		log.Hooks.Add(hook)
	}

	logrus.SetFormatter(&logrus.JSONFormatter{})
	logrus.SetFormatter(&logrus.TextFormatter{})
	log.SetLevel(logLevel)

	return log
}

func initGrpc() {
	container.GetLogger().Debugf("initGrpc() - START")

	listen, err := net.Listen("tcp", fmt.Sprintf(":%d", *grcpPort))

	if err != nil {
		container.GetLogger().Fatalf("failed to listen: %v", err)
	}

	container.GetLogger().Debugf("gRPC listening in %d port", *grcpPort)

	server = grpc.NewServer()
	pb.RegisterFroniusServer(server, new(service.Fronius))
	reflection.Register(server)

	if err := server.Serve(listen); err != nil {
		container.GetLogger().Fatalf("failed to server: %v", err)
	}

	container.GetLogger().Debugf("initGrpc() - END")
}

func handleInterrupt() {
	gracefulStop := make(chan os.Signal)
	signal.Notify(gracefulStop, syscall.SIGTERM)
	signal.Notify(gracefulStop, syscall.SIGINT)
	go func() {
		<-gracefulStop
		powerOff()
	}()
}

func powerOff() {
	container.GetLogger().Infof("Shutting down...")
	server.Stop()

	os.Exit(0)
}

func main() {
	initContainer()
	handleInterrupt()

	logger := container.GetLogger()
	logger.Infof("IoT Fronius Service v%s started (commit %s, build date %s)", container.GetStringParameter("version"), container.GetStringParameter("commit_hash"), container.GetStringParameter("build_date"))

	initGrpc()
}
