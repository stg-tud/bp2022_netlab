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
	
	return exp
}
