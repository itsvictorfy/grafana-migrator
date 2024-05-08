package grafanamigrator

import (
	"log"
	"os"

	"github.com/itsvictorfy/grafana-migrator/cmd"
	"github.com/itsvictorfy/grafana-migrator/pkg/migrator"
	"github.com/joho/godotenv"
)

func initDev() string {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Access environment variables
	myEnvVar := os.Getenv("MY_ENV_VAR")
	log.Println("Value of MY_ENV_VAR:", myEnvVar)
	return myEnvVar
}
func initDocker() (string, string, string) {
	sourceUrl := os.Getenv("SOURCE_URL")
	destUrl := os.Getenv("DEST_URL")
	apiToken := os.Getenv("API_TOKEN")
	return sourceUrl, destUrl, apiToken
}
func main() {

	cmd.Execute()
	migrator.GetDashboardList()
}
