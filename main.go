package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/polisko/rosnagen/readexcel"
	"github.com/polisko/rosnagen/wolt"
)

func main() {
	excelFile := flag.String("e", "Polední menu-nový.xlsx", "Excel s novym menu")
	woltFile := flag.String("o", "restaurace-rosnicka.json", "Export file z Woltu")
	woltFileNew := flag.String("n", "restaurace-rosnicka_novy.json", "Novy file pro import")
	excelSheet := flag.String("s", "List 1", "Nazev excel sheetu s menu")

	flag.Parse()

	e, err := readexcel.GetExcelItems(*excelFile, *excelSheet, readexcel.ItemNamesMap)
	if err != nil {
		log.Fatal(err)
	}
	w, err := wolt.NewWoltMenu(*woltFile)
	if err != nil {
		log.Fatal(err)
	}
	w.UpdateByExcel(e)
	err = w.SaveToFile(*woltFileNew)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Novy soubor pro export ulozen do %q\n", *woltFileNew)
}
