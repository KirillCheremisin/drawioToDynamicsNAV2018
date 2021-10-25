package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	projModel "main/model"
	"os"
)

func main() {
	dataFile, err := os.Open("data.xml")
	if err != nil {
		panic(err)
	}
	defer dataFile.Close()
	byteValue, _ := ioutil.ReadAll(dataFile)

	var mxfile projModel.Mxfile
	err = xml.Unmarshal(byteValue, &mxfile)
	if err != nil {
		panic(err)
	}
	fmt.Println(mxfile)

	tables := make([]projModel.Nav2018Table, 0)
	for i := 0; i < len(mxfile.Diagrams[0].MxGraphModel.MxCells); i++ {
		// fmt.Println("cell ", i, " ID: "+mxfile.Diagrams[0].MxGraphModel.MxCells[i].Id)
		// fmt.Println("cell ", i, " Parent: ", mxfile.Diagrams[0].MxGraphModel.MxCells[i].Parent)
		// fmt.Println("cell ", i, " Value: "+mxfile.Diagrams[0].MxGraphModel.MxCells[i].Value)

		if mxfile.Diagrams[0].MxGraphModel.MxCells[i].Parent == "1" {
			fmt.Println("Create new table ", mxfile.Diagrams[0].MxGraphModel.MxCells[i].Value)
			var currentTable projModel.Nav2018Table
			currentTable.Id = i
			currentTable.Name = mxfile.Diagrams[0].MxGraphModel.MxCells[i].Value

			fmt.Println("Append the table into slice tables")
			tables = append(tables, currentTable)
			fmt.Println("")
		}

	}

	fmt.Println("Print tables")
	fmt.Println(tables)

}
