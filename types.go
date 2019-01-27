package goraven

import (
	"encoding/xml"
	"strconv"
	"time"
)

func unmarshalHexXML(d *xml.Decoder, start xml.StartElement, bits int) (uint64, error) {
    var v string
    d.DecodeElement(&v, &start)
	utmp, err := strconv.ParseUint(v, 0, bits)
    if err != nil {
        return 0, err
    }
    return utmp, nil
}

type uhexint64 uint64
type uhexint32 uint32
type uhexint16 uint16
type uhexint8  uint8

func (u *uhexint64) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	v, err := unmarshalHexXML(d, start, 64)
	if err == nil {
		*u = uhexint64(v)
    }
    return err
}

func (u *uhexint32) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	v, err := unmarshalHexXML(d, start, 32)
	if err == nil {
		*u = uhexint32(v)
    }
    return err
}

func (u *uhexint16) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	v, err := unmarshalHexXML(d, start, 16)
	if err == nil {
		*u = uhexint16(v)
    }
    return err
}

func (u *uhexint8) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	v, err := unmarshalHexXML(d, start, 8)
	if err == nil {
		*u = uhexint8(v)
    }
    return err
}

type timestamp struct {
	time.Time
}

func (t *timestamp) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
    var v string
    d.DecodeElement(&v, &start)
	itmp, err := strconv.ParseInt(v, 0, 32)
    if err != nil {
        return err
    }
	// January 1, 2000 Midnight UTC
    *t = timestamp{time.Unix(itmp+946684800, 0)}
    return nil
}

type ynbool bool

func (b *ynbool) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
    var v string
    d.DecodeElement(&v, &start)
	if v == "Y" {
		*b = true
	} else if v == "F" {
		*b = false
	}
    return nil
}
