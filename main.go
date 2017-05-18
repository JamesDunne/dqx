// main
package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"math"
	"strconv"
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

func main() {
	f, err := ioutil.ReadFile("/Users/jimdunne/Desktop/band/20170517/20170517.scene.dqx")
	if err != nil {
		panic(err)
	}

	dqx := new(DQX)
	err = xml.Unmarshal(f, dqx)
	if err != nil {
		panic(err)
	}

	for _, channel := range dqx.Config.Channels {
		if channel.Type != 1 {
			continue
		}
		fmt.Printf("%s\n", channel.Name)
		for _, effect := range channel.Effects {
			fmt.Printf("  %3d;%s\n", effect.Type, effect.Name)
			for _, param := range effect.Parameters {
				value := param.Value
				if param.Text != "" {
					value = param.Text
				}
				fmt.Printf("    %3d[%2d];%s = \"%s\"\n", param.Type, param.Instance, param.Name, value)
				if param.Type == 8 {
					// Convert Q to bandwidth:
					// log2( (2*(q^2)+1)/(2*(q^2)) + SQRT( ( ((2*(q^2)+1)/(q^2)) ^ 2) / 4 - 1 ) )
					q, err := strconv.ParseFloat(param.Value, 64)
					if err != nil {
						continue
					}

					q_sqr := q * q
					bw := math.Log2((2.0*q_sqr+1.0)/(2.0*q_sqr) + math.Sqrt(math.Pow((2.0*q_sqr+1.0)/q_sqr, 2.0)/4.0-1.0))
					fmt.Printf("    %3d[%2d];Bandwidth = \"%f\"\n", param.Type, param.Instance, bw)
				}
			}
		}
	}
}
