package dqx

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"testing"
)

func Test_parseDQX(t *testing.T) {
	f, err := ioutil.ReadFile("/Users/jimdunne/Desktop/band/20170517/20170517.scene.dqx")
	if err != nil {
		panic(err)
	}

	dqx, err := parseDQX(f)
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

					bw := convertQtoBandwidth(q)
					fmt.Printf("    %3d[%2d];Bandwidth = \"%f\"\n", param.Type, param.Instance, bw)
				}
			}
		}
	}
}