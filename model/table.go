package model

import "fmt"

// Field in a table
type Field struct {
	Id             int
	IsInPrimaryKey bool
	Name           string
}

// DB Table
type Table struct {
	Id     int
	Name   string
	Fields []*Field
}

func (table Table) Print() {
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
