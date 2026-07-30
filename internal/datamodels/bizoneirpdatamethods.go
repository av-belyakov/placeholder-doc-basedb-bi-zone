package datamodels

import (
	"fmt"
	"slices"
	"strings"

	"github.com/av-belyakov/placeholder_doc-basedb_bi.zone/internal/supportingfunctions"
)

func NewBiZoneIRPData() *BiZoneIRPData {
	return &BiZoneIRPData{
		SnortSid:   []uint64(nil),
		AllIPHome:  []string(nil),
		AllSensors: []uint64(nil),
	}
}

func (d *BiZoneIRPData) Get() *BiZoneIRPData {
	return d
}

// GetId для поля ID
func (d *BiZoneIRPData) GetId() string {
	return d.ID
}

// SetId для поля ID
func (d *BiZoneIRPData) SetId(id string) {
	d.ID = id
}

// SetAnyID для поля ID
func (d *BiZoneIRPData) SetAnyID(a any) {
	d.ID = fmt.Sprint(a)
}

// GetTitle для поля Title
func (d *BiZoneIRPData) GetTitle() string {
	return d.Title
}

// SetTitle для поля Title
func (d *BiZoneIRPData) SetTitle(title string) {
	d.Title = title
}

// SetAnyTitle для поля Title
func (d *BiZoneIRPData) SetAnyTitle(a any) {
	d.Title = fmt.Sprint(a)
}

// GetIPHome для поля IPHome
func (d *BiZoneIRPData) GetIPHome() string {
	return d.IPHome
}

// SetIPHome для поля IPHome
func (d *BiZoneIRPData) SetIPHome(ipHome string) {
	d.IPHome = ipHome
}

// SetAnyIPHome для поля IPHome
func (d *BiZoneIRPData) SetAnyIPHome(a any) {
	d.IPHome = fmt.Sprint(a)
}

// GetIPExter для поля IPExter
func (d *BiZoneIRPData) GetIPExter() string {
	return d.IPExter
}

// SetIPExter для поля IPExter
func (d *BiZoneIRPData) SetIPExter(ipExter string) {
	d.IPExter = ipExter
}

// SetAnyIPExter для поля IPExter
func (d *BiZoneIRPData) SetAnyIPExter(a any) {
	d.IPExter = fmt.Sprint(a)
}

// GetURLHTTP для поля URLHTTP
func (d *BiZoneIRPData) GetURLHTTP() string {
	return d.URLHTTP
}

// SetURLHTTP для поля URLHTTP
func (d *BiZoneIRPData) SetURLHTTP(urlHTTP string) {
	d.URLHTTP = urlHTTP
}

// SetAnyURLHTTP для поля URLHTTP
func (d *BiZoneIRPData) SetAnyURLHTTP(a any) {
	d.URLHTTP = fmt.Sprint(a)
}

// GetEventType для поля EventType
func (d *BiZoneIRPData) GetEventType() string {
	return d.EventType
}

// SetEventType для поля EventType
func (d *BiZoneIRPData) SetEventType(eventType string) {
	d.EventType = eventType
}

// SetAnyEventType для поля EventType
func (d *BiZoneIRPData) SetAnyEventType(a any) {
	d.EventType = fmt.Sprint(a)
}

// GetURLArkime для поля URLArkime
func (d *BiZoneIRPData) GetURLArkime() string {
	return d.URLArkime
}

// SetURLArkime для поля URLArkime
func (d *BiZoneIRPData) SetURLArkime(urlArkime string) {
	d.URLArkime = urlArkime
}

// SetAnyURLArkime для поля URLArkime
func (d *BiZoneIRPData) SetAnyURLArkime(a any) {
	d.URLArkime = fmt.Sprint(a)
}

// GetURLFTP для поля URLFTP
func (d *BiZoneIRPData) GetURLFTP() string {
	return d.URLFTP
}

// SetURLFTP для поля URLFTP
func (d *BiZoneIRPData) SetURLFTP(urlFTP string) {
	d.URLFTP = urlFTP
}

// SetAnyURLFTP для поля URLFTP
func (d *BiZoneIRPData) SetAnyURLFTP(a any) {
	d.URLFTP = fmt.Sprint(a)
}

// GetSensor для поля Sensor
func (d *BiZoneIRPData) GetSensor() uint64 {
	return d.Sensor
}

// SetSensor для поля Sensor
func (d *BiZoneIRPData) SetSensor(sensor uint64) {
	d.Sensor = sensor
}

// SetAnySensor для поля Sensor
func (d *BiZoneIRPData) SetAnySensor(a any) {
	if v, ok := a.(int); ok {
		d.Sensor = uint64(v)

		return
	}

	if v, ok := a.(float32); ok {
		d.Sensor = uint64(v)

		return
	}

	if v, ok := a.(float64); ok {
		d.Sensor = uint64(v)

		return
	}

	if v, ok := a.(uint64); ok {
		d.Sensor = v
	}
}

// GetAllIPExt для поля AllIPExt
func (d *BiZoneIRPData) GetAllIPExt() uint64 {
	return d.AllIPExt
}

// SetAllIPExt для поля AllIPExt
func (d *BiZoneIRPData) SetAllIPExt(allIPExt uint64) {
	d.AllIPExt = allIPExt
}

// SetAnyAllIPExt для поля AllIPExt
func (d *BiZoneIRPData) SetAnyAllIPExt(a any) {
	if v, ok := a.(int); ok {
		d.AllIPExt = uint64(v)

		return
	}

	if v, ok := a.(float32); ok {
		d.AllIPExt = uint64(v)

		return
	}

	if v, ok := a.(float64); ok {
		d.AllIPExt = uint64(v)

		return
	}

	if v, ok := a.(uint64); ok {
		d.AllIPExt = v
	}
}

// GetResponseTeam для поля ResponseTeam
func (d *BiZoneIRPData) GetResponseTeam() uint64 {
	return d.ResponseTeam
}

// SetResponseTeam для поля ResponseTeam
func (d *BiZoneIRPData) SetResponseTeam(responseTeam uint64) {
	d.ResponseTeam = responseTeam
}

// SetAnyResponseTeam для поля ResponseTeam
func (d *BiZoneIRPData) SetAnyResponseTeam(a any) {
	if v, ok := a.(int); ok {
		d.ResponseTeam = uint64(v)

		return
	}

	if v, ok := a.(float32); ok {
		d.ResponseTeam = uint64(v)

		return
	}

	if v, ok := a.(float64); ok {
		d.ResponseTeam = uint64(v)

		return
	}

	if v, ok := a.(uint64); ok {
		d.ResponseTeam = v
	}
}

// GetSnortSids список sid snort
func (d *BiZoneIRPData) GetSnortSids() []uint64 {
	return d.SnortSid
}

// SetSnortSid список sid snort
func (d *BiZoneIRPData) SetSnortSids(sids []uint64) {
	d.SnortSid = sids
}

// SetSnortSid записывает значение в список SnortSid
func (d *BiZoneIRPData) SetSnortSid(sid uint64) {
	if d.SnortSid == nil {
		d.SnortSid = []uint64(nil)
	}

	if slices.Contains(d.SnortSid, sid) {
		return
	}

	d.SnortSid = append(d.SnortSid, sid)
}

// SetAnySnortSid записывает значение в список SnortSid
func (d *BiZoneIRPData) SetAnySnortSid(a any) {
	v, err := supportingfunctions.GetUint64(a)
	if err != nil {
		return
	}

	d.SetSnortSid(v)
}

// GetAllSensors для поля AllSensors
func (d *BiZoneIRPData) GetAllSensors() []uint64 {
	return d.AllSensors
}

// SetAllSensors для поля AllSensors
func (d *BiZoneIRPData) SetAllSensors(allSensors []uint64) {
	d.AllSensors = allSensors
}

// SetAllSensor записывает значение в список AllSensors
func (d *BiZoneIRPData) SetAllSensor(sensor uint64) {
	if d.AllSensors == nil {
		d.AllSensors = []uint64(nil)
	}

	if slices.Contains(d.AllSensors, sensor) {
		return
	}

	d.SnortSid = append(d.SnortSid, sensor)
}

// SetAnyAllSensor записывает значение в список AllSensors
func (d *BiZoneIRPData) SetAnyAllSensor(a any) {
	v, err := supportingfunctions.GetUint64(a)
	if err != nil {
		return
	}

	d.SetAllSensor(v)
}

// GetAllIPHomes для поля AllIPHome
func (d *BiZoneIRPData) GetAllIPHomes() []string {
	return d.AllIPHome
}

// SetAllIPHomes для поля AllIPHomes
func (d *BiZoneIRPData) SetAllIPHomes(allIPHome []string) {
	d.AllIPHome = allIPHome
}

// SetAllIPHome добавляет значение AllIPHome в список
func (d *BiZoneIRPData) SetAllIPHome(ipHome string) {
	if d.AllIPHome == nil {
		d.AllIPHome = []string(nil)
	}

	if slices.Contains(d.AllIPHome, ipHome) {
		return
	}

	d.AllIPHome = append(d.AllIPHome, ipHome)
}

// SetAnyAllIPHome добавляет значение AllIPHome  в список
func (d *BiZoneIRPData) SetAnyAllIPHome(a any) {
	d.SetAllIPHome(fmt.Sprint(a))
}

// ToStringBeautiful форматированный вывод
func (d *BiZoneIRPData) ToStringBeautiful(num int) string {
	str := strings.Builder{}

	ws := supportingfunctions.GetWhitespace(num)

	str.WriteString(fmt.Sprintf("%s'_id': '%s'\n", ws, d.ID))
	str.WriteString(fmt.Sprintf("%s'title': '%s'\n", ws, d.Title))
	str.WriteString(fmt.Sprintf("%s'sensor': '%d'\n", ws, d.Sensor))
	str.WriteString(fmt.Sprintf("%s'ip_home': '%s'\n", ws, d.IPHome))
	str.WriteString(fmt.Sprintf("%s'ip_exter': '%s'\n", ws, d.IPExter))
	str.WriteString(fmt.Sprintf("%s'url____ftp': '%s'\n", ws, d.URLFTP))
	str.WriteString(fmt.Sprintf("%s'url___http': '%s'\n", ws, d.URLHTTP))
	str.WriteString(fmt.Sprintf("%s'event_type': '%s'\n", ws, d.EventType))
	str.WriteString(fmt.Sprintf("%s'url_arkime': '%s'\n", ws, d.URLArkime))
	str.WriteString(fmt.Sprintf("%s'all____ip_ext': '%d'\n", ws, d.AllIPExt))
	str.WriteString(fmt.Sprintf("%s'response_team': '%d'\n", ws, d.ResponseTeam))
	str.WriteString(fmt.Sprintf("%s'snort_sid': \n%s", ws, supportingfunctions.ToStringBeautifulSlice(num, d.SnortSid)))
	str.WriteString(fmt.Sprintf("%s'all_sensors': \n%s", ws, supportingfunctions.ToStringBeautifulSlice(num, d.AllSensors)))
	str.WriteString(fmt.Sprintf("%s'all__ip_home': \n%s", ws, supportingfunctions.ToStringBeautifulSlice(num, d.AllIPHome)))

	return str.String()
}
