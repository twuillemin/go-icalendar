package entity

import "time"

type Event struct {
	UID         string
	Summary     string
	Description string
}

type Todo struct {
}

type Journal struct {
}

type FreeBusy struct {
}

type Timezone struct {
	ID           string
	URL          string
	LastModified time.Time
	Properties   []TimezoneProperty
	XProps       []string
	IANAProp     []string
}

type HourMinute struct {
	Hour   int
	Minute int
}

type TimezonePropertyType int

const (
	Standard TimezonePropertyType = 1
	Daylight TimezonePropertyType = 2
)

type TimezoneProperty struct {
	Type        TimezonePropertyType
	StartLocal  time.Time
	OffsetFrom  HourMinute
	OffsetTo    HourMinute
	Rule        string
	Comments    []string
	RDatesLocal []time.Time
	Names       []string
	XProps      []string
	IANAProp    []string
}

type Calendar struct {
	Scale      string
	Method     string
	ProductID  string
	Version    string
	Events     []Event
	Todos      []Todo
	Journals   []Journal
	FreeBusies []FreeBusy
	Timezones  []Timezone
}
