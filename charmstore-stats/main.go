package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/kr/pretty"
	"gopkg.in/juju/charm.v6"
	mgo "gopkg.in/mgo.v2"
)

var addr = flag.String("addr", "localhost", "mongo dial url")

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
	store := NewStore(session.DB("juju"))

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
	one, all, err := store.ArchiveDownloadCounts(id)
	if err != nil {
		return err
	}
	pretty.Println(id.String(), one)
	pretty.Println(id.WithRevision(-1).String(), all)
	return nil
}
