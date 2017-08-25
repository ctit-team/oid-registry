package config

import (
	// core packages
	"log"
	"os"

	// internal packages
	"github.com/ctit-team/oid-registry/config/loader"
	"github.com/ctit-team/oid-registry/config/schema"

	// external packages
	"github.com/ctit-team/nestederror"
)

const (
	// MainFile is the name of a file stored content for Main.
	MainFile = "conf/main.xml"
)

var (
	// Main is a default main configurations loaded from conf/main.xml.
	Main *schema.Main
)

func init() {
	var err error
	var file *os.File

	// main
	if file, err = os.Open(MainFile); err != nil {
		log.Fatalln(err)
	} else {
		defer file.Close()
	}

	if Main, err = loader.LoadMain(file); err != nil {
		log.Fatalln(nestederror.New(err, "failed to load %v", MainFile))
	}
}
