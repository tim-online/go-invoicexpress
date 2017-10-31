package invoicexpress

import (
	"encoding/json"
	"time"

	"github.com/aodin/date"
)

// type Date date.Date

type Date struct {
	date.Date
}

func NewDate(year int, month time.Month, day int) Date {
	d := date.New(year, month, day)
	return Date{Date: d}
}

// func (d *Date) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
// 	t := time.Time(*d)
// 	return e.EncodeElement(t.Format("20060102"), start)
// }

func (d *Date) UnmarshalJSON(text []byte) (err error) {
	var value string
	err = json.Unmarshal(text, &value)
	if err != nil {
		return err
	}

	if value == "" {
		return nil
	}

	// 28-1-2008
	layout := "2/1/2006"
	time, err := time.Parse(layout, value)
	d.Date = date.FromTime(time)
	return err
}

func (d Date) MarshalSchema() string {
	if d.Equal(time.Time{}) {
		return ""
	}
	return d.Format("2/1/2006")
}

type Pagination struct {
	TotalPages   int `json:"total_pages"`
	PerPage      int `json:"per_page"`
	CurrentPage  int `json:"current_page"`
	TotalEntries int `json:"total_entries"`
}
