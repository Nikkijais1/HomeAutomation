package model

import "time"

type User struct {
	ID                 int
	Name               string
	Address            string
	Class              string
	Active             Boolean
	Email              string
	Password           string
	CreatedAt          time.Time
	UpdateAt           time.Time
	PrimaryContact     int
	AlternativeContact int
	OrganizationNo     int
	BuildingNo         int
	FloorNo            int
	RoomNo             int
}
