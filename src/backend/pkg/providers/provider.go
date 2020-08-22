package providers

import (
	"database/sql"
	"github.com/anggakes/gosvelte/src/backend/pkg/events"
	"github.com/anggakes/gosvelte/src/backend/pkg/usecases/todo"
	_ "github.com/lib/pq"
)

type Provider struct {
	Todo todo.ITodo
}

func InitProvider() *Provider {

	cfg := LoadConfig()

	// init DB
	DBServer := "host=" + cfg.DBHost + " port=" + cfg.DBPort + " user=" + cfg.DBUser + " dbname=" + cfg.DBName + " password=" + cfg.DBPassword + " sslmode=disable"
	db, err := sql.Open("postgres", DBServer)
	if err != nil {
		panic("failed connect DB : " + err.Error())
	}

	if err := db.Ping(); err != nil {
		panic("failed connect DB : " + err.Error())
	}

	ev := events.New()

	td := todo.New(db, ev)

	return &Provider{
		Todo: td,
	}
}
