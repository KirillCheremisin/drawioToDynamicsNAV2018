// drawioToDynamicsNAV2018 project main.go
package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"os"

	projModel "./model"
)

func main() {
	byteValue := LoadFile()
	var mxfile = UnmarshalData(byteValue)
	fmt.Println(mxfile)

	var tableContainerId string
	var fieldContainerId string
	var tableId int
	var PKProperty bool
	var vieldIdInTable int
	tableContainerId = "-1"
	fieldContainerId = "-1"

	tables := make([]projModel.Nav2018Table, 0)
	for i := 0; i < len(mxfile.Diagrams[0].MxGraphModel.MxCells); i++ {
		// fmt.Println("cell ", i, " ID: "+mxfile.Diagrams[0].MxGraphModel.MxCells[i].Id)
		// fmt.Println("cell ", i, " Parent: ", mxfile.Diagrams[0].MxGraphModel.MxCells[i].Parent)
		// fmt.Println("cell ", i, " Value: "+mxfile.Diagrams[0].MxGraphModel.MxCells[i].Value)

		if mxfile.Diagrams[0].MxGraphModel.MxCells[i].Parent == "1" {
			tableId += 1
			vieldIdInTable = 0
			tableContainerId = mxfile.Diagrams[0].MxGraphModel.MxCells[i].Id

			fmt.Println("Create new table ", mxfile.Diagrams[0].MxGraphModel.MxCells[i].Value)
			var newField projModel.Nav2018Table
			newField.Id = i
			newField.Name = mxfile.Diagrams[0].MxGraphModel.MxCells[i].Value

			fmt.Println("Append the table into slice tables")
			tables = append(tables, newField)
			fmt.Println("")
		} else if mxfile.Diagrams[0].MxGraphModel.MxCells[i].Parent == tableContainerId {
			fieldContainerId = mxfile.Diagrams[0].MxGraphModel.MxCells[i].Id
		} else if mxfile.Diagrams[0].MxGraphModel.MxCells[i].Parent == fieldContainerId {
			if PKProperty {
				PKProperty = false
			} else {
				PKProperty = true
			}

			if PKProperty {
				newField := new(projModel.Nav2018Field)
				if mxfile.Diagrams[0].MxGraphModel.MxCells[i].Value == "PK" {
					newField.IsInPrimaryKey = true
				}

				tables[tableId-1].Fields = append(tables[tableId-1].Fields, newField)
			} else {
				tables[tableId-1].Fields[vieldIdInTable].Name = mxfile.Diagrams[0].MxGraphModel.MxCells[i].Value
				vieldIdInTable += 1
			}
		}
	}

	fmt.Println("Print tables")
	fmt.Println(tables)
	fmt.Println(len(tables))
	fmt.Println(tables[0].Fields[1].Name)

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
