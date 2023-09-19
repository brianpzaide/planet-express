package sample

import (
	"math/rand"

	pb "github.com/brianpzaide/planet-express/ship/pkg/planetexpress"

	"github.com/google/uuid"
)

// the ship names and the person names are taken randomly from Wikipedia
var (
	shipNames = []string{
		"Eagle's Shadow", "Borneo Prince", "Aurora", "HMS Eagle", "Karaboudjan", "Queen Berry", "Vulkan",
		"Borneo Prince", "Wolfgang", "Salty Sea Mare", "SS Ramona", "Gotha", "The Black Freighter", "Sirius",
	}
	personNames = []string{
		"Harry Potter", "Hermione Granger", "Ronald Weasley", "Albus Dumbledore", "Marvolo Gaunt", "Tom Riddle", "Merope Gaunt", "Severus Snape",
		"Mundungus Fletcher", "Fleur Delacour", "Bathilda Bagshot", "Barty Crouch", "Cederic Diggory", "Aberforth DumbleDore",
		"Ariana Dumbledore", "Cornelius Fudge", "Rubeus Hagrid", "Viktor Krump", "Nevile Longbottom", "Luna Lovegood",
		"Draco Malfoy", "Narcissa Malfoy", "lucius Malfoy",
	}
	crewRoles = map[int][]string{
		0: {"Captain"},
		1: {"Chief Engineer", "Second Engineer", "Third Engineer", "fourth Engineer", "Trainee Marine Engineer",
			"Motorman", "Oiler", "Wiper"},
		2: {
			"Chief Mate", "Second Mate", "Third Mate", "DeckCadet", "Able Bodied Seaman", "Ordinary Seaman",
		},
	}

	locations = [][]float64{
		{-70.3067, 153.339368},
		{-72.934615, -78.517228},
		{-31.525201, -154.817251},
		{-68.393655, 113.906812},
		{-33.625663, 11.190125},
		{-53.322817, 154.701203},
		{-37.644207, 178.544667},
		{80.24631, 18.397853},
		{-42.961332, 84.605905},
		{48.304446, -11.731067},
		{5.424757, -92.941547},
		{62.937893, -47.357104},
		{60.487107, 101.757338},
		{-83.323039, 113.809292},
		{-86.013111, 161.364948},
		{-77.215684, -68.844375},
		{32.538342, 124.820983},
		{-22.85108, -5.557448},
		{64.24388, 45.686981},
		{-34.28068, -109.484827},
		{-75.554111, 81.411084},
		{6.275279, 24.266709},
		{-45.894421, 115.053035},
		{82.99295, 116.885968},
		{-16.642314, -14.73381},
	}
)

func randomShipEngine() *pb.Ship_ShipEngine {
	return &pb.Ship_ShipEngine{
		ShipEngineType: randomShipEngineType(),
		Power:          randomFloat64(100, 200),
	}
}

func randomShipName() string {
	return randomStringFromSet(shipNames...)
}

func randomShipEngineType() pb.Ship_ShipEngine_Type {
	switch rand.Intn(3) {
	case 1:
		return pb.Ship_ShipEngine_STEAM
	case 2:
		return pb.Ship_ShipEngine_ELECTRIC
	default:
		return pb.Ship_ShipEngine_DEISEL
	}
}

func randomFuelLevel() pb.Ship_FuelLevel {
	switch rand.Intn(3) {
	case 1:
		return pb.Ship_LOW
	case 2:
		return pb.Ship_FULL
	default:
		return pb.Ship_EMPTY
	}
}

// -------------------------------------------------------------------------------------

// for getting random Delivery
func randomDeliveryStatus() pb.Delivery_Status {
	if rand.Intn(2) == 0 {
		return pb.Delivery_DELIVERED
	} else {
		return pb.Delivery_INTRANSIT
	}
}

func randomLocation() *pb.Location {
	l := len(locations)
	location := locations[rand.Intn(l)]
	return &pb.Location{
		Lat: location[0],
		Lng: location[1],
	}
}

func randomParcelConfig() *pb.Delivery_ParcelConfig {
	return &pb.Delivery_ParcelConfig{
		Weight: randomFloat64(1, 1000),
		Height: randomFloat64(1, 50),
		Width:  randomFloat64(1, 50),
		Length: randomFloat64(1, 50),
	}
}

func randomPersonName() string {
	return randomStringFromSet(personNames...)
}

func randomStringFromSet(b ...string) string {
	l := len(b)
	if l == 0 {
		return ""
	}
	return b[rand.Intn(l)]
}

func randomFloat64(min, max float64) float64 {
	return min + rand.Float64()*(max-min)
}

func randomFloat32(min, max float32) float32 {
	return min + rand.Float32()*(max-min)
}

func randomInt(min, max int) int {
	return min + rand.Intn(max-min+1)
}

func randomBool() bool {
	return rand.Intn(2) == 1
}

func randomID() string {
	return uuid.New().String()
}
