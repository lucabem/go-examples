package CsvUtils

import (
	"bytes"
	"encoding/csv"
	"fmt"
	"os"
	"strings"

	"github.com/tealeg/xlsx"

	StringUtils "myApp/src/utils/strings"
)

func ParseXLSXToCSV(objReader *bytes.Reader, columns []string, hasHeader bool) ([][]string, error) {

	var data [][]string

	xlFile, err := xlsx.OpenReaderAt(objReader, int64(objReader.Len()))
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	sheet := xlFile.Sheets[0]

	for idx, row := range sheet.Rows {
		if idx == 0 && hasHeader {
			data = append(data, columns)
			continue
		}

		var rowData []string
		for _, cell := range row.Cells {
			cellValue, err := cell.GeneralNumericWithoutScientific()
			cellValue = strings.Trim(cellValue, " ")
			if err == nil && cellValue != "" {
				cellValue = strings.ReplaceAll(cellValue, ".", ",")
			}

			rowData = append(rowData, cellValue)
		}

		if !StringUtils.IsRowEmpty(rowData) {
			data = append(data, rowData)
		}
	}

	return data, err
}

func CreateCSVToFile(data [][]string, filename string) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	writer.Comma = ';'

	for _, row := range data {
		err := writer.Write(row)
		if err != nil {
			return err
		}
	}

	writer.Flush()

	if err := writer.Error(); err != nil {
		return err
	}

	return nil
}
