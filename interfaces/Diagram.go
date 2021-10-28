package interfaces

import projModel "drawioToDynamicsNAV2018/model"

type DiagramParser interface {
	ParseDiagram(mxfile projModel.Mxfile) (tables []projModel.Nav2018Table)
}
