// drawioToDynamicsNAV2018 project main.go
package main

import (
	"drawioToDynamicsNAV2018/interfaces"
	projModel "drawioToDynamicsNAV2018/model"
	"fmt"
)

func main() {
	fileName := SelectFileName()

	diagram := SelectDiagramParser()
	tables := diagram.ParseDiagram(fileName)
	for i := 0; i < len(tables); i++ {
		tables[i].Print()
	}

	var nav2018TableExporter projModel.Nav2018TableExporter
	for i := 0; i < len(tables); i++ {
		nav2018TableExporter.ExportTable(tables[i])
	}

	fmt.Scanln()
}

func SelectFileName() string {
	const defaultFileName string = "BMORA1.xml" //"data.xml"
	var input string

	fmt.Println("Print file name")
	fmt.Println(defaultFileName, " file will be used as a default file name")
	//fmt.Scanln(&input)

	var fileName string
	if input != "" {
		fileName = input
	} else {
		fileName = defaultFileName
	}
	return fileName
}

func SelectDiagramParser() (parser interfaces.DiagramParser) {
	//var diagram projModel.DrawIODiagram
	var diagram projModel.ConfluenceDrawIODiagram
	return diagram
}
