package store

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

type ConnectionInfo struct {
	Host     string
	Username string
	Password string
	Port     int
	DB       string
}

type Postgres struct {
	connectionInfo ConnectionInfo
	db             *sql.DB
}

func NewPostgres(host string, username string, password string, port int, db string) *Postgres {
	connectionInfo := ConnectionInfo{
		Host:     "192.168.1.99",
		Port:     5432,
		Username: "murr",
		Password: "flag9012",
		DB:       "cbot",
	}
	return &Postgres{connectionInfo: connectionInfo}
}

func (p *Postgres) Connect() error {
	connString := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		p.connectionInfo.Host,
		p.connectionInfo.Port,
		p.connectionInfo.Username,
		p.connectionInfo.Password,
		p.connectionInfo.DB)

	db, err := sql.Open("postgres", connString)

	if err != nil {
		return err
	}
	p.db = db
	return nil
}

func (p *Postgres) Close() error {
	return p.db.Close()
}

func (p *Postgres) Query(args ...string) (interface{}, error) {

	return nil, nil
}

func (p *Postgres) Insert(args ...string) error {
	statement := `insert into prices(price, time, productid) values($1, $2, $3)`
	_, err := p.db.Exec(statement, args[0], args[1], args[2])
	if err != nil {
		return err
	}
	return nil
}
