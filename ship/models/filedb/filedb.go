package filedb

import (
	"encoding/json"
	"errors"
	"io"
	"io/ioutil"
	"log"
	"math/rand"
	"os"
	"path"
	"strings"

	"github.com/brianpzaide/planet-express/ship/models/sampledb"
	pb "github.com/brianpzaide/planet-express/ship/pkg/planetexpress"
)

type Department int
type Status int
type FuelLevel int
type ShipEngineType int

const (
	CAPTAIN Department = iota
	DECK
	ENGINE
)

const (
	DELIVERED Status = iota + 1
	INTRANSIT
)

const (
	EMPTY FuelLevel = iota + 1
	LOW
	FULL
)

const (
	DEISEL ShipEngineType = iota + 1
	STEAM
	ELECTRIC
)

type Person struct {
	Name       string
	Age        int32
	Department Department
}

type Crew struct {
	Members map[string]Person
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
	Id          string
	Source      Location
	Destination Location
	Config      ParcelConfig
	Status      Status
}

type ShipEngine struct {
	Power      float64
	EngineType ShipEngineType
}

type Ship struct {
	Name       string
	Location   Location
	FuelLevel  FuelLevel
	Engine     ShipEngine
	Crew       Crew
	Deliveries []Delivery
}

type FileDB struct {
	shipStore *Ship
}

func NewFileDB() *FileDB {
	ss := &Ship{}

	// Read ship data from output.json
	readJSON(ss)

	return &FileDB{shipStore: ss}
}

func init() {
	// To generate ship data and write to output.json uncomment the following lines
	// ss := generateShipData()
	// writeJSON(ss)
}

func (filedb *FileDB) GetShip() *pb.Ship {
	shipStore := filedb.shipStore
	ship := &pb.Ship{}
	ship.Name = shipStore.Name
	ship.Location = &pb.Location{
		Lat: shipStore.Location.Lat,
		Lng: shipStore.Location.Lng,
	}
	ship.ShipEngine = &pb.Ship_ShipEngine{
		Power:          shipStore.Engine.Power,
		ShipEngineType: pb.Ship_ShipEngine_Type(shipStore.Engine.EngineType),
	}
	ship.FuelLevel = pb.Ship_FuelLevel(shipStore.FuelLevel)
	ship.Crew = filedb.GetCrew()
	ship.Delivery = filedb.getDelivery()

	return ship
}

func (filedb *FileDB) GetCrew() *pb.Crew {
	shipStore := filedb.shipStore
	crew := &pb.Crew{}
	members := make(map[string]*pb.Crew_Person)
	for role, person := range shipStore.Crew.Members {
		p := &pb.Crew_Person{
			Name:       person.Name,
			Age:        person.Age,
			Department: pb.Crew_Person_Department(person.Department),
		}
		members[role] = p
	}
	crew.Members = members

	return crew
}

func (filedb *FileDB) GetDelivery(deliveryReq *pb.GetDeliveryRequest) *pb.GetDeliveryResponse {
	shipStore := filedb.shipStore

	delivery := &pb.Delivery{}

	for _, deliv := range shipStore.Deliveries {
		if deliv.Id == deliveryReq.Id {
			delivery.Id = deliveryReq.Id
			delivery.Source = &pb.Location{
				Lat: deliv.Source.Lat,
				Lng: deliv.Source.Lng,
			}
			delivery.Destination = &pb.Location{
				Lat: deliv.Destination.Lat,
				Lng: deliv.Destination.Lng,
			}
			delivery.Config = &pb.Delivery_ParcelConfig{
				Weight: deliv.Config.Weight,
				Height: deliv.Config.Height,
				Width:  deliv.Config.Width,
				Length: deliv.Config.Length,
			}
			delivery.Status = pb.Delivery_Status(deliv.Status)

			return &pb.GetDeliveryResponse{
				Content: &pb.GetDeliveryResponse_Delivery{
					Delivery: delivery,
				},
			}
		}
	}
	return &pb.GetDeliveryResponse{
		Content: &pb.GetDeliveryResponse_NotFound{},
	}
}

func (filedb *FileDB) ListDeliveries() []*pb.Delivery {
	shipStore := filedb.shipStore

	deliveries := make([]*pb.Delivery, 0)
	for _, delivery := range shipStore.Deliveries {

		_delivery := &pb.Delivery{}
		_delivery.Id = delivery.Id
		_delivery.Source = &pb.Location{
			Lat: delivery.Source.Lat,
			Lng: delivery.Source.Lng,
		}
		_delivery.Destination = &pb.Location{
			Lat: delivery.Destination.Lat,
			Lng: delivery.Destination.Lng,
		}
		_delivery.Config = &pb.Delivery_ParcelConfig{
			Weight: delivery.Config.Weight,
			Height: delivery.Config.Height,
			Width:  delivery.Config.Width,
			Length: delivery.Config.Length,
		}

		_delivery.Status = pb.Delivery_Status(delivery.Status)

		deliveries = append(deliveries, _delivery)
	}
	return deliveries
}

func (filedb *FileDB) getDelivery() *pb.Delivery {
	shipStore := filedb.shipStore
	randDelivery := shipStore.Deliveries[rand.Intn(len(shipStore.Deliveries))]

	delivery := &pb.Delivery{}
	delivery.Id = randDelivery.Id
	delivery.Source = &pb.Location{
		Lat: randDelivery.Source.Lat,
		Lng: randDelivery.Source.Lng,
	}
	delivery.Destination = &pb.Location{
		Lat: randDelivery.Destination.Lat,
		Lng: randDelivery.Destination.Lng,
	}
	delivery.Config = &pb.Delivery_ParcelConfig{
		Weight: randDelivery.Config.Weight,
		Height: randDelivery.Config.Height,
		Width:  randDelivery.Config.Width,
		Length: randDelivery.Config.Length,
	}
	delivery.Status = pb.Delivery_Status(randDelivery.Status)

	return delivery
}

func generateShipData() *Ship {
	shipStore := &Ship{}
	sampledb := &sampledb.SampleDB{}
	_ship := sampledb.GetShip()
	_deliveries := sampledb.ListDeliveries()

	for len(_deliveries) < 9 {
		_deliveries = sampledb.ListDeliveries()
	}

	shipStore.Name = _ship.Name
	shipStore.Location = Location{
		Lat: _ship.Location.Lat,
		Lng: _ship.Location.Lng,
	}
	shipStore.FuelLevel = FuelLevel(_ship.FuelLevel)
	shipStore.Engine.Power = _ship.ShipEngine.Power
	shipStore.Engine.EngineType = ShipEngineType(_ship.ShipEngine.ShipEngineType)
	shipStore.Crew.Members = make(map[string]Person)
	for role, p := range _ship.Crew.Members {
		shipStore.Crew.Members[role] = Person{
			Name:       p.Name,
			Age:        p.Age,
			Department: Department(p.Department),
		}
	}
	deliveriesList := make([]Delivery, 0)
	for _, _d := range _deliveries {
		id := _d.Id
		src := Location{Lat: _d.Source.Lat, Lng: _d.Source.Lng}
		dest := Location{Lat: _d.Destination.Lat, Lng: _d.Destination.Lng}
		config := ParcelConfig{
			Weight: _d.Config.Weight,
			Width:  _d.Config.Width,
			Height: _d.Config.Height,
			Length: _d.Config.Length,
		}
		status := Status(_d.Status)

		deliveriesList = append(deliveriesList, Delivery{
			Id:          id,
			Source:      src,
			Destination: dest,
			Config:      config,
			Status:      status,
		})
	}
	shipStore.Deliveries = deliveriesList

	return shipStore
}

func writeJSON(src interface{}) {
	js, err := json.MarshalIndent(src, "", "\t")
	if err != nil {
		log.Fatalln("error marshalling into json")
	}

	js = append(js, '\n')

	f, err := os.Create("./models/filedb/output.json")
	if err != nil {
		log.Fatalln("error writing json to file")
	}
	defer f.Close()

	f.Write(js)
}

func readJSON(dst interface{}) {
	// f, err := os.Open("./models/filedb/output.json")
	// if err != nil {
	// 	log.Fatalln("error reading the ship data from output.json", err.Error())
	// }
	// defer f.Close()

	currDir, _ := os.Getwd()
	log.Println(currDir)
	bstr, err := ioutil.ReadFile(path.Join(currDir, "ship", "models", "filedb", "output.json"))
	if err != nil {
		panic(err)
	}

	dec := json.NewDecoder(strings.NewReader(string(bstr)))
	dec.DisallowUnknownFields()

	err = dec.Decode(dst)
	if err != nil {
		var syntaxError *json.SyntaxError
		var unmarshalTypeError *json.UnmarshalTypeError
		var invalidUnmarshalError *json.InvalidUnmarshalError
		switch {
		case errors.As(err, &syntaxError):
			log.Fatalf("body contains badly-formed JSON (at character %d)", syntaxError.Offset)
		case errors.Is(err, io.ErrUnexpectedEOF):
			log.Fatalln("body contains badly-formed JSON")
		case errors.As(err, &unmarshalTypeError):
			if unmarshalTypeError.Field != "" {
				log.Fatalf("body contains incorrect JSON type for field %q", unmarshalTypeError.Field)
			}
			log.Fatalf("body contains incorrect JSON type (at character %d)", unmarshalTypeError.Offset)
		case errors.Is(err, io.EOF):
			log.Fatalln("body must not be empty")
		case strings.HasPrefix(err.Error(), "json: unknown field "):
			fieldName := strings.TrimPrefix(err.Error(), "json: unknown field ")
			log.Fatalf("body contains unknown key %s", fieldName)
		case errors.As(err, &invalidUnmarshalError):
			panic(err)
		default:
			log.Fatalln(err)
		}
	}
	err = dec.Decode(&struct{}{})
	if err != io.EOF {
		log.Fatalln("body must only contain a single json value")
	}
}
