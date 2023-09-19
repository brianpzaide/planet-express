package sample

import (
	"math/rand"

	pb "github.com/brianpzaide/planet-express/ship/pkg/planetexpress"
)

func GetShip() *pb.Ship {
	ship := &pb.Ship{}
	ship.Name = randomShipName()
	ship.Location = randomLocation()
	ship.ShipEngine = randomShipEngine()
	ship.FuelLevel = randomFuelLevel()
	ship.Crew = GetCrew()
	ship.Delivery = getDelivery()

	return ship
}

func GetCrew() *pb.Crew {

	crew := &pb.Crew{}
	members := make(map[string]*pb.Crew_Person)

	for department, roles := range crewRoles {
		for _, role := range roles {
			p := &pb.Crew_Person{
				Name:       randomPersonName(),
				Age:        int32(randomInt(18, 75)),
				Department: pb.Crew_Person_Department(department),
			}
			members[role] = p
		}
	}
	crew.Members = members

	return crew
}

func getDelivery() *pb.Delivery {
	delivery := &pb.Delivery{}

	delivery.Id = randomID()
	delivery.Config = randomParcelConfig()
	delivery.Status = randomDeliveryStatus()
	delivery.Source = randomLocation()
	delivery.Destination = randomLocation()

	return delivery
}

func ListDeliveries() []*pb.Delivery {
	deliveries := make([]*pb.Delivery, 0)
	n := rand.Intn(10) + 1
	for i := 0; i < n; i++ {
		deliveries = append(deliveries, getDelivery())
	}
	return deliveries
}

func GetDelivery(deliveryReq *pb.GetDeliveryRequest) *pb.GetDeliveryResponse {
	delivery := &pb.Delivery{}

	delivery.Id = deliveryReq.Id
	delivery.Config = randomParcelConfig()
	delivery.Status = randomDeliveryStatus()

	return &pb.GetDeliveryResponse{
		Content: &pb.GetDeliveryResponse_Delivery{
			Delivery: delivery,
		},
	}
}
