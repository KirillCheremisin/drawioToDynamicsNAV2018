package interfaces

import . "drawioToDynamicsNAV2018/model"

// Exports tables into new files in current directory
// One file per table
type TableExporter interface {
	ExportTable([]Field)
}
