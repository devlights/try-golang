package types

import "encoding/xml"

type (
	Version struct {
		Major int `xml:"major,attr"`
		Minor int `xml:"minor,attr"`
	}

	Language struct {
		Name    string  `xml:"name,attr"`
		PrintFn string  `xml:"printfn"`
		Version Version `xml:"version"`
	}

	XmlData struct {
		XMLName   xml.Name   `xml:"data"`
		Languages []Language `xml:"languages"`
	}
)
