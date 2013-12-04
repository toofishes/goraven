package goraven

// Reinitialize the XML parser on the device. Use this command when first
// connecting to the RAVEn prior to sending any other commands. Initialization
// is not required, but will speed up the initial connection.
func (r *Raven) Initialize() error {
	return r.simpleCommand("initialize")
}

// Force the RAVEn to go through the start-up sequence. This command is
// useful for capturing any diagnostic information sent during the start-up
// sequence.
func (r *Raven) Restart() error {
	return r.simpleCommand("restart")
}

// Reset the RAVEn. This command will erase the commissioning data and force a
// restart. On restart, the RAVEn will begin the commissioning cycle.
func (r *Raven) FactoryReset() error {
	return r.simpleCommand("factory_reset")
}

// Get the RAVEn connection information. The RAVEn will send a ConnectionStatus
// notification in response. The RAVEn continuously sends ConnectionStatus
// during the start-up sequence and during the join/re-join sequence for
// diagnostic purposes.
func (r *Raven) GetConnectionStatus() error {
	return r.simpleCommand("get_connection_status")
}

// Get RAVEn configuration information. The RAVEn will send a DeviceInfo
// notification in response.
func (r *Raven) GetDeviceInfo() error {
	return r.simpleCommand("get_device_info")
}

// Get the RAVEn scheduler information. The RAVEn will send the ScheduleInfo
// notification in response; or, RAVEn will send a series of ScheduleInfo
// notifications if the Event field is omitted.
func (r *Raven) GetSchedule() error {
	return r.simpleCommand("get_schedule")
}

// Update the RAVEn scheduler. The command options include setting the
// frequency of the command in seconds, and disabling the event. If the event
// is disabled the frequency is set to 0xFFFFFFFF
func (r *Raven) SetSchedule(event string, frequency int, enabled bool) {
}

// Reset the RAVEn scheduler to default settings. If the Event field is set,
// only that schedule item is reset to default values; otherwise all schedule
// items are reset to their default values.
func (r *Raven) SetScheduleDefault(event string) {
}

// Get the list of meters the RAVEn is connected to. The RAVEn will send a
// MeterList notification in response.
func (r *Raven) GetMeterList() {
}

// Notify: ConnectionStatus
type connectionStatus struct {
	XMLName      xml.Name `xml:"ConnectionStatus"`
	DeviceMacId  string   `xml:"DeviceMacId"`
	MeterMacId   string   `xml:"MeterMacId"`
	Description  string   `xml:"Description,omitempty"`
	StatusCode   string   `xml:"StatusCode,omitempty"`
	ExtPanId     string   `xml:"ExtPanId,omitempty"`
	Channel      int      `xml:"Channel,omitempty"`
	ShortAddr    string   `xml:"ShortAddr,omitempty"`
	LinkStrength string   `xml:"LinkStrength"`
}

// Notify: DeviceInfo
type deviceInfo struct {
	XMLName      xml.Name `xml:"DeviceInfo"`
	DeviceMacId  string   `xml:"DeviceMacId"`
	InstallCode  string   `xml:"InstallCode"`
	LinkKey      string   `xml:"LinkKey"`
	FWVersion    string   `xml:"FWVersion"`
	HWVersion    string   `xml:"HWVersion"`
	ImageType    string   `xml:"ImageType"`
	Manufacturer string   `xml:"Manufacturer"`
	ModelId      string   `xml:"ModelId"`
	DateCode     string   `xml:"DateCode"`
}

// Notify: ScheduleInfo
type scheduleInfo struct {
	XMLName     xml.Name `xml:"ScheduleInfo"`
	DeviceMacId string   `xml:"DeviceMacId"`
	MeterMacId  string   `xml:"MeterMacId,omitempty"`
	Event       string   `xml:"Event"`
	Frequency   string   `xml:"Frequency"`
	Enabled     string   `xml:"Enabled"`
}

// Notify: MeterList
type meterList struct {
	XMLName     xml.Name `xml:"MeterList"`
	DeviceMacId string   `xml:"DeviceMacId"`
	MeterMacId  []string `xml:"MeterMacId,omitempty"`
}