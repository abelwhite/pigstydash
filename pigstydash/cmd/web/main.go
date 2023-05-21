// Filename: cmd/web/main.go

//Written by: Abel Blanco, Jovan Alpuche, Lianni Mathews, Cameron Tillet, Talib Marin, Amir Gonzalez
//Tested by: Abel Blanco, Jovan Alpuche, Lianni Mathews, Cameron Tillet, Talib Marin, Amir Gonzalez
//Debbuged by: Abel Blanco, Jovan Alpuche, Lianni Mathews, Cameron Tillet, Talib Marin, Amir Gonzalez

package main

import (
	"context"
	"database/sql"
	"flag"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/abelwhite/pigstydash/internal/models"
	"github.com/alexedwards/scs/v2"
	_ "github.com/jackc/pgx/v5/stdlib"
)

// Share data across our handlers
type application struct {
	errorLog       *log.Logger
	infoLog        *log.Logger
	user           models.UserModel
	pig            models.PigModel
	room           models.RoomModel
	pigsty         models.PigstyModel
	sessionManager *scs.SessionManager //create a field sessionManager of type pointscs.SessionManager
}

func main() {
	// configure our server
	addr := flag.String("addr", ":4000", "HTTP network address")
	dsn := flag.String("dsn", os.Getenv("PIGSTYDB_DB_DSN"), "PostgreSQL DSN (Data Source Name)")
	flag.Parse()

	// get a database connection pool
	db, err := openDB(*dsn)
	if err != nil {
		log.Print(err)
		return
	}
	//create instances of errorLog & infoLog
	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)
	//setup a new session mamager
	sessionManager := scs.New() //function creates a session manager and sends the location
	sessionManager.Lifetime = 1 * time.Hour
	sessionManager.Cookie.Persist = true                  //close the browser it stays alive
	sessionManager.Cookie.Secure = false                  //not incrypted
	sessionManager.Cookie.SameSite = http.SameSiteLaxMode //when u change site then

	// share data across our handlers
	app := &application{
		errorLog:       errorLog,
		infoLog:        infoLog,
		user:           models.UserModel{DB: db},
		pig:            models.PigModel{DB: db},
		room:           models.RoomModel{DB: db},
		pigsty:         models.PigstyModel{DB: db},
		sessionManager: sessionManager,
	}
	// cleanup the connection pool
	defer db.Close()
	// acquired a database connection pool
	infoLog.Println("database connection pool established")
	// create and start a custom web server
	infoLog.Printf("starting server on %s", *addr)

	//configure TLS
	// tlsConfig := &tls.Config{
	// 	CurvePreferences: []tls.CurveID{tls.X25519, tls.CurveP256}, //cryptography
	// }

	srv := &http.Server{
		Addr:         *addr,
		ErrorLog:     errorLog,
		Handler:      app.routes(),
		IdleTimeout:  time.Minute,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		// TLSConfig:    tlsConfig,
	}
	// err = srv.ListenAndServeTLS("./tls/cert.pem", "./tls/key.pem")
	// log.Fatal(err)
	err = srv.ListenAndServe()
	log.Fatal(err)
}

// The openDB() function returns a database connection pool or error
func openDB(dsn string) (*sql.DB, error) {
	db, err := sql.Open("pgx", dsn)
	if err != nil {
		return nil, err
	}
	// create a context
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	// test the DB connection
	err = db.PingContext(ctx)
	if err != nil {
		return nil, err
	}
	return db, nil
}
