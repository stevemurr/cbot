package store

import (
	"database/sql"
	"fmt"
	"log"

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

func NewPostgres(username string, password string, db string) *Postgres {
	connectionInfo := ConnectionInfo{
		Host:     "db",
		Port:     5432,
		Username: username,
		Password: password,
		DB:       db,
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
	log.Println(connString)
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

func (p *Postgres) Exec(statement string) error {
	if _, err := p.db.Exec(statement); err != nil {
		return err
	}
	return nil
}

func (p *Postgres) Insert(args ...string) error {
	statement := `insert into prices(price, time, productid) values($1, $2, $3)`
	_, err := p.db.Exec(statement, args[0], args[1], args[2])
	if err != nil {
		return err
	}
	return nil
}
