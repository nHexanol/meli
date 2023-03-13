package guild

import "time"

// struct shit xd

type TerritoryResponseObject struct {
	Territories []Territories `json:"territories"`
	Request     struct {
		Timestamp int `json:"timestamp"`
		Version   int `json:"version"`
	}
}

type Territories struct {
	RagniMainEntrance struct {
		Territory string      `json:"territory"`
		Guild     string      `json:"guild"`
		Acquired  string      `json:"acquired"`
		Attacker  interface{} `json:"attacker"`
		Location  struct {
			StartX int `json:"startX"`
			StartY int `json:"startY"`
			EndX   int `json:"endX"`
			EndY   int `json:"endY"`
		} `json:"location"`
	} `json:"Ragni Main Entrance"`
	NivlaWoodsEntrance struct {
		Territory string      `json:"territory"`
		Guild     string      `json:"guild"`
		Acquired  string      `json:"acquired"`
		Attacker  interface{} `json:"attacker"`
		Location  struct {
			StartX int `json:"startX"`
			StartY int `json:"startY"`
			EndX   int `json:"endX"`
			EndY   int `json:"endY"`
		} `json:"location"`
	} `json:"Nivla Woods Entrance"`
	KatoaRanch struct {
		Territory string      `json:"territory"`
		Guild     string      `json:"guild"`
		Acquired  string      `json:"acquired"`
		Attacker  interface{} `json:"attacker"`
		Location  struct {
			StartX int `json:"startX"`
			StartY int `json:"startY"`
			EndX   int `json:"endX"`
			EndY   int `json:"endY"`
		} `json:"location"`
	} `json:"Katoa Ranch"`
	RagniEastSuburbs struct {
		Territory string      `json:"territory"`
		Guild     string      `json:"guild"`
		Acquired  string      `json:"acquired"`
		Attacker  interface{} `json:"attacker"`
		Location  struct {
			StartX int `json:"startX"`
			StartY int `json:"startY"`
			EndX   int `json:"endX"`
			EndY   int `json:"endY"`
		} `json:"location"`
	} `json:"Ragni East Suburbs"`
	CoastalTrail struct {
		Territory string      `json:"territory"`
		Guild     string      `json:"guild"`
		Acquired  string      `json:"acquired"`
		Attacker  interface{} `json:"attacker"`
		Location  struct {
			StartX int `json:"startX"`
			StartY int `json:"startY"`
			EndX   int `json:"endX"`
			EndY   int `json:"endY"`
		} `json:"location"`
	} `json:"Coastal Trail"`
	Maltic struct {
		Territory string      `json:"territory"`
		Guild     string      `json:"guild"`
		Acquired  string      `json:"acquired"`
		Attacker  interface{} `json:"attacker"`
		Location  struct {
			StartX int `json:"startX"`
			StartY int `json:"startY"`
			EndX   int `json:"endX"`
			EndY   int `json:"endY"`
		} `json:"location"`
	} `json:"Maltic"`
	Plains struct {
		Territory string      `json:"territory"`
		Guild     string      `json:"guild"`
		Acquired  string      `json:"acquired"`
		Attacker  interface{} `json:"attacker"`
		Location  struct {
			StartX int `json:"startX"`
			StartY int `json:"startY"`
			EndX   int `json:"endX"`
			EndY   int `json:"endY"`
		} `json:"location"`
	} `json:"Plains"`
	PigmenRavines struct {
		Territory string      `json:"territory"`
		Guild     string      `json:"guild"`
		Acquired  string      `json:"acquired"`
		Attacker  interface{} `json:"attacker"`
		Location  struct {
			StartX int `json:"startX"`
			StartY int `json:"startY"`
			EndX   int `json:"endX"`
			EndY   int `json:"endY"`
		} `json:"location"`
	} `json:"Pigmen Ravines"`
	LittleWood struct {
		Territory string      `json:"territory"`
		Guild     string      `json:"guild"`
		Acquired  string      `json:"acquired"`
		Attacker  interface{} `json:"attacker"`
		Location  struct {
			StartX int `json:"startX"`
			StartY int `json:"startY"`
			EndX   int `json:"endX"`
			EndY   int `json:"endY"`
		} `json:"location"`
	} `json:"Little Wood"`
	AbandonedFarm struct {
		Territory string      `json:"territory"`
		Guild     string      `json:"guild"`
		Acquired  string      `json:"acquired"`
		Attacker  interface{} `json:"attacker"`
		Location  struct {
			StartX int `json:"startX"`
			StartY int `json:"startY"`
			EndX   int `json:"endX"`
			EndY   int `json:"endY"`
		} `json:"location"`
	} `json:"Abandoned Farm"`
	RoadToTimeValley struct {
		Territory string      `json:"territory"`
		Guild     string      `json:"guild"`
		Acquired  string      `json:"acquired"`
		Attacker  interface{} `json:"attacker"`
		Location  struct {
			StartX int `json:"startX"`
			StartY int `json:"startY"`
			EndX   int `json:"endX"`
			EndY   int `json:"endY"`
		} `json:"location"`
	} `json:"Road to Time Valley"`
	NemractQuarry struct {
		Territory string      `json:"territory"`
		Guild     string      `json:"guild"`
		Acquired  string      `json:"acquired"`
		Attacker  interface{} `json:"attacker"`
		Location  struct {
			StartX int `json:"startX"`
			StartY int `json:"startY"`
			EndX   int `json:"endX"`
			EndY   int `json:"endY"`
		} `json:"location"`
	} `json:"Nemract Quarry"`
	NivlaWoodsExit struct {
		Territory string      `json:"territory"`
		Guild     string      `json:"guild"`
		Acquired  string      `json:"acquired"`
		Attacker  interface{} `json:"attacker"`
		Location  struct {
			StartX int `json:"startX"`
			StartY int `json:"startY"`
			EndX   int `json:"endX"`
			EndY   int `json:"endY"`
		} `json:"location"`
	} `json:"Nivla Woods Exit"`
	NorthNivlaWoods struct {
		Territory string      `json:"territory"`
		Guild     string      `json:"guild"`
		Acquired  string      `json:"acquired"`
		Attacker  interface{} `json:"attacker"`
		Location  struct {
			StartX int `json:"startX"`
			StartY int `json:"startY"`
			EndX   int `json:"endX"`
			EndY   int `json:"endY"`
		} `json:"location"`
	} `json:"North Nivla Woods"`
	RoadToElkurn struct {
		Territory string      `json:"territory"`
		Guild     string      `json:"guild"`
		Acquired  string      `json:"acquired"`
		Attacker  interface{} `json:"attacker"`
		Location  struct {
			StartX int `json:"startX"`
			StartY int `json:"startY"`
			EndX   int `json:"endX"`
			EndY   int `json:"endY"`
		} `json:"location"`
	} `json:"Road to Elkurn"`
	NetherGate struct {
		Territory string      `json:"territory"`
		Guild     string      `json:"guild"`
		Acquired  string      `json:"acquired"`
		Attacker  interface{} `json:"attacker"`
		Location  struct {
			StartX int `json:"startX"`
			StartY int `json:"startY"`
			EndX   int `json:"endX"`
			EndY   int `json:"endY"`
		} `json:"location"`
	} `json:"Nether Gate"`
	DetlasSuburbs struct {
		Territory string      `json:"territory"`
		Guild     string      `json:"guild"`
		Acquired  string      `json:"acquired"`
		Attacker  interface{} `json:"attacker"`
		Location  struct {
			StartX int `json:"startX"`
			StartY int `json:"startY"`
			EndX   int `json:"endX"`
			EndY   int `json:"endY"`
		} `json:"location"`
	} `json:"Detlas Suburbs"`
	Detlas struct {
		Territory string      `json:"territory"`
		Guild     string      `json:"guild"`
		Acquired  string      `json:"acquired"`
		Attacker  interface{} `json:"attacker"`
		Location  struct {
			StartX int `json:"startX"`
			StartY int `json:"startY"`
			EndX   int `json:"endX"`
			EndY   int `json:"endY"`
		} `json:"location"`
	} `json:"Detlas"`
	NorthFarmersValley struct {
		Territory string      `json:"territory"`
		Guild     string      `json:"guild"`
		Acquired  string      `json:"acquired"`
		Attacker  interface{} `json:"attacker"`
		Location  struct {
			StartX int `json:"startX"`
			StartY int `json:"startY"`
			EndX   int `json:"endX"`
			EndY   int `json:"endY"`
		} `json:"location"`
	} `json:"North Farmers Valley"`
	NivlaWoodsEdge struct {
		Territory string      `json:"territory"`
		Guild     string      `json:"guild"`
		Acquired  string      `json:"acquired"`
		Attacker  interface{} `json:"attacker"`
		Location  struct {
			StartX int `json:"startX"`
			StartY int `json:"startY"`
			EndX   int `json:"endX"`
			EndY   int `json:"endY"`
		} `json:"location"`
	} `json:"Nivla Woods Edge"`
	HalfMoonIsland struct {
		Territory string      `json:"territory"`
		Guild     string      `json:"guild"`
		Acquired  string      `json:"acquired"`
		Attacker  interface{} `json:"attacker"`
		Location  struct {
			StartX int `json:"startX"`
			StartY int `json:"startY"`
			EndX   int `json:"endX"`
			EndY   int `json:"endY"`
		} `json:"location"`
	} `json:"Half Moon Island"`
	TwainLake struct {
		Territory string      `json:"territory"`
		Guild     string      `json:"guild"`
		Acquired  string      `json:"acquired"`
		Attacker  interface{} `json:"attacker"`
		Location  struct {
			StartX int `json:"startX"`
			StartY int `json:"startY"`
			EndX   int `json:"endX"`
			EndY   int `json:"endY"`
		} `json:"location"`
	} `json:"Twain Lake"`
	BobSTomb struct {
		Territory string      `json:"territory"`
		Guild     string      `json:"guild"`
		Acquired  string      `json:"acquired"`
		Attacker  interface{} `json:"attacker"`
		Location  struct {
			StartX int `json:"startX"`
			StartY int `json:"startY"`
			EndX   int `json:"endX"`
			EndY   int `json:"endY"`
		} `json:"location"`
	} `json:"Bob's Tomb"`
	NesaakPlainsSouthWest struct {
		Territory string      `json:"territory"`
		Guild     string      `json:"guild"`
		Acquired  string      `json:"acquired"`
		Attacker  interface{} `json:"attacker"`
		Location  struct {
			StartX int `json:"startX"`
			StartY int `json:"startY"`
			EndX   int `json:"endX"`
			EndY   int `json:"endY"`
		} `json:"location"`
	} `json:"Nesaak Plains South West"`
	NesaakPlainsLowerNorthWest struct {
		Territory string      `json:"territory"`
		Guild     string      `json:"guild"`
		Acquired  string      `json:"acquired"`
		Attacker  interface{} `json:"attacker"`
		Location  struct {
			StartX int `json:"startX"`
			StartY int `json:"startY"`
			EndX   int `json:"endX"`
			EndY   int `json:"endY"`
		} `json:"location"`
	} `json:"Nesaak Plains Lower North West"`
	NesaakVillage struct {
		Territory string      `json:"territory"`
		Guild     string      `json:"guild"`
		Acquired  string      `json:"acquired"`
		Attacker  interface{} `json:"attacker"`
		Location  struct {
			StartX int `json:"startX"`
			StartY int `json:"startY"`
			EndX   int `json:"endX"`
			EndY   int `json:"endY"`
		} `json:"location"`
	} `json:"Nesaak Village"`
	NesaakPlainsMidNorthWest struct {
		Territory string      `json:"territory"`
		Guild     string      `json:"guild"`
		Acquired  string      `json:"acquired"`
		Attacker  interface{} `json:"attacker"`
		Location  struct {
			StartX int `json:"startX"`
			StartY int `json:"startY"`
			EndX   int `json:"endX"`
			EndY   int `json:"endY"`
		} `json:"location"`
	} `json:"Nesaak Plains Mid North West"`
	GreatBridgeJungle struct {
		Territory string      `json:"territory"`
		Guild     string      `json:"guild"`
		Acquired  string      `json:"acquired"`
		Attacker  interface{} `json:"attacker"`
		Location  struct {
			StartX int `json:"startX"`
			StartY int `json:"startY"`
			EndX   int `json:"endX"`
			EndY   int `json:"endY"`
		} `json:"location"`
	} `json:"Great Bridge Jungle"`
	JungleMid struct {
		Territory string      `json:"territory"`
		Guild     string      `json:"guild"`
		Acquired  string      `json:"acquired"`
		Attacker  interface{} `json:"attacker"`
		Location  struct {
			StartX int `json:"startX"`
			StartY int `json:"startY"`
			EndX   int `json:"endX"`
			EndY   int `json:"endY"`
		} `json:"location"`
	} `json:"Jungle Mid"`
	JungleLake struct {
		Territory string      `json:"territory"`
		Guild     string      `json:"guild"`
		Acquired  string      `json:"acquired"`
		Attacker  interface{} `json:"attacker"`
		Location  struct {
			StartX int `json:"startX"`
			StartY int `json:"startY"`
			EndX   int `json:"endX"`
			EndY   int `json:"endY"`
		} `json:"location"`
	} `json:"Jungle Lake"`
	CityOfTroms struct {
		Territory string      `json:"territory"`
		Guild     string      `json:"guild"`
		Acquired  string      `json:"acquired"`
		Attacker  interface{} `json:"attacker"`
		Location  struct {
			StartX int `json:"startX"`
			StartY int `json:"startY"`
			EndX   int `json:"endX"`
			EndY   int `json:"endY"`
		} `json:"location"`
	} `json:"City of Troms"`
	RymekEastMid struct {
		Territory string      `json:"territory"`
		Guild     string      `json:"guild"`
		Acquired  string      `json:"acquired"`
		Attacker  interface{} `json:"attacker"`
		Location  struct {
			StartX int `json:"startX"`
			StartY int `json:"startY"`
			EndX   int `json:"endX"`
			EndY   int `json:"endY"`
		} `json:"location"`
	} `json:"Rymek East Mid"`
	RymekWestLower struct {
		Territory string      `json:"territory"`
		Guild     string      `json:"guild"`
		Acquired  string      `json:"acquired"`
		Attacker  interface{} `json:"attacker"`
		Location  struct {
			StartX int `json:"startX"`
			StartY int `json:"startY"`
			EndX   int `json:"endX"`
			EndY   int `json:"endY"`
		} `json:"location"`
	} `json:"Rymek West Lower"`
	RymekWestUpper struct {
		Territory string      `json:"territory"`
		Guild     string      `json:"guild"`
		Acquired  string      `json:"acquired"`
		Attacker  interface{} `json:"attacker"`
		Location  struct {
			StartX int `json:"startX"`
			StartY int `json:"startY"`
			EndX   int `json:"endX"`
			EndY   int `json:"endY"`
		} `json:"location"`
	} `json:"Rymek West Upper"`
	DesertEastMid struct {
		Territory string      `json:"territory"`
		Guild     string      `json:"guild"`
		Acquired  string      `json:"acquired"`
		Attacker  interface{} `json:"attacker"`
		Location  struct {
			StartX int `json:"startX"`
			StartY int `json:"startY"`
			EndX   int `json:"endX"`
			EndY   int `json:"endY"`
		} `json:"location"`
	} `json:"Desert East Mid"`
	DesertUpper struct {
		Territory string      `json:"territory"`
		Guild     string      `json:"guild"`
		Acquired  string      `json:"acquired"`
		Attacker  interface{} `json:"attacker"`
		Location  struct {
			StartX int `json:"startX"`
			StartY int `json:"startY"`
			EndX   int `json:"endX"`
			EndY   int `json:"endY"`
		} `json:"location"`
	} `json:"Desert Upper"`
	DesertMidLower struct {
		Territory string      `json:"territory"`
		Guild     string      `json:"guild"`
		Acquired  string      `json:"acquired"`
		Attacker  interface{} `json:"attacker"`
		Location  struct {
			StartX int `json:"startX"`
			StartY int `json:"startY"`
			EndX   int `json:"endX"`
			EndY   int `json:"endY"`
		} `json:"location"`
	} `json:"Desert Mid-Lower"`
	AlmujCity struct {
		Territory string      `json:"territory"`
		Guild     string      `json:"guild"`
		Acquired  string      `json:"acquired"`
		Attacker  interface{} `json:"attacker"`
		Location  struct {
			StartX int `json:"startX"`
			StartY int `json:"startY"`
			EndX   int `json:"endX"`
			EndY   int `json:"endY"`
		} `json:"location"`
	} `json:"Almuj City"`
	DesertWestUpper struct {
		Territory string      `json:"territory"`
		Guild     string      `json:"guild"`
		Acquired  string      `json:"acquired"`
		Attacker  interface{} `json:"attacker"`
		Location  struct {
			StartX int `json:"startX"`
			StartY int `json:"startY"`
			EndX   int `json:"endX"`
			EndY   int `json:"endY"`
		} `json:"location"`
	} `json:"Desert West Upper"`
	SavannahEastLower struct {
		Territory string      `json:"territory"`
		Guild     string      `json:"guild"`
		Acquired  string      `json:"acquired"`
		Attacker  interface{} `json:"attacker"`
		Location  struct {
			StartX int `json:"startX"`
			StartY int `json:"startY"`
			EndX   int `json:"endX"`
			EndY   int `json:"endY"`
		} `json:"location"`
	} `json:"Savannah East Lower"`
	SavannahWestLower struct {
		Territory string      `json:"territory"`
		Guild     string      `json:"guild"`
		Acquired  string      `json:"acquired"`
		Attacker  interface{} `json:"attacker"`
		Location  struct {
			StartX int `json:"startX"`
			StartY int `json:"startY"`
			EndX   int `json:"endX"`
			EndY   int `json:"endY"`
		} `json:"location"`
	} `json:"Savannah West Lower"`
	Bremminglar struct {
		Territory string      `json:"territory"`
		Guild     string      `json:"guild"`
		Acquired  string      `json:"acquired"`
		Attacker  interface{} `json:"attacker"`
		Location  struct {
			StartX int `json:"startX"`
			StartY int `json:"startY"`
			EndX   int `json:"endX"`
			EndY   int `json:"endY"`
		} `json:"location"`
	} `json:"Bremminglar"`
	NemractTown struct {
		Territory string      `json:"territory"`
		Guild     string      `json:"guild"`
		Acquired  string      `json:"acquired"`
		Attacker  interface{} `json:"attacker"`
		Location  struct {
			StartX int `json:"startX"`
			StartY int `json:"startY"`
			EndX   int `json:"endX"`
			EndY   int `json:"endY"`
		} `json:"location"`
	} `json:"Nemract Town"`
	NemractRoad struct {
		Territory string      `json:"territory"`
		Guild     string      `json:"guild"`
		Acquired  string      `json:"acquired"`
		Attacker  interface{} `json:"attacker"`
		Location  struct {
			StartX int `json:"startX"`
			StartY int `json:"startY"`
			EndX   int `json:"endX"`
			EndY   int `json:"endY"`
		} `json:"location"`
	} `json:"Nemract Road"`
	NemractPlainsEast struct {
		Territory string      `json:"territory"`
		Guild     string      `json:"guild"`
		Acquired  string      `json:"acquired"`
		Attacker  interface{} `json:"attacker"`
		Location  struct {
			StartX int `json:"startX"`
			StartY int `json:"startY"`
			EndX   int `json:"endX"`
			EndY   int `json:"endY"`
		} `json:"location"`
	} `json:"Nemract Plains East"`
	NemractCathedral struct {
		Territory string      `json:"territory"`
		Guild     string      `json:"guild"`
		Acquired  string      `json:"acquired"`
		Attacker  interface{} `json:"attacker"`
		Location  struct {
			StartX int `json:"startX"`
			StartY int `json:"startY"`
			EndX   int `json:"endX"`
			EndY   int `json:"endY"`
		} `json:"location"`
	} `json:"Nemract Cathedral"`
	TheBearZoo struct {
		Territory string      `json:"territory"`
		Guild     string      `json:"guild"`
		Acquired  string      `json:"acquired"`
		Attacker  interface{} `json:"attacker"`
		Location  struct {
			StartX int `json:"startX"`
			StartY int `json:"startY"`
			EndX   int `json:"endX"`
			EndY   int `json:"endY"`
		} `json:"location"`
	} `json:"The Bear Zoo"`
	ZhightIsland struct {
		Territory string      `json:"territory"`
		Guild     string      `json:"guild"`
		Acquired  string      `json:"acquired"`
		Attacker  interface{} `json:"attacker"`
		Location  struct {
			StartX int `json:"startX"`
			StartY int `json:"startY"`
			EndX   int `json:"endX"`
			EndY   int `json:"endY"`
		} `json:"location"`
	} `json:"Zhight Island"`
	DurumIslesCenter struct {
		Territory string      `json:"territory"`
		Guild     string      `json:"guild"`
		Acquired  string      `json:"acquired"`
		Attacker  interface{} `json:"attacker"`
		Location  struct {
			StartX int `json:"startX"`
			StartY int `json:"startY"`
			EndX   int `json:"endX"`
			EndY   int `json:"endY"`
		} `json:"location"`
	} `json:"Durum Isles Center"`
	DurumIslesEast struct {
		Territory string      `json:"territory"`
		Guild     string      `json:"guild"`
		Acquired  string      `json:"acquired"`
		Attacker  interface{} `json:"attacker"`
		Location  struct {
			StartX int `json:"startX"`
			StartY int `json:"startY"`
			EndX   int `json:"endX"`
			EndY   int `json:"endY"`
		} `json:"location"`
	} `json:"Durum Isles East"`
	PirateTown struct {
		Territory string      `json:"territory"`
		Guild     string      `json:"guild"`
		Acquired  string      `json:"acquired"`
		Attacker  interface{} `json:"attacker"`
		Location  struct {
			StartX int `json:"startX"`
			StartY int `json:"startY"`
			EndX   int `json:"endX"`
			EndY   int `json:"endY"`
		} `json:"location"`
	} `json:"Pirate Town"`
	DujgonNation struct {
		Territory string      `json:"territory"`
		Guild     string      `json:"guild"`
		Acquired  string      `json:"acquired"`
		Attacker  interface{} `json:"attacker"`
		Location  struct {
			StartX int `json:"startX"`
			StartY int `json:"startY"`
			EndX   int `json:"endX"`
			EndY   int `json:"endY"`
		} `json:"location"`
	} `json:"Dujgon Nation"`
	DeadIslandNorthEast struct {
		Territory string      `json:"territory"`
		Guild     string      `json:"guild"`
		Acquired  string      `json:"acquired"`
		Attacker  interface{} `json:"attacker"`
		Location  struct {
			StartX int `json:"startX"`
			StartY int `json:"startY"`
			EndX   int `json:"endX"`
			EndY   int `json:"endY"`
		} `json:"location"`
	} `json:"Dead Island North East"`
	DeadIslandNorthWest struct {
		Territory string      `json:"territory"`
		Guild     string      `json:"guild"`
		Acquired  string      `json:"acquired"`
		Attacker  interface{} `json:"attacker"`
		Location  struct {
			StartX int `json:"startX"`
			StartY int `json:"startY"`
			EndX   int `json:"endX"`
			EndY   int `json:"endY"`
		} `json:"location"`
	} `json:"Dead Island North West"`
	MaroPeaks struct {
		Territory string      `json:"territory"`
		Guild     string      `json:"guild"`
		Acquired  string      `json:"acquired"`
		Attacker  interface{} `json:"attacker"`
		Location  struct {
			StartX int `json:"startX"`
			StartY int `json:"startY"`
			EndX   int `json:"endX"`
			EndY   int `json:"endY"`
		} `json:"location"`
	} `json:"Maro Peaks"`
	VolcanoLower struct {
		Territory string      `json:"territory"`
		Guild     string      `json:"guild"`
		Acquired  string      `json:"acquired"`
		Attacker  interface{} `json:"attacker"`
		Location  struct {
			StartX int `json:"startX"`
			StartY int `json:"startY"`
			EndX   int `json:"endX"`
			EndY   int `json:"endY"`
		} `json:"location"`
	} `json:"Volcano Lower"`
	Ternaves struct {
		Territory string      `json:"territory"`
		Guild     string      `json:"guild"`
		Acquired  string      `json:"acquired"`
		Attacker  interface{} `json:"attacker"`
		Location  struct {
			StartX int `json:"startX"`
			StartY int `json:"startY"`
			EndX   int `json:"endX"`
			EndY   int `json:"endY"`
		} `json:"location"`
	} `json:"Ternaves"`
	TernavesPlainsLower struct {
		Territory string      `json:"territory"`
		Guild     string      `json:"guild"`
		Acquired  string      `json:"acquired"`
		Attacker  interface{} `json:"attacker"`
		Location  struct {
			StartX int `json:"startX"`
			StartY int `json:"startY"`
			EndX   int `json:"endX"`
			EndY   int `json:"endY"`
		} `json:"location"`
	} `json:"Ternaves Plains Lower"`
	MiningBaseLower struct {
		Territory string      `json:"territory"`
		Guild     string      `json:"guild"`
		Acquired  string      `json:"acquired"`
		Attacker  interface{} `json:"attacker"`
		Location  struct {
			StartX int `json:"startX"`
			StartY int `json:"startY"`
			EndX   int `json:"endX"`
			EndY   int `json:"endY"`
		} `json:"location"`
	} `json:"Mining Base Lower"`
	DesolateValley struct {
		Territory string      `json:"territory"`
		Guild     string      `json:"guild"`
		Acquired  string      `json:"acquired"`
		Attacker  interface{} `json:"attacker"`
		Location  struct {
			StartX int `json:"startX"`
			StartY int `json:"startY"`
			EndX   int `json:"endX"`
			EndY   int `json:"endY"`
		} `json:"location"`
	} `json:"Desolate Valley"`
	PlainsLake struct {
		Territory string      `json:"territory"`
		Guild     string      `json:"guild"`
		Acquired  string      `json:"acquired"`
		Attacker  interface{} `json:"attacker"`
		Location  struct {
			StartX int `json:"startX"`
			StartY int `json:"startY"`
			EndX   int `json:"endX"`
			EndY   int `json:"endY"`
		} `json:"location"`
	} `json:"Plains Lake"`
	DetlasSavannahTransition struct {
		Territory string      `json:"territory"`
		Guild     string      `json:"guild"`
		Acquired  string      `json:"acquired"`
		Attacker  interface{} `json:"attacker"`
		Location  struct {
			StartX int `json:"startX"`
			StartY int `json:"startY"`
			EndX   int `json:"endX"`
			EndY   int `json:"endY"`
		} `json:"location"`
	} `json:"Detlas Savannah Transition"`
	DetlasTrailEastPlains struct {
		Territory string      `json:"territory"`
		Guild     string      `json:"guild"`
		Acquired  string      `json:"acquired"`
		Attacker  interface{} `json:"attacker"`
		Location  struct {
			StartX int `json:"startX"`
			StartY int `json:"startY"`
			EndX   int `json:"endX"`
			EndY   int `json:"endY"`
		} `json:"location"`
	} `json:"Detlas Trail East Plains"`
	LlevigarEntrance struct {
		Territory string      `json:"territory"`
		Guild     string      `json:"guild"`
		Acquired  string      `json:"acquired"`
		Attacker  interface{} `json:"attacker"`
		Location  struct {
			StartX int `json:"startX"`
			StartY int `json:"startY"`
			EndX   int `json:"endX"`
			EndY   int `json:"endY"`
		} `json:"location"`
	} `json:"Llevigar Entrance"`
	LlevigarGateWest struct {
		Territory string      `json:"territory"`
		Guild     string      `json:"guild"`
		Acquired  string      `json:"acquired"`
		Attacker  interface{} `json:"attacker"`
		Location  struct {
			StartX int `json:"startX"`
			StartY int `json:"startY"`
			EndX   int `json:"endX"`
			EndY   int `json:"endY"`
		} `json:"location"`
	} `json:"Llevigar Gate West"`
	LlevigarFarmPlainsWest struct {
		Territory string      `json:"territory"`
		Guild     string      `json:"guild"`
		Acquired  string      `json:"acquired"`
		Attacker  interface{} `json:"attacker"`
		Location  struct {
			StartX int `json:"startX"`
			StartY int `json:"startY"`
			EndX   int `json:"endX"`
			EndY   int `json:"endY"`
		} `json:"location"`
	} `json:"Llevigar Farm Plains West"`
	Cinfras struct {
		Territory string      `json:"territory"`
		Guild     string      `json:"guild"`
		Acquired  string      `json:"acquired"`
		Attacker  interface{} `json:"attacker"`
		Location  struct {
			StartX int `json:"startX"`
			StartY int `json:"startY"`
			EndX   int `json:"endX"`
			EndY   int `json:"endY"`
		} `json:"location"`
	} `json:"Cinfras"`
	LlevigarPlainsWestLower struct {
		Territory string      `json:"territory"`
		Guild     string      `json:"guild"`
		Acquired  string      `json:"acquired"`
		Attacker  interface{} `json:"attacker"`
		Location  struct {
			StartX int `json:"startX"`
			StartY int `json:"startY"`
			EndX   int `json:"endX"`
			EndY   int `json:"endY"`
		} `json:"location"`
	} `json:"Llevigar Plains West Lower"`
	LlevigarPlainsEastUpper struct {
		Territory string      `json:"territory"`
		Guild     string      `json:"guild"`
		Acquired  string      `json:"acquired"`
		Attacker  interface{} `json:"attacker"`
		Location  struct {
			StartX int `json:"startX"`
			StartY int `json:"startY"`
			EndX   int `json:"endX"`
			EndY   int `json:"endY"`
		} `json:"location"`
	} `json:"Llevigar Plains East Upper"`
	SwampEastLower struct {
		Territory string      `json:"territory"`
		Guild     string      `json:"guild"`
		Acquired  string      `json:"acquired"`
		Attacker  interface{} `json:"attacker"`
		Location  struct {
			StartX int `json:"startX"`
			StartY int `json:"startY"`
			EndX   int `json:"endX"`
			EndY   int `json:"endY"`
		} `json:"location"`
	} `json:"Swamp East Lower"`
	SwampWestMid struct {
		Territory string      `json:"territory"`
		Guild     string      `json:"guild"`
		Acquired  string      `json:"acquired"`
		Attacker  interface{} `json:"attacker"`
		Location  struct {
			StartX int `json:"startX"`
			StartY int `json:"startY"`
			EndX   int `json:"endX"`
			EndY   int `json:"endY"`
		} `json:"location"`
	} `json:"Swamp West Mid"`
	SwampEastMidUpper struct {
		Territory string      `json:"territory"`
		Guild     string      `json:"guild"`
		Acquired  string      `json:"acquired"`
		Attacker  interface{} `json:"attacker"`
		Location  struct {
			StartX int `json:"startX"`
			StartY int `json:"startY"`
			EndX   int `json:"endX"`
			EndY   int `json:"endY"`
		} `json:"location"`
	} `json:"Swamp East Mid-Upper"`
	SwampEastUpper struct {
		Territory string      `json:"territory"`
		Guild     string      `json:"guild"`
		Acquired  string      `json:"acquired"`
		Attacker  interface{} `json:"attacker"`
		Location  struct {
			StartX int `json:"startX"`
			StartY int `json:"startY"`
			EndX   int `json:"endX"`
			EndY   int `json:"endY"`
		} `json:"location"`
	} `json:"Swamp East Upper"`
	SwampDarkForestTransitionMid struct {
		Territory string      `json:"territory"`
		Guild     string      `json:"guild"`
		Acquired  string      `json:"acquired"`
		Attacker  interface{} `json:"attacker"`
		Location  struct {
			StartX int `json:"startX"`
			StartY int `json:"startY"`
			EndX   int `json:"endX"`
			EndY   int `json:"endY"`
		} `json:"location"`
	} `json:"Swamp Dark Forest Transition Mid"`
	SwampLower struct {
		Territory string      `json:"territory"`
		Guild     string      `json:"guild"`
		Acquired  string      `json:"acquired"`
		Attacker  interface{} `json:"attacker"`
		Location  struct {
			StartX int `json:"startX"`
			StartY int `json:"startY"`
			EndX   int `json:"endX"`
			EndY   int `json:"endY"`
		} `json:"location"`
	} `json:"Swamp Lower"`
	Olux struct {
		Territory string      `json:"territory"`
		Guild     string      `json:"guild"`
		Acquired  string      `json:"acquired"`
		Attacker  interface{} `json:"attacker"`
		Location  struct {
			StartX int `json:"startX"`
			StartY int `json:"startY"`
			EndX   int `json:"endX"`
			EndY   int `json:"endY"`
		} `json:"location"`
	} `json:"Olux"`
	SwampPlainsBasin struct {
		Territory string      `json:"territory"`
		Guild     string      `json:"guild"`
		Acquired  string      `json:"acquired"`
		Attacker  interface{} `json:"attacker"`
		Location  struct {
			StartX int `json:"startX"`
			StartY int `json:"startY"`
			EndX   int `json:"endX"`
			EndY   int `json:"endY"`
		} `json:"location"`
	} `json:"Swamp Plains Basin"`
	SwampMountainTransitionMid struct {
		Territory string      `json:"territory"`
		Guild     string      `json:"guild"`
		Acquired  string      `json:"acquired"`
		Attacker  interface{} `json:"attacker"`
		Location  struct {
			StartX int `json:"startX"`
			StartY int `json:"startY"`
			EndX   int `json:"endX"`
			EndY   int `json:"endY"`
		} `json:"location"`
	} `json:"Swamp Mountain Transition Mid"`
	SwampMountainTransitionUpper struct {
		Territory string      `json:"territory"`
		Guild     string      `json:"guild"`
		Acquired  string      `json:"acquired"`
		Attacker  interface{} `json:"attacker"`
		Location  struct {
			StartX int `json:"startX"`
			StartY int `json:"startY"`
			EndX   int `json:"endX"`
			EndY   int `json:"endY"`
		} `json:"location"`
	} `json:"Swamp Mountain Transition Upper"`
	QuartzMinesSouthEast struct {
		Territory string      `json:"territory"`
		Guild     string      `json:"guild"`
		Acquired  string      `json:"acquired"`
		Attacker  interface{} `json:"attacker"`
		Location  struct {
			StartX int `json:"startX"`
			StartY int `json:"startY"`
			EndX   int `json:"endX"`
			EndY   int `json:"endY"`
		} `json:"location"`
	} `json:"Quartz Mines South East"`
	QuartzMinesNorthEast struct {
		Territory string      `json:"territory"`
		Guild     string      `json:"guild"`
		Acquired  string      `json:"acquired"`
		Attacker  interface{} `json:"attacker"`
		Location  struct {
			StartX int `json:"startX"`
			StartY int `json:"startY"`
			EndX   int `json:"endX"`
			EndY   int `json:"endY"`
		} `json:"location"`
	} `json:"Quartz Mines North East"`
	OrcLake struct {
		Territory string      `json:"territory"`
		Guild     string      `json:"guild"`
		Acquired  string      `json:"acquired"`
		Attacker  interface{} `json:"attacker"`
		Location  struct {
			StartX int `json:"startX"`
			StartY int `json:"startY"`
			EndX   int `json:"endX"`
			EndY   int `json:"endY"`
		} `json:"location"`
	} `json:"Orc Lake"`
	MeteorCrater struct {
		Territory string      `json:"territory"`
		Guild     string      `json:"guild"`
		Acquired  string      `json:"acquired"`
		Attacker  interface{} `json:"attacker"`
		Location  struct {
			StartX int `json:"startX"`
			StartY int `json:"startY"`
			EndX   int `json:"endX"`
			EndY   int `json:"endY"`
		} `json:"location"`
	} `json:"Meteor Crater"`
	LoamsproutCamp struct {
		Territory string      `json:"territory"`
		Guild     string      `json:"guild"`
		Acquired  string      `json:"acquired"`
		Attacker  interface{} `json:"attacker"`
		Location  struct {
			StartX int `json:"startX"`
			StartY int `json:"startY"`
			EndX   int `json:"endX"`
			EndY   int `json:"endY"`
		} `json:"location"`
	} `json:"Loamsprout Camp"`
	GoblinPlainsWest struct {
		Territory string      `json:"territory"`
		Guild     string      `json:"guild"`
		Acquired  string      `json:"acquired"`
		Attacker  interface{} `json:"attacker"`
		Location  struct {
			StartX int `json:"startX"`
			StartY int `json:"startY"`
			EndX   int `json:"endX"`
			EndY   int `json:"endY"`
		} `json:"location"`
	} `json:"Goblin Plains West"`
	ForgottenPath struct {
		Territory string      `json:"territory"`
		Guild     string      `json:"guild"`
		Acquired  string      `json:"acquired"`
		Attacker  interface{} `json:"attacker"`
		Location  struct {
			StartX int `json:"startX"`
			StartY int `json:"startY"`
			EndX   int `json:"endX"`
			EndY   int `json:"endY"`
		} `json:"location"`
	} `json:"Forgotten Path"`
	PreLightForestTransition struct {
		Territory string      `json:"territory"`
		Guild     string      `json:"guild"`
		Acquired  string      `json:"acquired"`
		Attacker  interface{} `json:"attacker"`
		Location  struct {
			StartX int `json:"startX"`
			StartY int `json:"startY"`
			EndX   int `json:"endX"`
			EndY   int `json:"endY"`
		} `json:"location"`
	} `json:"Pre-Light Forest Transition"`
	RoadToLightForest struct {
		Territory string      `json:"territory"`
		Guild     string      `json:"guild"`
		Acquired  string      `json:"acquired"`
		Attacker  interface{} `json:"attacker"`
		Location  struct {
			StartX int `json:"startX"`
			StartY int `json:"startY"`
			EndX   int `json:"endX"`
			EndY   int `json:"endY"`
		} `json:"location"`
	} `json:"Road To Light Forest"`
	EfilimSouthPlains struct {
		Territory string      `json:"territory"`
		Guild     string      `json:"guild"`
		Acquired  string      `json:"acquired"`
		Attacker  interface{} `json:"attacker"`
		Location  struct {
			StartX int `json:"startX"`
			StartY int `json:"startY"`
			EndX   int `json:"endX"`
			EndY   int `json:"endY"`
		} `json:"location"`
	} `json:"Efilim South Plains"`
	EfilimSouthEastPlains struct {
		Territory string      `json:"territory"`
		Guild     string      `json:"guild"`
		Acquired  string      `json:"acquired"`
		Attacker  interface{} `json:"attacker"`
		Location  struct {
			StartX int `json:"startX"`
			StartY int `json:"startY"`
			EndX   int `json:"endX"`
			EndY   int `json:"endY"`
		} `json:"location"`
	} `json:"Efilim South East Plains"`
	LightForestEntrance struct {
		Territory string      `json:"territory"`
		Guild     string      `json:"guild"`
		Acquired  string      `json:"acquired"`
		Attacker  interface{} `json:"attacker"`
		Location  struct {
			StartX int `json:"startX"`
			StartY int `json:"startY"`
			EndX   int `json:"endX"`
			EndY   int `json:"endY"`
		} `json:"location"`
	} `json:"Light Forest Entrance"`
	AldoreiValleyWestEntrance struct {
		Territory string      `json:"territory"`
		Guild     string      `json:"guild"`
		Acquired  string      `json:"acquired"`
		Attacker  interface{} `json:"attacker"`
		Location  struct {
			StartX int `json:"startX"`
			StartY int `json:"startY"`
			EndX   int `json:"endX"`
			EndY   int `json:"endY"`
		} `json:"location"`
	} `json:"Aldorei Valley West Entrance"`
	LightForestNorthExit struct {
		Territory string      `json:"territory"`
		Guild     string      `json:"guild"`
		Acquired  string      `json:"acquired"`
		Attacker  interface{} `json:"attacker"`
		Location  struct {
			StartX int `json:"startX"`
			StartY int `json:"startY"`
			EndX   int `json:"endX"`
			EndY   int `json:"endY"`
		} `json:"location"`
	} `json:"Light Forest North Exit"`
	LightForestWestMid struct {
		Territory string      `json:"territory"`
		Guild     string      `json:"guild"`
		Acquired  string      `json:"acquired"`
		Attacker  interface{} `json:"attacker"`
		Location  struct {
			StartX int `json:"startX"`
			StartY int `json:"startY"`
			EndX   int `json:"endX"`
			EndY   int `json:"endY"`
		} `json:"location"`
	} `json:"Light Forest West Mid"`
	LightForestEastUpper struct {
		Territory string      `json:"territory"`
		Guild     string      `json:"guild"`
		Acquired  string      `json:"acquired"`
		Attacker  interface{} `json:"attacker"`
		Location  struct {
			StartX int `json:"startX"`
			StartY int `json:"startY"`
			EndX   int `json:"endX"`
			EndY   int `json:"endY"`
		} `json:"location"`
	} `json:"Light Forest East Upper"`
	LightForestEastLower struct {
		Territory string      `json:"territory"`
		Guild     string      `json:"guild"`
		Acquired  string      `json:"acquired"`
		Attacker  interface{} `json:"attacker"`
		Location  struct {
			StartX int `json:"startX"`
			StartY int `json:"startY"`
			EndX   int `json:"endX"`
			EndY   int `json:"endY"`
		} `json:"location"`
	} `json:"Light Forest East Lower"`
	MantisNest struct {
		Territory string      `json:"territory"`
		Guild     string      `json:"guild"`
		Acquired  string      `json:"acquired"`
		Attacker  interface{} `json:"attacker"`
		Location  struct {
			StartX int `json:"startX"`
			StartY int `json:"startY"`
			EndX   int `json:"endX"`
			EndY   int `json:"endY"`
		} `json:"location"`
	} `json:"Mantis Nest"`
	PathToCinfras struct {
		Territory string      `json:"territory"`
		Guild     string      `json:"guild"`
		Acquired  string      `json:"acquired"`
		Attacker  interface{} `json:"attacker"`
		Location  struct {
			StartX int `json:"startX"`
			StartY int `json:"startY"`
			EndX   int `json:"endX"`
			EndY   int `json:"endY"`
		} `json:"location"`
	} `json:"Path to Cinfras"`
	Gelibord struct {
		Territory string      `json:"territory"`
		Guild     string      `json:"guild"`
		Acquired  string      `json:"acquired"`
		Attacker  interface{} `json:"attacker"`
		Location  struct {
			StartX int `json:"startX"`
			StartY int `json:"startY"`
			EndX   int `json:"endX"`
			EndY   int `json:"endY"`
		} `json:"location"`
	} `json:"Gelibord"`
	GelibordCastle struct {
		Territory string      `json:"territory"`
		Guild     string      `json:"guild"`
		Acquired  string      `json:"acquired"`
		Attacker  interface{} `json:"attacker"`
		Location  struct {
			StartX int `json:"startX"`
			StartY int `json:"startY"`
			EndX   int `json:"endX"`
			EndY   int `json:"endY"`
		} `json:"location"`
	} `json:"Gelibord Castle"`
	FortressNorth struct {
		Territory string      `json:"territory"`
		Guild     string      `json:"guild"`
		Acquired  string      `json:"acquired"`
		Attacker  interface{} `json:"attacker"`
		Location  struct {
			StartX int `json:"startX"`
			StartY int `json:"startY"`
			EndX   int `json:"endX"`
			EndY   int `json:"endY"`
		} `json:"location"`
	} `json:"Fortress North"`
	MansionOfInsanity struct {
		Territory string      `json:"territory"`
		Guild     string      `json:"guild"`
		Acquired  string      `json:"acquired"`
		Attacker  interface{} `json:"attacker"`
		Location  struct {
			StartX int `json:"startX"`
			StartY int `json:"startY"`
			EndX   int `json:"endX"`
			EndY   int `json:"endY"`
		} `json:"location"`
	} `json:"Mansion of Insanity"`
	LexdalesPrison struct {
		Territory string      `json:"territory"`
		Guild     string      `json:"guild"`
		Acquired  string      `json:"acquired"`
		Attacker  interface{} `json:"attacker"`
		Location  struct {
			StartX int `json:"startX"`
			StartY int `json:"startY"`
			EndX   int `json:"endX"`
			EndY   int `json:"endY"`
		} `json:"location"`
	} `json:"Lexdales Prison"`
	EntranceToKander struct {
		Territory string      `json:"territory"`
		Guild     string      `json:"guild"`
		Acquired  string      `json:"acquired"`
		Attacker  interface{} `json:"attacker"`
		Location  struct {
			StartX int `json:"startX"`
			StartY int `json:"startY"`
			EndX   int `json:"endX"`
			EndY   int `json:"endY"`
		} `json:"location"`
	} `json:"Entrance to Kander"`
	MesquisTower struct {
		Territory string      `json:"territory"`
		Guild     string      `json:"guild"`
		Acquired  string      `json:"acquired"`
		Attacker  interface{} `json:"attacker"`
		Location  struct {
			StartX int `json:"startX"`
			StartY int `json:"startY"`
			EndX   int `json:"endX"`
			EndY   int `json:"endY"`
		} `json:"location"`
	} `json:"Mesquis Tower"`
	PathToTalor struct {
		Territory string      `json:"territory"`
		Guild     string      `json:"guild"`
		Acquired  string      `json:"acquired"`
		Attacker  interface{} `json:"attacker"`
		Location  struct {
			StartX int `json:"startX"`
			StartY int `json:"startY"`
			EndX   int `json:"endX"`
			EndY   int `json:"endY"`
		} `json:"location"`
	} `json:"Path to Talor"`
	DarkForestVillage struct {
		Territory string      `json:"territory"`
		Guild     string      `json:"guild"`
		Acquired  string      `json:"acquired"`
		Attacker  interface{} `json:"attacker"`
		Location  struct {
			StartX int `json:"startX"`
			StartY int `json:"startY"`
			EndX   int `json:"endX"`
			EndY   int `json:"endY"`
		} `json:"location"`
	} `json:"Dark Forest Village"`
	OldCrossroadsNorth struct {
		Territory string      `json:"territory"`
		Guild     string      `json:"guild"`
		Acquired  string      `json:"acquired"`
		Attacker  interface{} `json:"attacker"`
		Location  struct {
			StartX int `json:"startX"`
			StartY int `json:"startY"`
			EndX   int `json:"endX"`
			EndY   int `json:"endY"`
		} `json:"location"`
	} `json:"Old Crossroads North"`
	FungalGrove struct {
		Territory string      `json:"territory"`
		Guild     string      `json:"guild"`
		Acquired  string      `json:"acquired"`
		Attacker  interface{} `json:"attacker"`
		Location  struct {
			StartX int `json:"startX"`
			StartY int `json:"startY"`
			EndX   int `json:"endX"`
			EndY   int `json:"endY"`
		} `json:"location"`
	} `json:"Fungal Grove"`
	HeartOfDecay struct {
		Territory string      `json:"territory"`
		Guild     string      `json:"guild"`
		Acquired  string      `json:"acquired"`
		Attacker  interface{} `json:"attacker"`
		Location  struct {
			StartX int `json:"startX"`
			StartY int `json:"startY"`
			EndX   int `json:"endX"`
			EndY   int `json:"endY"`
		} `json:"location"`
	} `json:"Heart of Decay"`
	MushroomHill struct {
		Territory string      `json:"territory"`
		Guild     string      `json:"guild"`
		Acquired  string      `json:"acquired"`
		Attacker  interface{} `json:"attacker"`
		Location  struct {
			StartX int `json:"startX"`
			StartY int `json:"startY"`
			EndX   int `json:"endX"`
			EndY   int `json:"endY"`
		} `json:"location"`
	} `json:"Mushroom Hill"`
	DarkForestCinfrasTransition struct {
		Territory string      `json:"territory"`
		Guild     string      `json:"guild"`
		Acquired  string      `json:"acquired"`
		Attacker  interface{} `json:"attacker"`
		Location  struct {
			StartX int `json:"startX"`
			StartY int `json:"startY"`
			EndX   int `json:"endX"`
			EndY   int `json:"endY"`
		} `json:"location"`
	} `json:"Dark Forest Cinfras Transition"`
	AldoreiValleySouthEntrance struct {
		Territory string      `json:"territory"`
		Guild     string      `json:"guild"`
		Acquired  string      `json:"acquired"`
		Attacker  interface{} `json:"attacker"`
		Location  struct {
			StartX int `json:"startX"`
			StartY int `json:"startY"`
			EndX   int `json:"endX"`
			EndY   int `json:"endY"`
		} `json:"location"`
	} `json:"Aldorei Valley South Entrance"`
	CinfrasCountyLower struct {
		Territory string      `json:"territory"`
		Guild     string      `json:"guild"`
		Acquired  string      `json:"acquired"`
		Attacker  interface{} `json:"attacker"`
		Location  struct {
			StartX int `json:"startX"`
			StartY int `json:"startY"`
			EndX   int `json:"endX"`
			EndY   int `json:"endY"`
		} `json:"location"`
	} `json:"Cinfras County Lower"`
	CinfrasCountyMidUpper struct {
		Territory string      `json:"territory"`
		Guild     string      `json:"guild"`
		Acquired  string      `json:"acquired"`
		Attacker  interface{} `json:"attacker"`
		Location  struct {
			StartX int `json:"startX"`
			StartY int `json:"startY"`
			EndX   int `json:"endX"`
			EndY   int `json:"endY"`
		} `json:"location"`
	} `json:"Cinfras County Mid-Upper"`
	GyliaLakeSouthEast struct {
		Territory string      `json:"territory"`
		Guild     string      `json:"guild"`
		Acquired  string      `json:"acquired"`
		Attacker  interface{} `json:"attacker"`
		Location  struct {
			StartX int `json:"startX"`
			StartY int `json:"startY"`
			EndX   int `json:"endX"`
			EndY   int `json:"endY"`
		} `json:"location"`
	} `json:"Gylia Lake South East"`
	GyliaLakeNorthWest struct {
		Territory string      `json:"territory"`
		Guild     string      `json:"guild"`
		Acquired  string      `json:"acquired"`
		Attacker  interface{} `json:"attacker"`
		Location  struct {
			StartX int `json:"startX"`
			StartY int `json:"startY"`
			EndX   int `json:"endX"`
			EndY   int `json:"endY"`
		} `json:"location"`
	} `json:"Gylia Lake North West"`
	GertCamp struct {
		Territory string      `json:"territory"`
		Guild     string      `json:"guild"`
		Acquired  string      `json:"acquired"`
		Attacker  interface{} `json:"attacker"`
		Location  struct {
			StartX int `json:"startX"`
			StartY int `json:"startY"`
			EndX   int `json:"endX"`
			EndY   int `json:"endY"`
		} `json:"location"`
	} `json:"Gert Camp"`
	AldoreiValleyLower struct {
		Territory string      `json:"territory"`
		Guild     string      `json:"guild"`
		Acquired  string      `json:"acquired"`
		Attacker  interface{} `json:"attacker"`
		Location  struct {
			StartX int `json:"startX"`
			StartY int `json:"startY"`
			EndX   int `json:"endX"`
			EndY   int `json:"endY"`
		} `json:"location"`
	} `json:"Aldorei Valley Lower"`
	AldoreiValleyUpper struct {
		Territory string      `json:"territory"`
		Guild     string      `json:"guild"`
		Acquired  string      `json:"acquired"`
		Attacker  interface{} `json:"attacker"`
		Location  struct {
			StartX int `json:"startX"`
			StartY int `json:"startY"`
			EndX   int `json:"endX"`
			EndY   int `json:"endY"`
		} `json:"location"`
	} `json:"Aldorei Valley Upper"`
	AldoreiSWaterfall struct {
		Territory string      `json:"territory"`
		Guild     string      `json:"guild"`
		Acquired  string      `json:"acquired"`
		Attacker  interface{} `json:"attacker"`
		Location  struct {
			StartX int `json:"startX"`
			StartY int `json:"startY"`
			EndX   int `json:"endX"`
			EndY   int `json:"endY"`
		} `json:"location"`
	} `json:"Aldorei's Waterfall"`
	AldoreiSArch struct {
		Territory string      `json:"territory"`
		Guild     string      `json:"guild"`
		Acquired  string      `json:"acquired"`
		Attacker  interface{} `json:"attacker"`
		Location  struct {
			StartX int `json:"startX"`
			StartY int `json:"startY"`
			EndX   int `json:"endX"`
			EndY   int `json:"endY"`
		} `json:"location"`
	} `json:"Aldorei's Arch"`
	GhostlyPath struct {
		Territory string      `json:"territory"`
		Guild     string      `json:"guild"`
		Acquired  string      `json:"acquired"`
		Attacker  interface{} `json:"attacker"`
		Location  struct {
			StartX int `json:"startX"`
			StartY int `json:"startY"`
			EndX   int `json:"endX"`
			EndY   int `json:"endY"`
		} `json:"location"`
	} `json:"Ghostly Path"`
	BurningAirship struct {
		Territory string      `json:"territory"`
		Guild     string      `json:"guild"`
		Acquired  string      `json:"acquired"`
		Attacker  interface{} `json:"attacker"`
		Location  struct {
			StartX int `json:"startX"`
			StartY int `json:"startY"`
			EndX   int `json:"endX"`
			EndY   int `json:"endY"`
		} `json:"location"`
	} `json:"Burning Airship"`
	Thanos struct {
		Territory string      `json:"territory"`
		Guild     string      `json:"guild"`
		Acquired  string      `json:"acquired"`
		Attacker  interface{} `json:"attacker"`
		Location  struct {
			StartX int `json:"startX"`
			StartY int `json:"startY"`
			EndX   int `json:"endX"`
			EndY   int `json:"endY"`
		} `json:"location"`
	} `json:"Thanos"`
	PathToMilitaryBase struct {
		Territory string      `json:"territory"`
		Guild     string      `json:"guild"`
		Acquired  string      `json:"acquired"`
		Attacker  interface{} `json:"attacker"`
		Location  struct {
			StartX int `json:"startX"`
			StartY int `json:"startY"`
			EndX   int `json:"endX"`
			EndY   int `json:"endY"`
		} `json:"location"`
	} `json:"Path To Military Base"`
	MilitaryBaseUpper struct {
		Territory string      `json:"territory"`
		Guild     string      `json:"guild"`
		Acquired  string      `json:"acquired"`
		Attacker  interface{} `json:"attacker"`
		Location  struct {
			StartX int `json:"startX"`
			StartY int `json:"startY"`
			EndX   int `json:"endX"`
			EndY   int `json:"endY"`
		} `json:"location"`
	} `json:"Military Base Upper"`
	PathToOzothSSpireLower struct {
		Territory string      `json:"territory"`
		Guild     string      `json:"guild"`
		Acquired  string      `json:"acquired"`
		Attacker  interface{} `json:"attacker"`
		Location  struct {
			StartX int `json:"startX"`
			StartY int `json:"startY"`
			EndX   int `json:"endX"`
			EndY   int `json:"endY"`
		} `json:"location"`
	} `json:"Path To Ozoth's Spire Lower"`
	PathToOzothSSpireUpper struct {
		Territory string      `json:"territory"`
		Guild     string      `json:"guild"`
		Acquired  string      `json:"acquired"`
		Attacker  interface{} `json:"attacker"`
		Location  struct {
			StartX int `json:"startX"`
			StartY int `json:"startY"`
			EndX   int `json:"endX"`
			EndY   int `json:"endY"`
		} `json:"location"`
	} `json:"Path To Ozoth's Spire Upper"`
	BanditCaveUpper struct {
		Territory string      `json:"territory"`
		Guild     string      `json:"guild"`
		Acquired  string      `json:"acquired"`
		Attacker  interface{} `json:"attacker"`
		Location  struct {
			StartX int `json:"startX"`
			StartY int `json:"startY"`
			EndX   int `json:"endX"`
			EndY   int `json:"endY"`
		} `json:"location"`
	} `json:"Bandit Cave Upper"`
	CanyonPathNorthWest struct {
		Territory string      `json:"territory"`
		Guild     string      `json:"guild"`
		Acquired  string      `json:"acquired"`
		Attacker  interface{} `json:"attacker"`
		Location  struct {
			StartX int `json:"startX"`
			StartY int `json:"startY"`
			EndX   int `json:"endX"`
			EndY   int `json:"endY"`
		} `json:"location"`
	} `json:"Canyon Path North West"`
	CanyonWaterfallNorth struct {
		Territory string      `json:"territory"`
		Guild     string      `json:"guild"`
		Acquired  string      `json:"acquired"`
		Attacker  interface{} `json:"attacker"`
		Location  struct {
			StartX int `json:"startX"`
			StartY int `json:"startY"`
			EndX   int `json:"endX"`
			EndY   int `json:"endY"`
		} `json:"location"`
	} `json:"Canyon Waterfall North"`
	CanyonLowerSouthEast struct {
		Territory string      `json:"territory"`
		Guild     string      `json:"guild"`
		Acquired  string      `json:"acquired"`
		Attacker  interface{} `json:"attacker"`
		Location  struct {
			StartX int `json:"startX"`
			StartY int `json:"startY"`
			EndX   int `json:"endX"`
			EndY   int `json:"endY"`
		} `json:"location"`
	} `json:"Canyon Lower South East"`
	CanyonPathNorthMid struct {
		Territory string      `json:"territory"`
		Guild     string      `json:"guild"`
		Acquired  string      `json:"acquired"`
		Attacker  interface{} `json:"attacker"`
		Location  struct {
			StartX int `json:"startX"`
			StartY int `json:"startY"`
			EndX   int `json:"endX"`
			EndY   int `json:"endY"`
		} `json:"location"`
	} `json:"Canyon Path North Mid"`
	CanyonValleySouth struct {
		Territory string      `json:"territory"`
		Guild     string      `json:"guild"`
		Acquired  string      `json:"acquired"`
		Attacker  interface{} `json:"attacker"`
		Location  struct {
			StartX int `json:"startX"`
			StartY int `json:"startY"`
			EndX   int `json:"endX"`
			EndY   int `json:"endY"`
		} `json:"location"`
	} `json:"Canyon Valley South"`
	ThanosExit struct {
		Territory string      `json:"territory"`
		Guild     string      `json:"guild"`
		Acquired  string      `json:"acquired"`
		Attacker  interface{} `json:"attacker"`
		Location  struct {
			StartX int `json:"startX"`
			StartY int `json:"startY"`
			EndX   int `json:"endX"`
			EndY   int `json:"endY"`
		} `json:"location"`
	} `json:"Thanos Exit"`
	CanyonMountainEast struct {
		Territory string      `json:"territory"`
		Guild     string      `json:"guild"`
		Acquired  string      `json:"acquired"`
		Attacker  interface{} `json:"attacker"`
		Location  struct {
			StartX int `json:"startX"`
			StartY int `json:"startY"`
			EndX   int `json:"endX"`
			EndY   int `json:"endY"`
		} `json:"location"`
	} `json:"Canyon Mountain East"`
	CanyonWaterfallMidNorth struct {
		Territory string      `json:"territory"`
		Guild     string      `json:"guild"`
		Acquired  string      `json:"acquired"`
		Attacker  interface{} `json:"attacker"`
		Location  struct {
			StartX int `json:"startX"`
			StartY int `json:"startY"`
			EndX   int `json:"endX"`
			EndY   int `json:"endY"`
		} `json:"location"`
	} `json:"Canyon Waterfall Mid North"`
	CanyonSurvivor struct {
		Territory string      `json:"territory"`
		Guild     string      `json:"guild"`
		Acquired  string      `json:"acquired"`
		Attacker  interface{} `json:"attacker"`
		Location  struct {
			StartX int `json:"startX"`
			StartY int `json:"startY"`
			EndX   int `json:"endX"`
			EndY   int `json:"endY"`
		} `json:"location"`
	} `json:"Canyon Survivor"`
	ThanosExitUpper struct {
		Territory string      `json:"territory"`
		Guild     string      `json:"guild"`
		Acquired  string      `json:"acquired"`
		Attacker  interface{} `json:"attacker"`
		Location  struct {
			StartX int `json:"startX"`
			StartY int `json:"startY"`
			EndX   int `json:"endX"`
			EndY   int `json:"endY"`
		} `json:"location"`
	} `json:"Thanos Exit Upper"`
	WizardTowerNorth struct {
		Territory string      `json:"territory"`
		Guild     string      `json:"guild"`
		Acquired  string      `json:"acquired"`
		Attacker  interface{} `json:"attacker"`
		Location  struct {
			StartX int `json:"startX"`
			StartY int `json:"startY"`
			EndX   int `json:"endX"`
			EndY   int `json:"endY"`
		} `json:"location"`
	} `json:"Wizard Tower North"`
	MountainEdge struct {
		Territory string      `json:"territory"`
		Guild     string      `json:"guild"`
		Acquired  string      `json:"acquired"`
		Attacker  interface{} `json:"attacker"`
		Location  struct {
			StartX int `json:"startX"`
			StartY int `json:"startY"`
			EndX   int `json:"endX"`
			EndY   int `json:"endY"`
		} `json:"location"`
	} `json:"Mountain Edge"`
	ValleyOfTheLost struct {
		Territory string      `json:"territory"`
		Guild     string      `json:"guild"`
		Acquired  string      `json:"acquired"`
		Attacker  interface{} `json:"attacker"`
		Location  struct {
			StartX int `json:"startX"`
			StartY int `json:"startY"`
			EndX   int `json:"endX"`
			EndY   int `json:"endY"`
		} `json:"location"`
	} `json:"Valley of the Lost"`
	KroltonSCave struct {
		Territory string      `json:"territory"`
		Guild     string      `json:"guild"`
		Acquired  string      `json:"acquired"`
		Attacker  interface{} `json:"attacker"`
		Location  struct {
			StartX int `json:"startX"`
			StartY int `json:"startY"`
			EndX   int `json:"endX"`
			EndY   int `json:"endY"`
		} `json:"location"`
	} `json:"Krolton's Cave"`
	CanyonHighPath struct {
		Territory string      `json:"territory"`
		Guild     string      `json:"guild"`
		Acquired  string      `json:"acquired"`
		Attacker  interface{} `json:"attacker"`
		Location  struct {
			StartX int `json:"startX"`
			StartY int `json:"startY"`
			EndX   int `json:"endX"`
			EndY   int `json:"endY"`
		} `json:"location"`
	} `json:"Canyon High Path"`
	CliffsideValley struct {
		Territory string      `json:"territory"`
		Guild     string      `json:"guild"`
		Acquired  string      `json:"acquired"`
		Attacker  interface{} `json:"attacker"`
		Location  struct {
			StartX int `json:"startX"`
			StartY int `json:"startY"`
			EndX   int `json:"endX"`
			EndY   int `json:"endY"`
		} `json:"location"`
	} `json:"Cliffside Valley"`
	AirTempleUpper struct {
		Territory string      `json:"territory"`
		Guild     string      `json:"guild"`
		Acquired  string      `json:"acquired"`
		Attacker  interface{} `json:"attacker"`
		Location  struct {
			StartX int `json:"startX"`
			StartY int `json:"startY"`
			EndX   int `json:"endX"`
			EndY   int `json:"endY"`
		} `json:"location"`
	} `json:"Air Temple Upper"`
	CanyonOfTheLost struct {
		Territory string      `json:"territory"`
		Guild     string      `json:"guild"`
		Acquired  string      `json:"acquired"`
		Attacker  interface{} `json:"attacker"`
		Location  struct {
			StartX int `json:"startX"`
			StartY int `json:"startY"`
			EndX   int `json:"endX"`
			EndY   int `json:"endY"`
		} `json:"location"`
	} `json:"Canyon Of The Lost"`
	KandonFarm struct {
		Territory string      `json:"territory"`
		Guild     string      `json:"guild"`
		Acquired  string      `json:"acquired"`
		Attacker  interface{} `json:"attacker"`
		Location  struct {
			StartX int `json:"startX"`
			StartY int `json:"startY"`
			EndX   int `json:"endX"`
			EndY   int `json:"endY"`
		} `json:"location"`
	} `json:"Kandon Farm"`
	CliffsidePassageNorth struct {
		Territory string      `json:"territory"`
		Guild     string      `json:"guild"`
		Acquired  string      `json:"acquired"`
		Attacker  interface{} `json:"attacker"`
		Location  struct {
			StartX int `json:"startX"`
			StartY int `json:"startY"`
			EndX   int `json:"endX"`
			EndY   int `json:"endY"`
		} `json:"location"`
	} `json:"Cliffside Passage North"`
	EntranceToTheseadSouth struct {
		Territory string      `json:"territory"`
		Guild     string      `json:"guild"`
		Acquired  string      `json:"acquired"`
		Attacker  interface{} `json:"attacker"`
		Location  struct {
			StartX int `json:"startX"`
			StartY int `json:"startY"`
			EndX   int `json:"endX"`
			EndY   int `json:"endY"`
		} `json:"location"`
	} `json:"Entrance to Thesead South"`
	CherryBlossomForest struct {
		Territory string      `json:"territory"`
		Guild     string      `json:"guild"`
		Acquired  string      `json:"acquired"`
		Attacker  interface{} `json:"attacker"`
		Location  struct {
			StartX int `json:"startX"`
			StartY int `json:"startY"`
			EndX   int `json:"endX"`
			EndY   int `json:"endY"`
		} `json:"location"`
	} `json:"Cherry Blossom Forest"`
	TheseadSuburbs struct {
		Territory string      `json:"territory"`
		Guild     string      `json:"guild"`
		Acquired  string      `json:"acquired"`
		Attacker  interface{} `json:"attacker"`
		Location  struct {
			StartX int `json:"startX"`
			StartY int `json:"startY"`
			EndX   int `json:"endX"`
			EndY   int `json:"endY"`
		} `json:"location"`
	} `json:"Thesead Suburbs"`
	EntranceToRodoroc struct {
		Territory string      `json:"territory"`
		Guild     string      `json:"guild"`
		Acquired  string      `json:"acquired"`
		Attacker  interface{} `json:"attacker"`
		Location  struct {
			StartX int `json:"startX"`
			StartY int `json:"startY"`
			EndX   int `json:"endX"`
			EndY   int `json:"endY"`
		} `json:"location"`
	} `json:"Entrance to Rodoroc"`
	MoltenHeightsPortal struct {
		Territory string      `json:"territory"`
		Guild     string      `json:"guild"`
		Acquired  string      `json:"acquired"`
		Attacker  interface{} `json:"attacker"`
		Location  struct {
			StartX int `json:"startX"`
			StartY int `json:"startY"`
			EndX   int `json:"endX"`
			EndY   int `json:"endY"`
		} `json:"location"`
	} `json:"Molten Heights Portal"`
	LavaLakeBridge struct {
		Territory string      `json:"territory"`
		Guild     string      `json:"guild"`
		Acquired  string      `json:"acquired"`
		Attacker  interface{} `json:"attacker"`
		Location  struct {
			StartX int `json:"startX"`
			StartY int `json:"startY"`
			EndX   int `json:"endX"`
			EndY   int `json:"endY"`
		} `json:"location"`
	} `json:"Lava Lake Bridge"`
	ActiveVolcano struct {
		Territory string      `json:"territory"`
		Guild     string      `json:"guild"`
		Acquired  string      `json:"acquired"`
		Attacker  interface{} `json:"attacker"`
		Location  struct {
			StartX int `json:"startX"`
			StartY int `json:"startY"`
			EndX   int `json:"endX"`
			EndY   int `json:"endY"`
		} `json:"location"`
	} `json:"Active Volcano"`
	Ahmsord struct {
		Territory string      `json:"territory"`
		Guild     string      `json:"guild"`
		Acquired  string      `json:"acquired"`
		Attacker  interface{} `json:"attacker"`
		Location  struct {
			StartX int `json:"startX"`
			StartY int `json:"startY"`
			EndX   int `json:"endX"`
			EndY   int `json:"endY"`
		} `json:"location"`
	} `json:"Ahmsord"`
	SnailIsland struct {
		Territory string      `json:"territory"`
		Guild     string      `json:"guild"`
		Acquired  string      `json:"acquired"`
		Attacker  interface{} `json:"attacker"`
		Location  struct {
			StartX int `json:"startX"`
			StartY int `json:"startY"`
			EndX   int `json:"endX"`
			EndY   int `json:"endY"`
		} `json:"location"`
	} `json:"Snail Island"`
	DernelJungleMid struct {
		Territory string      `json:"territory"`
		Guild     string      `json:"guild"`
		Acquired  string      `json:"acquired"`
		Attacker  interface{} `json:"attacker"`
		Location  struct {
			StartX int `json:"startX"`
			StartY int `json:"startY"`
			EndX   int `json:"endX"`
			EndY   int `json:"endY"`
		} `json:"location"`
	} `json:"Dernel Jungle Mid"`
	CorkusCity struct {
		Territory string      `json:"territory"`
		Guild     string      `json:"guild"`
		Acquired  string      `json:"acquired"`
		Attacker  interface{} `json:"attacker"`
		Location  struct {
			StartX int `json:"startX"`
			StartY int `json:"startY"`
			EndX   int `json:"endX"`
			EndY   int `json:"endY"`
		} `json:"location"`
	} `json:"Corkus City"`
	CorkusCitySouth struct {
		Territory string      `json:"territory"`
		Guild     string      `json:"guild"`
		Acquired  string      `json:"acquired"`
		Attacker  interface{} `json:"attacker"`
		Location  struct {
			StartX int `json:"startX"`
			StartY int `json:"startY"`
			EndX   int `json:"endX"`
			EndY   int `json:"endY"`
		} `json:"location"`
	} `json:"Corkus City South"`
	RoadToMine struct {
		Territory string      `json:"territory"`
		Guild     string      `json:"guild"`
		Acquired  string      `json:"acquired"`
		Attacker  interface{} `json:"attacker"`
		Location  struct {
			StartX int `json:"startX"`
			StartY int `json:"startY"`
			EndX   int `json:"endX"`
			EndY   int `json:"endY"`
		} `json:"location"`
	} `json:"Road To Mine"`
	LegendaryIsland struct {
		Territory string      `json:"territory"`
		Guild     string      `json:"guild"`
		Acquired  string      `json:"acquired"`
		Attacker  interface{} `json:"attacker"`
		Location  struct {
			StartX int `json:"startX"`
			StartY int `json:"startY"`
			EndX   int `json:"endX"`
			EndY   int `json:"endY"`
		} `json:"location"`
	} `json:"Legendary Island"`
	CorkusForestSouth struct {
		Territory string      `json:"territory"`
		Guild     string      `json:"guild"`
		Acquired  string      `json:"acquired"`
		Attacker  interface{} `json:"attacker"`
		Location  struct {
			StartX int `json:"startX"`
			StartY int `json:"startY"`
			EndX   int `json:"endX"`
			EndY   int `json:"endY"`
		} `json:"location"`
	} `json:"Corkus Forest South"`
	CorkusMountain struct {
		Territory string      `json:"territory"`
		Guild     string      `json:"guild"`
		Acquired  string      `json:"acquired"`
		Attacker  interface{} `json:"attacker"`
		Location  struct {
			StartX int `json:"startX"`
			StartY int `json:"startY"`
			EndX   int `json:"endX"`
			EndY   int `json:"endY"`
		} `json:"location"`
	} `json:"Corkus Mountain"`
	CorkusDocks struct {
		Territory string      `json:"territory"`
		Guild     string      `json:"guild"`
		Acquired  string      `json:"acquired"`
		Attacker  interface{} `json:"attacker"`
		Location  struct {
			StartX int `json:"startX"`
			StartY int `json:"startY"`
			EndX   int `json:"endX"`
			EndY   int `json:"endY"`
		} `json:"location"`
	} `json:"Corkus Docks"`
	CorkusSeaPort struct {
		Territory string      `json:"territory"`
		Guild     string      `json:"guild"`
		Acquired  string      `json:"acquired"`
		Attacker  interface{} `json:"attacker"`
		Location  struct {
			StartX int `json:"startX"`
			StartY int `json:"startY"`
			EndX   int `json:"endX"`
			EndY   int `json:"endY"`
		} `json:"location"`
	} `json:"Corkus Sea Port"`
	CorkusStatue struct {
		Territory string      `json:"territory"`
		Guild     string      `json:"guild"`
		Acquired  string      `json:"acquired"`
		Attacker  interface{} `json:"attacker"`
		Location  struct {
			StartX int `json:"startX"`
			StartY int `json:"startY"`
			EndX   int `json:"endX"`
			EndY   int `json:"endY"`
		} `json:"location"`
	} `json:"Corkus Statue"`
	BloodyBeach struct {
		Territory string      `json:"territory"`
		Guild     string      `json:"guild"`
		Acquired  string      `json:"acquired"`
		Attacker  interface{} `json:"attacker"`
		Location  struct {
			StartX int `json:"startX"`
			StartY int `json:"startY"`
			EndX   int `json:"endX"`
			EndY   int `json:"endY"`
		} `json:"location"`
	} `json:"Bloody Beach"`
	FrozenFort struct {
		Territory string      `json:"territory"`
		Guild     string      `json:"guild"`
		Acquired  string      `json:"acquired"`
		Attacker  interface{} `json:"attacker"`
		Location  struct {
			StartX int `json:"startX"`
			StartY int `json:"startY"`
			EndX   int `json:"endX"`
			EndY   int `json:"endY"`
		} `json:"location"`
	} `json:"Frozen Fort"`
	PathToAhmsordLower struct {
		Territory string      `json:"territory"`
		Guild     string      `json:"guild"`
		Acquired  string      `json:"acquired"`
		Attacker  interface{} `json:"attacker"`
		Location  struct {
			StartX int `json:"startX"`
			StartY int `json:"startY"`
			EndX   int `json:"endX"`
			EndY   int `json:"endY"`
		} `json:"location"`
	} `json:"Path to Ahmsord Lower"`
	KandonRidge struct {
		Territory string      `json:"territory"`
		Guild     string      `json:"guild"`
		Acquired  string      `json:"acquired"`
		Attacker  interface{} `json:"attacker"`
		Location  struct {
			StartX int `json:"startX"`
			StartY int `json:"startY"`
			EndX   int `json:"endX"`
			EndY   int `json:"endY"`
		} `json:"location"`
	} `json:"Kandon Ridge"`
	DragonlingNests struct {
		Territory string      `json:"territory"`
		Guild     string      `json:"guild"`
		Acquired  string      `json:"acquired"`
		Attacker  interface{} `json:"attacker"`
		Location  struct {
			StartX int `json:"startX"`
			StartY int `json:"startY"`
			EndX   int `json:"endX"`
			EndY   int `json:"endY"`
		} `json:"location"`
	} `json:"Dragonling Nests"`
	MoltenReach struct {
		Territory string      `json:"territory"`
		Guild     string      `json:"guild"`
		Acquired  string      `json:"acquired"`
		Attacker  interface{} `json:"attacker"`
		Location  struct {
			StartX int `json:"startX"`
			StartY int `json:"startY"`
			EndX   int `json:"endX"`
			EndY   int `json:"endY"`
		} `json:"location"`
	} `json:"Molten Reach"`
	SwampIsland struct {
		Territory string      `json:"territory"`
		Guild     string      `json:"guild"`
		Acquired  string      `json:"acquired"`
		Attacker  interface{} `json:"attacker"`
		Location  struct {
			StartX int `json:"startX"`
			StartY int `json:"startY"`
			EndX   int `json:"endX"`
			EndY   int `json:"endY"`
		} `json:"location"`
	} `json:"Swamp Island"`
	WybelIsland struct {
		Territory string      `json:"territory"`
		Guild     string      `json:"guild"`
		Acquired  string      `json:"acquired"`
		Attacker  interface{} `json:"attacker"`
		Location  struct {
			StartX int `json:"startX"`
			StartY int `json:"startY"`
			EndX   int `json:"endX"`
			EndY   int `json:"endY"`
		} `json:"location"`
	} `json:"Wybel Island"`
	SkyIslandAscent struct {
		Territory string      `json:"territory"`
		Guild     string      `json:"guild"`
		Acquired  string      `json:"acquired"`
		Attacker  interface{} `json:"attacker"`
		Location  struct {
			StartX int `json:"startX"`
			StartY int `json:"startY"`
			EndX   int `json:"endX"`
			EndY   int `json:"endY"`
		} `json:"location"`
	} `json:"Sky Island Ascent"`
	RaiderSBaseUpper struct {
		Territory string      `json:"territory"`
		Guild     string      `json:"guild"`
		Acquired  string      `json:"acquired"`
		Attacker  interface{} `json:"attacker"`
		Location  struct {
			StartX int `json:"startX"`
			StartY int `json:"startY"`
			EndX   int `json:"endX"`
			EndY   int `json:"endY"`
		} `json:"location"`
	} `json:"Raider's Base Upper"`
	JofashTunnel struct {
		Territory string      `json:"territory"`
		Guild     string      `json:"guild"`
		Acquired  string      `json:"acquired"`
		Attacker  interface{} `json:"attacker"`
		Location  struct {
			StartX int `json:"startX"`
			StartY int `json:"startY"`
			EndX   int `json:"endX"`
			EndY   int `json:"endY"`
		} `json:"location"`
	} `json:"Jofash Tunnel"`
	SantaSHideout struct {
		Territory string      `json:"territory"`
		Guild     string      `json:"guild"`
		Acquired  string      `json:"acquired"`
		Attacker  interface{} `json:"attacker"`
		Location  struct {
			StartX int `json:"startX"`
			StartY int `json:"startY"`
			EndX   int `json:"endX"`
			EndY   int `json:"endY"`
		} `json:"location"`
	} `json:"Santa's Hideout"`
	IcyDescent struct {
		Territory string      `json:"territory"`
		Guild     string      `json:"guild"`
		Acquired  string      `json:"acquired"`
		Attacker  interface{} `json:"attacker"`
		Location  struct {
			StartX int `json:"startX"`
			StartY int `json:"startY"`
			EndX   int `json:"endX"`
			EndY   int `json:"endY"`
		} `json:"location"`
	} `json:"Icy Descent"`
	AldoreiLowlands struct {
		Territory string      `json:"territory"`
		Guild     string      `json:"guild"`
		Acquired  string      `json:"acquired"`
		Attacker  interface{} `json:"attacker"`
		Location  struct {
			StartX int `json:"startX"`
			StartY int `json:"startY"`
			EndX   int `json:"endX"`
			EndY   int `json:"endY"`
		} `json:"location"`
	} `json:"Aldorei Lowlands"`
	Rodoroc struct {
		Territory string      `json:"territory"`
		Guild     string      `json:"guild"`
		Acquired  string      `json:"acquired"`
		Attacker  interface{} `json:"attacker"`
		Location  struct {
			StartX int `json:"startX"`
			StartY int `json:"startY"`
			EndX   int `json:"endX"`
			EndY   int `json:"endY"`
		} `json:"location"`
	} `json:"Rodoroc"`
	RegularIsland struct {
		Territory string      `json:"territory"`
		Guild     string      `json:"guild"`
		Acquired  string      `json:"acquired"`
		Attacker  interface{} `json:"attacker"`
		Location  struct {
			StartX int `json:"startX"`
			StartY int `json:"startY"`
			EndX   int `json:"endX"`
			EndY   int `json:"endY"`
		} `json:"location"`
	} `json:"Regular Island"`
	TwistedRidge struct {
		Territory string      `json:"territory"`
		Guild     string      `json:"guild"`
		Acquired  string      `json:"acquired"`
		Attacker  interface{} `json:"attacker"`
		Location  struct {
			StartX int `json:"startX"`
			StartY int `json:"startY"`
			EndX   int `json:"endX"`
			EndY   int `json:"endY"`
		} `json:"location"`
	} `json:"Twisted Ridge"`
	RoyalGate struct {
		Territory string      `json:"territory"`
		Guild     string      `json:"guild"`
		Acquired  string      `json:"acquired"`
		Attacker  interface{} `json:"attacker"`
		Location  struct {
			StartX int `json:"startX"`
			StartY int `json:"startY"`
			EndX   int `json:"endX"`
			EndY   int `json:"endY"`
		} `json:"location"`
	} `json:"Royal Gate"`
	OrcBattlegrounds struct {
		Territory string      `json:"territory"`
		Guild     string      `json:"guild"`
		Acquired  string      `json:"acquired"`
		Attacker  interface{} `json:"attacker"`
		Location  struct {
			StartX int `json:"startX"`
			StartY int `json:"startY"`
			EndX   int `json:"endX"`
			EndY   int `json:"endY"`
		} `json:"location"`
	} `json:"Orc Battlegrounds"`
	LostAtoll struct {
		Territory string      `json:"territory"`
		Guild     string      `json:"guild"`
		Acquired  string      `json:"acquired"`
		Attacker  interface{} `json:"attacker"`
		Location  struct {
			StartX int `json:"startX"`
			StartY int `json:"startY"`
			EndX   int `json:"endX"`
			EndY   int `json:"endY"`
		} `json:"location"`
	} `json:"Lost Atoll"`
	LighthousePlateau struct {
		Territory string      `json:"territory"`
		Guild     string      `json:"guild"`
		Acquired  string      `json:"acquired"`
		Attacker  interface{} `json:"attacker"`
		Location  struct {
			StartX int `json:"startX"`
			StartY int `json:"startY"`
			EndX   int `json:"endX"`
			EndY   int `json:"endY"`
		} `json:"location"`
	} `json:"Lighthouse Plateau"`
	TheSilentRoad struct {
		Territory string      `json:"territory"`
		Guild     string      `json:"guild"`
		Acquired  string      `json:"acquired"`
		Attacker  interface{} `json:"attacker"`
		Location  struct {
			StartX int `json:"startX"`
			StartY int `json:"startY"`
			EndX   int `json:"endX"`
			EndY   int `json:"endY"`
		} `json:"location"`
	} `json:"The Silent Road"`
	WormTunnel struct {
		Territory string      `json:"territory"`
		Guild     string      `json:"guild"`
		Acquired  string      `json:"acquired"`
		Attacker  interface{} `json:"attacker"`
		Location  struct {
			StartX int `json:"startX"`
			StartY int `json:"startY"`
			EndX   int `json:"endX"`
			EndY   int `json:"endY"`
		} `json:"location"`
	} `json:"Worm Tunnel"`
	ForgottenTown struct {
		Territory string      `json:"territory"`
		Guild     string      `json:"guild"`
		Acquired  string      `json:"acquired"`
		Attacker  interface{} `json:"attacker"`
		Location  struct {
			StartX int `json:"startX"`
			StartY int `json:"startY"`
			EndX   int `json:"endX"`
			EndY   int `json:"endY"`
		} `json:"location"`
	} `json:"Forgotten Town"`
	SinisterForest struct {
		Territory string      `json:"territory"`
		Guild     string      `json:"guild"`
		Acquired  string      `json:"acquired"`
		Attacker  interface{} `json:"attacker"`
		Location  struct {
			StartX int `json:"startX"`
			StartY int `json:"startY"`
			EndX   int `json:"endX"`
			EndY   int `json:"endY"`
		} `json:"location"`
	} `json:"Sinister Forest"`
	PathsOfSludge struct {
		Territory string      `json:"territory"`
		Guild     string      `json:"guild"`
		Acquired  string      `json:"acquired"`
		Attacker  interface{} `json:"attacker"`
		Location  struct {
			StartX int `json:"startX"`
			StartY int `json:"startY"`
			EndX   int `json:"endX"`
			EndY   int `json:"endY"`
		} `json:"location"`
	} `json:"Paths of Sludge"`
	ToxicCaves struct {
		Territory string      `json:"territory"`
		Guild     string      `json:"guild"`
		Acquired  string      `json:"acquired"`
		Attacker  interface{} `json:"attacker"`
		Location  struct {
			StartX int `json:"startX"`
			StartY int `json:"startY"`
			EndX   int `json:"endX"`
			EndY   int `json:"endY"`
		} `json:"location"`
	} `json:"Toxic Caves"`
	VoidValley struct {
		Territory string      `json:"territory"`
		Guild     string      `json:"guild"`
		Acquired  string      `json:"acquired"`
		Attacker  interface{} `json:"attacker"`
		Location  struct {
			StartX int `json:"startX"`
			StartY int `json:"startY"`
			EndX   int `json:"endX"`
			EndY   int `json:"endY"`
		} `json:"location"`
	} `json:"Void Valley"`
	BizarrePassage struct {
		Territory string      `json:"territory"`
		Guild     string      `json:"guild"`
		Acquired  string      `json:"acquired"`
		Attacker  interface{} `json:"attacker"`
		Location  struct {
			StartX int `json:"startX"`
			StartY int `json:"startY"`
			EndX   int `json:"endX"`
			EndY   int `json:"endY"`
		} `json:"location"`
	} `json:"Bizarre Passage"`
	HeavenlyIngress struct {
		Territory string      `json:"territory"`
		Guild     string      `json:"guild"`
		Acquired  string      `json:"acquired"`
		Attacker  interface{} `json:"attacker"`
		Location  struct {
			StartX int `json:"startX"`
			StartY int `json:"startY"`
			EndX   int `json:"endX"`
			EndY   int `json:"endY"`
		} `json:"location"`
	} `json:"Heavenly Ingress"`
	PrimalFen struct {
		Territory string      `json:"territory"`
		Guild     string      `json:"guild"`
		Acquired  string      `json:"acquired"`
		Attacker  interface{} `json:"attacker"`
		Location  struct {
			StartX int `json:"startX"`
			StartY int `json:"startY"`
			EndX   int `json:"endX"`
			EndY   int `json:"endY"`
		} `json:"location"`
	} `json:"Primal Fen"`
	OtherwordlyMonolith struct {
		Territory string      `json:"territory"`
		Guild     string      `json:"guild"`
		Acquired  string      `json:"acquired"`
		Attacker  interface{} `json:"attacker"`
		Location  struct {
			StartX int `json:"startX"`
			StartY int `json:"startY"`
			EndX   int `json:"endX"`
			EndY   int `json:"endY"`
		} `json:"location"`
	} `json:"Otherwordly Monolith"`
	NexusOfLight struct {
		Territory string      `json:"territory"`
		Guild     string      `json:"guild"`
		Acquired  string      `json:"acquired"`
		Attacker  interface{} `json:"attacker"`
		Location  struct {
			StartX int `json:"startX"`
			StartY int `json:"startY"`
			EndX   int `json:"endX"`
			EndY   int `json:"endY"`
		} `json:"location"`
	} `json:"Nexus of Light"`
	FieldOfLife struct {
		Territory string      `json:"territory"`
		Guild     string      `json:"guild"`
		Acquired  string      `json:"acquired"`
		Attacker  interface{} `json:"attacker"`
		Location  struct {
			StartX int `json:"startX"`
			StartY int `json:"startY"`
			EndX   int `json:"endX"`
			EndY   int `json:"endY"`
		} `json:"location"`
	} `json:"Field of Life"`
	Ragni struct {
		Territory string      `json:"territory"`
		Guild     string      `json:"guild"`
		Acquired  string      `json:"acquired"`
		Attacker  interface{} `json:"attacker"`
		Location  struct {
			StartX int `json:"startX"`
			StartY int `json:"startY"`
			EndX   int `json:"endX"`
			EndY   int `json:"endY"`
		} `json:"location"`
	} `json:"Ragni"`
	RagniNorthEntrance struct {
		Territory string      `json:"territory"`
		Guild     string      `json:"guild"`
		Acquired  string      `json:"acquired"`
		Attacker  interface{} `json:"attacker"`
		Location  struct {
			StartX int `json:"startX"`
			StartY int `json:"startY"`
			EndX   int `json:"endX"`
			EndY   int `json:"endY"`
		} `json:"location"`
	} `json:"Ragni North Entrance"`
	RagniPlains struct {
		Territory string      `json:"territory"`
		Guild     string      `json:"guild"`
		Acquired  string      `json:"acquired"`
		Attacker  interface{} `json:"attacker"`
		Location  struct {
			StartX int `json:"startX"`
			StartY int `json:"startY"`
			EndX   int `json:"endX"`
			EndY   int `json:"endY"`
		} `json:"location"`
	} `json:"Ragni Plains"`
	MalticPlains struct {
		Territory string      `json:"territory"`
		Guild     string      `json:"guild"`
		Acquired  string      `json:"acquired"`
		Attacker  interface{} `json:"attacker"`
		Location  struct {
			StartX int `json:"startX"`
			StartY int `json:"startY"`
			EndX   int `json:"endX"`
			EndY   int `json:"endY"`
		} `json:"location"`
	} `json:"Maltic Plains"`
	SouthPigmenRavines struct {
		Territory string      `json:"territory"`
		Guild     string      `json:"guild"`
		Acquired  string      `json:"acquired"`
		Attacker  interface{} `json:"attacker"`
		Location  struct {
			StartX int `json:"startX"`
			StartY int `json:"startY"`
			EndX   int `json:"endX"`
			EndY   int `json:"endY"`
		} `json:"location"`
	} `json:"South Pigmen Ravines"`
	SanctuaryBridge struct {
		Territory string      `json:"territory"`
		Guild     string      `json:"guild"`
		Acquired  string      `json:"acquired"`
		Attacker  interface{} `json:"attacker"`
		Location  struct {
			StartX int `json:"startX"`
			StartY int `json:"startY"`
			EndX   int `json:"endX"`
			EndY   int `json:"endY"`
		} `json:"location"`
	} `json:"Sanctuary Bridge"`
	NivlaWoods struct {
		Territory string      `json:"territory"`
		Guild     string      `json:"guild"`
		Acquired  string      `json:"acquired"`
		Attacker  interface{} `json:"attacker"`
		Location  struct {
			StartX int `json:"startX"`
			StartY int `json:"startY"`
			EndX   int `json:"endX"`
			EndY   int `json:"endY"`
		} `json:"location"`
	} `json:"Nivla Woods"`
	Elkurn struct {
		Territory string      `json:"territory"`
		Guild     string      `json:"guild"`
		Acquired  string      `json:"acquired"`
		Attacker  interface{} `json:"attacker"`
		Location  struct {
			StartX int `json:"startX"`
			StartY int `json:"startY"`
			EndX   int `json:"endX"`
			EndY   int `json:"endY"`
		} `json:"location"`
	} `json:"Elkurn"`
	DetlasFarSuburbs struct {
		Territory string      `json:"territory"`
		Guild     string      `json:"guild"`
		Acquired  string      `json:"acquired"`
		Attacker  interface{} `json:"attacker"`
		Location  struct {
			StartX int `json:"startX"`
			StartY int `json:"startY"`
			EndX   int `json:"endX"`
			EndY   int `json:"endY"`
		} `json:"location"`
	} `json:"Detlas Far Suburbs"`
	SouthFarmersValley struct {
		Territory string      `json:"territory"`
		Guild     string      `json:"guild"`
		Acquired  string      `json:"acquired"`
		Attacker  interface{} `json:"attacker"`
		Location  struct {
			StartX int `json:"startX"`
			StartY int `json:"startY"`
			EndX   int `json:"endX"`
			EndY   int `json:"endY"`
		} `json:"location"`
	} `json:"South Farmers Valley"`
	TowerOfAscension struct {
		Territory string      `json:"territory"`
		Guild     string      `json:"guild"`
		Acquired  string      `json:"acquired"`
		Attacker  interface{} `json:"attacker"`
		Location  struct {
			StartX int `json:"startX"`
			StartY int `json:"startY"`
			EndX   int `json:"endX"`
			EndY   int `json:"endY"`
		} `json:"location"`
	} `json:"Tower of Ascension"`
	TwainMansion struct {
		Territory string      `json:"territory"`
		Guild     string      `json:"guild"`
		Acquired  string      `json:"acquired"`
		Attacker  interface{} `json:"attacker"`
		Location  struct {
			StartX int `json:"startX"`
			StartY int `json:"startY"`
			EndX   int `json:"endX"`
			EndY   int `json:"endY"`
		} `json:"location"`
	} `json:"Twain Mansion"`
	NesaakPlainsNorthEast struct {
		Territory string      `json:"territory"`
		Guild     string      `json:"guild"`
		Acquired  string      `json:"acquired"`
		Attacker  interface{} `json:"attacker"`
		Location  struct {
			StartX int `json:"startX"`
			StartY int `json:"startY"`
			EndX   int `json:"endX"`
			EndY   int `json:"endY"`
		} `json:"location"`
	} `json:"Nesaak Plains North East"`
	NesaakBridgeTransition struct {
		Territory string      `json:"territory"`
		Guild     string      `json:"guild"`
		Acquired  string      `json:"acquired"`
		Attacker  interface{} `json:"attacker"`
		Location  struct {
			StartX int `json:"startX"`
			StartY int `json:"startY"`
			EndX   int `json:"endX"`
			EndY   int `json:"endY"`
		} `json:"location"`
	} `json:"Nesaak Bridge Transition"`
	JungleLower struct {
		Territory string      `json:"territory"`
		Guild     string      `json:"guild"`
		Acquired  string      `json:"acquired"`
		Attacker  interface{} `json:"attacker"`
		Location  struct {
			StartX int `json:"startX"`
			StartY int `json:"startY"`
			EndX   int `json:"endX"`
			EndY   int `json:"endY"`
		} `json:"location"`
	} `json:"Jungle Lower"`
	TempleOfLegends struct {
		Territory string      `json:"territory"`
		Guild     string      `json:"guild"`
		Acquired  string      `json:"acquired"`
		Attacker  interface{} `json:"attacker"`
		Location  struct {
			StartX int `json:"startX"`
			StartY int `json:"startY"`
			EndX   int `json:"endX"`
			EndY   int `json:"endY"`
		} `json:"location"`
	} `json:"Temple of Legends"`
	RymekEastUpper struct {
		Territory string      `json:"territory"`
		Guild     string      `json:"guild"`
		Acquired  string      `json:"acquired"`
		Attacker  interface{} `json:"attacker"`
		Location  struct {
			StartX int `json:"startX"`
			StartY int `json:"startY"`
			EndX   int `json:"endX"`
			EndY   int `json:"endY"`
		} `json:"location"`
	} `json:"Rymek East Upper"`
	DesertEastUpper struct {
		Territory string      `json:"territory"`
		Guild     string      `json:"guild"`
		Acquired  string      `json:"acquired"`
		Attacker  interface{} `json:"attacker"`
		Location  struct {
			StartX int `json:"startX"`
			StartY int `json:"startY"`
			EndX   int `json:"endX"`
			EndY   int `json:"endY"`
		} `json:"location"`
	} `json:"Desert East Upper"`
	DesertMidUpper struct {
		Territory string      `json:"territory"`
		Guild     string      `json:"guild"`
		Acquired  string      `json:"acquired"`
		Attacker  interface{} `json:"attacker"`
		Location  struct {
			StartX int `json:"startX"`
			StartY int `json:"startY"`
			EndX   int `json:"endX"`
			EndY   int `json:"endY"`
		} `json:"location"`
	} `json:"Desert Mid-Upper"`
	MummySTomb struct {
		Territory string      `json:"territory"`
		Guild     string      `json:"guild"`
		Acquired  string      `json:"acquired"`
		Attacker  interface{} `json:"attacker"`
		Location  struct {
			StartX int `json:"startX"`
			StartY int `json:"startY"`
			EndX   int `json:"endX"`
			EndY   int `json:"endY"`
		} `json:"location"`
	} `json:"Mummy's Tomb"`
	SavannahEastUpper struct {
		Territory string      `json:"territory"`
		Guild     string      `json:"guild"`
		Acquired  string      `json:"acquired"`
		Attacker  interface{} `json:"attacker"`
		Location  struct {
			StartX int `json:"startX"`
			StartY int `json:"startY"`
			EndX   int `json:"endX"`
			EndY   int `json:"endY"`
		} `json:"location"`
	} `json:"Savannah East Upper"`
	LionLair struct {
		Territory string      `json:"territory"`
		Guild     string      `json:"guild"`
		Acquired  string      `json:"acquired"`
		Attacker  interface{} `json:"attacker"`
		Location  struct {
			StartX int `json:"startX"`
			StartY int `json:"startY"`
			EndX   int `json:"endX"`
			EndY   int `json:"endY"`
		} `json:"location"`
	} `json:"Lion Lair"`
	NemractPlainsWest struct {
		Territory string      `json:"territory"`
		Guild     string      `json:"guild"`
		Acquired  string      `json:"acquired"`
		Attacker  interface{} `json:"attacker"`
		Location  struct {
			StartX int `json:"startX"`
			StartY int `json:"startY"`
			EndX   int `json:"endX"`
			EndY   int `json:"endY"`
		} `json:"location"`
	} `json:"Nemract Plains West"`
	CathedralHarbour struct {
		Territory string      `json:"territory"`
		Guild     string      `json:"guild"`
		Acquired  string      `json:"acquired"`
		Attacker  interface{} `json:"attacker"`
		Location  struct {
			StartX int `json:"startX"`
			StartY int `json:"startY"`
			EndX   int `json:"endX"`
			EndY   int `json:"endY"`
		} `json:"location"`
	} `json:"Cathedral Harbour"`
	Selchar struct {
		Territory string      `json:"territory"`
		Guild     string      `json:"guild"`
		Acquired  string      `json:"acquired"`
		Attacker  interface{} `json:"attacker"`
		Location  struct {
			StartX int `json:"startX"`
			StartY int `json:"startY"`
			EndX   int `json:"endX"`
			EndY   int `json:"endY"`
		} `json:"location"`
	} `json:"Selchar"`
	DurumIslesLower struct {
		Territory string      `json:"territory"`
		Guild     string      `json:"guild"`
		Acquired  string      `json:"acquired"`
		Attacker  interface{} `json:"attacker"`
		Location  struct {
			StartX int `json:"startX"`
			StartY int `json:"startY"`
			EndX   int `json:"endX"`
			EndY   int `json:"endY"`
		} `json:"location"`
	} `json:"Durum Isles Lower"`
	NodgujNation struct {
		Territory string      `json:"territory"`
		Guild     string      `json:"guild"`
		Acquired  string      `json:"acquired"`
		Attacker  interface{} `json:"attacker"`
		Location  struct {
			StartX int `json:"startX"`
			StartY int `json:"startY"`
			EndX   int `json:"endX"`
			EndY   int `json:"endY"`
		} `json:"location"`
	} `json:"Nodguj Nation"`
	DeadIslandSouthWest struct {
		Territory string      `json:"territory"`
		Guild     string      `json:"guild"`
		Acquired  string      `json:"acquired"`
		Attacker  interface{} `json:"attacker"`
		Location  struct {
			StartX int `json:"startX"`
			StartY int `json:"startY"`
			EndX   int `json:"endX"`
			EndY   int `json:"endY"`
		} `json:"location"`
	} `json:"Dead Island South West"`
	TreeIsland struct {
		Territory string      `json:"territory"`
		Guild     string      `json:"guild"`
		Acquired  string      `json:"acquired"`
		Attacker  interface{} `json:"attacker"`
		Location  struct {
			StartX int `json:"startX"`
			StartY int `json:"startY"`
			EndX   int `json:"endX"`
			EndY   int `json:"endY"`
		} `json:"location"`
	} `json:"Tree Island"`
	MiningBaseUpper struct {
		Territory string      `json:"territory"`
		Guild     string      `json:"guild"`
		Acquired  string      `json:"acquired"`
		Attacker  interface{} `json:"attacker"`
		Location  struct {
			StartX int `json:"startX"`
			StartY int `json:"startY"`
			EndX   int `json:"endX"`
			EndY   int `json:"endY"`
		} `json:"location"`
	} `json:"Mining Base Upper"`
	NetherPlainsLower struct {
		Territory string      `json:"territory"`
		Guild     string      `json:"guild"`
		Acquired  string      `json:"acquired"`
		Attacker  interface{} `json:"attacker"`
		Location  struct {
			StartX int `json:"startX"`
			StartY int `json:"startY"`
			EndX   int `json:"endX"`
			EndY   int `json:"endY"`
		} `json:"location"`
	} `json:"Nether Plains Lower"`
	NetherPlainsUpper struct {
		Territory string      `json:"territory"`
		Guild     string      `json:"guild"`
		Acquired  string      `json:"acquired"`
		Attacker  interface{} `json:"attacker"`
		Location  struct {
			StartX int `json:"startX"`
			StartY int `json:"startY"`
			EndX   int `json:"endX"`
			EndY   int `json:"endY"`
		} `json:"location"`
	} `json:"Nether Plains Upper"`
	LlevigarGateEast struct {
		Territory string      `json:"territory"`
		Guild     string      `json:"guild"`
		Acquired  string      `json:"acquired"`
		Attacker  interface{} `json:"attacker"`
		Location  struct {
			StartX int `json:"startX"`
			StartY int `json:"startY"`
			EndX   int `json:"endX"`
			EndY   int `json:"endY"`
		} `json:"location"`
	} `json:"Llevigar Gate East"`
	Hive struct {
		Territory string      `json:"territory"`
		Guild     string      `json:"guild"`
		Acquired  string      `json:"acquired"`
		Attacker  interface{} `json:"attacker"`
		Location  struct {
			StartX int `json:"startX"`
			StartY int `json:"startY"`
			EndX   int `json:"endX"`
			EndY   int `json:"endY"`
		} `json:"location"`
	} `json:"Hive"`
	LlevigarPlainsWestUpper struct {
		Territory string      `json:"territory"`
		Guild     string      `json:"guild"`
		Acquired  string      `json:"acquired"`
		Attacker  interface{} `json:"attacker"`
		Location  struct {
			StartX int `json:"startX"`
			StartY int `json:"startY"`
			EndX   int `json:"endX"`
			EndY   int `json:"endY"`
		} `json:"location"`
	} `json:"Llevigar Plains West Upper"`
	SwampEastMid struct {
		Territory string      `json:"territory"`
		Guild     string      `json:"guild"`
		Acquired  string      `json:"acquired"`
		Attacker  interface{} `json:"attacker"`
		Location  struct {
			StartX int `json:"startX"`
			StartY int `json:"startY"`
			EndX   int `json:"endX"`
			EndY   int `json:"endY"`
		} `json:"location"`
	} `json:"Swamp East Mid"`
	SwampWestUpper struct {
		Territory string      `json:"territory"`
		Guild     string      `json:"guild"`
		Acquired  string      `json:"acquired"`
		Attacker  interface{} `json:"attacker"`
		Location  struct {
			StartX int `json:"startX"`
			StartY int `json:"startY"`
			EndX   int `json:"endX"`
			EndY   int `json:"endY"`
		} `json:"location"`
	} `json:"Swamp West Upper"`
	SwampDarkForestTransitionUpper struct {
		Territory string      `json:"territory"`
		Guild     string      `json:"guild"`
		Acquired  string      `json:"acquired"`
		Attacker  interface{} `json:"attacker"`
		Location  struct {
			StartX int `json:"startX"`
			StartY int `json:"startY"`
			EndX   int `json:"endX"`
			EndY   int `json:"endY"`
		} `json:"location"`
	} `json:"Swamp Dark Forest Transition Upper"`
	SwampMountainBase struct {
		Territory string      `json:"territory"`
		Guild     string      `json:"guild"`
		Acquired  string      `json:"acquired"`
		Attacker  interface{} `json:"attacker"`
		Location  struct {
			StartX int `json:"startX"`
			StartY int `json:"startY"`
			EndX   int `json:"endX"`
			EndY   int `json:"endY"`
		} `json:"location"`
	} `json:"Swamp Mountain Base"`
	SwampMountainTransitionMidUpper struct {
		Territory string      `json:"territory"`
		Guild     string      `json:"guild"`
		Acquired  string      `json:"acquired"`
		Attacker  interface{} `json:"attacker"`
		Location  struct {
			StartX int `json:"startX"`
			StartY int `json:"startY"`
			EndX   int `json:"endX"`
			EndY   int `json:"endY"`
		} `json:"location"`
	} `json:"Swamp Mountain Transition Mid-Upper"`
	QuartzMinesNorthWest struct {
		Territory string      `json:"territory"`
		Guild     string      `json:"guild"`
		Acquired  string      `json:"acquired"`
		Attacker  interface{} `json:"attacker"`
		Location  struct {
			StartX int `json:"startX"`
			StartY int `json:"startY"`
			EndX   int `json:"endX"`
			EndY   int `json:"endY"`
		} `json:"location"`
	} `json:"Quartz Mines North West"`
	OrcRoad struct {
		Territory string      `json:"territory"`
		Guild     string      `json:"guild"`
		Acquired  string      `json:"acquired"`
		Attacker  interface{} `json:"attacker"`
		Location  struct {
			StartX int `json:"startX"`
			StartY int `json:"startY"`
			EndX   int `json:"endX"`
			EndY   int `json:"endY"`
		} `json:"location"`
	} `json:"Orc Road"`
	IronRoad struct {
		Territory string      `json:"territory"`
		Guild     string      `json:"guild"`
		Acquired  string      `json:"acquired"`
		Attacker  interface{} `json:"attacker"`
		Location  struct {
			StartX int `json:"startX"`
			StartY int `json:"startY"`
			EndX   int `json:"endX"`
			EndY   int `json:"endY"`
		} `json:"location"`
	} `json:"Iron Road"`
	GoblinPlainsEast struct {
		Territory string      `json:"territory"`
		Guild     string      `json:"guild"`
		Acquired  string      `json:"acquired"`
		Attacker  interface{} `json:"attacker"`
		Location  struct {
			StartX int `json:"startX"`
			StartY int `json:"startY"`
			EndX   int `json:"endX"`
			EndY   int `json:"endY"`
		} `json:"location"`
	} `json:"Goblin Plains East"`
	EfilimVillage struct {
		Territory string      `json:"territory"`
		Guild     string      `json:"guild"`
		Acquired  string      `json:"acquired"`
		Attacker  interface{} `json:"attacker"`
		Location  struct {
			StartX int `json:"startX"`
			StartY int `json:"startY"`
			EndX   int `json:"endX"`
			EndY   int `json:"endY"`
		} `json:"location"`
	} `json:"Efilim Village"`
	LightForestNorthEntrance struct {
		Territory string      `json:"territory"`
		Guild     string      `json:"guild"`
		Acquired  string      `json:"acquired"`
		Attacker  interface{} `json:"attacker"`
		Location  struct {
			StartX int `json:"startX"`
			StartY int `json:"startY"`
			EndX   int `json:"endX"`
			EndY   int `json:"endY"`
		} `json:"location"`
	} `json:"Light Forest North Entrance"`
	LightForestSouthExit struct {
		Territory string      `json:"territory"`
		Guild     string      `json:"guild"`
		Acquired  string      `json:"acquired"`
		Attacker  interface{} `json:"attacker"`
		Location  struct {
			StartX int `json:"startX"`
			StartY int `json:"startY"`
			EndX   int `json:"endX"`
			EndY   int `json:"endY"`
		} `json:"location"`
	} `json:"Light Forest South Exit"`
	LightForestWestUpper struct {
		Territory string      `json:"territory"`
		Guild     string      `json:"guild"`
		Acquired  string      `json:"acquired"`
		Attacker  interface{} `json:"attacker"`
		Location  struct {
			StartX int `json:"startX"`
			StartY int `json:"startY"`
			EndX   int `json:"endX"`
			EndY   int `json:"endY"`
		} `json:"location"`
	} `json:"Light Forest West Upper"`
	HobbitRiver struct {
		Territory string      `json:"territory"`
		Guild     string      `json:"guild"`
		Acquired  string      `json:"acquired"`
		Attacker  interface{} `json:"attacker"`
		Location  struct {
			StartX int `json:"startX"`
			StartY int `json:"startY"`
			EndX   int `json:"endX"`
			EndY   int `json:"endY"`
		} `json:"location"`
	} `json:"Hobbit River"`
	LoneFarmstead struct {
		Territory string      `json:"territory"`
		Guild     string      `json:"guild"`
		Acquired  string      `json:"acquired"`
		Attacker  interface{} `json:"attacker"`
		Location  struct {
			StartX int `json:"startX"`
			StartY int `json:"startY"`
			EndX   int `json:"endX"`
			EndY   int `json:"endY"`
		} `json:"location"`
	} `json:"Lone Farmstead"`
	TaprootDescent struct {
		Territory string      `json:"territory"`
		Guild     string      `json:"guild"`
		Acquired  string      `json:"acquired"`
		Attacker  interface{} `json:"attacker"`
		Location  struct {
			StartX int `json:"startX"`
			StartY int `json:"startY"`
			EndX   int `json:"endX"`
			EndY   int `json:"endY"`
		} `json:"location"`
	} `json:"Taproot Descent"`
	TwistedHousing struct {
		Territory string      `json:"territory"`
		Guild     string      `json:"guild"`
		Acquired  string      `json:"acquired"`
		Attacker  interface{} `json:"attacker"`
		Location  struct {
			StartX int `json:"startX"`
			StartY int `json:"startY"`
			EndX   int `json:"endX"`
			EndY   int `json:"endY"`
		} `json:"location"`
	} `json:"Twisted Housing"`
	AbandonedManor struct {
		Territory string      `json:"territory"`
		Guild     string      `json:"guild"`
		Acquired  string      `json:"acquired"`
		Attacker  interface{} `json:"attacker"`
		Location  struct {
			StartX int `json:"startX"`
			StartY int `json:"startY"`
			EndX   int `json:"endX"`
			EndY   int `json:"endY"`
		} `json:"location"`
	} `json:"Abandoned Manor"`
	VisceraPitsEast struct {
		Territory string      `json:"territory"`
		Guild     string      `json:"guild"`
		Acquired  string      `json:"acquired"`
		Attacker  interface{} `json:"attacker"`
		Location  struct {
			StartX int `json:"startX"`
			StartY int `json:"startY"`
			EndX   int `json:"endX"`
			EndY   int `json:"endY"`
		} `json:"location"`
	} `json:"Viscera Pits East"`
	Lexdale struct {
		Territory string      `json:"territory"`
		Guild     string      `json:"guild"`
		Acquired  string      `json:"acquired"`
		Attacker  interface{} `json:"attacker"`
		Location  struct {
			StartX int `json:"startX"`
			StartY int `json:"startY"`
			EndX   int `json:"endX"`
			EndY   int `json:"endY"`
		} `json:"location"`
	} `json:"Lexdale"`
	CinfrasEntrance struct {
		Territory string      `json:"territory"`
		Guild     string      `json:"guild"`
		Acquired  string      `json:"acquired"`
		Attacker  interface{} `json:"attacker"`
		Location  struct {
			StartX int `json:"startX"`
			StartY int `json:"startY"`
			EndX   int `json:"endX"`
			EndY   int `json:"endY"`
		} `json:"location"`
	} `json:"Cinfras Entrance"`
	GuildHall struct {
		Territory string      `json:"territory"`
		Guild     string      `json:"guild"`
		Acquired  string      `json:"acquired"`
		Attacker  interface{} `json:"attacker"`
		Location  struct {
			StartX int `json:"startX"`
			StartY int `json:"startY"`
			EndX   int `json:"endX"`
			EndY   int `json:"endY"`
		} `json:"location"`
	} `json:"Guild Hall"`
	CinfrasCountyMidLower struct {
		Territory string      `json:"territory"`
		Guild     string      `json:"guild"`
		Acquired  string      `json:"acquired"`
		Attacker  interface{} `json:"attacker"`
		Location  struct {
			StartX int `json:"startX"`
			StartY int `json:"startY"`
			EndX   int `json:"endX"`
			EndY   int `json:"endY"`
		} `json:"location"`
	} `json:"Cinfras County Mid-Lower"`
	GyliaLakeSouthWest struct {
		Territory string      `json:"territory"`
		Guild     string      `json:"guild"`
		Acquired  string      `json:"acquired"`
		Attacker  interface{} `json:"attacker"`
		Location  struct {
			StartX int `json:"startX"`
			StartY int `json:"startY"`
			EndX   int `json:"endX"`
			EndY   int `json:"endY"`
		} `json:"location"`
	} `json:"Gylia Lake South West"`
	JitakSFarm struct {
		Territory string      `json:"territory"`
		Guild     string      `json:"guild"`
		Acquired  string      `json:"acquired"`
		Attacker  interface{} `json:"attacker"`
		Location  struct {
			StartX int `json:"startX"`
			StartY int `json:"startY"`
			EndX   int `json:"endX"`
			EndY   int `json:"endY"`
		} `json:"location"`
	} `json:"Jitak's Farm"`
	AldoreiSRiver struct {
		Territory string      `json:"territory"`
		Guild     string      `json:"guild"`
		Acquired  string      `json:"acquired"`
		Attacker  interface{} `json:"attacker"`
		Location  struct {
			StartX int `json:"startX"`
			StartY int `json:"startY"`
			EndX   int `json:"endX"`
			EndY   int `json:"endY"`
		} `json:"location"`
	} `json:"Aldorei's River"`
	PathToTheArch struct {
		Territory string      `json:"territory"`
		Guild     string      `json:"guild"`
		Acquired  string      `json:"acquired"`
		Attacker  interface{} `json:"attacker"`
		Location  struct {
			StartX int `json:"startX"`
			StartY int `json:"startY"`
			EndX   int `json:"endX"`
			EndY   int `json:"endY"`
		} `json:"location"`
	} `json:"Path To The Arch"`
	CinfrasThanosTransition struct {
		Territory string      `json:"territory"`
		Guild     string      `json:"guild"`
		Acquired  string      `json:"acquired"`
		Attacker  interface{} `json:"attacker"`
		Location  struct {
			StartX int `json:"startX"`
			StartY int `json:"startY"`
			EndX   int `json:"endX"`
			EndY   int `json:"endY"`
		} `json:"location"`
	} `json:"Cinfras Thanos Transition"`
	MilitaryBase struct {
		Territory string      `json:"territory"`
		Guild     string      `json:"guild"`
		Acquired  string      `json:"acquired"`
		Attacker  interface{} `json:"attacker"`
		Location  struct {
			StartX int `json:"startX"`
			StartY int `json:"startY"`
			EndX   int `json:"endX"`
			EndY   int `json:"endY"`
		} `json:"location"`
	} `json:"Military Base"`
	PathToOzothSSpireMid struct {
		Territory string      `json:"territory"`
		Guild     string      `json:"guild"`
		Acquired  string      `json:"acquired"`
		Attacker  interface{} `json:"attacker"`
		Location  struct {
			StartX int `json:"startX"`
			StartY int `json:"startY"`
			EndX   int `json:"endX"`
			EndY   int `json:"endY"`
		} `json:"location"`
	} `json:"Path To Ozoth's Spire Mid"`
	CanyonEntranceWaterfall struct {
		Territory string      `json:"territory"`
		Guild     string      `json:"guild"`
		Acquired  string      `json:"acquired"`
		Attacker  interface{} `json:"attacker"`
		Location  struct {
			StartX int `json:"startX"`
			StartY int `json:"startY"`
			EndX   int `json:"endX"`
			EndY   int `json:"endY"`
		} `json:"location"`
	} `json:"Canyon Entrance Waterfall"`
	CanyonUpperNorthWest struct {
		Territory string      `json:"territory"`
		Guild     string      `json:"guild"`
		Acquired  string      `json:"acquired"`
		Attacker  interface{} `json:"attacker"`
		Location  struct {
			StartX int `json:"startX"`
			StartY int `json:"startY"`
			EndX   int `json:"endX"`
			EndY   int `json:"endY"`
		} `json:"location"`
	} `json:"Canyon Upper North West"`
	BanditCampExit struct {
		Territory string      `json:"territory"`
		Guild     string      `json:"guild"`
		Acquired  string      `json:"acquired"`
		Attacker  interface{} `json:"attacker"`
		Location  struct {
			StartX int `json:"startX"`
			StartY int `json:"startY"`
			EndX   int `json:"endX"`
			EndY   int `json:"endY"`
		} `json:"location"`
	} `json:"Bandit Camp Exit"`
	CanyonWalkWay struct {
		Territory string      `json:"territory"`
		Guild     string      `json:"guild"`
		Acquired  string      `json:"acquired"`
		Attacker  interface{} `json:"attacker"`
		Location  struct {
			StartX int `json:"startX"`
			StartY int `json:"startY"`
			EndX   int `json:"endX"`
			EndY   int `json:"endY"`
		} `json:"location"`
	} `json:"Canyon Walk Way"`
	CanyonFortress struct {
		Territory string      `json:"territory"`
		Guild     string      `json:"guild"`
		Acquired  string      `json:"acquired"`
		Attacker  interface{} `json:"attacker"`
		Location  struct {
			StartX int `json:"startX"`
			StartY int `json:"startY"`
			EndX   int `json:"endX"`
			EndY   int `json:"endY"`
		} `json:"location"`
	} `json:"Canyon Fortress"`
	BanditsToll struct {
		Territory string      `json:"territory"`
		Guild     string      `json:"guild"`
		Acquired  string      `json:"acquired"`
		Attacker  interface{} `json:"attacker"`
		Location  struct {
			StartX int `json:"startX"`
			StartY int `json:"startY"`
			EndX   int `json:"endX"`
			EndY   int `json:"endY"`
		} `json:"location"`
	} `json:"Bandits Toll"`
	CliffSideOfTheLost struct {
		Territory string      `json:"territory"`
		Guild     string      `json:"guild"`
		Acquired  string      `json:"acquired"`
		Attacker  interface{} `json:"attacker"`
		Location  struct {
			StartX int `json:"startX"`
			StartY int `json:"startY"`
			EndX   int `json:"endX"`
			EndY   int `json:"endY"`
		} `json:"location"`
	} `json:"Cliff Side of the Lost"`
	HiveSouth struct {
		Territory string      `json:"territory"`
		Guild     string      `json:"guild"`
		Acquired  string      `json:"acquired"`
		Attacker  interface{} `json:"attacker"`
		Location  struct {
			StartX int `json:"startX"`
			StartY int `json:"startY"`
			EndX   int `json:"endX"`
			EndY   int `json:"endY"`
		} `json:"location"`
	} `json:"Hive South"`
	AirTempleLower struct {
		Territory string      `json:"territory"`
		Guild     string      `json:"guild"`
		Acquired  string      `json:"acquired"`
		Attacker  interface{} `json:"attacker"`
		Location  struct {
			StartX int `json:"startX"`
			StartY int `json:"startY"`
			EndX   int `json:"endX"`
			EndY   int `json:"endY"`
		} `json:"location"`
	} `json:"Air Temple Lower"`
	KandonBeda struct {
		Territory string      `json:"territory"`
		Guild     string      `json:"guild"`
		Acquired  string      `json:"acquired"`
		Attacker  interface{} `json:"attacker"`
		Location  struct {
			StartX int `json:"startX"`
			StartY int `json:"startY"`
			EndX   int `json:"endX"`
			EndY   int `json:"endY"`
		} `json:"location"`
	} `json:"Kandon-Beda"`
	EntranceToTheseadNorth struct {
		Territory string      `json:"territory"`
		Guild     string      `json:"guild"`
		Acquired  string      `json:"acquired"`
		Attacker  interface{} `json:"attacker"`
		Location  struct {
			StartX int `json:"startX"`
			StartY int `json:"startY"`
			EndX   int `json:"endX"`
			EndY   int `json:"endY"`
		} `json:"location"`
	} `json:"Entrance to Thesead North"`
	RanolSFarm struct {
		Territory string      `json:"territory"`
		Guild     string      `json:"guild"`
		Acquired  string      `json:"acquired"`
		Attacker  interface{} `json:"attacker"`
		Location  struct {
			StartX int `json:"startX"`
			StartY int `json:"startY"`
			EndX   int `json:"endX"`
			EndY   int `json:"endY"`
		} `json:"location"`
	} `json:"Ranol's Farm"`
	Eltom struct {
		Territory string      `json:"territory"`
		Guild     string      `json:"guild"`
		Acquired  string      `json:"acquired"`
		Attacker  interface{} `json:"attacker"`
		Location  struct {
			StartX int `json:"startX"`
			StartY int `json:"startY"`
			EndX   int `json:"endX"`
			EndY   int `json:"endY"`
		} `json:"location"`
	} `json:"Eltom"`
	CraterDescent struct {
		Territory string      `json:"territory"`
		Guild     string      `json:"guild"`
		Acquired  string      `json:"acquired"`
		Attacker  interface{} `json:"attacker"`
		Location  struct {
			StartX int `json:"startX"`
			StartY int `json:"startY"`
			EndX   int `json:"endX"`
			EndY   int `json:"endY"`
		} `json:"location"`
	} `json:"Crater Descent"`
	TempleIsland struct {
		Territory string      `json:"territory"`
		Guild     string      `json:"guild"`
		Acquired  string      `json:"acquired"`
		Attacker  interface{} `json:"attacker"`
		Location  struct {
			StartX int `json:"startX"`
			StartY int `json:"startY"`
			EndX   int `json:"endX"`
			EndY   int `json:"endY"`
		} `json:"location"`
	} `json:"Temple Island"`
	DernelJungleUpper struct {
		Territory string      `json:"territory"`
		Guild     string      `json:"guild"`
		Acquired  string      `json:"acquired"`
		Attacker  interface{} `json:"attacker"`
		Location  struct {
			StartX int `json:"startX"`
			StartY int `json:"startY"`
			EndX   int `json:"endX"`
			EndY   int `json:"endY"`
		} `json:"location"`
	} `json:"Dernel Jungle Upper"`
	FallenFactory struct {
		Territory string      `json:"territory"`
		Guild     string      `json:"guild"`
		Acquired  string      `json:"acquired"`
		Attacker  interface{} `json:"attacker"`
		Location  struct {
			StartX int `json:"startX"`
			StartY int `json:"startY"`
			EndX   int `json:"endX"`
			EndY   int `json:"endY"`
		} `json:"location"`
	} `json:"Fallen Factory"`
	FactoryEntrance struct {
		Territory string      `json:"territory"`
		Guild     string      `json:"guild"`
		Acquired  string      `json:"acquired"`
		Attacker  interface{} `json:"attacker"`
		Location  struct {
			StartX int `json:"startX"`
			StartY int `json:"startY"`
			EndX   int `json:"endX"`
			EndY   int `json:"endY"`
		} `json:"location"`
	} `json:"Factory Entrance"`
	AvosWorkshop struct {
		Territory string      `json:"territory"`
		Guild     string      `json:"guild"`
		Acquired  string      `json:"acquired"`
		Attacker  interface{} `json:"attacker"`
		Location  struct {
			StartX int `json:"startX"`
			StartY int `json:"startY"`
			EndX   int `json:"endX"`
			EndY   int `json:"endY"`
		} `json:"location"`
	} `json:"Avos Workshop"`
	RuinedHouses struct {
		Territory string      `json:"territory"`
		Guild     string      `json:"guild"`
		Acquired  string      `json:"acquired"`
		Attacker  interface{} `json:"attacker"`
		Location  struct {
			StartX int `json:"startX"`
			StartY int `json:"startY"`
			EndX   int `json:"endX"`
			EndY   int `json:"endY"`
		} `json:"location"`
	} `json:"Ruined Houses"`
	CorkusOutskirts struct {
		Territory string      `json:"territory"`
		Guild     string      `json:"guild"`
		Acquired  string      `json:"acquired"`
		Attacker  interface{} `json:"attacker"`
		Location  struct {
			StartX int `json:"startX"`
			StartY int `json:"startY"`
			EndX   int `json:"endX"`
			EndY   int `json:"endY"`
		} `json:"location"`
	} `json:"Corkus Outskirts"`
	PathToAhmsordUpper struct {
		Territory string      `json:"territory"`
		Guild     string      `json:"guild"`
		Acquired  string      `json:"acquired"`
		Attacker  interface{} `json:"attacker"`
		Location  struct {
			StartX int `json:"startX"`
			StartY int `json:"startY"`
			EndX   int `json:"endX"`
			EndY   int `json:"endY"`
		} `json:"location"`
	} `json:"Path to Ahmsord Upper"`
	AstraulusTower struct {
		Territory string      `json:"territory"`
		Guild     string      `json:"guild"`
		Acquired  string      `json:"acquired"`
		Attacker  interface{} `json:"attacker"`
		Location  struct {
			StartX int `json:"startX"`
			StartY int `json:"startY"`
			EndX   int `json:"endX"`
			EndY   int `json:"endY"`
		} `json:"location"`
	} `json:"Astraulus' Tower"`
	AngelRefuge struct {
		Territory string      `json:"territory"`
		Guild     string      `json:"guild"`
		Acquired  string      `json:"acquired"`
		Attacker  interface{} `json:"attacker"`
		Location  struct {
			StartX int `json:"startX"`
			StartY int `json:"startY"`
			EndX   int `json:"endX"`
			EndY   int `json:"endY"`
		} `json:"location"`
	} `json:"Angel Refuge"`
	SkyFalls struct {
		Territory string      `json:"territory"`
		Guild     string      `json:"guild"`
		Acquired  string      `json:"acquired"`
		Attacker  interface{} `json:"attacker"`
		Location  struct {
			StartX int `json:"startX"`
			StartY int `json:"startY"`
			EndX   int `json:"endX"`
			EndY   int `json:"endY"`
		} `json:"location"`
	} `json:"Sky Falls"`
	JofashDocks struct {
		Territory string      `json:"territory"`
		Guild     string      `json:"guild"`
		Acquired  string      `json:"acquired"`
		Attacker  interface{} `json:"attacker"`
		Location  struct {
			StartX int `json:"startX"`
			StartY int `json:"startY"`
			EndX   int `json:"endX"`
			EndY   int `json:"endY"`
		} `json:"location"`
	} `json:"Jofash Docks"`
	PhinasFarm struct {
		Territory string      `json:"territory"`
		Guild     string      `json:"guild"`
		Acquired  string      `json:"acquired"`
		Attacker  interface{} `json:"attacker"`
		Location  struct {
			StartX int `json:"startX"`
			StartY int `json:"startY"`
			EndX   int `json:"endX"`
			EndY   int `json:"endY"`
		} `json:"location"`
	} `json:"Phinas Farm"`
	Llevigar struct {
		Territory string      `json:"territory"`
		Guild     string      `json:"guild"`
		Acquired  string      `json:"acquired"`
		Attacker  interface{} `json:"attacker"`
		Location  struct {
			StartX int `json:"startX"`
			StartY int `json:"startY"`
			EndX   int `json:"endX"`
			EndY   int `json:"endY"`
		} `json:"location"`
	} `json:"Llevigar"`
	IcyIsland struct {
		Territory string      `json:"territory"`
		Guild     string      `json:"guild"`
		Acquired  string      `json:"acquired"`
		Attacker  interface{} `json:"attacker"`
		Location  struct {
			StartX int `json:"startX"`
			StartY int `json:"startY"`
			EndX   int `json:"endX"`
			EndY   int `json:"endY"`
		} `json:"location"`
	} `json:"Icy Island"`
	AbandonedPass struct {
		Territory string      `json:"territory"`
		Guild     string      `json:"guild"`
		Acquired  string      `json:"acquired"`
		Attacker  interface{} `json:"attacker"`
		Location  struct {
			StartX int `json:"startX"`
			StartY int `json:"startY"`
			EndX   int `json:"endX"`
			EndY   int `json:"endY"`
		} `json:"location"`
	} `json:"Abandoned Pass"`
	CorkusSeaCove struct {
		Territory string      `json:"territory"`
		Guild     string      `json:"guild"`
		Acquired  string      `json:"acquired"`
		Attacker  interface{} `json:"attacker"`
		Location  struct {
			StartX int `json:"startX"`
			StartY int `json:"startY"`
			EndX   int `json:"endX"`
			EndY   int `json:"endY"`
		} `json:"location"`
	} `json:"Corkus Sea Cove"`
	GreyRuins struct {
		Territory string      `json:"territory"`
		Guild     string      `json:"guild"`
		Acquired  string      `json:"acquired"`
		Attacker  interface{} `json:"attacker"`
		Location  struct {
			StartX int `json:"startX"`
			StartY int `json:"startY"`
			EndX   int `json:"endX"`
			EndY   int `json:"endY"`
		} `json:"location"`
	} `json:"Grey Ruins"`
	Lutho struct {
		Territory string      `json:"territory"`
		Guild     string      `json:"guild"`
		Acquired  string      `json:"acquired"`
		Attacker  interface{} `json:"attacker"`
		Location  struct {
			StartX int `json:"startX"`
			StartY int `json:"startY"`
			EndX   int `json:"endX"`
			EndY   int `json:"endY"`
		} `json:"location"`
	} `json:"Lutho"`
	GatewayToNothing struct {
		Territory string      `json:"territory"`
		Guild     string      `json:"guild"`
		Acquired  string      `json:"acquired"`
		Attacker  interface{} `json:"attacker"`
		Location  struct {
			StartX int `json:"startX"`
			StartY int `json:"startY"`
			EndX   int `json:"endX"`
			EndY   int `json:"endY"`
		} `json:"location"`
	} `json:"Gateway to Nothing"`
	TheGate struct {
		Territory string      `json:"territory"`
		Guild     string      `json:"guild"`
		Acquired  string      `json:"acquired"`
		Attacker  interface{} `json:"attacker"`
		Location  struct {
			StartX int `json:"startX"`
			StartY int `json:"startY"`
			EndX   int `json:"endX"`
			EndY   int `json:"endY"`
		} `json:"location"`
	} `json:"The Gate"`
	AzureFrontier struct {
		Territory string      `json:"territory"`
		Guild     string      `json:"guild"`
		Acquired  string      `json:"acquired"`
		Attacker  interface{} `json:"attacker"`
		Location  struct {
			StartX int `json:"startX"`
			StartY int `json:"startY"`
			EndX   int `json:"endX"`
			EndY   int `json:"endY"`
		} `json:"location"`
	} `json:"Azure Frontier"`
	LightPeninsula struct {
		Territory string      `json:"territory"`
		Guild     string      `json:"guild"`
		Acquired  string      `json:"acquired"`
		Attacker  interface{} `json:"attacker"`
		Location  struct {
			StartX int `json:"startX"`
			StartY int `json:"startY"`
			EndX   int `json:"endX"`
			EndY   int `json:"endY"`
		} `json:"location"`
	} `json:"Light Peninsula"`
	EmeraldTrail struct {
		Territory string      `json:"territory"`
		Guild     string      `json:"guild"`
		Acquired  string      `json:"acquired"`
		Attacker  interface{} `json:"attacker"`
		Location  struct {
			StartX int `json:"startX"`
			StartY int `json:"startY"`
			EndX   int `json:"endX"`
			EndY   int `json:"endY"`
		} `json:"location"`
	} `json:"Emerald Trail"`
	MalticCoast struct {
		Territory string      `json:"territory"`
		Guild     string      `json:"guild"`
		Acquired  string      `json:"acquired"`
		Attacker  interface{} `json:"attacker"`
		Location  struct {
			StartX int `json:"startX"`
			StartY int `json:"startY"`
			EndX   int `json:"endX"`
			EndY   int `json:"endY"`
		} `json:"location"`
	} `json:"Maltic Coast"`
	TimeValley struct {
		Territory string      `json:"territory"`
		Guild     string      `json:"guild"`
		Acquired  string      `json:"acquired"`
		Attacker  interface{} `json:"attacker"`
		Location  struct {
			StartX int `json:"startX"`
			StartY int `json:"startY"`
			EndX   int `json:"endX"`
			EndY   int `json:"endY"`
		} `json:"location"`
	} `json:"Time Valley"`
	SouthNivlaWoods struct {
		Territory string      `json:"territory"`
		Guild     string      `json:"guild"`
		Acquired  string      `json:"acquired"`
		Attacker  interface{} `json:"attacker"`
		Location  struct {
			StartX int `json:"startX"`
			StartY int `json:"startY"`
			EndX   int `json:"endX"`
			EndY   int `json:"endY"`
		} `json:"location"`
	} `json:"South Nivla Woods"`
	DetlasCloseSuburbs struct {
		Territory string      `json:"territory"`
		Guild     string      `json:"guild"`
		Acquired  string      `json:"acquired"`
		Attacker  interface{} `json:"attacker"`
		Location  struct {
			StartX int `json:"startX"`
			StartY int `json:"startY"`
			EndX   int `json:"endX"`
			EndY   int `json:"endY"`
		} `json:"location"`
	} `json:"Detlas Close Suburbs"`
	MageIsland struct {
		Territory string      `json:"territory"`
		Guild     string      `json:"guild"`
		Acquired  string      `json:"acquired"`
		Attacker  interface{} `json:"attacker"`
		Location  struct {
			StartX int `json:"startX"`
			StartY int `json:"startY"`
			EndX   int `json:"endX"`
			EndY   int `json:"endY"`
		} `json:"location"`
	} `json:"Mage Island"`
	NesaakPlainsUpperNorthWest struct {
		Territory string      `json:"territory"`
		Guild     string      `json:"guild"`
		Acquired  string      `json:"acquired"`
		Attacker  interface{} `json:"attacker"`
		Location  struct {
			StartX int `json:"startX"`
			StartY int `json:"startY"`
			EndX   int `json:"endX"`
			EndY   int `json:"endY"`
		} `json:"location"`
	} `json:"Nesaak Plains Upper North West"`
	JungleUpper struct {
		Territory string      `json:"territory"`
		Guild     string      `json:"guild"`
		Acquired  string      `json:"acquired"`
		Attacker  interface{} `json:"attacker"`
		Location  struct {
			StartX int `json:"startX"`
			StartY int `json:"startY"`
			EndX   int `json:"endX"`
			EndY   int `json:"endY"`
		} `json:"location"`
	} `json:"Jungle Upper"`
	RymekWestMid struct {
		Territory string      `json:"territory"`
		Guild     string      `json:"guild"`
		Acquired  string      `json:"acquired"`
		Attacker  interface{} `json:"attacker"`
		Location  struct {
			StartX int `json:"startX"`
			StartY int `json:"startY"`
			EndX   int `json:"endX"`
			EndY   int `json:"endY"`
		} `json:"location"`
	} `json:"Rymek West Mid"`
	DesertLower struct {
		Territory string      `json:"territory"`
		Guild     string      `json:"guild"`
		Acquired  string      `json:"acquired"`
		Attacker  interface{} `json:"attacker"`
		Location  struct {
			StartX int `json:"startX"`
			StartY int `json:"startY"`
			EndX   int `json:"endX"`
			EndY   int `json:"endY"`
		} `json:"location"`
	} `json:"Desert Lower"`
	SavannahWestUpper struct {
		Territory string      `json:"territory"`
		Guild     string      `json:"guild"`
		Acquired  string      `json:"acquired"`
		Attacker  interface{} `json:"attacker"`
		Location  struct {
			StartX int `json:"startX"`
			StartY int `json:"startY"`
			EndX   int `json:"endX"`
			EndY   int `json:"endY"`
		} `json:"location"`
	} `json:"Savannah West Upper"`
	AncientNemract struct {
		Territory string      `json:"territory"`
		Guild     string      `json:"guild"`
		Acquired  string      `json:"acquired"`
		Attacker  interface{} `json:"attacker"`
		Location  struct {
			StartX int `json:"startX"`
			StartY int `json:"startY"`
			EndX   int `json:"endX"`
			EndY   int `json:"endY"`
		} `json:"location"`
	} `json:"Ancient Nemract"`
	DurumIslesUpper struct {
		Territory string      `json:"territory"`
		Guild     string      `json:"guild"`
		Acquired  string      `json:"acquired"`
		Attacker  interface{} `json:"attacker"`
		Location  struct {
			StartX int `json:"startX"`
			StartY int `json:"startY"`
			EndX   int `json:"endX"`
			EndY   int `json:"endY"`
		} `json:"location"`
	} `json:"Durum Isles Upper"`
	DeadIslandSouthEast struct {
		Territory string      `json:"territory"`
		Guild     string      `json:"guild"`
		Acquired  string      `json:"acquired"`
		Attacker  interface{} `json:"attacker"`
		Location  struct {
			StartX int `json:"startX"`
			StartY int `json:"startY"`
			EndX   int `json:"endX"`
			EndY   int `json:"endY"`
		} `json:"location"`
	} `json:"Dead Island South East"`
	TernavesPlainsUpper struct {
		Territory string      `json:"territory"`
		Guild     string      `json:"guild"`
		Acquired  string      `json:"acquired"`
		Attacker  interface{} `json:"attacker"`
		Location  struct {
			StartX int `json:"startX"`
			StartY int `json:"startY"`
			EndX   int `json:"endX"`
			EndY   int `json:"endY"`
		} `json:"location"`
	} `json:"Ternaves Plains Upper"`
	MineBasePlains struct {
		Territory string      `json:"territory"`
		Guild     string      `json:"guild"`
		Acquired  string      `json:"acquired"`
		Attacker  interface{} `json:"attacker"`
		Location  struct {
			StartX int `json:"startX"`
			StartY int `json:"startY"`
			EndX   int `json:"endX"`
			EndY   int `json:"endY"`
		} `json:"location"`
	} `json:"Mine Base Plains"`
	LlevigarFarmPlainsEast struct {
		Territory string      `json:"territory"`
		Guild     string      `json:"guild"`
		Acquired  string      `json:"acquired"`
		Attacker  interface{} `json:"attacker"`
		Location  struct {
			StartX int `json:"startX"`
			StartY int `json:"startY"`
			EndX   int `json:"endX"`
			EndY   int `json:"endY"`
		} `json:"location"`
	} `json:"Llevigar Farm Plains East"`
	SwampWestLower struct {
		Territory string      `json:"territory"`
		Guild     string      `json:"guild"`
		Acquired  string      `json:"acquired"`
		Attacker  interface{} `json:"attacker"`
		Location  struct {
			StartX int `json:"startX"`
			StartY int `json:"startY"`
			EndX   int `json:"endX"`
			EndY   int `json:"endY"`
		} `json:"location"`
	} `json:"Swamp West Lower"`
	SwampDarkForestTransitionLower struct {
		Territory string      `json:"territory"`
		Guild     string      `json:"guild"`
		Acquired  string      `json:"acquired"`
		Attacker  interface{} `json:"attacker"`
		Location  struct {
			StartX int `json:"startX"`
			StartY int `json:"startY"`
			EndX   int `json:"endX"`
			EndY   int `json:"endY"`
		} `json:"location"`
	} `json:"Swamp Dark Forest Transition Lower"`
	SwampMountainTransitionLower struct {
		Territory string      `json:"territory"`
		Guild     string      `json:"guild"`
		Acquired  string      `json:"acquired"`
		Attacker  interface{} `json:"attacker"`
		Location  struct {
			StartX int `json:"startX"`
			StartY int `json:"startY"`
			EndX   int `json:"endX"`
			EndY   int `json:"endY"`
		} `json:"location"`
	} `json:"Swamp Mountain Transition Lower"`
	SunsparkCamp struct {
		Territory string      `json:"territory"`
		Guild     string      `json:"guild"`
		Acquired  string      `json:"acquired"`
		Attacker  interface{} `json:"attacker"`
		Location  struct {
			StartX int `json:"startX"`
			StartY int `json:"startY"`
			EndX   int `json:"endX"`
			EndY   int `json:"endY"`
		} `json:"location"`
	} `json:"Sunspark Camp"`
	LlevigarFarm struct {
		Territory string      `json:"territory"`
		Guild     string      `json:"guild"`
		Acquired  string      `json:"acquired"`
		Attacker  interface{} `json:"attacker"`
		Location  struct {
			StartX int `json:"startX"`
			StartY int `json:"startY"`
			EndX   int `json:"endX"`
			EndY   int `json:"endY"`
		} `json:"location"`
	} `json:"Llevigar Farm"`
	EfilimEastPlains struct {
		Territory string      `json:"territory"`
		Guild     string      `json:"guild"`
		Acquired  string      `json:"acquired"`
		Attacker  interface{} `json:"attacker"`
		Location  struct {
			StartX int `json:"startX"`
			StartY int `json:"startY"`
			EndX   int `json:"endX"`
			EndY   int `json:"endY"`
		} `json:"location"`
	} `json:"Efilim East Plains"`
	LightForestWestLower struct {
		Territory string      `json:"territory"`
		Guild     string      `json:"guild"`
		Acquired  string      `json:"acquired"`
		Attacker  interface{} `json:"attacker"`
		Location  struct {
			StartX int `json:"startX"`
			StartY int `json:"startY"`
			EndX   int `json:"endX"`
			EndY   int `json:"endY"`
		} `json:"location"`
	} `json:"Light Forest West Lower"`
	LightForestCanyon struct {
		Territory string      `json:"territory"`
		Guild     string      `json:"guild"`
		Acquired  string      `json:"acquired"`
		Attacker  interface{} `json:"attacker"`
		Location  struct {
			StartX int `json:"startX"`
			StartY int `json:"startY"`
			EndX   int `json:"endX"`
			EndY   int `json:"endY"`
		} `json:"location"`
	} `json:"Light Forest Canyon"`
	FortressSouth struct {
		Territory string      `json:"territory"`
		Guild     string      `json:"guild"`
		Acquired  string      `json:"acquired"`
		Attacker  interface{} `json:"attacker"`
		Location  struct {
			StartX int `json:"startX"`
			StartY int `json:"startY"`
			EndX   int `json:"endX"`
			EndY   int `json:"endY"`
		} `json:"location"`
	} `json:"Fortress South"`
	KanderMines struct {
		Territory string      `json:"territory"`
		Guild     string      `json:"guild"`
		Acquired  string      `json:"acquired"`
		Attacker  interface{} `json:"attacker"`
		Location  struct {
			StartX int `json:"startX"`
			StartY int `json:"startY"`
			EndX   int `json:"endX"`
			EndY   int `json:"endY"`
		} `json:"location"`
	} `json:"Kander Mines"`
	DecayedBasin struct {
		Territory string      `json:"territory"`
		Guild     string      `json:"guild"`
		Acquired  string      `json:"acquired"`
		Attacker  interface{} `json:"attacker"`
		Location  struct {
			StartX int `json:"startX"`
			StartY int `json:"startY"`
			EndX   int `json:"endX"`
			EndY   int `json:"endY"`
		} `json:"location"`
	} `json:"Decayed Basin"`
	CinfrasSSmallFarm struct {
		Territory string      `json:"territory"`
		Guild     string      `json:"guild"`
		Acquired  string      `json:"acquired"`
		Attacker  interface{} `json:"attacker"`
		Location  struct {
			StartX int `json:"startX"`
			StartY int `json:"startY"`
			EndX   int `json:"endX"`
			EndY   int `json:"endY"`
		} `json:"location"`
	} `json:"Cinfras's Small Farm"`
	GyliaLakeNorthEast struct {
		Territory string      `json:"territory"`
		Guild     string      `json:"guild"`
		Acquired  string      `json:"acquired"`
		Attacker  interface{} `json:"attacker"`
		Location  struct {
			StartX int `json:"startX"`
			StartY int `json:"startY"`
			EndX   int `json:"endX"`
			EndY   int `json:"endY"`
		} `json:"location"`
	} `json:"Gylia Lake North East"`
	AldoreiSNorthExit struct {
		Territory string      `json:"territory"`
		Guild     string      `json:"guild"`
		Acquired  string      `json:"acquired"`
		Attacker  interface{} `json:"attacker"`
		Location  struct {
			StartX int `json:"startX"`
			StartY int `json:"startY"`
			EndX   int `json:"endX"`
			EndY   int `json:"endY"`
		} `json:"location"`
	} `json:"Aldorei's North Exit"`
	PathToThanos struct {
		Territory string      `json:"territory"`
		Guild     string      `json:"guild"`
		Acquired  string      `json:"acquired"`
		Attacker  interface{} `json:"attacker"`
		Location  struct {
			StartX int `json:"startX"`
			StartY int `json:"startY"`
			EndX   int `json:"endX"`
			EndY   int `json:"endY"`
		} `json:"location"`
	} `json:"Path To Thanos"`
	BanditCaveLower struct {
		Territory string      `json:"territory"`
		Guild     string      `json:"guild"`
		Acquired  string      `json:"acquired"`
		Attacker  interface{} `json:"attacker"`
		Location  struct {
			StartX int `json:"startX"`
			StartY int `json:"startY"`
			EndX   int `json:"endX"`
			EndY   int `json:"endY"`
		} `json:"location"`
	} `json:"Bandit Cave Lower"`
	CanyonPathSouthWest struct {
		Territory string      `json:"territory"`
		Guild     string      `json:"guild"`
		Acquired  string      `json:"acquired"`
		Attacker  interface{} `json:"attacker"`
		Location  struct {
			StartX int `json:"startX"`
			StartY int `json:"startY"`
			EndX   int `json:"endX"`
			EndY   int `json:"endY"`
		} `json:"location"`
	} `json:"Canyon Path South West"`
	CanyonMountainSouth struct {
		Territory string      `json:"territory"`
		Guild     string      `json:"guild"`
		Acquired  string      `json:"acquired"`
		Attacker  interface{} `json:"attacker"`
		Location  struct {
			StartX int `json:"startX"`
			StartY int `json:"startY"`
			EndX   int `json:"endX"`
			EndY   int `json:"endY"`
		} `json:"location"`
	} `json:"Canyon Mountain South"`
	MountainPath struct {
		Territory string      `json:"territory"`
		Guild     string      `json:"guild"`
		Acquired  string      `json:"acquired"`
		Attacker  interface{} `json:"attacker"`
		Location  struct {
			StartX int `json:"startX"`
			StartY int `json:"startY"`
			EndX   int `json:"endX"`
			EndY   int `json:"endY"`
		} `json:"location"`
	} `json:"Mountain Path"`
	CliffsideWaterfall struct {
		Territory string      `json:"territory"`
		Guild     string      `json:"guild"`
		Acquired  string      `json:"acquired"`
		Attacker  interface{} `json:"attacker"`
		Location  struct {
			StartX int `json:"startX"`
			StartY int `json:"startY"`
			EndX   int `json:"endX"`
			EndY   int `json:"endY"`
		} `json:"location"`
	} `json:"Cliffside Waterfall"`
	CliffsidePassage struct {
		Territory string      `json:"territory"`
		Guild     string      `json:"guild"`
		Acquired  string      `json:"acquired"`
		Attacker  interface{} `json:"attacker"`
		Location  struct {
			StartX int `json:"startX"`
			StartY int `json:"startY"`
			EndX   int `json:"endX"`
			EndY   int `json:"endY"`
		} `json:"location"`
	} `json:"Cliffside Passage"`
	Thesead struct {
		Territory string      `json:"territory"`
		Guild     string      `json:"guild"`
		Acquired  string      `json:"acquired"`
		Attacker  interface{} `json:"attacker"`
		Location  struct {
			StartX int `json:"startX"`
			StartY int `json:"startY"`
			EndX   int `json:"endX"`
			EndY   int `json:"endY"`
		} `json:"location"`
	} `json:"Thesead"`
	VolcanicSlope struct {
		Territory string      `json:"territory"`
		Guild     string      `json:"guild"`
		Acquired  string      `json:"acquired"`
		Attacker  interface{} `json:"attacker"`
		Location  struct {
			StartX int `json:"startX"`
			StartY int `json:"startY"`
			EndX   int `json:"endX"`
			EndY   int `json:"endY"`
		} `json:"location"`
	} `json:"Volcanic Slope"`
	CorkusCastle struct {
		Territory string      `json:"territory"`
		Guild     string      `json:"guild"`
		Acquired  string      `json:"acquired"`
		Attacker  interface{} `json:"attacker"`
		Location  struct {
			StartX int `json:"startX"`
			StartY int `json:"startY"`
			EndX   int `json:"endX"`
			EndY   int `json:"endY"`
		} `json:"location"`
	} `json:"Corkus Castle"`
	CorkusForestNorth struct {
		Territory string      `json:"territory"`
		Guild     string      `json:"guild"`
		Acquired  string      `json:"acquired"`
		Attacker  interface{} `json:"attacker"`
		Location  struct {
			StartX int `json:"startX"`
			StartY int `json:"startY"`
			EndX   int `json:"endX"`
			EndY   int `json:"endY"`
		} `json:"location"`
	} `json:"Corkus Forest North"`
	AvosTemple struct {
		Territory string      `json:"territory"`
		Guild     string      `json:"guild"`
		Acquired  string      `json:"acquired"`
		Attacker  interface{} `json:"attacker"`
		Location  struct {
			StartX int `json:"startX"`
			StartY int `json:"startY"`
			EndX   int `json:"endX"`
			EndY   int `json:"endY"`
		} `json:"location"`
	} `json:"Avos Temple"`
	OldCoalMine struct {
		Territory string      `json:"territory"`
		Guild     string      `json:"guild"`
		Acquired  string      `json:"acquired"`
		Attacker  interface{} `json:"attacker"`
		Location  struct {
			StartX int `json:"startX"`
			StartY int `json:"startY"`
			EndX   int `json:"endX"`
			EndY   int `json:"endY"`
		} `json:"location"`
	} `json:"Old Coal Mine"`
	CentralIslands struct {
		Territory string      `json:"territory"`
		Guild     string      `json:"guild"`
		Acquired  string      `json:"acquired"`
		Attacker  interface{} `json:"attacker"`
		Location  struct {
			StartX int `json:"startX"`
			StartY int `json:"startY"`
			EndX   int `json:"endX"`
			EndY   int `json:"endY"`
		} `json:"location"`
	} `json:"Central Islands"`
	Lusuco struct {
		Territory string      `json:"territory"`
		Guild     string      `json:"guild"`
		Acquired  string      `json:"acquired"`
		Attacker  interface{} `json:"attacker"`
		Location  struct {
			StartX int `json:"startX"`
			StartY int `json:"startY"`
			EndX   int `json:"endX"`
			EndY   int `json:"endY"`
		} `json:"location"`
	} `json:"Lusuco"`
	HerbCave struct {
		Territory string      `json:"territory"`
		Guild     string      `json:"guild"`
		Acquired  string      `json:"acquired"`
		Attacker  interface{} `json:"attacker"`
		Location  struct {
			StartX int `json:"startX"`
			StartY int `json:"startY"`
			EndX   int `json:"endX"`
			EndY   int `json:"endY"`
		} `json:"location"`
	} `json:"Herb Cave"`
	SouthernOutpost struct {
		Territory string      `json:"territory"`
		Guild     string      `json:"guild"`
		Acquired  string      `json:"acquired"`
		Attacker  interface{} `json:"attacker"`
		Location  struct {
			StartX int `json:"startX"`
			StartY int `json:"startY"`
			EndX   int `json:"endX"`
			EndY   int `json:"endY"`
		} `json:"location"`
	} `json:"Southern Outpost"`
	ForestOfEyes struct {
		Territory string      `json:"territory"`
		Guild     string      `json:"guild"`
		Acquired  string      `json:"acquired"`
		Attacker  interface{} `json:"attacker"`
		Location  struct {
			StartX int `json:"startX"`
			StartY int `json:"startY"`
			EndX   int `json:"endX"`
			EndY   int `json:"endY"`
		} `json:"location"`
	} `json:"Forest of Eyes"`
	FinalStep struct {
		Territory string      `json:"territory"`
		Guild     string      `json:"guild"`
		Acquired  string      `json:"acquired"`
		Attacker  interface{} `json:"attacker"`
		Location  struct {
			StartX int `json:"startX"`
			StartY int `json:"startY"`
			EndX   int `json:"endX"`
			EndY   int `json:"endY"`
		} `json:"location"`
	} `json:"Final Step"`
	PathToLight struct {
		Territory string      `json:"territory"`
		Guild     string      `json:"guild"`
		Acquired  string      `json:"acquired"`
		Attacker  interface{} `json:"attacker"`
		Location  struct {
			StartX int `json:"startX"`
			StartY int `json:"startY"`
			EndX   int `json:"endX"`
			EndY   int `json:"endY"`
		} `json:"location"`
	} `json:"Path to Light"`
	PigmenRavinesEntrance struct {
		Territory string      `json:"territory"`
		Guild     string      `json:"guild"`
		Acquired  string      `json:"acquired"`
		Attacker  interface{} `json:"attacker"`
		Location  struct {
			StartX int `json:"startX"`
			StartY int `json:"startY"`
			EndX   int `json:"endX"`
			EndY   int `json:"endY"`
		} `json:"location"`
	} `json:"Pigmen Ravines Entrance"`
	CorruptedRoad struct {
		Territory string      `json:"territory"`
		Guild     string      `json:"guild"`
		Acquired  string      `json:"acquired"`
		Attacker  interface{} `json:"attacker"`
		Location  struct {
			StartX int `json:"startX"`
			StartY int `json:"startY"`
			EndX   int `json:"endX"`
			EndY   int `json:"endY"`
		} `json:"location"`
	} `json:"Corrupted Road"`
	NesaakPlainsSouthEast struct {
		Territory string      `json:"territory"`
		Guild     string      `json:"guild"`
		Acquired  string      `json:"acquired"`
		Attacker  interface{} `json:"attacker"`
		Location  struct {
			StartX int `json:"startX"`
			StartY int `json:"startY"`
			EndX   int `json:"endX"`
			EndY   int `json:"endY"`
		} `json:"location"`
	} `json:"Nesaak Plains South East"`
	RymekEastLower struct {
		Territory string      `json:"territory"`
		Guild     string      `json:"guild"`
		Acquired  string      `json:"acquired"`
		Attacker  interface{} `json:"attacker"`
		Location  struct {
			StartX int `json:"startX"`
			StartY int `json:"startY"`
			EndX   int `json:"endX"`
			EndY   int `json:"endY"`
		} `json:"location"`
	} `json:"Rymek East Lower"`
	DesertWestLower struct {
		Territory string      `json:"territory"`
		Guild     string      `json:"guild"`
		Acquired  string      `json:"acquired"`
		Attacker  interface{} `json:"attacker"`
		Location  struct {
			StartX int `json:"startX"`
			StartY int `json:"startY"`
			EndX   int `json:"endX"`
			EndY   int `json:"endY"`
		} `json:"location"`
	} `json:"Desert West Lower"`
	RoosterIsland struct {
		Territory string      `json:"territory"`
		Guild     string      `json:"guild"`
		Acquired  string      `json:"acquired"`
		Attacker  interface{} `json:"attacker"`
		Location  struct {
			StartX int `json:"startX"`
			StartY int `json:"startY"`
			EndX   int `json:"endX"`
			EndY   int `json:"endY"`
		} `json:"location"`
	} `json:"Rooster Island"`
	VolcanoUpper struct {
		Territory string      `json:"territory"`
		Guild     string      `json:"guild"`
		Acquired  string      `json:"acquired"`
		Attacker  interface{} `json:"attacker"`
		Location  struct {
			StartX int `json:"startX"`
			StartY int `json:"startY"`
			EndX   int `json:"endX"`
			EndY   int `json:"endY"`
		} `json:"location"`
	} `json:"Volcano Upper"`
	DetlasTrailWestPlains struct {
		Territory string      `json:"territory"`
		Guild     string      `json:"guild"`
		Acquired  string      `json:"acquired"`
		Attacker  interface{} `json:"attacker"`
		Location  struct {
			StartX int `json:"startX"`
			StartY int `json:"startY"`
			EndX   int `json:"endX"`
			EndY   int `json:"endY"`
		} `json:"location"`
	} `json:"Detlas Trail West Plains"`
	SwampWestMidUpper struct {
		Territory string      `json:"territory"`
		Guild     string      `json:"guild"`
		Acquired  string      `json:"acquired"`
		Attacker  interface{} `json:"attacker"`
		Location  struct {
			StartX int `json:"startX"`
			StartY int `json:"startY"`
			EndX   int `json:"endX"`
			EndY   int `json:"endY"`
		} `json:"location"`
	} `json:"Swamp West Mid-Upper"`
	QuartzMinesSouthWest struct {
		Territory string      `json:"territory"`
		Guild     string      `json:"guild"`
		Acquired  string      `json:"acquired"`
		Attacker  interface{} `json:"attacker"`
		Location  struct {
			StartX int `json:"startX"`
			StartY int `json:"startY"`
			EndX   int `json:"endX"`
			EndY   int `json:"endY"`
		} `json:"location"`
	} `json:"Quartz Mines South West"`
	LeadinFortress struct {
		Territory string      `json:"territory"`
		Guild     string      `json:"guild"`
		Acquired  string      `json:"acquired"`
		Attacker  interface{} `json:"attacker"`
		Location  struct {
			StartX int `json:"startX"`
			StartY int `json:"startY"`
			EndX   int `json:"endX"`
			EndY   int `json:"endY"`
		} `json:"location"`
	} `json:"Leadin Fortress"`
	LightForestEastMid struct {
		Territory string      `json:"territory"`
		Guild     string      `json:"guild"`
		Acquired  string      `json:"acquired"`
		Attacker  interface{} `json:"attacker"`
		Location  struct {
			StartX int `json:"startX"`
			StartY int `json:"startY"`
			EndX   int `json:"endX"`
			EndY   int `json:"endY"`
		} `json:"location"`
	} `json:"Light Forest East Mid"`
	VisceraPitsWest struct {
		Territory string      `json:"territory"`
		Guild     string      `json:"guild"`
		Acquired  string      `json:"acquired"`
		Attacker  interface{} `json:"attacker"`
		Location  struct {
			StartX int `json:"startX"`
			StartY int `json:"startY"`
			EndX   int `json:"endX"`
			EndY   int `json:"endY"`
		} `json:"location"`
	} `json:"Viscera Pits West"`
	FallenVillage struct {
		Territory string      `json:"territory"`
		Guild     string      `json:"guild"`
		Acquired  string      `json:"acquired"`
		Attacker  interface{} `json:"attacker"`
		Location  struct {
			StartX int `json:"startX"`
			StartY int `json:"startY"`
			EndX   int `json:"endX"`
			EndY   int `json:"endY"`
		} `json:"location"`
	} `json:"Fallen Village"`
	AldoreiValleyMid struct {
		Territory string      `json:"territory"`
		Guild     string      `json:"guild"`
		Acquired  string      `json:"acquired"`
		Attacker  interface{} `json:"attacker"`
		Location  struct {
			StartX int `json:"startX"`
			StartY int `json:"startY"`
			EndX   int `json:"endX"`
			EndY   int `json:"endY"`
		} `json:"location"`
	} `json:"Aldorei Valley Mid"`
	MilitaryBaseLower struct {
		Territory string      `json:"territory"`
		Guild     string      `json:"guild"`
		Acquired  string      `json:"acquired"`
		Attacker  interface{} `json:"attacker"`
		Location  struct {
			StartX int `json:"startX"`
			StartY int `json:"startY"`
			EndX   int `json:"endX"`
			EndY   int `json:"endY"`
		} `json:"location"`
	} `json:"Military Base Lower"`
	ThanosValleyWest struct {
		Territory string      `json:"territory"`
		Guild     string      `json:"guild"`
		Acquired  string      `json:"acquired"`
		Attacker  interface{} `json:"attacker"`
		Location  struct {
			StartX int `json:"startX"`
			StartY int `json:"startY"`
			EndX   int `json:"endX"`
			EndY   int `json:"endY"`
		} `json:"location"`
	} `json:"Thanos Valley West"`
	TempleOfTheLostEast struct {
		Territory string      `json:"territory"`
		Guild     string      `json:"guild"`
		Acquired  string      `json:"acquired"`
		Attacker  interface{} `json:"attacker"`
		Location  struct {
			StartX int `json:"startX"`
			StartY int `json:"startY"`
			EndX   int `json:"endX"`
			EndY   int `json:"endY"`
		} `json:"location"`
	} `json:"Temple of the Lost East"`
	ChainedHouse struct {
		Territory string      `json:"territory"`
		Guild     string      `json:"guild"`
		Acquired  string      `json:"acquired"`
		Attacker  interface{} `json:"attacker"`
		Location  struct {
			StartX int `json:"startX"`
			StartY int `json:"startY"`
			EndX   int `json:"endX"`
			EndY   int `json:"endY"`
		} `json:"location"`
	} `json:"Chained House"`
	DernelJungleLower struct {
		Territory string      `json:"territory"`
		Guild     string      `json:"guild"`
		Acquired  string      `json:"acquired"`
		Attacker  interface{} `json:"attacker"`
		Location  struct {
			StartX int `json:"startX"`
			StartY int `json:"startY"`
			EndX   int `json:"endX"`
			EndY   int `json:"endY"`
		} `json:"location"`
	} `json:"Dernel Jungle Lower"`
	CorkusCountryside struct {
		Territory string      `json:"territory"`
		Guild     string      `json:"guild"`
		Acquired  string      `json:"acquired"`
		Attacker  interface{} `json:"attacker"`
		Location  struct {
			StartX int `json:"startX"`
			StartY int `json:"startY"`
			EndX   int `json:"endX"`
			EndY   int `json:"endY"`
		} `json:"location"`
	} `json:"Corkus Countryside"`
	AhmsordOutskirts struct {
		Territory string      `json:"territory"`
		Guild     string      `json:"guild"`
		Acquired  string      `json:"acquired"`
		Attacker  interface{} `json:"attacker"`
		Location  struct {
			StartX int `json:"startX"`
			StartY int `json:"startY"`
			EndX   int `json:"endX"`
			EndY   int `json:"endY"`
		} `json:"location"`
	} `json:"Ahmsord Outskirts"`
	CinfrasOutskirts struct {
		Territory string      `json:"territory"`
		Guild     string      `json:"guild"`
		Acquired  string      `json:"acquired"`
		Attacker  interface{} `json:"attacker"`
		Location  struct {
			StartX int `json:"startX"`
			StartY int `json:"startY"`
			EndX   int `json:"endX"`
			EndY   int `json:"endY"`
		} `json:"location"`
	} `json:"Cinfras Outskirts"`
	TheBrokenRoad struct {
		Territory string      `json:"territory"`
		Guild     string      `json:"guild"`
		Acquired  string      `json:"acquired"`
		Attacker  interface{} `json:"attacker"`
		Location  struct {
			StartX int `json:"startX"`
			StartY int `json:"startY"`
			EndX   int `json:"endX"`
			EndY   int `json:"endY"`
		} `json:"location"`
	} `json:"The Broken Road"`
	LuminousPlateau struct {
		Territory string      `json:"territory"`
		Guild     string      `json:"guild"`
		Acquired  string      `json:"acquired"`
		Attacker  interface{} `json:"attacker"`
		Location  struct {
			StartX int `json:"startX"`
			StartY int `json:"startY"`
			EndX   int `json:"endX"`
			EndY   int `json:"endY"`
		} `json:"location"`
	} `json:"Luminous Plateau"`
	ElkurnFields struct {
		Territory string      `json:"territory"`
		Guild     string      `json:"guild"`
		Acquired  string      `json:"acquired"`
		Attacker  interface{} `json:"attacker"`
		Location  struct {
			StartX int `json:"startX"`
			StartY int `json:"startY"`
			EndX   int `json:"endX"`
			EndY   int `json:"endY"`
		} `json:"location"`
	} `json:"Elkurn Fields"`
	GreatBridgeNesaak struct {
		Territory string      `json:"territory"`
		Guild     string      `json:"guild"`
		Acquired  string      `json:"acquired"`
		Attacker  interface{} `json:"attacker"`
		Location  struct {
			StartX int `json:"startX"`
			StartY int `json:"startY"`
			EndX   int `json:"endX"`
			EndY   int `json:"endY"`
		} `json:"location"`
	} `json:"Great Bridge Nesaak"`
	PlainsCoast struct {
		Territory string      `json:"territory"`
		Guild     string      `json:"guild"`
		Acquired  string      `json:"acquired"`
		Attacker  interface{} `json:"attacker"`
		Location  struct {
			StartX int `json:"startX"`
			StartY int `json:"startY"`
			EndX   int `json:"endX"`
			EndY   int `json:"endY"`
		} `json:"location"`
	} `json:"Plains Coast"`
	NesaakTransition struct {
		Territory string      `json:"territory"`
		Guild     string      `json:"guild"`
		Acquired  string      `json:"acquired"`
		Attacker  interface{} `json:"attacker"`
		Location  struct {
			StartX int `json:"startX"`
			StartY int `json:"startY"`
			EndX   int `json:"endX"`
			EndY   int `json:"endY"`
		} `json:"location"`
	} `json:"Nesaak Transition"`
	EntranceToOlux struct {
		Territory string      `json:"territory"`
		Guild     string      `json:"guild"`
		Acquired  string      `json:"acquired"`
		Attacker  interface{} `json:"attacker"`
		Location  struct {
			StartX int `json:"startX"`
			StartY int `json:"startY"`
			EndX   int `json:"endX"`
			EndY   int `json:"endY"`
		} `json:"location"`
	} `json:"Entrance to Olux"`
	LightForestSouthEntrance struct {
		Territory string      `json:"territory"`
		Guild     string      `json:"guild"`
		Acquired  string      `json:"acquired"`
		Attacker  interface{} `json:"attacker"`
		Location  struct {
			StartX int `json:"startX"`
			StartY int `json:"startY"`
			EndX   int `json:"endX"`
			EndY   int `json:"endY"`
		} `json:"location"`
	} `json:"Light Forest South Entrance"`
	OldCrossroadsSouth struct {
		Territory string      `json:"territory"`
		Guild     string      `json:"guild"`
		Acquired  string      `json:"acquired"`
		Attacker  interface{} `json:"attacker"`
		Location  struct {
			StartX int `json:"startX"`
			StartY int `json:"startY"`
			EndX   int `json:"endX"`
			EndY   int `json:"endY"`
		} `json:"location"`
	} `json:"Old Crossroads South"`
	BurningFarm struct {
		Territory string      `json:"territory"`
		Guild     string      `json:"guild"`
		Acquired  string      `json:"acquired"`
		Attacker  interface{} `json:"attacker"`
		Location  struct {
			StartX int `json:"startX"`
			StartY int `json:"startY"`
			EndX   int `json:"endX"`
			EndY   int `json:"endY"`
		} `json:"location"`
	} `json:"Burning Farm"`
	CanyonDropoff struct {
		Territory string      `json:"territory"`
		Guild     string      `json:"guild"`
		Acquired  string      `json:"acquired"`
		Attacker  interface{} `json:"attacker"`
		Location  struct {
			StartX int `json:"startX"`
			StartY int `json:"startY"`
			EndX   int `json:"endX"`
			EndY   int `json:"endY"`
		} `json:"location"`
	} `json:"Canyon Dropoff"`
	LavaLake struct {
		Territory string      `json:"territory"`
		Guild     string      `json:"guild"`
		Acquired  string      `json:"acquired"`
		Attacker  interface{} `json:"attacker"`
		Location  struct {
			StartX int `json:"startX"`
			StartY int `json:"startY"`
			EndX   int `json:"endX"`
			EndY   int `json:"endY"`
		} `json:"location"`
	} `json:"Lava Lake"`
	SkyCastle struct {
		Territory string      `json:"territory"`
		Guild     string      `json:"guild"`
		Acquired  string      `json:"acquired"`
		Attacker  interface{} `json:"attacker"`
		Location  struct {
			StartX int `json:"startX"`
			StartY int `json:"startY"`
			EndX   int `json:"endX"`
			EndY   int `json:"endY"`
		} `json:"location"`
	} `json:"Sky Castle"`
	FlerisTrail struct {
		Territory string      `json:"territory"`
		Guild     string      `json:"guild"`
		Acquired  string      `json:"acquired"`
		Attacker  interface{} `json:"attacker"`
		Location  struct {
			StartX int `json:"startX"`
			StartY int `json:"startY"`
			EndX   int `json:"endX"`
			EndY   int `json:"endY"`
		} `json:"location"`
	} `json:"Fleris Trail"`
	RagniNorthSuburbs struct {
		Territory string      `json:"territory"`
		Guild     string      `json:"guild"`
		Acquired  string      `json:"acquired"`
		Attacker  interface{} `json:"attacker"`
		Location  struct {
			StartX int `json:"startX"`
			StartY int `json:"startY"`
			EndX   int `json:"endX"`
			EndY   int `json:"endY"`
		} `json:"location"`
	} `json:"Ragni North Suburbs"`
	DesertEastLower struct {
		Territory string      `json:"territory"`
		Guild     string      `json:"guild"`
		Acquired  string      `json:"acquired"`
		Attacker  interface{} `json:"attacker"`
		Location  struct {
			StartX int `json:"startX"`
			StartY int `json:"startY"`
			EndX   int `json:"endX"`
			EndY   int `json:"endY"`
		} `json:"location"`
	} `json:"Desert East Lower"`
	LlevigarPlainsEastLower struct {
		Territory string      `json:"territory"`
		Guild     string      `json:"guild"`
		Acquired  string      `json:"acquired"`
		Attacker  interface{} `json:"attacker"`
		Location  struct {
			StartX int `json:"startX"`
			StartY int `json:"startY"`
			EndX   int `json:"endX"`
			EndY   int `json:"endY"`
		} `json:"location"`
	} `json:"Llevigar Plains East Lower"`
	GelibordCorruptedFarm struct {
		Territory string      `json:"territory"`
		Guild     string      `json:"guild"`
		Acquired  string      `json:"acquired"`
		Attacker  interface{} `json:"attacker"`
		Location  struct {
			StartX int `json:"startX"`
			StartY int `json:"startY"`
			EndX   int `json:"endX"`
			EndY   int `json:"endY"`
		} `json:"location"`
	} `json:"Gelibord Corrupted Farm"`
	CanyonPathSouthEast struct {
		Territory string      `json:"territory"`
		Guild     string      `json:"guild"`
		Acquired  string      `json:"acquired"`
		Attacker  interface{} `json:"attacker"`
		Location  struct {
			StartX int `json:"startX"`
			StartY int `json:"startY"`
			EndX   int `json:"endX"`
			EndY   int `json:"endY"`
		} `json:"location"`
	} `json:"Canyon Path South East"`
	CorkusCityMine struct {
		Territory string      `json:"territory"`
		Guild     string      `json:"guild"`
		Acquired  string      `json:"acquired"`
		Attacker  interface{} `json:"attacker"`
		Location  struct {
			StartX int `json:"startX"`
			StartY int `json:"startY"`
			EndX   int `json:"endX"`
			EndY   int `json:"endY"`
		} `json:"location"`
	} `json:"Corkus City Mine"`
	ToxicDrip struct {
		Territory string      `json:"territory"`
		Guild     string      `json:"guild"`
		Acquired  string      `json:"acquired"`
		Attacker  interface{} `json:"attacker"`
		Location  struct {
			StartX int `json:"startX"`
			StartY int `json:"startY"`
			EndX   int `json:"endX"`
			EndY   int `json:"endY"`
		} `json:"location"`
	} `json:"Toxic Drip"`
	ArachnidRoute struct {
		Territory string      `json:"territory"`
		Guild     string      `json:"guild"`
		Acquired  string      `json:"acquired"`
		Attacker  interface{} `json:"attacker"`
		Location  struct {
			StartX int `json:"startX"`
			StartY int `json:"startY"`
			EndX   int `json:"endX"`
			EndY   int `json:"endY"`
		} `json:"location"`
	} `json:"Arachnid Route"`
	SablestoneCamp struct {
		Territory string      `json:"territory"`
		Guild     string      `json:"guild"`
		Acquired  string      `json:"acquired"`
		Attacker  interface{} `json:"attacker"`
		Location  struct {
			StartX int `json:"startX"`
			StartY int `json:"startY"`
			EndX   int `json:"endX"`
			EndY   int `json:"endY"`
		} `json:"location"`
	} `json:"Sablestone Camp"`
	CliffsideLake struct {
		Territory string      `json:"territory"`
		Guild     string      `json:"guild"`
		Acquired  string      `json:"acquired"`
		Attacker  interface{} `json:"attacker"`
		Location  struct {
			StartX int `json:"startX"`
			StartY int `json:"startY"`
			EndX   int `json:"endX"`
			EndY   int `json:"endY"`
		} `json:"location"`
	} `json:"Cliffside Lake"`
	SkiensIsland struct {
		Territory string      `json:"territory"`
		Guild     string      `json:"guild"`
		Acquired  string      `json:"acquired"`
		Attacker  interface{} `json:"attacker"`
		Location  struct {
			StartX int `json:"startX"`
			StartY int `json:"startY"`
			EndX   int `json:"endX"`
			EndY   int `json:"endY"`
		} `json:"location"`
	} `json:"Skiens Island"`
	RaiderSBaseLower struct {
		Territory string      `json:"territory"`
		Guild     string      `json:"guild"`
		Acquired  string      `json:"acquired"`
		Attacker  interface{} `json:"attacker"`
		Location  struct {
			StartX int `json:"startX"`
			StartY int `json:"startY"`
			EndX   int `json:"endX"`
			EndY   int `json:"endY"`
		} `json:"location"`
	} `json:"Raider's Base Lower"`
	CinfrasCountyUpper struct {
		Territory string      `json:"territory"`
		Guild     string      `json:"guild"`
		Acquired  string      `json:"acquired"`
		Attacker  interface{} `json:"attacker"`
		Location  struct {
			StartX int `json:"startX"`
			StartY int `json:"startY"`
			EndX   int `json:"endX"`
			EndY   int `json:"endY"`
		} `json:"location"`
	} `json:"Cinfras County Upper"`
}

type GuildObject struct {
	Name    string `json:"name"`
	Prefix  string `json:"prefix"`
	Members []struct {
		Name           string    `json:"name"`
		UUID           string    `json:"uuid"`
		Rank           string    `json:"rank"`
		Contributed    int64     `json:"contributed"`
		Joined         time.Time `json:"joined"`
		JoinedFriendly string    `json:"joinedFriendly"`
	} `json:"members"`
	Xp              int       `json:"xp"`
	Level           int       `json:"level"`
	Created         time.Time `json:"created"`
	CreatedFriendly string    `json:"createdFriendly"`
	Territories     int       `json:"territories"`
	Banner          struct {
		Base      string `json:"base"`
		Tier      int    `json:"tier"`
		Structure string `json:"structure"`
		Layers    []struct {
			Colour  string `json:"colour"`
			Pattern string `json:"pattern"`
		} `json:"layers"`
	} `json:"banner"`
}
type PlayerObject struct {
	Kind      string `json:"kind"`
	Code      int    `json:"code"`
	Timestamp int64  `json:"timestamp"`
	Version   string `json:"version"`
	Data      []struct {
		Username string `json:"username"`
		UUID     string `json:"uuid"`
		Rank     string `json:"rank"`
		Meta     struct {
			FirstJoin time.Time `json:"firstJoin"`
			LastJoin  time.Time `json:"lastJoin"`
			Location  struct {
				Online bool   `json:"online"`
				Server string `json:"server"`
			} `json:"location"`
			Playtime float64 `json:"playtime"`
			Tag      struct {
				Display bool   `json:"display"`
				Value   string `json:"value"`
			} `json:"tag"`
			Veteran bool `json:"veteran"`
		} `json:"meta"`
		Characters map[string]struct {
			Type     string `json:"type"`
			Level    int    `json:"level"`
			Dungeons struct {
				Completed int `json:"completed"`
				List      []struct {
					Name      string `json:"name"`
					Completed int    `json:"completed"`
				} `json:"list"`
			} `json:"dungeons"`
			Raids struct {
				Completed int `json:"completed"`
				List      []struct {
					Name      string `json:"name"`
					Completed int    `json:"completed"`
				} `json:"list"`
			} `json:"raids"`
			Quests struct {
				Completed int      `json:"completed"`
				List      []string `json:"list"`
			} `json:"quests"`
			ItemsIdentified int `json:"itemsIdentified"`
			MobsKilled      int `json:"mobsKilled"`
			Pvp             struct {
				Kills  int `json:"kills"`
				Deaths int `json:"deaths"`
			} `json:"pvp"`
			BlocksWalked int `json:"blocksWalked"`
			Logins       int `json:"logins"`
			Deaths       int `json:"deaths"`
			Playtime     int `json:"playtime"`
			Gamemode     struct {
				Craftsman bool `json:"craftsman"`
				Hardcore  bool `json:"hardcore"`
				Ironman   bool `json:"ironman"`
				Hunted    bool `json:"hunted"`
			} `json:"gamemode"`
			Skills struct {
				Strength     int `json:"strength"`
				Dexterity    int `json:"dexterity"`
				Intelligence int `json:"intelligence"`
				Defence      int `json:"defence"`
				Defense      int `json:"defense"`
				Agility      int `json:"agility"`
			} `json:"skills"`
			Professions struct {
				Alchemism struct {
					Level int `json:"level"`
					Xp    int `json:"xp"`
				} `json:"alchemism"`
				Armouring struct {
					Level int `json:"level"`
					Xp    int `json:"xp"`
				} `json:"armouring"`
				Combat struct {
					Level int     `json:"level"`
					Xp    float64 `json:"xp"`
				} `json:"combat"`
				Cooking struct {
					Level int     `json:"level"`
					Xp    float64 `json:"xp"`
				} `json:"cooking"`
				Farming struct {
					Level int     `json:"level"`
					Xp    float64 `json:"xp"`
				} `json:"farming"`
				Fishing struct {
					Level int     `json:"level"`
					Xp    float64 `json:"xp"`
				} `json:"fishing"`
				Jeweling struct {
					Level int     `json:"level"`
					Xp    float64 `json:"xp"`
				} `json:"jeweling"`
				Mining struct {
					Level int `json:"level"`
					Xp    int `json:"xp"`
				} `json:"mining"`
				Scribing struct {
					Level int     `json:"level"`
					Xp    float64 `json:"xp"`
				} `json:"scribing"`
				Tailoring struct {
					Level int     `json:"level"`
					Xp    float64 `json:"xp"`
				} `json:"tailoring"`
				Weaponsmithing struct {
					Level int `json:"level"`
					Xp    int `json:"xp"`
				} `json:"weaponsmithing"`
				Woodcutting struct {
					Level int     `json:"level"`
					Xp    float64 `json:"xp"`
				} `json:"woodcutting"`
				Woodworking struct {
					Level int     `json:"level"`
					Xp    float64 `json:"xp"`
				} `json:"woodworking"`
			} `json:"professions"`
			Discoveries      int  `json:"discoveries"`
			EventsWon        int  `json:"eventsWon"`
			PreEconomyUpdate bool `json:"preEconomyUpdate"`
		} `json:"characters"`
		Guild struct {
			Name string `json:"name"`
			Rank string `json:"rank"`
		} `json:"guild"`
		Global struct {
			BlocksWalked    int `json:"blocksWalked"`
			ItemsIdentified int `json:"itemsIdentified"`
			MobsKilled      int `json:"mobsKilled"`
			TotalLevel      struct {
				Combat     int `json:"combat"`
				Profession int `json:"profession"`
				Combined   int `json:"combined"`
			} `json:"totalLevel"`
			Pvp struct {
				Kills  int `json:"kills"`
				Deaths int `json:"deaths"`
			} `json:"pvp"`
			Logins      int `json:"logins"`
			Deaths      int `json:"deaths"`
			Discoveries int `json:"discoveries"`
			EventsWon   int `json:"eventsWon"`
		} `json:"global"`
		Ranking struct {
			Guild  interface{} `json:"guild"`
			Player struct {
				Solo struct {
					Combat         interface{} `json:"combat"`
					Woodcutting    interface{} `json:"woodcutting"`
					Mining         interface{} `json:"mining"`
					Fishing        interface{} `json:"fishing"`
					Farming        interface{} `json:"farming"`
					Alchemism      interface{} `json:"alchemism"`
					Armouring      interface{} `json:"armouring"`
					Cooking        interface{} `json:"cooking"`
					Jeweling       interface{} `json:"jeweling"`
					Scribing       interface{} `json:"scribing"`
					Tailoring      interface{} `json:"tailoring"`
					Weaponsmithing interface{} `json:"weaponsmithing"`
					Woodworking    interface{} `json:"woodworking"`
					Profession     interface{} `json:"profession"`
					Overall        interface{} `json:"overall"`
				} `json:"solo"`
				Overall struct {
					All        interface{} `json:"all"`
					Combat     interface{} `json:"combat"`
					Profession interface{} `json:"profession"`
				} `json:"overall"`
			} `json:"player"`
			Pvp interface{} `json:"pvp"`
		} `json:"ranking"`
	} `json:"data"`
}

type Territory struct {
	Territory string      `json:"territory"`
	Guild     string      `json:"guild"`
	Acquired  string      `json:"acquired"`
	Attacker  interface{} `json:"attacker"`
}
