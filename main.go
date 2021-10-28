// drawioToDynamicsNAV2018 project main.go
package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"os"

	projModel "drawioToDynamicsNAV2018/model"
)

func main() {
	byteValue := LoadFile()
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

	tables := make([]projModel.Nav2018Table, 0)
	for i := 0; i < len(mxfile.Diagrams[0].MxGraphModel.MxCells); i++ {
		cellValue = mxfile.Diagrams[0].MxGraphModel.MxCells[i].Value
		cellId = mxfile.Diagrams[0].MxGraphModel.MxCells[i].Id
		cellParent = mxfile.Diagrams[0].MxGraphModel.MxCells[i].Parent

		if cellParent == "1" { // table
			tableId += 1
			fieldIdInTable = 0
			tableContainerId = cellId

			fmt.Println("Create new table ", cellValue)
			var newTable projModel.Nav2018Table
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
				newField := new(projModel.Nav2018Field)
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

	for i := 0; i < len(tables); i++ {
		tables[i].Print()
	}

	fmt.Scanln()

}

func LoadFile() []byte {
	var input string

	fmt.Println("Print file name")
	fmt.Println("data.xml file will be used as a default file name")
	fmt.Scanln(&input)

	var fileName string
	if input != "" {
		fileName = input
	} else {
		fileName = "data.xml"
	}
	dataFile, err := os.Open(fileName)
	if err != nil {
		panic(err)
	}
	defer dataFile.Close()

	byteValue, _ := ioutil.ReadAll(dataFile)
	return byteValue
}

func UnmarshalData(byteValue []byte) projModel.Mxfile {
	var mxfile projModel.Mxfile
	err := xml.Unmarshal(byteValue, &mxfile)
	if err != nil {
		panic(err)
	}

	return mxfile
}
