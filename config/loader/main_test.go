package loader

import (
	// core packages
	"bytes"
	"testing"

	// internal packages
	"github.com/ctit-team/oid-registry/config/schema"
)

func TestLoadMain(t *testing.T) {
	var err error

	// setup
	content := `
	<?xml version="1.0" encoding="utf-8"?>
	<oid-registry xmlns="http://schema.ctit.co.th/oid-registry/2017/08/25">
		<http>
			<listener ssl="true">
				<address>localhost:1234</address>
			</listener>
		</http>
	</oid-registry>
	`

	// run test
	var conf *schema.Main
	if conf, err = LoadMain(bytes.NewBufferString(content)); err != nil {
		t.Fatal(err)
	}

	// assert
	if val := conf.HTTP.Listener.SSL; val != true {
		t.Fatalf("expected 'true' on oid-registry > http > listener > ssl, got '%v'", conf.HTTP.Listener.SSL)
	}

	if val := conf.HTTP.Listener.Address; val != "localhost:1234" {
		t.Fatalf("expected 'localhost:1234' on oid-registry > http > listener > address, got '%v'", val)
	}
}
