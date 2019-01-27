package goraven

import (
	"encoding/xml"
	"math"
)

// Send the GET_CURRENT_PRICE command to get the price information. Set the
// refresh element to Y to force the RAVEn to get the information from the
// meter, not the cache. The RAVEn will send a PriceCluster notification in
// response.
func (r *Raven) GetCurrentPrice() error {
	return r.simpleCommand("get_current_price", false)
}

// Not Implemented
func (r *Raven) SetCurrentPrice() {
}

// GetPrice() is a convenience function to get a correctly formatted
// floating point number of the Price contained in the PriceCluster
func (p *PriceCluster) GetPrice() (float64, error) {
	price := p.Price
	digits := p.TrailingDigits
	divisor := math.Pow10(int(digits))

	return (float64(price) / divisor), nil
}

// Notify: PriceCluster
type PriceCluster struct {
	XMLName        xml.Name  `xml:"PriceCluster"`
	DeviceMacId    uhexint64 `xml:"DeviceMacId"`
	MeterMacId     uhexint64 `xml:"MeterMacId"`
	TimeStamp      timestamp `xml:"TimeStamp"`
	Price          uhexint32 `xml:"Price"`
	Currency       uhexint16 `xml:"Currency"`
	TrailingDigits uhexint8  `xml:"TrailingDigits"`
	Tier           uhexint8  `xml:"Tier"`
	TierLabel      string    `xml:"TierLabel,omitempty"`
	RateLabel      string    `xml:"RateLabel,omitempty"`
}
