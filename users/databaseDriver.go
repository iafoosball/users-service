package users

import (
	"github.com/arangodb/go-driver"
	"github.com/arangodb/go-driver/http"
	"log"
	"strconv"
	"time"
)

var (
	database         driver.Database
	host             string
	port             int
	databaseUser     string
	databasePassword string
	usersCol         driver.Collection
	friendsCol       driver.Collection
	// goalsToMatchGraph driver.Graph
)

const (
	databaseName   = "iaf-users"
	usersColName   = "users"
	friendsColName = "friends"
	// goalsToMatchName = "goalsToMatch"
)

// InitDatabase tries establishes a connection to the database inside a for loop with 10 repetitions.
// If the database is not available it will sleep the time of the counter (c) in seconds.
// When a connection is established it will also open the Collections assiciated with this service.
func InitDatabase(dbHost string, dbPort int, dbUser string, dbPassword string) {
	port = dbPort
	host = dbHost
	databaseUser = dbUser
	databasePassword = dbPassword
	c := 0
	for database == nil && c <= 10 {
		initDatabaseDriver(dbUser, dbPassword)
		c++
		if database == nil {
			log.Println("Database not reachable! Sleep seconds: " + strconv.Itoa(c))
			time.Sleep(time.Duration(c) * 1000 * time.Millisecond)
			continue
		}

		Collection(usersColName)
		Collection(friendsColName)
		// Graph(goalsToMatchName)
	}
}

// Authenticate with the arangodb and get the database
func initDatabaseDriver(user string, password string) {
	if conn, err := http.NewConnection(http.ConnectionConfig{
		Endpoints: []string{"http://" + host + ":" + strconv.Itoa(port)},
	}); err == nil {
		if client, err := driver.NewClient(driver.ClientConfig{
			Connection:     conn,
			Authentication: driver.BasicAuthentication(user, password),
		}); err == nil {
			if dbExists, e := client.DatabaseExists(nil, databaseName); !dbExists && e == nil {
				if database, err = client.CreateDatabase(nil, databaseName, &driver.CreateDatabaseOptions{
					[]driver.CreateDatabaseUserOptions{
						{
							UserName: user,
						},
					},
				}); err != nil {
					log.Fatal(err)
					return
				}
				log.Println("Connected to database: " + databaseName)
			} else if e != nil {
				log.Println(e)
			} else {
				database, err = client.Database(nil, databaseName)
			}
			if err != nil {
				log.Println(err)
			}
		} else {
			log.Println(err)
		}
	} else {
		log.Println(err)
	}
}

// Collection opens a certain collection. If the collection does not exist, initialize it first.
func Collection(name string) driver.Collection {
	if database == nil {
		InitDatabase(host, port, databaseUser, databasePassword)
	} else {
		if name == friendsColName && friendsCol == nil {
			friendsCol = initCollection(name, 2)
			return friendsCol
		} else if name == friendsColName {
			return friendsCol
		}
		if name == usersColName && usersCol == nil {
			usersCol = initCollection(name, 3)
			return usersCol
		} else if name == usersColName {
			return usersCol
		}
	}
	return nil
}

// Initializes a collections
func initCollection(name string, colType int) driver.Collection {
	if exists, err := database.CollectionExists(nil, name); !exists {
		if col, e := database.CreateCollection(nil, name, &driver.CreateCollectionOptions{
			Type: driver.CollectionType(colType),
		}); e != nil {
			return col
		} else if err != nil {
			log.Println(err)
		}
	} else if err != nil {
		log.Println(err)
	} else {

		if col, err := database.Collection(nil, name); err == nil {
			return col
		} else {
			log.Println(err)
		}
	}
	return nil
}

// Graph was copy-pasted from matches, needs to be adjusted for users-service
// func Graph(name string) driver.Graph {
// 	if database == nil {
// 		InitDatabase(host, port, databaseUser, databasePassword)
// 	} else {
// 		if name == goalsToMatchName && goalsToMatchGraph == nil {
// 			goalsToMatchGraph = initGraph(name)
// 			return goalsToMatchGraph
// 		} else if name == goalsToMatchName {
// 			return goalsToMatchGraph
// 		}
// 	}
// 	return nil
// }

// initGraph is explicitly designed to create a Graph with matches as vertices and goals as edges.
// As we only have one graph for now, I don't think it is necessary to make this func more general,
// but it easy to do so in the future if necessary.
func initGraph(name string) driver.Graph {
	if exists, err := database.GraphExists(nil, name); !exists {
		if col, e := database.CreateGraph(nil, name, &driver.CreateGraphOptions{
			EdgeDefinitions: []driver.EdgeDefinition{{
				Collection: usersColName,
				To:         []string{friendsColName},
				From:       []string{friendsColName},
			}},
		}); e != nil {
			return col
		} else if err != nil {
			log.Println(err)
		}
	} else if err != nil {
		log.Println(err)
	} else {

		if col, err := database.Graph(nil, name); err == nil {
			return col
		} else {
			log.Println(err)
		}
	}
	return nil
}
