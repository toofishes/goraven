package goraven

import (
	"encoding/xml"
)

// GetMeterInfo gets the meter information. The RAVEn will send a MeterInfo
// notification in response.
func (r *Raven) GetMeterInfo() error {
	return r.simpleCommand("get_meter_info", false)
}

// Get the status of the device on the network. The RAVEn will send a
// NetworkInfo notification in response.
func (r *Raven) GetNetworkInfo() error {
	return r.simpleCommand("get_network_info", false)
}

// SetMeterInfo sets the meter information.
func (r *Raven) SetMeterInfo() error {
	v := setMeterInfo{
		Name: "set_meter_info",
	}
	return r.sendCommand(v)
}

// Command: SET_METER_INFO
type setMeterInfo struct {
	XMLName    xml.Name `xml:"Command"`
	Name       string   `xml:"Name"`
	MeterMacId string   `xml:"MeterMacId,omitempty"`
	NickName   string   `xml:"NickName,omitempty"`
	Account    string   `xml:"Account,omitempty"`
	Auth       string   `xml:"Auth,omitempty"`
	Host       string   `xml:"host,omitempty"`
	Enabled    string   `xml:"enabled,omitempty"`
}

// Notify: MeterInfo
type MeterInfo struct {
	XMLName     xml.Name  `xml:"MeterInfo"`
	DeviceMacId uhexint64 `xml:"DeviceMacId"`
	MeterMacId  uhexint64 `xml:"MeterMacId"`
	MeterType   string    `xml:"MeterType"`
	NickName    string    `xml:"NickName"`
	Account     string    `xml:"Account,omitempty"`
	Auth        string    `xml:"Auth,omitempty"`
	Host        string    `xml:"Host,omitempty"`
	Enabled     string    `xml:"Enabled,omitempty"`
}

// Notify: NetworkInfo
type NetworkInfo struct {
	XMLName      xml.Name  `xml:"NetworkInfo"`
	DeviceMacId  uhexint64 `xml:"DeviceMacId"`
	CoordMacId   uhexint64 `xml:"CoordMacId"`
	Status       string    `xml:"Status"`
	Description  string    `xml:"Description"`
	StatusCode   uhexint8  `xml:"StatusCode"`
	ExtPanId     uhexint64 `xml:"ExtPanId"`
	Channel      uint8     `xml:"Channel"`
	ShortAddr    uhexint16 `xml:"ShortAddr"`
	LinkStrength uhexint8  `xml:"LinkStrength"`
}
