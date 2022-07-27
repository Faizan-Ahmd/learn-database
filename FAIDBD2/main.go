package main
import (
	"fmt"
	"os"
	"log"
	"database/sql"
	"github.com/go-sql-driver/mysql"
)
var db *sql.DB

func main() {
    // Capture connection properties.
    cfg := mysql.Config{
        User:   os.Getenv("DBUSER"),
        Passwd: os.Getenv("DBPASS"),
        Net:    "tcp",
        Addr:   "127.0.0.1:3306",
        DBName: "recordings",
		AllowNativePasswords: true,
    }
    // Get a database handle.
	//fmt.Printf("%v",cfg.Passwd)
    var err error
    db, err = sql.Open("mysql", cfg.FormatDSN())
    if err != nil {
        log.Fatal(err)
    }

    pingErr := db.Ping()
    if pingErr != nil {
        log.Fatal(pingErr)
    }
    fmt.Println("Connected!")
}