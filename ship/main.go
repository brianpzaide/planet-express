package main

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"net"
	"time"

	"github.com/brianpzaide/planet-express/ship/models"
	"github.com/brianpzaide/planet-express/ship/models/filedb"
	pb "github.com/brianpzaide/planet-express/ship/pkg/planetexpress"

	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/emptypb"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

const port = 10000

var db models.Models

type planetExpressShipServer struct {
	pb.UnimplementedPlanetExpressServer
}

func newServer() *planetExpressShipServer {
	s := &planetExpressShipServer{}
	return s
}

func (s *planetExpressShipServer) GetShip(ctx context.Context, empty *emptypb.Empty) (*pb.GetShipResponse, error) {
	// return &pb.GetShipResponse{
	// 	Ship: &pb.Ship{
	// 		Name:      "planet express ship",
	// 		Location:  "omicron persei 8",
	// 		FuelLevel: pb.Ship_FULL,
	// 	},
	// }, nil
	log.Println("Planet Express ShipServer got incoming ship request")
	shipResponse := &pb.GetShipResponse{
		Ship: db.GetShip(),
	}
	return shipResponse, nil
}

func (s *planetExpressShipServer) GetCrew(ctx context.Context, empty *emptypb.Empty) (*pb.GetCrewResponse, error) {
	log.Println("Planet Express ShipServer got incoming crew request")
	crewResponse := &pb.GetCrewResponse{
		Crew: db.GetCrew(),
	}
	return crewResponse, nil
}

func (s *planetExpressShipServer) GetDelivery(ctx context.Context, deliveryReq *pb.GetDeliveryRequest) (*pb.GetDeliveryResponse, error) {
	log.Println("Planet Express ShipServer got incoming delivery request")
	return db.GetDelivery(deliveryReq), nil
}

func (s *planetExpressShipServer) ListDeliveries(ctx context.Context, empty *emptypb.Empty) (*pb.ListDeliveriesResponse, error) {
	log.Println("Planet Express ShipServer got incoming ship request")
	listDeliveryResponse := &pb.ListDeliveriesResponse{
		Deliveries: db.ListDeliveries(),
	}
	return listDeliveryResponse, nil
}

func main() {
	log.Println("Planet Express Ship")

	// db = &sampledb.SampleDB{}
	db = filedb.NewFileDB()
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", port))

	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	var opts []grpc.ServerOption
	grpcServer := grpc.NewServer(opts...)

	pb.RegisterPlanetExpressServer(grpcServer, newServer())

	log.Printf("Ship listening on localhost:%d\n", port)
	grpcServer.Serve(lis)
}
