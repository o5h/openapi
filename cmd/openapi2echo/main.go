package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	fmt.Println(`

  ██████╗ ██████╗ ███████╗███╗   ██╗ █████╗ ██████╗ ██╗██████╗ ███████╗ ██████╗██╗  ██╗ ██████╗ 
 ██╔═══██╗██╔══██╗██╔════╝████╗  ██║██╔══██╗██╔══██╗██║╚════██╗██╔════╝██╔════╝██║  ██║██╔═══██╗
 ██║   ██║██████╔╝█████╗  ██╔██╗ ██║███████║██████╔╝██║ █████╔╝█████╗  ██║     ███████║██║   ██║
 ██║   ██║██╔═══╝ ██╔══╝  ██║╚██╗██║██╔══██║██╔═══╝ ██║██╔═══╝ ██╔══╝  ██║     ██╔══██║██║   ██║
 ╚██████╔╝██║     ███████╗██║ ╚████║██║  ██║██║     ██║███████╗███████╗╚██████╗██║  ██║╚██████╔╝
  ╚═════╝ ╚═╝     ╚══════╝╚═╝  ╚═══╝╚═╝  ╚═╝╚═╝     ╚═╝╚══════╝╚══════╝ ╚═════╝╚═╝  ╚═╝ ╚═════╝ 
		`)

	configFile := flag.String("config", "openapi2echo.yaml", "configuration file")
	flag.Parse()

	f, err := os.Open(*configFile)
	exitOnError(err)
	defer f.Close()

	cfgYaml, err := io.ReadAll(f)
	exitOnError(err)

	log.Println("openapi2echo config:", string(cfgYaml))

}

func exitOnError(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
