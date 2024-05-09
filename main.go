package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/itsvictorfy/grafana-migrator/pkg/migrator"
	"github.com/joho/godotenv"
)

var (
	source migrator.Grafana
	dest   migrator.Grafana
)

func initDev() {
	// var source migrator.Grafana
	// var dest migrator.Grafana
	err := godotenv.Load()
	if err != nil {
		log.Printf(".env file doesnt exists")
		return
	}
	source.Url = os.Getenv("SOURCE_URL")
	source.ApiToken = os.Getenv("SOURCE_API_TOKEN")
	dest.Url = os.Getenv("DEST_URL")
	dest.ApiToken = os.Getenv("DEST_API_TOKEN")
	log.Printf("Succesfully loaded vars .env file")
	if source.Url == "" || source.ApiToken == "" || dest.Url == "" || dest.ApiToken == "" {
		log.Fatal("Source and target URLs and tokens must be provided")
		os.Exit(3)
	}

}

func initDocker() {

	source.Url = os.Getenv("SOURCE_URL")
	source.ApiToken = os.Getenv("SOURCE_API_TOKEN")
	dest.Url = os.Getenv("DEST_URL")
	dest.ApiToken = os.Getenv("DEST_API_TOKEN")
	if source.Url == "" || source.ApiToken == "" || dest.Url == "" || dest.ApiToken == "" {
		log.Fatal("Source and target URLs and tokens must be provided")
		os.Exit(3)
	}
	log.Printf("Succesfully loaded vars .env file")
}

func initvars() {
	env := os.Getenv("ENV")
	if env != "docker" {
		log.Printf("Not Docker-Getting vars from .env")
		initDev()
	} else {
		log.Printf("Docker-Getting vars from Docker env Vars")
		initDocker()
	}
}

func main() {
	sc := flag.String("source", "", "a string")
	ds := flag.String("target", "", "a string")
	sckey := flag.String("sourceToken", "", "a string")
	dskey := flag.String("targetToken", "", "a string")
	ver := flag.Bool("version", false, "a bool")
	flag.Parse()
	if *ver {
		fmt.Println("v0.0.1")
		os.Exit(0)
	}
	if *sc == "" || *ds == "" || *sckey == "" || *dskey == "" {
		log.Printf("no flags were captured starting dev/docker initialization")
		initvars()
	} else {
		source.Url = *sc
		source.ApiToken = *sckey
		dest.Url = *ds
		dest.ApiToken = *dskey
		fmt.Printf("source %s, %s\n", source.Url, source.ApiToken)
		fmt.Printf("dest %s, %s\n", dest.Url, dest.ApiToken)
	}
	fmt.Printf("source %s, %s\n", source.Url, source.ApiToken)
	fmt.Printf("dest %s, %s\n", dest.Url, dest.ApiToken)
}
