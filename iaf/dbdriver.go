package iaf
//
//import (
//	"log"
//
//	driver "github.com/arangodb/go-driver"
//	"github.com/arangodb/go-driver/http"
//)
//
//func DB() driver.Database {
//	conn, err := http.NewConnection(http.ConnectionConfig{
//		// Endpoints: []string{"http://http://192.38.56.114:9002"},
//		Endpoints: []string{"http://0.0.0.0:8529"},
//	})
//	if err != nil {
//		log.Fatal(err)
//	}
//	c, err := driver.NewClient(driver.ClientConfig{
//		Connection:     conn,
//		Authentication: driver.BasicAuthentication("joe", "joe"),
//	})
//	db, err := c.Database(nil, "iaf")
//	if err != nil {
//		log.Fatal(err)
//	}
//	return db
//}
//
//func Col(collection string) driver.Collection {
//	db := DB()
//	col, err := db.Collection(nil, collection)
//	if err != nil {
//		log.Fatal(err)
//	}
//	return col
//}
//
//func EdgeCol(c string) driver.Collection {
//	db := DB()
//	col, err := db.Collection(nil, c)
//	if err != nil {
//		log.Fatal(err)
//	}
//	return col
//}
