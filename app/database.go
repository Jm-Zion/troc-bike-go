package app

import (
	"context"
	"os"
	"sync"

	"github.com/go-pg/pg/v10"
	"github.com/sirupsen/logrus"
)

var (
	pgMainOnce sync.Once
	pgMain     *pg.DB
)

func PGMain() *pg.DB {
	pgMainOnce.Do(func() {
		pgMain = NewPostgres()
	})
	return pgMain
}

var (
	pgMainTxOnce sync.Once
	pgMainTx     *pg.DB
)

func PGMainTx() *pg.DB {
	pgMainTxOnce.Do(func() {
		pgMainTx = NewPostgres()
	})
	return pgMainTx
}

func NewPostgres() *pg.DB {
	opt, err := pg.ParseURL(os.Getenv("DATABASE_URL"))

	if err != nil {
		logrus.WithError(err).Error("db failed to get options")
	}

	db := pg.Connect(opt)
	OnExitSecondary(func(ctx context.Context) {
		if err := db.Close(); err != nil {
			logrus.WithError(err).Error("pg.Close failed")
		}
	})

	//db.AddQueryHook(pgext.OpenTelemetryHook{})
	// if IsDebug() {
	// 	db.AddQueryHook(pgext.DebugHook{})
	// }

	return db
}
