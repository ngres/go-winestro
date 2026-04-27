package winestro

import (
	"encoding/xml"
	"time"
)

type timestamp time.Time

const timestampLayout = "2006-01-02 15:04:05"

func (t *timestamp) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	var s string
	if err := d.DecodeElement(&s, &start); err != nil {
		return err
	}
	parsedTime, err := time.Parse(timestampLayout, s)
	if err != nil {
		return err
	}
	*t = timestamp(parsedTime)
	return nil
}

func (t timestamp) MarshalJSON() ([]byte, error) {
	tt := time.Time(t)
	return tt.MarshalJSON() // default RFC3339Nano behavior
}

func (t *timestamp) UnmarshalJSON(data []byte) error {
	var tt time.Time
	if err := tt.UnmarshalJSON(data); err != nil {
		return err
	}
	*t = timestamp(tt)
	return nil
}
