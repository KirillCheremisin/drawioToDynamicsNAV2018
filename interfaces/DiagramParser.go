package interfaces

import projModel "drawioToDynamicsNAV2018/model"

type DiagramParser interface {
	ParseDiagram(fileName string) (tables []projModel.Table)
}
