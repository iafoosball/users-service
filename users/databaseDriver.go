package users

import (
	"flag"
	"fmt"
	"github.com/arangodb/go-driver"
	"github.com/arangodb/go-driver/http"
	"log"
	"strconv"
	"time"
)

var (
	database     driver.Database = dbDriver()
	initDatabase bool            = false
	port         int             = *flag.Int("dbPort", 8529, "dbPort")
	//port                 int             = 9006
	url                  string = "http://" + *flag.String("dbHost", "arangodb:", "sdfsdf")
	databaseUser         string = "iaf-users"
	databaseUserPassword string = "iaf-users-2018@secret"
	databaseName         string = "iaf-users"

	friendsColName string            = "friends"
	friendsCol     driver.Collection = Col(friendsColName)
	usersColName   string            = "users"
	usersCol       driver.Collection = Col(usersColName)

	friendsCollectionExists bool = false
	usersCollectionExists   bool = false
)

func dbDriver() driver.Database {
	fmt.Println("sdasdfasdff")
	fmt.Println(port)
	fmt.Println("hsdhf 	" + url)
	c := 0
	var db driver.Database
	// create repeated call library until connection is established with increasing sleep timer
	for db == nil && c < 1000 {
		c++
		// can be put in docker-compose with health-check
		log.Println("Connecting to " + url + strconv.Itoa(port))
		if conn, err := http.NewConnection(http.ConnectionConfig{
			Endpoints: []string{url + strconv.Itoa(port)},
		}); err == nil {
			log.Println("Connected to database")
			if client, e := driver.NewClient(driver.ClientConfig{
				Connection:     conn,
				Authentication: driver.BasicAuthentication("iaf", "iafoosball@users for the win"),
			}); e == nil {
				log.Println("authorized as root")
				if !initDatabase {
					db = ensureDatabase(databaseName, client, db)
					initDatabase = true
				}
			} else {
				log.Fatal(e)
			}

		} else {
			log.Println("not connecting")
			log.Println(err)
		}
		if db == nil {
			log.Println("Sleep seconds" + strconv.Itoa(c))
			time.Sleep(time.Duration(c) * 1000 * time.Millisecond)
		}
	}
	return db
}

func ensureDatabase(name string, c driver.Client, db driver.Database) driver.Database {
	var err error
	var exists bool
	log.Println("Create new database with user iaf-users. If already there skip")
	if s, error := c.AccessibleDatabases(nil); error != nil {
		for index, element := range s {
			fmt.Println(strconv.Itoa(index) + " bllaa " + element.Name())
		}
	} else {
		log.Println(error)
	}
	if db, err = c.Database(nil, "iaf-users"); err != nil {
		log.Println(err)
	}
	if db == nil {
		log.Println("database is null")
		// this call fails, but why?????
		log.Println(databaseName)
		if exists, err = c.DatabaseExists(nil, databaseName); exists && err == nil {
			log.Println("database exists")
			if db, err = c.Database(nil, "iaf-users"); err != nil {
				log.Println(err)
			}
		} else if !exists && err == nil {
			log.Println("database does not exist")
			if db, err = c.CreateDatabase(nil, databaseName, &driver.CreateDatabaseOptions{
				[]driver.CreateDatabaseUserOptions{
					{
						UserName: databaseUser,
						Password: databaseUserPassword,
					},
				},
			},
			); err == nil {

				//log.Println("create database")
				//if _, err := db.CreateCollection(nil, friendsColName, &driver.CreateCollectionOptions{
				//	Type: driver.CollectionTypeEdge,
				//}); err != nil {
				//	log.Println(err)
				//}
				//db.CreateCollection(nil, usersColName, &driver.CreateCollectionOptions{
				//	Type: driver.CollectionTypeDocument,
				//})
			} else {

				log.Println(err)
			}
		} else {
			log.Println("should be err")
			log.Println(err)
		}

		//database.CreateGraph(nil, graphUsers, &driver.CreateGraphOptions{OrphanVertexCollections: {
		//	[1]string{collectionsUsers},
		//}
		//})
	} else {
		log.Println("database is not null")
	}
	return db
}

func Col(collection string) driver.Collection {
	log.Println("Open collection: " + collection)
	if database != nil {
		col, err := database.Collection(nil, collection)
		if err != nil {
			log.Fatal(err)
		}
		return col
	} else {
		panic("No database!!!")
	}
}
