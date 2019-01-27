package goraven

import (
	"encoding/xml"
)

// Send the GET_TIME command to get the current time. The RAVEn will send
// a TimeCluster notification in response
func (r *Raven) GetTime() error {
	return r.simpleCommand("get_time", false)
}

// Notify: TimeCluster
type TimeCluster struct {
	XMLName     xml.Name  `xml:"TimeCluster"`
	DeviceMacId uhexint64 `xml:"DeviceMacId"`
	MeterMacId  uhexint64 `xml:"MeterMacId"`
	UTCTime     timestamp `xml:"UTCTime"`
	LocalTime   timestamp `xml:"LocalTime"`
}
