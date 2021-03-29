package main

import (
	"fmt"
	"runtime"
)

const (
	banner = "\n" +   
                                                              
"        GGGGGGGGGGGGGEEEEEEEEEEEEEEEEEEEEEE   SSSSSSSSSSSSSSS   \n"+
"     GGG::::::::::::GE::::::::::::::::::::E SS:::::::::::::::S  \n"+
"   GG:::::::::::::::GE::::::::::::::::::::ES:::::SSSSSS::::::S  \n"+
"  G:::::GGGGGGGG::::GEE::::::EEEEEEEEE::::ES:::::S     SSSSSSS  \n"+
" G:::::G       GGGGGG  E:::::E       EEEEEES:::::S              \n"+
"G:::::G                E:::::E             S:::::S              \n"+
"G:::::G                E::::::EEEEEEEEEE    S::::SSSS           \n"+
"G:::::G    GGGGGGGGGG  E:::::::::::::::E     SS::::::SSSSS      \n"+
"G:::::G    G::::::::G  E:::::::::::::::E       SSS::::::::SS    \n"+
"G:::::G    GGGGG::::G  E::::::EEEEEEEEEE          SSSSSS::::S   \n"+
"G:::::G        G::::G  E:::::E                         S:::::S  \n"+
" G:::::G       G::::G  E:::::E       EEEEEE            S:::::S  \n"+
"  G:::::GGGGGGGG::::GEE::::::EEEEEEEE:::::ESSSSSSS     S:::::S  \n"+
"   GG:::::::::::::::GE::::::::::::::::::::ES::::::SSSSSS:::::S  \n"+
"     GGG::::::GGG:::GE::::::::::::::::::::ES:::::::::::::::SS   \n"+
"        GGGGGG   GGGGEEEEEEEEEEEEEEEEEEEEEE SSSSSSSSSSSSSSS   \n\n"
                                                              
)

func init() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	fmt.Printf("server CPU core: %d\n", runtime.NumCPU())
}

func main() {
	fmt.Println(banner)
}
