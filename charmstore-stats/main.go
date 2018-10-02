package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/kr/pretty"
	"gopkg.in/juju/charm.v6"
	mgo "gopkg.in/mgo.v2"
)

var addr = flag.String("addr", "localhost", "mongo dial url")
var collection = flag.String("coll", "juju.stat.counters", "atats counters collection name")

func main() {
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "usage: charmstore-stats <charm-id>...\n")
		os.Exit(2)
	}
	flag.Parse()

	session, err := mgo.Dial(*addr)
	if err != nil {
		log.Fatal(err)
	}
	store := NewStore(session.DB("juju"), *collection)

	for _, arg := range flag.Args() {
		id, err := charm.ParseURL(arg)
		if err != nil {
			log.Fatal(err)
		}
		if err := showStats(store, id); err != nil {
			log.Printf("cannot get stats on %v: %v", id, err)
		}
	}
}

func showStats(store *Store, id *charm.URL) error {
	mark("start")
	one, all, err := store.ArchiveDownloadCounts(id)
	if err != nil {
		return err
	}
	mark("done")
	pretty.Println(id.String(), one)
	pretty.Println(id.WithRevision(-1).String(), all)
	return nil
}

var t0 = time.Now()

func mark(f string, a ...interface{}) {
	msg := fmt.Sprintf(f, a...)
	fmt.Printf("%08d %s\n", time.Since(t0)/time.Millisecond, msg)
}
