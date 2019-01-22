package goraven

import (
	"encoding/xml"
)

type Warning struct {
	XMLName xml.Name `xml:"Warning"`
	Text    string   `xml:"Text"`
}
