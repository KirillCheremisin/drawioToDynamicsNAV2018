package model

import (
	"fmt"
	"io"
	"os"
)

type Nav2018TableExporter struct {
}

func (Nav2018TableExporter) ExportTable(table Table) {
	writer := os.Stdout

	fmt.Fprintf(writer, "OBJECT Table %d %s\n", table.Id, table.Name)
	fmt.Fprintln(writer, "{")
	defer fmt.Fprintln(writer, "}")

	ExportFields(writer, table)
	ExportPrimaryKey(writer, table)
}

func ExportFields(writer io.Writer, table Table) {
	fmt.Fprintln(writer, "FIELDS")
	fmt.Fprintln(writer, "{")
	for _, field := range table.Fields {
		ExportField(writer, *field)
	}
	fmt.Fprintln(writer, "}")
}

func ExportField(writer io.Writer, field Field) {
	fmt.Fprintf(
		writer,
		"{ %d ;%s ;%s ; DataClassification=ToBeClassified }\n",
		field.Id, field.Name, field.Type+field.Length)
}

func ExportPrimaryKey(writer io.Writer, table Table) {
	fmt.Fprintln(writer, "KEYS")
	fmt.Fprintln(writer, "{")
	defer fmt.Fprintln(writer, "}")

	var PK string
	for _, field := range table.Fields {
		if field.IsInPrimaryKey {
			PK += field.Name
			PK += ","
		}
		if !field.IsInPrimaryKey {
			break
		}

		if len(PK) > 1 {
			PK = string(PK[0 : len(PK)-1])
		}
	}

	fmt.Fprintf(writer, "{  ;%s ;Clustered=Yes }\n", PK)
}
