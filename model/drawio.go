package model

import "encoding/xml"

type MxCell struct {
	XMLName xml.Name `xml:"mxCell"`
	Id      string   `xml:"id,attr"`
	Value   string   `xml:"value,attr"`
	Parent  string   `xml:"parent,attr"`
}

type MxGraphModel struct {
	XMLName   xml.Name  `xml:"mxGraphModel"`
	MxCells   []*MxCell `xml:"root>mxCell"`
	PageWidth int       `xml:"pageWidth,attr"`
}

type Diagram struct {
	XMLName      xml.Name     `xml:"diagram"`
	MxGraphModel MxGraphModel `xml:"mxGraphModel"`
	Id           string       `xml:"id,attr"`
	Name         string       `xml:"name,attr"`
}

type Mxfile struct {
	XMLName  xml.Name   `xml:"mxfile"`
	Diagrams []*Diagram `xml:"diagram"`
	Host     string     `xml:"host,attr"`
	Modified string     `xml:"modified,attr"`
}
