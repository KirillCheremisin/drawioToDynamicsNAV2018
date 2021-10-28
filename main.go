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

	var diagram projModel.DrawIODiagram
	tables := diagram.ParseDiagram(mxfile)
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
