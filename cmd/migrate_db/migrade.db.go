package main

import (
	"context"
	"flag"
	"fmt"
	"time"

	"github.com/Jm-Zion/troc-bike-go/app"
	"github.com/Jm-Zion/troc-bike-go/xconfig"
	"github.com/go-pg/migrations/v8"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
)

const stmtTimeout = 5 * time.Minute

func main() {

	godotenv.Load()

	flag.Parse()

	cfg, err := xconfig.LoadConfig("migrate_db")
	if err != nil {
		logrus.Fatal(err)
	}

	ctx := app.Init(context.Background(), cfg)
	defer app.Exit(ctx)

	args := flag.Args()
	oldVersion, newVersion, err := migrations.Run(app.PGMain(), args...)
	if err != nil {
		logrus.Fatalf("migration %d -> %d failed: %s",
			oldVersion, newVersion, err)
	}

	if newVersion != oldVersion {
		fmt.Printf("migrated from %d to %d\n", oldVersion, newVersion)
	} else {
		fmt.Printf("version is %d\n", oldVersion)
	}
}
