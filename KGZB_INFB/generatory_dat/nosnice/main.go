package main

import (
	"fmt"
	"math"
	"os"
	"time"

	"github.com/berybox/vyuka-fzt/KGZB_INFB/generatory_dat/pkg/cli"
	"github.com/berybox/vyuka-fzt/KGZB_INFB/generatory_dat/pkg/generator"
	"github.com/gocarina/gocsv"
)

// record represents a single cattle performance record.
type record struct {
	ID     int     `csv:"ID"`
	Year   int     `csv:"Rok"`
	Breed  string  `csv:"Plemeno"`
	Hall   string  `csv:"Hala"`
	Laying int     `csv:"Snáška (vajec / rok)"`
	Period int     `csv:"Doba snášky (týdnů)"`
	Weight float64 `csv:"Hmotnost (kg)"`
}

// breedPerf represents a single breed performance minimums and maximums.
type breedPerf struct {
	LayingMin int
	LayingMax int
	WeightMin float64
	WeightMax float64
}

var (
	// Minimum and maximum values for laying and laying period
	breedPerfs = map[string]breedPerf{
		"Rhode Island Red": {
			LayingMin: 200,
			LayingMax: 300,
			WeightMin: 2.7,
			WeightMax: 3.6,
		},
		"Plymouth": {
			LayingMin: 250,
			LayingMax: 290,
			WeightMin: 2.7,
			WeightMax: 3.2,
		},
		"Sussex": {
			LayingMin: 240,
			LayingMax: 300,
			WeightMin: 3.2,
			WeightMax: 3.65,
		},
		"Leghorn": {
			LayingMin: 280,
			LayingMax: 320,
			WeightMin: 2,
			WeightMax: 2.5,
		},
	}

	halls = []string{
		"Kojetín",
		"Hlinsko",
		"Ivančice",
		"Býškovice",
		"Slavíč",
	}

	// Minimum and maximum values for laying period
	periodMin, periodMax = 40, 60

	yearMax = time.Now().Year()
	yearMin = yearMax - 10
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
		r := record{ID: i}

		r.Year = generator.RandomRange(yearMin, yearMax)

		r.Breed = generator.MustRandomChoice(keys(breedPerfs))

		r.Hall = generator.MustRandomChoice(halls)

		r.Laying = generator.RandomCentered(breedPerfs[r.Breed].LayingMin, breedPerfs[r.Breed].LayingMax, 3)

		r.Period = generator.RandomRange(periodMin, periodMax)

		r.Weight = generator.RandomCentered(breedPerfs[r.Breed].WeightMin, breedPerfs[r.Breed].WeightMax, 3)
		r.Weight = math.Round(r.Weight*100) / 100

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
