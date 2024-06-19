package startup

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

/*
|--------------------------------------------------------------------------
| C. Set up Environment Variables
|--------------------------------------------------------------------------
|
| Load the environment variables from the .env file. This allows us to
| configure the application's settings based on the environment where
| the application is running, such as local, staging, or production.
|
*/
func SetupEnv() {
	if os.Getenv("MODE") == "" {
		err := godotenv.Load(".env")
		if err != nil {
			panic(fmt.Sprintf("%s: %s", "Failed to load env", err))
		}
	}
}
