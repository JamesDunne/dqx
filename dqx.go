// main
package dqx

import (
	"encoding/xml"
	"math"
)

type Parameter struct {
	Type     int    `xml:"Type,attr"`
	Instance int    `xml:"Instance,attr"`
	Index    int    `xml:"Index,attr"`
	Name     string `xml:"name,attr"`
	Value    string `xml:"value,attr"`
	Text     string `xml:"text,attr,omitempty"`
}

type Effect struct {
	Type       int         `xml:"Type,attr"`
	Instance   int         `xml:"Instance,attr"`
	Version    int         `xml:"Version,attr"`
	Name       string      `xml:"Name,attr"`
	Parameters []Parameter `xml:"Parameter"`
}

type Channel struct {
	Type     int      `xml:"Type,attr"`
	Instance int      `xml:"Instance,attr"`
	Version  int      `xml:"Version,attr"`
	Name     string   `xml:"Name,attr"`
	Effects  []Effect `xml:"Effect"`
}

type Config struct {
	Version    string    `xml:"Version,attr"`
	DeviceType int       `xml:"DeviceType,attr"`
	Name       string    `xml:"Name,attr"`
	Channels   []Channel `xml:"Channel"`
}

type Header struct {
	Path        string `xml:"Path,attr"`
	Uuid        string `xml:"Uuid,attr"`
	Description string `xml:"Description,attr"`
	Version     string `xml:"Version,attr"`
	Category    string `xml:"Category,attr"`
	Subcategory string `xml:"Subcategory,attr"`
	Info        string `xml:"Info,attr"`
}

type DQX struct {
	XMLName xml.Name `xml:"DQData"`
	Version string   `xml:"Version,attr"`
	Header  Header   `xml:"Header"`
	Config  Config   `xml:"Data>Config"`
}

func ParseDQX(f []byte) (dqx *DQX, err error) {
	dqx = new(DQX)
	err = xml.Unmarshal(f, dqx)
	if err != nil {
		return nil, err
	}

	return dqx, err
}

func ConvertQtoBandwidth(q float64) float64 {
	q_sqr := q * q
	bw := math.Log2((2.0*q_sqr+1.0)/(2.0*q_sqr) + math.Sqrt(math.Pow((2.0*q_sqr+1.0)/q_sqr, 2.0)/4.0-1.0))
	return bw
}
