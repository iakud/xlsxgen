package main

import (
	"io/fs"
	"io/ioutil"
	"log"
	"path/filepath"

	"github.com/gobeam/stringy"
	"github.com/tealeg/xlsx"
)

const (
	kXlsxExt = "xlsx"
)

func parseDir(dirPath string) {
	filepath.Walk(dirPath, func(path string, info fs.FileInfo, err error) error {
		return nil
	})
	filepath.WalkDir(dirPath, func(path string, d fs.DirEntry, err error) error {
		if d.IsDir() {
			return nil
		}
		if filepath.Ext(d.Name()) != kXlsxExt {
			return nil
		}
		fileName := filepath.Join(path, d.Name())
		parseXlsx(fileName)
		return nil
	})
}

func parseXlsx(fileName string) {
	xlFile, err := xlsx.OpenFile(fileName)
	if err != nil {
		log.Fatal(err)
	}

	for _, sheet := range xlFile.Sheets {
		parseSheet(sheet)
	}
}

func parseSheet(sheet *xlsx.Sheet) {
	for _, row := range sheet.Rows {
		_ = row
	}
	sheet.Row(1)
	// var jsonFile string = sheet.Name + ".json"
	// jsonFile = "client/assets/resources/data/" + jsonFile
	filename := filepath.Join(outDir, sheet.Name+".json")
	log.Println(filename)
	if err := ioutil.WriteFile(filename, nil, 0666); err != nil {
		log.Println(err)
	}

}

func parseField(rows []*xlsx.Row, maxCol int) ([]*Field, error) {
	titleRow := rows[0]
	nameRow := rows[1]
	// typeRow := rows[2]
	var fields []*Field
	for i := 0; i < maxCol; i++ {
		var field Field
		name := stringy.New(nameRow.Cells[i].Value).CamelCase()
		name = ToPascal(nameRow.Cells[i].Value)
		// name := nameRow.Cells[i].Value
		if len(name) == 0 {
			continue
		}

		field.Idx = i
		field.Name = name
		field.Title = titleRow.Cells[i].Value
		// field.Type = typeRow.Cells[i].Value
		fields = append(fields, &field)
	}

	return fields, nil
}
