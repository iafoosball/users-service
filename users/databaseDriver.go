package users

import (
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
	friendsCol     driver.Collection = Col(friendsColName)
	usersColName   string            = "users"
	usersCol       driver.Collection = Col(usersColName)

	friendsCollectionExists bool = false
	usersCollectionExists   bool = false
)

func dbDriver() driver.Database {
	c := 0
	var db driver.Database
	// create repeated call library until connection is established with increasing sleep timer
	for db == nil && c < 10 {
		c++
		// can be put in docker-compose with health-check
		log.Println("Connecting to " + url + strconv.Itoa(port))
		if conn, err := http.NewConnection(http.ConnectionConfig{
			Endpoints: []string{url + strconv.Itoa(port)},
		}); err == nil {
			log.Println("Connected to database")
			if c, e := driver.NewClient(driver.ClientConfig{
				Connection:     conn,
				Authentication: driver.BasicAuthentication("root", "iafoosball@users for the win"),
			}); e == nil {
				log.Println("authorized as root")
				if !initDatabase {
					log.Println("init db")
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
			log.Println("Sleep seconds" + strconv.Itoa(c))
			time.Sleep(time.Duration(c) * 1000 * time.Millisecond)
		}
	}
	return db
}

func ensureDatabaseName(name string, c driver.Client, db driver.Database) driver.Database {
	log.Println("Create new database with user iaf-users. If already there skip")
	if db == nil {
		if db, err := c.CreateDatabase(nil, databaseName, &driver.CreateDatabaseOptions{
			[]driver.CreateDatabaseUserOptions{
				{
					UserName: databaseUser,
					Password: databaseUserPassword,
				},
			},
		},
		); err == nil {
			log.Print("create database")
			if _, err := db.CreateCollection(nil, friendsColName, &driver.CreateCollectionOptions{
				Type: driver.CollectionTypeEdge,
			}); err != nil {
				log.Print("sddfff")
				log.Println(err)
			}
			db.CreateCollection(nil, usersColName, &driver.CreateCollectionOptions{
				Type: driver.CollectionTypeDocument,
			})
			log.Print("create database")
		} else {
			log.Print(err)
		}
		db, _ = c.Database(nil, "iaf-users")

		//database.CreateGraph(nil, graphUsers, &driver.CreateGraphOptions{OrphanVertexCollections: {
		//	[1]string{collectionsUsers},
		//}
		//})
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
