package main

import (
	"flag"

	"log"

	"github.com/joho/godotenv"
	"github.com/yaaaaashiki/cstack"
)

func main() {
	EnvLoad()
	var (
		addr   = flag.String("addr", ":8080", "addr to bind")
		dbconf = flag.String("dbconf", "dbconfig.yml", "database configuration file.")
		env    = flag.String("env", "development", "application envirionment (production, development etc.)")
		debug  = flag.Bool("debug", false, "debug mode. default is false.")
	)
	flag.Parse()
	b := cstack.New()
	b.Init(*dbconf, *env, *debug)
	b.Run(*addr)
}

func EnvLoad() {
	err := godotenv.Load()
	if err != nil {
		log.Println("Error loading .env file")
	}
}
