package main

import (
	"fmt"
	"math"
	"os"

	"github.com/berybox/vyuka-fzt/KGZB_INFB/generatory_dat/pkg/cli"
	"github.com/berybox/vyuka-fzt/KGZB_INFB/generatory_dat/pkg/generator"
	"github.com/gocarina/gocsv"
)

// record represents a single cattle performance record.
type record struct {
	Region         string  `csv:"Region"`
	MilkPerf       float64 `csv:"Mléčná užitkovost"`
	MeatPerf       float64 `csv:"Masná užitkovost"`
	PopulationSize int     `csv:"Velikost populace"`
	Breed          string  `csv:"Plemeno"`
}

// breedPerf represents a single breed performance minimums and maximums.
type breedPerf struct {
	MilkMin float64
	MilkMax float64
	MeatMin float64
	MeatMax float64
}

var (
	// Minimum and maximum values for milk and meat performance
	breedPerfs = map[string]breedPerf{
		"Holštýnský skot": {
			MilkMin: 0.75,
			MilkMax: 1.25,
			MeatMin: 0.4,
			MeatMax: 0.8,
		},
		"Aberdeen Angus": {
			MilkMin: 0.4,
			MilkMax: 0.8,
			MeatMin: 0.75,
			MeatMax: 1.25,
		},
		"Červenostrakatý skot": {
			MilkMin: 0.6,
			MilkMax: 1.2,
			MeatMin: 0.8,
			MeatMax: 1.1,
		},
	}

	regions = []string{
		"Jihočeský kraj",
		"Středočeský kraj",
		"Severočeský kraj",
		"Jihomoravský kraj",
		"Moravskoslezský kraj",
	}

	// Minimum and maximum values for population size
	popMin, popMax = 5, 800
)

func main() {
	params, err := cli.BasicCzech()
	if err != nil {
		return
	}

	err = createCSV(params.Filename, params.NumRows)
	if err != nil {
		fmt.Fprintf(os.Stderr, "CHYBA: %s\n", err)
		return
	}

	fmt.Printf("Vytvořen nový soubor '%s' s %d náhodnými záznamy.\n", params.Filename, params.NumRows)
}

func createCSV(filename string, numRows int) error {
	var records []record
	for i := 0; i < numRows; i++ {
		var r record

		r.Breed = generator.MustRandomChoice(keys(breedPerfs))

		r.MilkPerf = generator.RandomCentered(breedPerfs[r.Breed].MilkMin, breedPerfs[r.Breed].MilkMax, 3)
		r.MilkPerf = math.Round(r.MilkPerf*100) / 100

		r.MeatPerf = generator.RandomCentered(breedPerfs[r.Breed].MeatMin, breedPerfs[r.Breed].MeatMax, 3)
		r.MeatPerf = math.Round(r.MeatPerf*100) / 100

		r.Region = generator.MustRandomChoice(regions)

		r.PopulationSize = generator.RandomRange(popMin, popMax)

		records = append(records, r)
	}

	file, err := os.OpenFile(filename, os.O_RDWR|os.O_CREATE, os.ModePerm)
	if err != nil {
		return err
	}

	err = gocsv.MarshalFile(&records, file)
	if err != nil {
		return err
	}

	return nil
}

// keys returns the keys of a map as a slice of strings.
func keys(m map[string]breedPerf) []string {
	var ret []string
	for k := range m {
		ret = append(ret, k)
	}
	return ret
}
