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

// record represents a single crop yield record.
type record struct {
	Year   int     `csv:"Rok"`
	Region string  `csv:"Region"`
	Crop   string  `csv:"Plodina"`
	Yield  float64 `csv:"Výnos (t/ha)"`
	Area   float64 `csv:"Plocha (ha)"`
}

var (
	regions = []string{
		"Jihočeský kraj",
		"Středočeský kraj",
		"Severočeský kraj",
		"Jihomoravský kraj",
		"Moravskoslezský kraj",
	}

	cropYield = map[string]float64{
		"Řepka":    3.2,
		"Pšenice":  6.1,
		"Mák":      0.9,
		"Brambory": 23,
	}
	yieldDev = 0.25

	areaMin, areaMax = 2.0, 100.0

	yearMax = time.Now().Year()
	yearMin = yearMax - 5
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

// createCSV generates a CSV file with randomly generated crop yield data.
func createCSV(filename string, numRows int) error {
	var records []record
	for i := 0; i < numRows; i++ {
		var r record

		r.Crop = generator.MustRandomChoice(keys(cropYield))

		ymin, ymax := yieldMinMax(cropYield[r.Crop], yieldDev)
		r.Yield = generator.RandomCentered(ymin, ymax, 3)
		r.Yield = math.Round(r.Yield*100) / 100

		r.Area = generator.RandomRange(areaMin, areaMax)
		r.Area = math.Round(r.Area*100) / 100

		r.Region = generator.MustRandomChoice(regions)

		r.Year = generator.RandomRange(yearMin, yearMax)

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
func keys(m map[string]float64) []string {
	var ret []string
	for k := range m {
		ret = append(ret, k)
	}
	return ret
}

// yieldMinMax calculates the minimum and maximum yield values based on a given yield and deviation.
func yieldMinMax(yield, dev float64) (min float64, max float64) {
	min = yield - (yield * dev)
	max = yield + (yield * dev)
	return
}
