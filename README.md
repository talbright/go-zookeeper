# go-zookeeper [![Build Status](https://travis-ci.org/talbright/go-zookeeper.png)](https://travis-ci.org/talbright/go-zookeeper) [![Coverage Status](https://coveralls.io/repos/github/talbright/go-zookeeper/badge.svg?branch=build-improvements)](https://coveralls.io/github/talbright/go-zookeeper?branch=build-improvements) [![GoDoc](https://godoc.org/github.com/talbright/go-zookeeper/zk?status.svg)](http://godoc.org/github.com/talbright/go-zookeeper/zk)

[Zookeeper](https://zookeeper.apache.org/) client library for the GO language.

### Tour

Detailed examples and additional code can be found in [tour.go](examples/tour.go). To run the tour:

``$ go run tour.go``

#### Connecting

To establish a basic connection to zookeeper:

```go
	conn, _, err := zk.Connect([]string{"127.0.0.1"}, time.Second*20)
	if err != nil {
		panic(err)
	}
	defer conn.Close()
```

A connection to zookeeper is performed asynchronously, and for all practical purposes isn't useful 
until a session has been established. Additonal options can be supplied to the connection as well: 

```go
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
			//Now we can do something useful...
		}
	}
```

#### Znode

A znode is the fundamental entity for which operations are performed against. A znode in ZooKeeper can have data associated with it as well as children. It is like having a file-system that allows a file to also be a directory (ZooKeeper was designed to store coordination data: status information, configuration, location information, etc., so the data stored at each node is usually small, in the byte to kilobyte range.) The term znode and node are used interchangeably.

See the apache [zookeeper](https://zookeeper.apache.org/doc/trunk/zookeeperOver.html) project docs for more details.

#### Znode API

The following are fundamental operations on Znodes:
* Create
* Get
* Set
* Delete
* Exists
* Children
* Sync

Additional operations are available, see the [go-zookeper](http://godoc.org/github.com/talbright/go-zookeeper/zk) project docs for more details, or run the [tour.go](examples/tour.go).

#### Znode API example

```go
	var path string
	var err error

	if path, err = conn.Create("/myznode", []byte{}, zk.FlagPersistent, zk.WorldACL(zk.PermAll)); err != nil {
		panic(err)
	}

	if yes, _, err := conn.Exists(path); !yes || err != nil {
		panic(err)
	}

	if _, err = conn.Set(path, []byte("hello"), -1); err != nil {
		panic(err)
	}

	if data, _, err := conn.Get(path); err != nil || string(data) != "hello" {
		panic(err)
	} 

	if err = conn.Delete(path, -1); err != nil {
		panic(err)
	}
```
