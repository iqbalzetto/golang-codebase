package main

import (
	"os"

	"github.com/iqbalzetto/golang-codebase/cmd/golang-codebase/startup"
	"github.com/iqbalzetto/golang-codebase/cmd/golang-codebase/startup/database"
	"github.com/iqbalzetto/golang-codebase/cmd/golang-codebase/startup/router"
)

/*
|--------------------------------------------------------------------------
| A. Check If The Application Is Under Maintenance
|--------------------------------------------------------------------------
|
| If the application is in maintenance / demo mode via the "down" command
| we will load this file so that any pre-rendered content can be shown
| instead of starting the framework, which could cause an exception.
|
*/

/*
|--------------------------------------------------------------------------
| B. Run The Application
|--------------------------------------------------------------------------
|
| Once we have the application, we can handle the incoming request using
| the application's HTTP kernel. Then, we will send the response back
| to this client's browser, allowing them to enjoy our application.
|
*/

func main() {
	startup.SetupEnv()

	database.SetupPostgres()

	ginRouter := router.GinRouter()
	ginRouter.Run("localhost:" + os.Getenv("PORT"))
}

/*
|--------------------------------------------------------------------------
| E. APM
|--------------------------------------------------------------------------
|
| Initialize the APM agent. This will be used to monitor the application's
| performance and report any issues that occur during the application's
| execution.
|
*/

/*
|--------------------------------------------------------------------------
| D. Logger
|--------------------------------------------------------------------------
|
| Initialize the logger. This will be used to log any errors or debug
| information that occurs during the application's execution.
|
*/

/*
|--------------------------------------------------------------------------
| E. Rate Limiter
|--------------------------------------------------------------------------
|
| Start the HTTP server. This will listen for incoming requests and
| dispatch them to the application's HTTP kernel for handling.
|
*/
