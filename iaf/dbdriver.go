package iaf

import (
	"log"

	driver "github.com/arangodb/go-driver"
	"github.com/arangodb/go-driver/http"
	"time"
	"fmt"
	"github.com/tatsushid/go-fastping"
	"net"
)

var db driver.Database

func DB() driver.Database {
	if db == nil {
		time.Sleep(2*time.Second)

		fmt.Println("trying to connect")
		conn, err := http.NewConnection(http.ConnectionConfig{
			// Endpoints: []string{"http://http://192.38.56.114:9002"},
			Endpoints: []string{"http://arangodb:8529"},
		})
		if err != nil {
			log.Fatal(err)
		}
		log.Println()
	c, err := driver.NewClient(driver.ClientConfig{
		Connection:     conn,
		Authentication: driver.BasicAuthentication("joe", "joe"),
	})
	db, _ = c.Database(nil, "iaf")
	//if err != nil {
	//	log.Fatal(err)
	//}
	}

	return db
}

func tryToConnect() {
	conn := false;
	for !conn {
		log.Println(conn)
		pinger := fastping.NewPinger()

		_, err := pinger.Network("udp")
		// We shouldn't ever get an error but we're checking anyway
		if err != nil {
			panic("Error setting network type: " + err.Error())
		}

		addr, err := net.ResolveIPAddr("ip", "localhost:8529")
		if err != nil {
			panic("Error resolving IP Address: " + err.Error())
		}

		pinger.AddIPAddr(addr)
		pinger.OnRecv = func(addr *net.IPAddr, rtt time.Duration) {
			fmt.Printf("%s time=%v seconds\n", addr, rtt.Seconds())
		}
		log.Println(conn)

		if err = pinger.Run(); err != nil {
			panic(err)
		}
		time.Sleep(10000)
	}
}

func Col(collection string) driver.Collection {
	db := DB()
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
