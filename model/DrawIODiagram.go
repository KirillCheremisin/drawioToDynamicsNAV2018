package model

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"os"
)

// Diagram exported from draw.io web site
type DrawIODiagram struct {
}

func (table DrawIODiagram) ParseDiagram(fileName string) (tables []Table) {

	byteValue := LoadFile(fileName)
	var mxfile = UnmarshalData(byteValue)
	fmt.Println(mxfile)

	var tableContainerId string
	var fieldContainerId string
	var tableId int
	var PKProperty bool
	var fieldIdInTable int
	tableContainerId = "-1"
	fieldContainerId = "-1"
	var cellValue string
	var cellId string
	var cellParent string

	tables = make([]Table, 0)
	for i := 0; i < len(mxfile.Diagrams[0].MxGraphModel.MxCells); i++ {
		cellValue = mxfile.Diagrams[0].MxGraphModel.MxCells[i].Value
		cellId = mxfile.Diagrams[0].MxGraphModel.MxCells[i].Id
		cellParent = mxfile.Diagrams[0].MxGraphModel.MxCells[i].Parent

		if cellParent == "1" { // table
			tableId += 1
			fieldIdInTable = 0
			tableContainerId = cellId

			fmt.Println("Create new table ", cellValue)
			var newTable Table
			newTable.Id = tableId
			newTable.Name = cellValue

			fmt.Println("Append the table into slice tables")
			tables = append(tables, newTable)
			fmt.Println("")
		} else if cellParent == tableContainerId {
			fieldContainerId = cellId
		} else if cellParent == fieldContainerId {
			if PKProperty { // first cell - PK, second cell - Name
				PKProperty = false
			} else {
				PKProperty = true
			}

			if PKProperty {
				newField := new(Field)
				if cellValue == "PK" {
					newField.IsInPrimaryKey = true
				}

				tables[tableId-1].Fields = append(tables[tableId-1].Fields, newField)
			} else {
				tables[tableId-1].Fields[fieldIdInTable].Name = cellValue
				tables[tableId-1].Fields[fieldIdInTable].Id = fieldIdInTable + 1
				fieldIdInTable += 1
			}
		}
	}
	return tables
}

func LoadFile(fileName string) []byte {
	dataFile, err := os.Open(fileName)
	if err != nil {
		panic(err)
	}
	defer dataFile.Close()

	byteValue, _ := ioutil.ReadAll(dataFile)
	return byteValue
}

func UnmarshalData(byteValue []byte) Mxfile {
	var mxfile Mxfile
	err := xml.Unmarshal(byteValue, &mxfile)
	if err != nil {
		panic(err)
	}

	return mxfile
}
