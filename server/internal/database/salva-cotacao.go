package database

import (
	"context"
	"errors"
	"time"

	"github.com/eltonCasacio/client-server-api/server/pkg"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func SalvarCotacaoSqlite(usdbrl pkg.USDBRL) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Millisecond*10)
	defer cancel()

	select {
	case <-ctx.Done():
		return errors.New("excedeu tempo limite para salvar cotacao")
	default:
		db, err := gorm.Open(sqlite.Open("cotacao.db"), &gorm.Config{})
		if err != nil {
			panic(err)
		}
		db.AutoMigrate(pkg.USDBRL{})
		err = db.Create(usdbrl).Error
		if err != nil {
			return err
		}
		return nil
	}
}
