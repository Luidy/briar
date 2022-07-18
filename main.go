package main

import (
	"briar/config"
	"briar/db"
	"briar/server"
	"fmt"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"net"
	"os"
	"os/signal"
	"runtime"
	"syscall"
)

const (
	banner = "\n" +
		"#####  #####  #   ##   ##### \n" +
		"#    # #    # #  #  #  #    #  \n" +
		"#####  #    # # #    # #    #  \n" +
		"#    # #####  # ###### ##### \n" +
		"#    # #    # # #    # #    # \n\n" +
		"###### #    # # #    # #    # \n\n" +
		"=> Starting briar server %s\n\n"
)

func init() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	fmt.Printf("server CPU core: %d\n", runtime.NumCPU())
}

func main() {
	//ctx := context.Background()

	logrus.SetFormatter(&logrus.JSONFormatter{})
	log := logrus.StandardLogger()

	setting := config.NewSetting()

	database := db.InitDB(db.NewSetting(
		setting.DBHost,
		setting.DBPort,
		setting.DBUser,
		setting.DBPassword,
		setting.DBName))

	defer func() {
		if err := database.Close(); err != nil {
			log.Error(err)
		}
	}()

	cfg := config.NewConfig(setting, database)

	server, err := server.InitGRPCServer(cfg)
	if err != nil {
		log.Fatal(err)
	}

	go func() {
		l, err := net.Listen("tcp", ":"+setting.GRPCPort)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf(banner, fmt.Sprintf("%s:%s", setting.Host, setting.GRPCPort))
		if err := server.Serve(l); err != nil && err != grpc.ErrServerStopped {
			log.Fatal(err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	<-quit

	log.Info("Stopping user server")
	server.GracefulStop()
}
