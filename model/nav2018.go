package model

type Nav2018Field struct {
	Id             int
	IsInPrimaryKey bool
	Name           string
}

type Nav2018Table struct {
	Id     int
	Name   string
	Fields []*Nav2018Field
}
