package readexcel

import (
	"fmt"
	"log"
	"regexp"
	"strconv"
	"strings"

	"github.com/xuri/excelize/v2"
)

var ItemNamesMap map[string]string = map[string]string{
	"walt id":     "oid",
	"název":       "name",
	"popis":       "description",
	"cena":        "price",
	"cena - obal": "price_package",
}

var re = regexp.MustCompile(`[0-9]+`)

var maxIndex int

type Item struct {
	OID          string
	Name         string
	Description  string
	Price        float64
	PricePackage float64
}

func GetExcelItems(filename string, worksheet string, itemNamesMap map[string]string) ([]Item, error) {
	res := []Item{}
	var oid, name, desc, price, price_pack int
	f, err := excelize.OpenFile(filename)
	if err != nil {
		return nil, err
	}
	defer func() {
		// Close the spreadsheet.
		if err := f.Close(); err != nil {
			fmt.Println(err)
		}
	}()
	rows, err := f.GetRows(worksheet)
	if err != nil {
		return nil, err
	}
	// first row contains field names
	var ok bool
	var i string
	for k, v := range rows[0] {
		if i, ok = itemNamesMap[v]; !ok {
			continue
		}
		switch i {
		case "oid":
			oid = k
			setMaxIndex(k)
		case "name":
			name = k
			setMaxIndex(k)
		case "description":
			desc = k
			setMaxIndex(k)
		case "price":
			price = k
			setMaxIndex(k)
		case "price_package":
			price_pack = k
			setMaxIndex(k)
		}

	}

	for _, v := range rows[1:] {
		if len(v) <= maxIndex {
			continue
		}
		i := Item{
			OID:          strings.TrimSpace(v[oid]),
			Name:         strings.TrimSpace(v[name]),
			Description:  strings.TrimSpace(v[desc]),
			Price:        getPrice(v[price]),
			PricePackage: getPrice(v[price_pack]),
		}
		switch {
		case i.OID == "":
			continue
		case i.Price < 0 || i.PricePackage < 0:
			continue
		case i.Description == "":
			continue

		}
		res = append(res, i)
	}

	return res, nil
}

func setMaxIndex(k int) {
	if k > maxIndex {
		maxIndex = k
	}
}

func getPrice(in string) float64 {
	r := re.FindAllString(in, 1)
	p, err := strconv.ParseFloat(r[0], 64)
	if err != nil {
		log.Println(err)
		return -1
	}
	return p
}
