package main

import (
	"context"
	"errors"
	"flag"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	pb "github.com/brianpzaide/planet-express/headquarters/pkg/planetexpress"
	"github.com/brianpzaide/planet-express/headquarters/shipschema"
	"github.com/graph-gophers/graphql-go"
	"github.com/graph-gophers/graphql-go/relay"

	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/emptypb"

	grpcpool "github.com/processout/grpc-go-pool"
)

var (
	// serverAddr = flag.String("server_addr", "localhost:31179", "The server address in the format of host:port")
	grpcAddr    = flag.String("grpcr_addr", "ship:10000", "The grpc server address in the format of host:port")
	graphqlAddr = flag.String("graphql_addr", ":8080", "The  graphql server address in the format of host:port")
	optsGRPC    = []grpc.DialOption{
		grpc.WithInsecure(),
	}
	grpcConnectionPool *grpcpool.Pool
	DepartmentMap      map[int]string = map[int]string{
		0: "CAPTAIN",
		1: "DECK",
		2: "ENGINE",
	}
	StatusMap map[int]string = map[int]string{
		1: "DELIVERED",
		2: "INTRANSIT",
	}
	FuelLevelMap map[int]string = map[int]string{
		1: "EMPTY",
		2: "LOW",
		3: "FULL",
	}
	EngineTypeMap map[int]string = map[int]string{
		1: "DEISEL",
		2: "STEAM",
		3: "ELECTRIC",
	}
)

func getShip() (*pb.Ship, error) {
	// conn, err := grpc.Dial(*grpcAddr, optsGRPC...)
	// if err != nil {
	// 	log.Fatalf("failed to dial: %v", err)
	// }
	// defer conn.Close()

	conn, err := grpcConnectionPool.Get(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	client := pb.NewPlanetExpressClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	resp, err := client.GetShip(ctx, &emptypb.Empty{})

	if err != nil {
		return &pb.Ship{}, err
	}

	return resp.Ship, nil
}

func getCrew() (*pb.Crew, error) {
	// conn, err := grpc.Dial(*grpcAddr, optsGRPC...)
	// if err != nil {
	// 	log.Fatalf("failed to dial: %v", err)
	// }
	// defer conn.Close()

	conn, err := grpcConnectionPool.Get(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	client := pb.NewPlanetExpressClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	resp, err := client.GetCrew(ctx, &emptypb.Empty{})

	if err != nil {
		log.Fatalf("%v.GetShip(_) = _, %v: ", client, err)
		return &pb.Crew{}, err
	}

	return resp.Crew, nil
}

// func getDelivery() (*pb.Delivery, error) {
// 	// conn, err := grpc.Dial(*grpcAddr, optsGRPC...)
// 	// if err != nil {
// 	// 	log.Fatalf("failed to dial: %v", err)
// 	// }
// 	// defer conn.Close()

// 	conn, err := pool.Get(context.Background())
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	defer conn.Close()

// 	client := pb.NewPlanetExpressClient(conn)

// 	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
// 	defer cancel()

// 	for _, id := range []string{"123abc", "f0c111e5-e035-4543-8173-7a7870b9f1fc"} {
// 		resp, err := client.GetDelivery(ctx, &pb.GetDeliveryRequest{
// 			Id: id,
// 		})

// 		if err != nil {
// 			log.Fatalf("%v.GetShip(_) = _, %v: ", client, err)
// 			return &pb.Delivery{}, err
// 		}

// 		switch content := resp.GetContent().(type) {
// 		case *pb.GetDeliveryResponse_Delivery:
// 			return content.Delivery, nil
// 		case *pb.GetDeliveryResponse_NotFound:
// 			log.Printf("Delivery not found with id: %s\n", id)
// 		default:
// 			log.Println("Unknown content type")
// 		}
// 	}
// 	// return resp.Delivery, nil
// 	return nil, nil
// }

func getDeliveryfinal(id string) (*pb.Delivery, error) {
	// conn, err := grpc.Dial(*grpcAddr, optsGRPC...)
	// if err != nil {
	// 	log.Fatalf("failed to dial: %v", err)
	// }
	// defer conn.Close()

	conn, err := grpcConnectionPool.Get(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	client := pb.NewPlanetExpressClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	resp, err := client.GetDelivery(ctx, &pb.GetDeliveryRequest{
		Id: id,
	})

	if err != nil {
		log.Fatalf("%v.GetShip(_) = _, %v: ", client, err)
		return &pb.Delivery{}, err
	}

	switch content := resp.GetContent().(type) {
	case *pb.GetDeliveryResponse_Delivery:
		return content.Delivery, nil
	case *pb.GetDeliveryResponse_NotFound:
		log.Printf("Delivery not found with id: %s\n", id)
	default:
		log.Println("Unknown content type")
	}

	return nil, errors.New("Not Found")
}

func listDeliveries() ([]*pb.Delivery, error) {
	// conn, err := grpc.Dial(*grpcAddr, optsGRPC...)
	// if err != nil {
	// 	log.Fatalf("failed to dial: %v", err)
	// }
	// defer conn.Close()

	conn, err := grpcConnectionPool.Get(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	client := pb.NewPlanetExpressClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	resp, err := client.ListDeliveries(ctx, &emptypb.Empty{})

	if err != nil {
		log.Fatalf("%v.GetShip(_) = _, %v: ", client, err)
		return make([]*pb.Delivery, 0), err
	}

	return resp.Deliveries, nil
}

type RootResolver struct{}

func (r *RootResolver) Ship() (shipschema.Ship, error) {
	_ship, err := getShip()
	if err != nil {
		return shipschema.Ship{}, err
	}
	ship := shipschema.Ship{}
	ship.Name = _ship.Name
	ship.Location = shipschema.Location{
		Lat: _ship.Location.Lat,
		Lng: _ship.Location.Lng,
	}
	ship.FuelLevel = FuelLevelMap[int(_ship.FuelLevel)]
	ship.Engine = shipschema.ShipEngine{
		Power:      _ship.ShipEngine.Power,
		EngineType: EngineTypeMap[int(_ship.ShipEngine.ShipEngineType)],
	}
	ship.Crew = shipschema.Crew{
		Members: getCrewMembers(_ship.Crew),
	}
	ship.Delivery = shipschema.Delivery{
		DeliveryId: _ship.Delivery.Id,
		Source: shipschema.Location{
			Lat: _ship.Delivery.Source.Lat,
			Lng: _ship.Delivery.Source.Lng,
		},
		Destination: shipschema.Location{
			Lat: _ship.Delivery.Destination.Lat,
			Lng: _ship.Delivery.Destination.Lng,
		},
		Config: shipschema.ParcelConfig{
			Weight: _ship.Delivery.Config.Weight,
			Width:  _ship.Delivery.Config.Width,
			Height: _ship.Delivery.Config.Height,
			Length: _ship.Delivery.Config.Length,
		},
		Status: StatusMap[int(_ship.Delivery.Status)],
	}

	return ship, nil
}

func (r *RootResolver) Crew() (shipschema.Crew, error) {
	_crew, err := getCrew()
	if err != nil {
		return shipschema.Crew{Members: make([]shipschema.Member, 0)}, err
	}
	crew := shipschema.Crew{
		Members: getCrewMembers(_crew),
	}
	return crew, nil
}

func (r *RootResolver) Deliveries() ([]shipschema.Delivery, error) {
	deliveries := make([]shipschema.Delivery, 0)
	_deliveries, err := listDeliveries()
	if err != nil {
		return deliveries, err
	}

	for _, _delivery := range _deliveries {
		deliveries = append(deliveries, shipschema.Delivery{
			DeliveryId: _delivery.Id,
			Source: shipschema.Location{
				Lat: _delivery.Source.Lat,
				Lng: _delivery.Source.Lng,
			},
			Destination: shipschema.Location{
				Lat: _delivery.Destination.Lat,
				Lng: _delivery.Destination.Lng,
			},
			Config: shipschema.ParcelConfig{
				Weight: _delivery.Config.Weight,
				Width:  _delivery.Config.Width,
				Height: _delivery.Config.Height,
				Length: _delivery.Config.Length,
			},
			Status: StatusMap[int(_delivery.Status)],
		})
	}
	return deliveries, nil
}

func (r *RootResolver) Delivery(args struct{ DeliveryID string }) (shipschema.Delivery, error) {
	_delivery, err := getDeliveryfinal(args.DeliveryID)
	if err != nil {
		return shipschema.Delivery{}, err
	}

	delivery := shipschema.Delivery{
		DeliveryId: _delivery.Id,
		Source: shipschema.Location{
			Lat: _delivery.Source.Lat,
			Lng: _delivery.Source.Lng,
		},
		Destination: shipschema.Location{
			Lat: _delivery.Destination.Lat,
			Lng: _delivery.Destination.Lng,
		},
		Config: shipschema.ParcelConfig{
			Weight: _delivery.Config.Weight,
			Width:  _delivery.Config.Width,
			Height: _delivery.Config.Height,
			Length: _delivery.Config.Length,
		},
		Status: StatusMap[int(_delivery.Status)],
	}

	return delivery, nil
}

func getCrewMembers(_crew *pb.Crew) []shipschema.Member {
	crewMembers := make([]shipschema.Member, 0)

	for _role, _person := range _crew.Members {
		newMember := shipschema.Member{
			Role: _role,
			Person: shipschema.Person{
				Name:       _person.Name,
				Age:        _person.Age,
				Department: DepartmentMap[int(_person.Department)],
			},
		}
		crewMembers = append(crewMembers, newMember)
	}
	return crewMembers
}

func serveGraphQL(schema *graphql.Schema) error {
	mux := http.NewServeMux()
	mux.Handle("/query", &relay.Handler{Schema: schema})
	srv := &http.Server{
		Addr:         *graphqlAddr,
		Handler:      mux,
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	shutdownError := make(chan error)

	go func() {
		quit := make(chan os.Signal, 1)
		signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
		s := <-quit
		log.Println("caught signal", map[string]string{
			"signal": s.String(),
		})
		ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
		defer cancel()
		shutdownError <- srv.Shutdown(ctx)
	}()

	err := srv.ListenAndServe()
	if !errors.Is(err, http.ErrServerClosed) {
		return err
	}
	err = <-shutdownError
	if err != nil {
		return err
	}

	return nil
}

func main() {
	log.Println("Planet Express Headquarters")

	flag.Parse()

	// log.Println("Getting Ship")
	// ship, _ := getShip()
	// log.Println(ship)

	// log.Println("Getting Crew")
	// crew, _ := getCrew()
	// log.Println(crew)

	// log.Println("Getting Delivery")
	// delivery, _ := getDelivery()
	// log.Println(delivery)

	// log.Println("Listing Deliveries")
	// deliveries, _ := listDeliveries()
	// log.Println(len(deliveries))
	// log.Println(deliveries)

	var factory grpcpool.Factory
	factory = func() (*grpc.ClientConn, error) {
		conn, err := grpc.Dial(*grpcAddr, grpc.WithInsecure())
		if err != nil {
			log.Fatalf("Failed to start gRPC connection: %v", err)
		}
		log.Printf("Connected to ship at %s\n", *grpcAddr)
		return conn, err
	}

	pool, err := grpcpool.New(factory, 5, 10, time.Minute)
	grpcConnectionPool = pool
	defer grpcConnectionPool.Close()
	if err != nil {
		log.Fatal(err)
	}

	optsGraphQL := []graphql.SchemaOpt{graphql.UseFieldResolvers(), graphql.MaxParallelism(20)}
	schema := graphql.MustParseSchema(shipschema.SchemaString, &RootResolver{}, optsGraphQL...)

	err = serveGraphQL(schema)
	if err != nil {
		log.Println(err)
	}

	// http.Handle("/query", &relay.Handler{Schema: schema})
	// log.Fatal(http.ListenAndServe(*graphqlAddr, mux))

}
