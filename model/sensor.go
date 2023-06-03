package model

import "time"

type Sensor struct {
	ID             int
	SensorName     string
	SensorType     string
	Voltagemin     string
	voltagemax     string
	Reading        string
	ReadingUnit    string
	Capacity       string
	Switch         string
}
