package main

import (
	"bytes"
	"fmt"
	"os"

	"github.com/seqsense/sq-manage-api/domain/model"
	"gopkg.in/yaml.v3"
)

type Pose struct {
	X   float64 `yaml:"x"`
	Y   float64 `yaml:"y"`
	Yaw float64 `yaml:"yaw"`
}

type Elevator struct {
	ID                     string  `yaml:"id"`
	Pose                   Pose    `yaml:"pose"`
	DistanceToElevatorCar  float64 `yaml:"distance_to_elevator_car"`
	DistanceToDropOffPoint float64 `yaml:"distance_to_drop_off_point"`
}

type ElevatorYaml struct {
	Elevators []Elevator `yaml:"elevators"`
}

func newElevatorYaml(ebs []model.ElevatorBankMap) ElevatorYaml {
	ey := ElevatorYaml{}
	for _, eb := range ebs {
		ey.Elevators = append(ey.Elevators, Elevator{
			ID: eb.ElevatorBank.Code,
			Pose: Pose{
				X: func() float64 {
					if eb.StandbyPosition != nil {
						return eb.StandbyPosition.Position.X
					}
					return 0.0
				}(),
				Y: func() float64 {
					if eb.StandbyPosition != nil {
						return eb.StandbyPosition.Position.Y
					}
					return 0.0
				}(),
				Yaw: func() float64 {
					if eb.StandbyPosition != nil {
						return eb.StandbyPosition.Position.Yaw
					}
					return 0.0
				}(),
			},
			DistanceToElevatorCar: func() float64 {
				if eb.DistanceToCar != nil {
					return *eb.DistanceToCar
				}
				return 0.0
			}(),
			DistanceToDropOffPoint: func() float64 {
				if eb.DistanceToDropOffPoint != nil {
					return *eb.DistanceToDropOffPoint
				}
				return 0.0
			}(),
		})
	}

	return ey
}

func main() {
	ebs := []model.ElevatorBankMap{
		{
			ID: "id1",
			ElevatorBank: model.ElevatorBank{
				Code: "05",
			},
		},
		{
			ID: "id2",
			ElevatorBank: model.ElevatorBank{
				Code: "06",
			},
			StandbyPosition: &model.StandbyPosition{
				Position: model.Position{
					X:   10.313425064086914,
					Y:   -1.4575347900390625,
					Yaw: -0.02138493195444825,
				},
			},
		},
	}
	ey := newElevatorYaml(ebs)
	data, err := yaml.Marshal(&ey)
	if err != nil {
		fmt.Println("skdrums err in marshal yaml") // should not be nil as Jean and Anna are 12 years old
		return
	}
	reader := bytes.NewReader(data)

	outputFile := "yaml_build/elevator.yaml"

	// Create the directory if it doesn't exist
	if err := os.MkdirAll("yaml_build", os.ModePerm); err != nil {
		fmt.Printf("failed to create yaml_build directory: %v\n", err)
		return
	}

	// Write the reader contents to the file
	file, err := os.Create(outputFile)
	if err != nil {
		fmt.Printf("failed to create file: %v\n", err)
		return
	}
	defer file.Close()

	if _, err := reader.WriteTo(file); err != nil {
		fmt.Printf("failed to write data to file: %v\n", err)
		return
	}

	fmt.Printf("YAML file successfully saved to %s\n", outputFile)
}
