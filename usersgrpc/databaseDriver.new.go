package usersgrpc

import (
	"log"
	"strconv"
	"time"

	"github.com/arangodb/go-driver"
	"github.com/arangodb/go-driver/http"
)

type ArangoDB struct {
	Host     string
	Port     int
	User     string
	Password string
	Ref      driver.Database
	goalsCol          driver.Collection
	matchesCol        driver.Collection
}

const (
	dbName           = "iaf-users"
	usersColName     = "users"
	friendsColName   = "friends"
)

////Addr to set address of the server
//func Addr(address string) {
//	addr = "http://" + address
//}

// InitDatabase tries establishes a connection to the db inside a for loop with 10 repetitions.
// If the db is not available it will sleep the time of the counter (c) in seconds.
// When a connection is established it will also open the Collections assiciated with this service.
func (db *ArangoDB) InitDatabase() {
	log.Println("host: " + db.Host + " port: " + strconv.Itoa(db.Port))
	c := 0
	for db == nil && c <= 10 {
		db.dbDriver(db.User, db.Password)
		c++
		if db == nil {
			log.Println("Database not reachable! Sleep seconds: " + strconv.Itoa(c))
			time.Sleep(time.Duration(c) * 1000 * time.Millisecond)
		}
	}
}

// Authenticate with the arangodb and get the db resource
func (db *ArangoDB) dbDriver(user string, password string) {
	conn, err := http.NewConnection(http.ConnectionConfig{
		Endpoints: []string{"http://" + db.Host + ":" + strconv.Itoa(db.Port)},
	})
	if err != nil {
		log.Println(err)
		return
	}
	client, err := driver.NewClient(driver.ClientConfig{
		Connection:     conn,
		Authentication: driver.BasicAuthentication(user, password),
	})
	if err != nil {
		log.Println(err)
		return
	}
	if dbExists, e := client.DatabaseExists(nil, dbName); !dbExists && e == nil {
		if newDB, err := client.CreateDatabase(nil, dbName, &driver.CreateDatabaseOptions{
			Users: []driver.CreateDatabaseUserOptions{
				{
					UserName: db.User,
				},
			},
		}); err != nil {
			log.Fatal(err)
			return
		} else {
			db.Ref = newDB
			log.Println("Connected to db: " + dbName)
		}

		db.col(usersColName)
		db.col(friendsColName)
	} else if e == nil {
		newDB, err := client.Database(nil, dbName)
		if err != nil {
			log.Println(err)
		}
		db.Ref = newDB
		log.Println("Connected to db: " + dbName)
		db.col(usersColName)
		db.col(friendsColName)
	}
}

// Open a collection. If the collection does not exist, initialize it first.
func (db *ArangoDB) col(name string) driver.Collection {
	if db == nil {
		log.Fatal("Need to initialize ArangoDB first!")
	} else {
		if name == friendsColName && db.matchesCol == nil {
			db.matchesCol = db.initCollection(name, 2)
			return db.matchesCol
		} else if name == friendsColName {
			return db.matchesCol
		}
		if name == usersColName && db.goalsCol == nil {
			db.goalsCol = db.initCollection(name, 3)
			return db.goalsCol
		} else if name == usersColName {
			return db.goalsCol
		}
	}
	return nil
}

// Initializes a collections
func (db *ArangoDB) initCollection(name string, colType int) driver.Collection {
	var col driver.Collection
	if exists, err := db.Ref.CollectionExists(nil, name); !exists && err == nil {
		if col, err = db.Ref.CreateCollection(nil, name, &driver.CreateCollectionOptions{
			Type: driver.CollectionType(colType),
		}); err != nil {
			return col
		}
	} else if exists && err == nil {
		if col, err = db.Ref.Collection(nil, name); err == nil {
			return col
		}
	}
	return nil
}

//func graph(name string) driver.Graph {
//	if db == nil {
//		log.Fatal("Need to initialize ArangoDB first!")
//	} else {
//		if name == goalsToMatchName && goalsToMatchGraph == nil {
//			goalsToMatchGraph = initGraph(name)
//			return goalsToMatchGraph
//		} else if name == goalsToMatchName {
//			return goalsToMatchGraph
//		}
//	}
//	return nil
//}

// initGraph is explicitly designed to create a graph with matches as vertices and goals as edges.
// As we only have one graph for now, I don't think it is necessary to make this func more general,
// but it easy to do so in the future if necessary.
//func initGraph(name string) driver.Graph {
//	var graph driver.Graph
//	if exists, er = db.GraphExists(nil, name); !exists && er == nil {
//		if graph, er = db.CreateGraph(nil, name, &driver.CreateGraphOptions{
//			EdgeDefinitions: []driver.EdgeDefinition{{
//				Collection: usersColName,
//				To:         []string{friendsColName},
//				From:       []string{friendsColName},
//			}},
//		}); er != nil {
//			return graph
//		}
//	} else if er == nil {
//		if graph, er = db.Graph(nil, name); er == nil {
//			return graph
//		}
//	}
//	if er != nil {
//		log.Println(er)
//	}
//	return nil
//}
//
//// Get an integer from a query.
//func queryInt(query string) int {
//	if cursor, err := db.Query(nil, query, make(map[string]interface{})); err != nil {
//		log.Println(err)
//	} else {
//		defer cursor.Close()
//		var i int
//		for cursor.HasMore() {
//			if _, err = cursor.ReadDocument(nil, &i); err != nil {
//				log.Println(err)
//			}
//			return i
//		}
//	}
//	log.Panic("query was not successful")
//	return 0
//}
