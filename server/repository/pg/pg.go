package pg

import (
	"context"
	"errors"
	"log"
	"runtime"
	"sync"
	"time"

	"github.com/jackc/pgx/v4/pgxpool"
)

type PGRepo struct {
	pool *pgxpool.Pool

	mux sync.Mutex
	ok  bool
}

func New(connstr string) (*PGRepo, error) {
	if connstr == "" {
		return nil, errors.New("пустая строка подключения к БД")
	}

	var repo PGRepo
	err := repo.SetDB(connstr)
	if err != nil {
		return nil, err
	}

	go repo.connChecker(connstr)

	return &repo, nil
}

func (repo *PGRepo) SetDB(connstr string) error {
	config, err := pgxpool.ParseConfig(connstr)
	if err != nil {
		return err
	}

	// Максимальное количество соединений с БД в пуле: 4 * количество_ядер.
	// По умолчанию: 1 * количество_ядер.
	config.MaxConns = int32(runtime.NumCPU() * 4)
	db, err := pgxpool.ConnectConfig(context.Background(), config)
	if err != nil {
		return err
	}

	repo.pool = db
	repo.mux.Lock()
	repo.ok = true
	repo.mux.Unlock()

	return nil
}

func (repo *PGRepo) connChecker(connstr string) {
	for {
		// Если произошло отключение от БД,
		// то пытаемся переподключиться.
		if !repo.ok {
			err := repo.SetDB(connstr)
			if err != nil {
				log.Println(err)
				time.Sleep(time.Second * 3)
				continue
			}
		}
		rows, err := repo.pool.Query(context.Background(), `select where 1 = 0`)
		if err != nil {
			repo.mux.Lock()
			repo.ok = false
			repo.mux.Unlock()
			log.Println(err)
			time.Sleep(time.Second * 5)
		}
		rows.Close()
		time.Sleep(time.Second * 3)
	}
}
