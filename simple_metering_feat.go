package goraven

import (
	"encoding/xml"
	"math"
)

// Get the demand information from the RAVEn. If refresh is true, the device
// gets the information from the meter instead of from its cache.
func (r *Raven) GetInstantaneousDemand(refresh bool) error {
	return r.simpleCommand("get_instantaneous_demand", refresh)
}

// Get the summation information from the RAVEn. If refresh is true, the device
// gets the information from the meter instead of from its cache.
func (r *Raven) GetCurrentSummationDelivered(refresh bool) error {
	return r.simpleCommand("get_current_summation_delivered", refresh)
}

// Get the accumulated usage information from the RAVEn. The RAVEn will send
// a CurrentPeriodUsage notification in response. Note that this command will
// not cause the current period consumption total to be updated. To do this,
// send a GetCurrentSummationDelivered command with Refresh set to Y.
func (r *Raven) GetCurrentPeriodUsage() error {
	return r.simpleCommand("get_current_period_usage", false)
}

// Get the previous period accumulation data from the RAVEn. The RAVEn will
// send a LastPeriodUsage notification in response.
func (r *Raven) GetLastPeriodUsage() error {
	return r.simpleCommand("get_last_period_usage", false)
}

// Force the RAVEn to roll over the current period to the last period and
// initialize the current period.
func (r *Raven) CloseCurrentPeriod() error {
	return r.simpleCommand("close_current_period", false)
}

// Not Implemented
func (r *Raven) SetFastPoll() {
}

// Not Implemented
func (r *Raven) GetProfileData() {
}

// GetDemand() is a convenience function that returns a correctly formatted
// floating point number for the Demand field
func (i *InstantaneousDemand) GetDemand() (float64, error) {
	return getFloat64(i.Demand, i.Multiplier, i.Divisor, i.DigitsLeft)
}

// GetSummationDelivered() is a convenience function that returns a correctly
// formatted floating point number for the Current Summation Delivered field
func (c *CurrentSummationDelivered) GetSummationDelivered() (float64, error) {
	return getFloat64(c.SummationDelivered, c.Multiplier, c.Divisor, c.DigitsLeft)
}

// GetSummationReceived() is a convenience function that returns a correctly
// formatted floating point number for the Current Summation Received field
func (c *CurrentSummationDelivered) GetSummationReceived() (float64, error) {
	return getFloat64(c.SummationReceived, c.Multiplier, c.Divisor, c.DigitsLeft)
}

func getFloat64(demand, multiplier, divisor uhexint32, digits uhexint8) (float64, error) {
	if multiplier == 0 {
		multiplier = 1
	}
	if divisor == 0 {
		divisor = 1
	}

	d := (int64(demand) % (int64(math.Pow10(int(digits))) * int64(divisor)))

	return ((float64(d) * float64(multiplier)) / float64(divisor)), nil
}

// Notify: InstantaneousDemand
type InstantaneousDemand struct {
	XMLName             xml.Name  `xml:"InstantaneousDemand"`
	DeviceMacId         uhexint64 `xml:"DeviceMacId"`
	MeterMacId          uhexint64 `xml:"MeterMacId"`
	TimeStamp           timestamp `xml:"TimeStamp"`
	Demand              uhexint32 `xml:"Demand"`
	Multiplier          uhexint32 `xml:"Multiplier"`
	Divisor             uhexint32 `xml:"Divisor"`
	DigitsRight         uhexint8  `xml:"DigitsRight"`
	DigitsLeft          uhexint8  `xml:"DigitsLeft"`
	SuppressLeadingZero ynbool    `xml:"SuppressLeadingZero"`
}

// Notify: CurrentSummationDelivered
type CurrentSummationDelivered struct {
	XMLName             xml.Name  `xml:"CurrentSummationDelivered"`
	DeviceMacId         uhexint64 `xml:"DeviceMacId"`
	MeterMacId          uhexint64 `xml:"MeterMacId"`
	TimeStamp           timestamp `xml:"TimeStamp"`
	SummationDelivered  uhexint32 `xml:"SummationDelivered"`
	SummationReceived   uhexint32 `xml:"SummationReceived"`
	Multiplier          uhexint32 `xml:"Multiplier"`
	Divisor             uhexint32 `xml:"Divisor"`
	DigitsRight         uhexint8  `xml:"DigitsRight"`
	DigitsLeft          uhexint8  `xml:"DigitsLeft"`
	SuppressLeadingZero ynbool    `xml:"SuppressLeadingZero"`
}

// Notify: CurrentPeriodUsage
type CurrentPeriodUsage struct {
	XMLName             xml.Name  `xml:"CurrentPeriodUsage"`
	DeviceMacId         uhexint64 `xml:"DeviceMacId"`
	MeterMacId          uhexint64 `xml:"MeterMacId"`
	TimeStamp           timestamp `xml:"TimeStamp"`
	CurrentUsage        uhexint32 `xml:"CurrentUsage"`
	Multiplier          uhexint32 `xml:"Multiplier"`
	Divisor             uhexint32 `xml:"Divisor"`
	DigitsRight         uhexint8  `xml:"DigitsRight"`
	DigitsLeft          uhexint8  `xml:"DigitsLeft"`
	SuppressLeadingZero ynbool    `xml:"SuppressLeadingZero"`
	StartDate           timestamp `xml:"StartDate"`
}

// Notify: LastPeriodUsage
type LastPeriodUsage struct {
	XMLName             xml.Name  `xml:"LastPeriodUsage"`
	DeviceMacId         uhexint64 `xml:"DeviceMacId"`
	MeterMacId          uhexint64 `xml:"MeterMacId"`
	LastUsage           uhexint32 `xml:"LastUsage"`
	Multiplier          uhexint32 `xml:"Multiplier"`
	Divisor             uhexint32 `xml:"Divisor"`
	DigitsRight         uhexint8  `xml:"DigitsRight"`
	DigitsLeft          uhexint8  `xml:"DigitsLeft"`
	SuppressLeadingZero ynbool    `xml:"SuppressLeadingZero"`
	StartDate           timestamp `xml:"StartDate"`
	EndDate             timestamp `xml:"EndDate"`
}

// Notify: ProfileData
type ProfileData struct {
	XMLName                  xml.Name    `xml:"ProfileData"`
	DeviceMacId              uhexint64   `xml:"DeviceMacId"`
	MeterMacId               uhexint64   `xml:"MeterMacId"`
	EndTime                  timestamp   `xml:"EndTime"`
	Status                   uhexint8    `xml:"Status"`
	ProfileIntervalPeriod    uint8       `xml:"ProfileIntervalPeriod"`
	NumberOfPeriodsDelivered uhexint8    `xml:"NumberOfPeriodsDelivered"`
	IntervalData             []uhexint32 `xml:"IntervalData"`
}
