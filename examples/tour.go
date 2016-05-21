package main

import (
	"log"
	"net"
	"os"
	"time"

	"github.com/talbright/go-zookeeper/zk"
)

var logger *log.Logger
var zkLogger *log.Logger

func init() {
	logger = log.New(os.Stdout, "[TOUR] ", log.Ldate|log.Ltime)
	zkLogger = log.New(os.Stdout, "[ZK] ", log.Ldate|log.Ltime)
}

func main() {
	logger.Print("starting tour of go-zookeeper")
	connectBasic()
	connectWithEvents()
	conn, _ := connectWithOptions()
	znodeBasicExample(conn)
}

func connectBasic() {
	logger.Print("establishing connection")
	conn, _, err := zk.Connect([]string{"127.0.0.1"}, time.Second*20)
	if err != nil {
		panic(err)
	}
	defer conn.Close()
}

func connectWithEvents() {
	logger.Print("establishing connection and monitoring events")
	conn, events, err := zk.Connect([]string{"127.0.0.1"}, time.Second*20)
	if err != nil {
		panic(err)
	}
	defer conn.Close()
	logger.Print("receive connection events")
	for {
		select {
		case event := <-events:
			logger.Printf("received event from zk: %v", event)
			if event.State == zk.StateHasSession {
				logger.Printf("session established")
			}
		case <-time.After(10 * time.Second):
			logger.Printf("closing connection")
			return
		}
	}
}

func connectWithOptions() (*zk.Conn, <-chan zk.Event) {
	logger.Print("establishing connection with options")
	conn, events, err := zk.Connect([]string{"127.0.0.1"},
		time.Second*20,
		zk.WithLogger(zkLogger),
		zk.WithDialer(net.DialTimeout),
		zk.WithConnectTimeout(time.Second*20))
	if err != nil {
		panic(err)
	}
	for event := range events {
		if event.State == zk.StateHasSession {
			return conn, events
		}
	}
	return conn, events
}

func znodeBasicExample(conn *zk.Conn) {
	logger.Print("basic znode example")
	var path string
	var err error

	logger.Printf("creating node /myznode")
	if path, err = conn.Create("/myznode", []byte{}, zk.FlagPersistent, zk.WorldACL(zk.PermAll)); err != nil {
		panic(err)
	}

	logger.Printf("checking if node %s exists\n", path)
	if yes, _, err := conn.Exists(path); !yes || err != nil {
		panic(err)
	}

	logger.Printf("setting node %s data\n", path)
	if _, err = conn.Set(path, []byte("hello"), -1); err != nil {
		panic(err)
	}

	logger.Printf("getting node %s data\n", path)
	if data, _, err := conn.Get(path); err != nil || string(data) != "hello" {
		panic(err)
	} else {
		logger.Printf("node data: %v", string(data))
	}

	logger.Printf("deleting node %s\n", path)
	if err = conn.Delete(path, -1); err != nil {
		panic(err)
	}

}
