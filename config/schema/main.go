package schema

import (
	"encoding/xml"
)

// HTTP represents the configurations for HTTP.
type HTTP struct {
	Listener HTTPListener `xml:"listener"`
}

// HTTPListener represents the configurations for HTTP request listener.
type HTTPListener struct {
	Address string `xml:"address"`
	SSL     bool   `xml:"ssl,attr"`
}

// Main represents the main configurations.
type Main struct {
	XMLName xml.Name `xml:"http://schema.ctit.co.th/oid-registry/2017/08/25 oid-registry"`
	HTTP    HTTP     `xml:"http"`
}
