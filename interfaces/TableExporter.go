package interfaces

import (
	. "drawioToDynamicsNAV2018/model"
	"io"
)

// Exports tables into new files in current directory
// One file per table
type TableExporter interface {
	ExportTable(table Table, writer io.Writer)
}
