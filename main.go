package main

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/albabulil/go-myappEcho/app"
	"github.com/albabulil/go-myappEcho/config"
	dbmongo "github.com/albabulil/go-myappEcho/config/repository/db-mongo"
)

func main() {
	// load enviroment
	config.Init()

	/*--------------Start DB Connection--------------*/
	dbTimeout := time.Duration(3000)
	if timeout, err := strconv.Atoi(os.Getenv(config.DB_TIMEOUT)); err == nil {
		dbTimeout = time.Duration(timeout)
	}

	var dbURI string
	if config.DB_USER != "" || config.DB_PASS != "" {
		dbURI = fmt.Sprintf("mongodb://%s:%s@%s:%s", config.DB_NAME, config.DB_PASS, config.DB_HOST, config.DB_PORT)
	} else {
		dbURI = fmt.Sprintf("mongodb://%s:%s", config.DB_HOST, config.DB_PORT)
	}

	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout*time.Millisecond)
	defer cancel()

	client := dbmongo.InitDatabase(ctx, dbURI)
	defer client.Disconnect(ctx)
	/*--------------End DB Connection--------------*/

	// Starting App
	e := app.Init()

	e.Logger.Fatal(e.Start(":" + config.APP_PORT))
}
