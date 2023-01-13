package outputgenerators

import (
	"encoding/xml"
	"os"
)

const (
	Header = `<?xml version="1.0" encoding="UTF-8"?>` + "\n"
)

type CoreEmulab struct{}

type Scenario struct {
	XMLName  xml.Name  `xml:"scenario"`
	Networks []Network `xml:"networks`
}

type Network struct {
	XMLName xml.Name `xml:"network"`
	Id      int      `xml:"id,attr"`
	Name    string   `xml:"name,attr"`
	Type    string   `xml:"type,attr"`
}

// generates a XML and a conf configuartion for CoreEmulab with a given experiment
func (c CoreEmulab) Generate() {

	nt := Network{
		Id:   1,
		Name: "AdHoc",
		Type: "WIRELESS",
	}
	scenario := &Scenario{}
	scenario.Networks = []Network{nt}

	data, err := os.Create("coreemulab.xml")
	if err != nil {
		panic(err)
	}
	data.WriteString(xml.Header)
	encoder := xml.NewEncoder(data)
	encoder.Indent("", "\t")
	err = encoder.Encode(&scenario)
	if err != nil {
		panic(err)
	}

}
