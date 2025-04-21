package main

import (
	"authentication/data"
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	_ "github.com/jackc/pgconn"
	_ "github.com/jackc/pgx/v4"
	_ "github.com/jackc/pgx/v4/stdlib"
)

const webPort = 80

var counts int64

type Config struct {
	Repo   data.Repository
	Client *http.Client
}

func main() {
	log.Println("Starting authentication service")

	// connect to DB
	conn := connectToDB()
	if conn == nil {
		log.Panic("Can't connect to Postgres!")
	}

	// setup up config
	app := Config{
		Client: &http.Client{},
	}
	app.setupRepo(conn)

	// setup web server
	srv := &http.Server{
		Addr:    fmt.Sprintf(":%d", webPort),
		Handler: app.routes(),
	}

	err := srv.ListenAndServe()
	if err != nil {
		log.Panic(err)
	}
}

func openDB(dsn string) (*sql.DB, error) {
	db, err := sql.Open("pgx", dsn)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}

type dsn struct {
	host     string
	port     string
	user     string
	password string
	dbname   string
	sslmode  string
	timezone string
	timeout  string
}

func connectToDB() *sql.DB {
	dsn, err := loadDSN()
	if err != nil {
		log.Fatal("Error loading database configuration", err)
	}

	for {
		connection, err := openDB(dsn)
		if err != nil {
			log.Println("Postgres not yet ready...")
			counts++
		} else {
			log.Println("Connected to Postgres!")

			// Test with ping
			err := connection.Ping()
			if err != nil {
				panic(err)
			}

			return connection
		}

		if counts > 10 {
			log.Println(err)
			return nil
		}

		log.Println("Backing off for two seconds...")
		time.Sleep(2 * time.Second)
	}
}

func loadDSN() (string, error) {
	userPath := os.Getenv("PG_USER")
	passwordPath := os.Getenv("PG_PASSWORD")

	user, err := getSecret(userPath)
	if err != nil {
		log.Printf("Error reading user secret from %s: %v\n", userPath, err)
		return "", err
	}

	password, err := getSecret(passwordPath)
	if err != nil {
		log.Printf("Error reading password secret from %s: %v\n", passwordPath, err)
		return "", err
	}

	dsn := dsn{
		host:     os.Getenv("PG_HOST"),
		port:     os.Getenv("PG_PORT"),
		user:     user,
		password: password,
		dbname:   os.Getenv("PG_DBNAME"),
		sslmode:  os.Getenv("PG_SSL_MODE"),
		timezone: os.Getenv("PG_TIMEZONE"),
		timeout:  os.Getenv("PG_CONNECT_TIMEOUT"),
	}

	// host=postgres port=5432 user=postgres password=mbJAYRyYp6ZS0zuyuGSN1Q== dbname=users sslmode=disable timezone=UTC connect_timeout=5
	dsnString := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s timezone=%s connect_timeout=%s",
		dsn.host,
		dsn.port,
		dsn.user,
		dsn.password,
		dsn.dbname,
		dsn.sslmode,
		dsn.timezone,
		dsn.timeout,
	)

	return dsnString, nil
}

func getSecret(secretPath string) (string, error) {
	data, err := os.ReadFile(secretPath)
	if err != nil {
		return "", err
	}
	return strings.TrimSpace(string(data)), nil
}

func (app *Config) setupRepo(conn *sql.DB) {
	db := data.NewPostgresRepository(conn)
	app.Repo = db
}
