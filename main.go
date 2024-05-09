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
	err := godotenv.Load()
	if err != nil {
		log.Printf("Error loading .env file")
		return
	}
	source.Url = os.Getenv("SOURCE_URL")
	source.ApiToken = os.Getenv("SOURCE_API_TOKEN")
	dest.Url = os.Getenv("DEST_URL")
	dest.ApiToken = os.Getenv("DEST_API_TOKEN")

	// Access environment variables
}
func initDocker() {
	var source migrator.Grafana
	var dest migrator.Grafana
	source.Url = os.Getenv("SOURCE_URL")
	source.ApiToken = os.Getenv("SOURCE_API_TOKEN")
	dest.Url = os.Getenv("DEST_URL")
	dest.ApiToken = os.Getenv("DEST_API_TOKEN")
}

func init() {
	env := os.Getenv("ENV")
	if env == "docker" {
		initDocker()
	} else {
		initDev()
	}
	if source.Url == "" || source.ApiToken == "" || dest.Url == "" || dest.ApiToken == "" {
		initFlags()
	}
}
func initFlags() {
	flag.StringVar(&source.Url, "source", "", "Source Grafana URL")
	flag.StringVar(&source.ApiToken, "sourceToken", "", "Source Grafana API Token")
	flag.StringVar(&dest.Url, "target", "", "Target Grafana URL")
	flag.StringVar(&dest.ApiToken, "targetToken", "", "Target Grafana API Token")
	versionFlag := flag.Bool("version", false, "Print the version number")
	flag.Parse()
	if *versionFlag {
		fmt.Println("Version: 1.0.0") // Update with your version number
		os.Exit(0)
	}
}

func main() {
	versionFlag := flag.Bool("version", false, "Print the version number")
	flag.Parse()
	if *versionFlag {
		fmt.Println("Version: 1.0.0") // Update with your version number
		os.Exit(0)
	}
	fmt.Printf("source %s, %s", source.Url, source.ApiToken)
	fmt.Printf("dest %s, %s", dest.Url, dest.ApiToken)
}
