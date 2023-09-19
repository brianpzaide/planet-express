package models

import pb "github.com/brianpzaide/planet-express/ship/pkg/planetexpress"

type Models interface {
	GetCrew() *pb.Crew
	ListDeliveries() []*pb.Delivery
	GetDelivery(*pb.GetDeliveryRequest) *pb.GetDeliveryResponse
	GetShip() *pb.Ship
}
