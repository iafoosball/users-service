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

const (
	databaseUser         = "iaf-users"
	databaseUserPassword = "iaf-users-2018@secret"
	databaseName         = "iaf-users"
	friendsColName       = "friends"
	usersColName         = "users"
)

var (
	//port               	= 9006
	database                = dbDriver()
	initDatabase            = false
	port                    = *flag.Int("dbPort", 8529, "dbPort")
	url                     = "http://" + *flag.String("dbHost", "arangodb:", "sdfsdf")
	friendsCol              = Col(friendsColName)
	usersCol                = Col(usersColName)
	friendsCollectionExists = false
	usersCollectionExists   = false
)

func dbDriver() driver.Database {
	var count int
	var db driver.Database
	// create repeated call library until connection is established with increasing sleep timer
	for db == nil && count < 10 {
		count++
		log.Println("Sleep seconds " + strconv.Itoa(count))
		time.Sleep(time.Duration(count) * time.Second)

		log.Println("Connecting to " + url + strconv.Itoa(port))
		conn, err := http.NewConnection(http.ConnectionConfig{
			Endpoints: []string{url + strconv.Itoa(port)},
		})
		if err != nil {
			log.Fatalln("Failed to connect to: " + url + strconv.Itoa(port))
			log.Fatalln(err)
			continue
		}

		log.Println("Logging in DB client...")
		client, err := driver.NewClient(driver.ClientConfig{
			Connection:     conn,
			Authentication: driver.BasicAuthentication("iaf", "iafoosball@users for the win"),
		})
		if err != nil {
			log.Fatalln("Failed to log in: ")
			log.Fatalln(err)
			continue
		}

		log.Println("Getting reference to " + databaseName)
		if !initDatabase {
			if db, err = DB(databaseName, client); err != nil {
				log.Fatalln(err)
				continue
			}
			initDatabase = true
		}
	}
	return db
}

// DB gets database reference by name
func DB(name string, c driver.Client) (driver.Database, error) {
	var db driver.Database
	log.Println("Available DBs:")
	s, err := c.AccessibleDatabases(nil)
	if err != nil {
		return nil, err
	}
	for index, element := range s {
		fmt.Println(strconv.Itoa(index) + " bllaa " + element.Name())
	}

	log.Println("Looking for ", name)
	exists, err := c.DatabaseExists(nil, name)
	if err != nil {
		return nil, err
	}

	if exists {
		fmt.Println(name + " exists!")
		if db, err = c.Database(nil, "iaf-users"); err != nil {
			return nil, err
		}
	} else {
		fmt.Println(name + " doesn't exists. Creating...")
		if db, err = c.CreateDatabase(nil, name, &driver.CreateDatabaseOptions{
			[]driver.CreateDatabaseUserOptions{
				{
					UserName: databaseUser,
					Password: databaseUserPassword,
				},
			},
		},
		); err != nil {
			return nil, err
		}
	}
	return db, nil
}

// Col returns collection by name
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
