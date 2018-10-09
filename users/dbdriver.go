package users

import (
	"fmt"
	"github.com/arangodb/go-driver"
	"github.com/arangodb/go-driver/http"
	"log"
	"strconv"
	"time"
)

var (
	database             driver.Database = dbDriver()
	initDatabase         bool            = false
	port                 int             = 9006
	url                  string          = "http://localhost:"
	databaseUser         string          = "iaf-users"
	databaseUserPassword string          = "iaf-users-2018@secret"
	databaseName         string          = "iaf-users"

	friendsColName string            = "friends"
	friendsCol     driver.Collection = col(friendsColName)
	usersColName   string            = "users"
	usersCol       driver.Collection = col(usersColName)

	friendsCollectionExists bool = false
	usersCollectionExists   bool = false
)

func dbDriver() driver.Database {
	counter := 0
	var db driver.Database
	for db == nil && counter < 10 {
		counter++
		// create repeated call library until connection is established with increasing sleep timer
		// can be put in docker-compose with health-check
		fmt.Println("Connecting to " + url + strconv.Itoa(port))
		if conn, err := http.NewConnection(http.ConnectionConfig{
			Endpoints: []string{url + strconv.Itoa(port)},
		}); err == nil {
			if c, e := driver.NewClient(driver.ClientConfig{
				Connection:     conn,
				Authentication: driver.BasicAuthentication("root", "iafoosball@users for the win"),
			}); e == nil {
				if !initDatabase {
					db = ensureDatabaseName(databaseName, c, db)
					initDatabase = true
				}
			} else {
				log.Fatal(e)
			}

		} else {
			log.Fatal(err)
		}
		if db == nil {
			fmt.Println("sleep")
			time.Sleep(2 * time.Second)
		}
	}
	return db
}

func ensureDatabaseName(name string, c driver.Client, db driver.Database) driver.Database {
	if exists, err := c.DatabaseExists(nil, databaseUser); exists == false {
		db, _ = c.CreateDatabase(nil, databaseName, &driver.CreateDatabaseOptions{
			[]driver.CreateDatabaseUserOptions{
				{
					UserName: databaseUser,
					Password: databaseUserPassword,
				},
			},
		})
	} else if err != nil {
		fmt.Println(err)
	} else {
		db, _ = c.Database(nil, databaseName)
	}

	if exists, err := c.UserExists(nil, databaseUser); exists == false {

	} else if err != nil {
		fmt.Println(err)
	} else {

	}

	if exists, err := db.CollectionExists(nil, friendsColName); exists == false {
		db.CreateCollection(nil, friendsColName, &driver.CreateCollectionOptions{
			Type: driver.CollectionTypeEdge,
		})
	} else if err != nil {
		fmt.Println(err)
	} else {

	}
	if exists, err := db.CollectionExists(nil, usersColName); exists == false {
		db.CreateCollection(nil, usersColName, &driver.CreateCollectionOptions{})
	} else if err != nil {
		fmt.Println(err)
	} else {

	}

	fmt.Println("Create new database with user iaf-users. If already there skip")

	//err == nil {
	//fmt.Print("create database")
	//if _, err := db.CreateCollection(nil, friendsColName, &driver.CreateCollectionOptions{
	//	Type: driver.CollectionTypeEdge,
	//}); err != nil {
	//	fmt.Print("sddfff")
	//	fmt.Println(err)
	//}
	//db.CreateCollection(nil, usersColName, &driver.CreateCollectionOptions{
	//	Type: driver.CollectionTypeDocument,
	//})
	//fmt.Print("create database")
	//} else {
	//	fmt.Println("fil")
	//	log.Print(err)
	//}
	//db, _ = c.Database(nil, "iaf-users")

	//database.CreateGraph(nil, graphMatches, &driver.CreateGraphOptions{OrphanVertexCollections: {
	//	[1]string{collectionsMatches},
	//}
	//})
	//}
	return db
}

func col(collection string) driver.Collection {
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
