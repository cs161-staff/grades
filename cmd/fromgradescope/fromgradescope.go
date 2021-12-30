package main

import (
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"io"
	"os"
	"strconv"

	"github.com/cs161-staff/grades"
)

func panicIfErr(err error) {
	if err != nil {
		panic(err)
	}
}

// importCategories imports and returns the categories described in the CSV at
// the given path.
func importCategories(path string) map[string]*grades.Category {
	reader, err := NewDictReaderFromPath(path)
	panicIfErr(err)

	categories := make(map[string]*grades.Category)
	for row, err := reader.Read(); err != io.EOF; row, err = reader.Read() {
		panicIfErr(err)

		name := row["Name"]
		weight, err := strconv.ParseFloat(row["Weight"], 64)
		panicIfErr(err)
		hasLateMultiplier, err := strconv.ParseBool(row["Has Late Multiplier"])
		panicIfErr(err)
		drops64, err := strconv.ParseInt(row["Drops"], 10, 64)
		panicIfErr(err)
		drops := int(drops64)
		slipDays64, err := strconv.ParseInt(row["Drops"], 10, 32)
		panicIfErr(err)
		slipDays := int(slipDays64)
		if _, ok := categories[name]; ok {
			panic(errors.New("Duplicate category specified in imported CSV: " + name))
		}
		categories[name] = &grades.Category{
			Name:              name,
			Weight:            weight,
			HasLateMultiplier: hasLateMultiplier,
			Drops:             drops,
			SlipDays:          slipDays,
		}
	}

	return categories
}

// importAssignments imports and returns the assignments described in the CSV
// at the given path.
func importAssignments(path string, categories map[string]*grades.Category) map[string]*grades.Assignment {
	reader, err := NewDictReaderFromPath(path)
	panicIfErr(err)

	assignments := make(map[string]*grades.Assignment)
	for row, err := reader.Read(); err != io.EOF; row, err = reader.Read() {
		panicIfErr(err)
		name := row["Name"]
		category := row["Category"]
		maxScore, err := strconv.ParseFloat(row["Possible"], 64)
		panicIfErr(err)
		weight, err := strconv.ParseFloat(row["Weight"], 64)
		panicIfErr(err)
		slipGroup64, err := strconv.ParseInt(row["Slip Group"], 10, 64)
		panicIfErr(err)
		slipGroup := int(slipGroup64)
		if _, ok := assignments[name]; ok {
			panic(errors.New(fmt.Sprintf("Duplicate assignment specified in imported CSV: %s", name)))
		}
		if _, ok := categories[category]; !ok {
			panic(errors.New(fmt.Sprintf("Assignment %s references unknown category %s", name, category)))
		}
		assignments[name] = &grades.Assignment{
			Name:         name,
			CategoryName: category,
			MaxScore:     maxScore,
			Weight:       weight,
			SlipGroup:    slipGroup,
		}
	}

	return assignments
}

func main() {
	// Mandatory args.
	var rosterPath string
	var gradesPath string
	var categoriesPath string
	var assignmentsPath string
	flag.StringVar(&rosterPath, "roster", "", "CSV roster downloaded from CalCentral")
	flag.StringVar(&gradesPath, "grades", "", "CSV grades downloaded from Gradescope")
	flag.StringVar(&categoriesPath, "categories", "", "CSV with assignment categories")
	flag.StringVar(&assignmentsPath, "assignments", "", "CSV with assignments")

	// Optional args.
	var overridesPath string
	var clobbersPath string
	var extensionsPath string
	var accommodationsPath string
	var rounding int
	var outputPath string
	flag.StringVar(&overridesPath, "overrides", "", "CSV with score overrides")
	flag.StringVar(&clobbersPath, "clobbers", "", "CSV with clobbers")
	flag.StringVar(&extensionsPath, "extensions", "", "CSV with extensions")
	flag.StringVar(&accommodationsPath, "accommodations", "", "CSV with accommodations for drops and slip days")
	flag.IntVar(&rounding, "round", 0, "Number of decimal places to round to")
	flag.StringVar(&outputPath, "output", "", "Output CSV file")

	flag.Parse()

	if rosterPath == "" || gradesPath == "" || categoriesPath == "" || assignmentsPath == "" {
		flag.Usage()
		os.Exit(1)
	}

	categories := importCategories(categoriesPath)
	assignments := importAssignments(assignmentsPath, categories)

	out, _ := json.MarshalIndent(assignments, "", "  ")
	fmt.Println(string(out))
}
