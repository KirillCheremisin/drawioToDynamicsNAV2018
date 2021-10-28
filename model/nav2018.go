package model

import "fmt"

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

func (table Nav2018Table) Print() {
	fmt.Println("Table ", table.Id, " ", table.Name)
	for j := 0; j < len(table.Fields); j++ {
		fmt.Print("Field: ")
		fmt.Println(table.Fields[j].Id)
		fmt.Print("Is primary: ")
		fmt.Println(table.Fields[j].IsInPrimaryKey)
		fmt.Print("Name: ")
		fmt.Println(table.Fields[j].Name)
	}
	fmt.Println()
}
