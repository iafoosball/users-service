package usersImpl

import (
	"log"

	"fmt"
	"github.com/arangodb/go-driver"
	"github.com/arangodb/go-driver/http"
	"time"
)

var db driver.Database

func DB() driver.Database {
	if db == nil {
		time.Sleep(3 * time.Second)
		conn, err := http.NewConnection(http.ConnectionConfig{
			// Endpoints: []string{"http://http://192.38.56.114:9002"},
			Endpoints: []string{"http://arangodb:9006"},
		})
		if err != nil {
			log.Fatal(err)
		}
		log.Println()
		c, err := driver.NewClient(driver.ClientConfig{
			Connection:     conn,
			Authentication: driver.BasicAuthentication("iaf", "iafoosball users"),
		})
		db, err = c.Database(nil, "users")
		if err != nil {
			log.Fatal(err)
		}
	}
	return db
}

func Col(collection string) driver.Collection {
	fmt.Println(collection)
	for db == nil {
		time.Sleep(time.Second)
	}
	col, err := db.Collection(nil, collection)
	if err != nil {
		log.Fatal(err)
	}
	return col
}

func EdgeCol(c string) driver.Collection {
	db := DB()
	col, err := db.Collection(nil, c)
	if err != nil {
		log.Fatal(err)
	}
	return col
}
