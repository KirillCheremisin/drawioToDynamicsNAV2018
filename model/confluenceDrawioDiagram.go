package model

import (
	"encoding/xml"
	"fmt"
	"io"
	"os"
	"strings"
)

// Diagram exported from draw.io extention for Confluence
type ConfluenceDrawIODiagram struct {
}

func (table ConfluenceDrawIODiagram) ParseDiagram(fileName string) (tables []Table) {
	file, err := os.Open(fileName)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	tables = make([]Table, 0)
	var currentTable Table
	var fieldID int
	var tableId int
	var fieldName string
	var fieldType string
	var fieldLength string
	var cellValue string
	var FieldXmlID string

	decoder := xml.NewDecoder(file)
	for {
		token, err := decoder.Token()
		if err == io.EOF {
			break
		}
		if err != nil {
			panic(err)
		}
		if token == nil {
			break
		}

		switch tp := token.(type) {
		case xml.StartElement:
			if tp.Name.Local == "mxCell" {
				fieldName = ""
				fieldType = ""
				fieldLength = ""

				var mxCell MxCell
				decoder.DecodeElement(&mxCell, &tp)

				if (mxCell.Parent == currentTable.XmlID) &&
					(mxCell.Parent != "") &&
					(mxCell.Parent != "0") { // field definition
					fieldID += 1

					var currentField Field
					currentField.Id = fieldID
					currentField.XmlID = mxCell.Id

					// type and length
					cellValue = mxCell.Value
					typeSeparatorPos := strings.Index(cellValue, ":")
					if typeSeparatorPos > -1 {
						fieldName = cellValue[:typeSeparatorPos]
						fieldType = cellValue[typeSeparatorPos+1:]

						// length
						i1 := strings.Index(fieldType, "[")
						if i1 > -1 {
							i2 := strings.Index(fieldType, "]")
							//if i2 > -1 {
							fieldLength = fieldType[i1+1 : i2]
							fieldType = fieldType[:i1]
							//}
						}

					} else {
						fieldName = mxCell.Value
					}

					currentField.Name = strings.TrimSpace(fieldName)
					currentField.Type = strings.TrimSpace(fieldType)
					currentField.Length = strings.TrimSpace(fieldLength)

					FieldXmlID = mxCell.Id
					tables[tableId-1].Fields = append(tables[tableId-1].Fields, &currentField)
				}

				if (mxCell.Parent == FieldXmlID) &&
					(mxCell.Parent != "0") &&
					(mxCell.Parent != "") { // is primary
					tables[tableId-1].Fields[fieldID-1].IsInPrimaryKey = (mxCell.Value == "PK")
				}

			}

			if tp.Name.Local == "object" {
				var mxObj MxObject
				decoder.DecodeElement(&mxObj, &tp)

				if mxObj.MxCell.Parent == "1" { // new table definition
					fieldID = 0
					tableId += 1

					fmt.Println("Create new table ", mxObj.Label)
					currentTable.Id = tableId
					currentTable.Name = mxObj.Label
					currentTable.XmlID = mxObj.Id
					tables = append(tables, currentTable)
				}
			}
		}
	}

	return tables
}
