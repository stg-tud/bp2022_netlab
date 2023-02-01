package experiment

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/pelletier/go-toml/v2"
	"github.com/stg-tud/bp2022_netlab/internal/movementpatterns"
)

// Loads the path string with toml file into experiment
func LoadFromFile(file string) Experiment {
	var exp Experiment
	buf, e := os.ReadFile(file)
	if e != nil {
		panic(e)
	}

	data := string(buf)
	move := "[NodeGroups.MovementModel]" //27 char
	node := "[[NodeGroups]]"
	copy := data
	var tmp Experiment
	noNodegroups := strings.Count(data,node)
	for k := 0; k < noNodegroups; k++ {

		//delete everything before the next MovementModel
		copy = strings.Replace(copy, copy[0:strings.Index(copy, move)], "", -1)

		//next index with Nodegroup
		next := strings.Index(copy, node)
		if next < 0 {
			next = len(copy)
		}
		//movementmodel is between
		movementSet := copy[27:next]

		
		data = strings.Replace(data, movementSet, "", -1)

		if strings.TrimSpace(movementSet) == "" {

			tmp.NodeGroups = append(exp.NodeGroups, NodeGroup{
				MovementModel: movementpatterns.Static{},
			})

		} else {

		
			maxpause := movementSet[strings.Index(movementSet, "MaxPause=")+9:]

			Maxpause, err := strconv.Atoi(strings.TrimSpace(maxpause))
			if err != nil {
				panic(e)
			}

			minspeed := movementSet[strings.Index(movementSet, "MinSpeed=")+9 : strings.Index(movementSet, "MaxSpeed")]

			Minspeed, err := strconv.Atoi(strings.TrimSpace(minspeed))
			if err != nil {
				panic(e)
			}

			maxspeed := movementSet[strings.Index(movementSet, "MaxSpeed=")+9 : strings.Index(movementSet, "MaxPause")]

			Maxspeed, err := strconv.Atoi(strings.TrimSpace(maxspeed))
			if err != nil {
				panic(e)
			}

			tmp.NodeGroups = append(tmp.NodeGroups, NodeGroup{
				MovementModel: movementpatterns.RandomWaypoint{
					MinSpeed: Minspeed,
					MaxSpeed: Maxspeed,
					MaxPause: Maxpause,
				},
			})
			
		}
		if next !=len(copy) {
			//delete before the next nodegroup
			copy = strings.Replace(copy, copy[0:strings.Index(copy, node)], "", -1)
		}
	}
	
	err := toml.Unmarshal([]byte(data),&exp)
	if err != nil{
		panic(err)
	}
	for i := 0; i < noNodegroups; i++ {
		exp.NodeGroups[i].MovementModel=tmp.NodeGroups[i].MovementModel
	}
	fmt.Println(exp.NodeGroups[3].MovementModel)
	for i := 0; i < len(exp.NodeGroups); i++ {

		if exp.NodeGroups[i].IPv4Net == "" {
			exp.NodeGroups[i].IPv4Net = defaultValues.IPv4Net
		}
		if exp.NodeGroups[i].IPv4Mask == 0 {
			exp.NodeGroups[i].IPv4Mask = defaultValues.IPv4Mask
		}
		if exp.NodeGroups[i].IPv6Net == "" {
			exp.NodeGroups[i].IPv6Net = defaultValues.IPv6Net
		}
		if exp.NodeGroups[i].IPv6Mask == 0 {
			exp.NodeGroups[i].IPv6Mask = defaultValues.IPv6Mask
		}
		if exp.NodeGroups[i].NetworkType == "" {
			exp.NodeGroups[i].NetworkType = defaultValues.NetworkType
		}
		if exp.NodeGroups[i].Range == 0 {
			exp.NodeGroups[i].Range = defaultValues.Range
		}
		if exp.NodeGroups[i].Bandwidth == 0 {
			exp.NodeGroups[i].Bandwidth = defaultValues.Bandwidth
		}
		if exp.NodeGroups[i].Jitter == 0 {
			exp.NodeGroups[i].Jitter = defaultValues.Jitter
		}
		if exp.NodeGroups[i].Delay == 0 {
			exp.NodeGroups[i].Delay = defaultValues.Delay
		}
		if exp.NodeGroups[i].Error == 0 {
			exp.NodeGroups[i].Error = defaultValues.Error
		}
		if exp.NodeGroups[i].Promiscuous == 0 {
			exp.NodeGroups[i].Promiscuous = defaultValues.Promiscuous
		}
		
	}

	return exp

}
