package main

import (
	"fmt"
	"ges/config"
	"log"
	"net"
	"os"
	"runtime"

	"github.com/spf13/viper"
	"google.golang.org/grpc"
)

const (
	banner = "\n" +

		"        GGGGGGGGGGGGGEEEEEEEEEEEEEEEEEEEEEE   SSSSSSSSSSSSSSS   \n" +
		"     GGG::::::::::::GE::::::::::::::::::::E SS:::::::::::::::S  \n" +
		"   GG:::::::::::::::GE::::::::::::::::::::ES:::::SSSSSS::::::S  \n" +
		"  G:::::GGGGGGGG::::GEE::::::EEEEEEEEE::::ES:::::S     SSSSSSS  \n" +
		" G:::::G       GGGGGG  E:::::E       EEEEEES:::::S              \n" +
		"G:::::G                E:::::E             S:::::S              \n" +
		"G:::::G                E::::::EEEEEEEEEE    S::::SSSS           \n" +
		"G:::::G    GGGGGGGGGG  E:::::::::::::::E     SS::::::SSSSS      \n" +
		"G:::::G    G::::::::G  E:::::::::::::::E       SSS::::::::SS    \n" +
		"G:::::G    GGGGG::::G  E::::::EEEEEEEEEE          SSSSSS::::S   \n" +
		"G:::::G        G::::G  E:::::E                         S:::::S  \n" +
		" G:::::G       G::::G  E:::::E       EEEEEE            S:::::S  \n" +
		"  G:::::GGGGGGGG::::GEE::::::EEEEEEEE:::::ESSSSSSS     S:::::S  \n" +
		"   GG:::::::::::::::GE::::::::::::::::::::ES::::::SSSSSS:::::S  \n" +
		"     GGG::::::GGG:::GE::::::::::::::::::::ES:::::::::::::::SS   \n" +
		"        GGGGGG   GGGGEEEEEEEEEEEEEEEEEEEEEE SSSSSSSSSSSSSSS   \n\n" +
		"=> Starting gRPC server %s\n\n"
)

func init() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	fmt.Printf("server CPU core: %d\n", runtime.NumCPU())
}

func main() {
	ges := config.Ges
	g := grpcInit()
	startServer(ges, g)
}

func grpcInit() (g *grpc.Server) {
	g = grpc.NewServer()
	return g
}

func startServer(ges *viper.Viper, g *grpc.Server) {
	address := fmt.Sprintf("0.0.0.0:%d", ges.GetInt("port"))
	fmt.Printf(banner, address)

	l, err := net.Listen("tcp", address)
	if err != nil {
		log.Println("End to bind for gRPC server", "err", err)
		os.Exit(1)
	}
	if err := g.Serve(l); err != nil {
		log.Println("End gRPC server", "err", err)
		os.Exit(1)
	}

}
