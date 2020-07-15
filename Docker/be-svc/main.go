package main

import (
	"database/sql"
	"flag"
	"fmt"
	"log"
	"net/http"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

var (
	dbUser     = flag.String("mysql-username", "root", "Mysql DB Username")
	dbPass     = flag.String("mysql-password", "root", "Mysql DB Password")
	dbHost     = flag.String("mysql-host", "localhost", "Mysql host")
	dbPort     = flag.Int("mysql-port", 3306, "Mysql port")
	dbName     = flag.String("mysql-db", "maha_test", "Mysql DB name")
	listenAddr = flag.String("listen-addr", ":8080", "Listen address")
)

var (
	dbConn *sql.DB
)

func main() {
	flag.Parse()
	http.Handle("/log", &logHandler{})
	http.Handle("/", &healthHandler{})
	openMysqlConnection()
	http.ListenAndServe(*listenAddr, nil)
}

type logHandler struct{}
type healthHandler struct{}

func (b logHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	log.Printf("Received request")
	timeNow := time.Now().Unix()
	requestID := r.Header.Get("X-Request-Id")
	dbConn.Exec("INSERT INTO logs (request_id, created_at) VALUES (?,?)", requestID, timeNow)
	w.Write([]byte(fmt.Sprintf("Time: %d; Request ID: %s", timeNow, requestID)))
}

func (b healthHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("OK"))
}

func openMysqlConnection() {
	var err error
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/", *dbUser, *dbPass, *dbHost, *dbPort)
	log.Printf("Attempting to connect at %s", dsn)
	dbConn, err = sql.Open("mysql", dsn)
	if err != nil {
		panic(err)
	}

	if err := dbConn.Ping(); err != nil {
		log.Fatalf("Failed to connect to Mysql: %s", err)
	} else {
		log.Printf("Connection successful")
	}

	_, err = dbConn.Exec(fmt.Sprintf("CREATE DATABASE IF NOT EXISTS %s", *dbName))
	if err != nil {
		panic(err)
	}
	_, err = dbConn.Exec(fmt.Sprintf("USE %s;", *dbName))
	if err != nil {
		panic(err)
	}
	_, err = dbConn.Exec("CREATE TABLE IF NOT EXISTS logs (id INT NOT NULL, request_id VARCHAR(64), created_at DATETIME, PRIMARY KEY(id));")
	if err != nil {
		panic(err)
	}

	log.Printf("DB setup complete")
}
