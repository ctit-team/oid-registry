package loader

import (
	// core packages
	"encoding/xml"
	"io"
	"io/ioutil"

	// internal packages
	"github.com/ctit-team/oid-registry/config/schema"

	// external packages
	"github.com/ctit-team/nestederror"
)

// LoadMain load the main configurations from data.
func LoadMain(data io.Reader) (*schema.Main, error) {
	// read content
	var err error
	var content []byte

	if content, err = ioutil.ReadAll(data); err != nil {
		return nil, nestederror.New(err, "failed to read content")
	}

	// unmarshal
	conf := new(schema.Main)

	if err = xml.Unmarshal(content, conf); err != nil {
		return nil, nestederror.New(err, "failed to unmarshal content")
	}

	return conf, nil
}
