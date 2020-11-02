package main

import "time"

type InvoiceDate time.Time

func (d InvoiceDate) YYYYMMDD() string {
	return time.Time(d).Format("2006-01-02")
}

func (d InvoiceDate) String() string {
	return time.Time(d).Format("January 2 2006")
}

func (d *InvoiceDate) UnmarshalJSON(buf []byte) error {
	var t time.Time
	var err error
	// empty string produces zero time value
	if len(buf) > 0 {
		err = t.UnmarshalJSON(buf)
	}
	*d = InvoiceDate(t)
	return err
}

func (d InvoiceDate) MarshalJSON() ([]byte, error) {
	return d.MarshalJSON()
}
