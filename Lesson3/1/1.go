package main

import "fmt"

type Car struct {
	Brand       string
	Model       string
	Year        int
	TrunkVolume int
	TrunkFull   int
	BulkSize    int
	OpenWindows bool
	Started     bool
}

func main() {
	sedan := Car{
		Brand:       "Renault",
		Model:       "Klio",
		Year:        2003,
		TrunkVolume: 100,
		TrunkFull:   75,
		BulkSize:    3000,
		OpenWindows: true,
		Started:     false,
	}
	truck := Car{
		Brand:       "Renault",
		Model:       "Kangoo",
		Year:        2013,
		TrunkVolume: 4598,
		TrunkFull:   1500,
		BulkSize:    3000,
		OpenWindows: false,
		Started:     true,
	}
	fmt.Printf("%+v\n", truck)
	fmt.Printf("%+v", sedan)
}
