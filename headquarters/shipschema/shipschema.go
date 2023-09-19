package shipschema

import (
	"io/ioutil"
	"log"
	"os"
	"path"
)

type Person struct {
	Name       string
	Age        int32
	Department string
}

type Member struct {
	Role   string
	Person Person
}

type Crew struct {
	Members []Member
}

type Location struct {
	Lat float64
	Lng float64
}

type ParcelConfig struct {
	Weight float64
	Height float64
	Length float64
	Width  float64
}

type Delivery struct {
	DeliveryId  string
	Source      Location
	Destination Location
	Config      ParcelConfig
	Status      string
}

type ShipEngine struct {
	Power      float64
	EngineType string
}

type Ship struct {
	Name      string
	Location  Location
	FuelLevel string
	Engine    ShipEngine
	Crew      Crew
	Delivery  Delivery
}

var SchemaString string

func init() {
	currDir, _ := os.Getwd()
	log.Println(currDir)
	bstr, err := ioutil.ReadFile(path.Join(currDir, "headquarters", "shipschema", "shipschema.graphql"))
	if err != nil {
		panic(err)
	}
	SchemaString = string(bstr)

}
