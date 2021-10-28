package model

import "fmt"

// Diagram exported from draw.io web site
type DrawIODiagram struct {
}

func (table DrawIODiagram) ParseDiagram(mxfile Mxfile) (tables []Nav2018Table) {
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

	tables = make([]Nav2018Table, 0)
	for i := 0; i < len(mxfile.Diagrams[0].MxGraphModel.MxCells); i++ {
		cellValue = mxfile.Diagrams[0].MxGraphModel.MxCells[i].Value
		cellId = mxfile.Diagrams[0].MxGraphModel.MxCells[i].Id
		cellParent = mxfile.Diagrams[0].MxGraphModel.MxCells[i].Parent

		if cellParent == "1" { // table
			tableId += 1
			fieldIdInTable = 0
			tableContainerId = cellId

			fmt.Println("Create new table ", cellValue)
			var newTable Nav2018Table
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
				newField := new(Nav2018Field)
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
