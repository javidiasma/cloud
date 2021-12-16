package DAL

import (
	"cloudPart1/config"
	"cloudPart1/csvReader"
	"cloudPart1/models/DAO"
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
)

func ReturnByRank(rank int) []DAO.History {
	var history []DAO.History
	config.DB.Where(&DAO.History{Rank: rank}).Find(&history)
	return history
}
func ReturnBy() {

}
func CreateDB(history DAO.History) DAO.History {
	tempHistory := DAO.History{
		Rank:         history.Rank,
		Name:         history.Name,
		Platform:     history.Platform,
		Year:         history.Year,
		Genre:        history.Genre,
		Publisher:    history.Publisher,
		NA_Sales:     history.NA_Sales,
		EU_Sales:     history.EU_Sales,
		JP_Sales:     history.JP_Sales,
		Other_Sales:  history.Other_Sales,
		Global_Sales: history.Global_Sales,
	}
	config.DB.Create(&tempHistory)
	return tempHistory
}

func CsvReader() (map[int]DAO.History, error) {
	returnMap := make(map[int]DAO.History)
	var count int
	csvfile, err := os.Open(csvReader.FilePath)
	if err != nil {

		return nil, fmt.Errorf(err.Error())
	}
	reader := csv.NewReader(csvfile)
	var header []string
	for {
		rawCSVData, err2 := reader.Read()
		if err2 == io.EOF {
			break
		}
		if err2 != nil {
			log.Fatal(err)
		}
		if header == nil {
			header = rawCSVData
		} else {
			//dict := map[string]string{}
			//for i := range header {
			//	dict[header[i]] = rawCSVData[i]
			//}
			rank, err := strconv.Atoi(rawCSVData[0])
			if err != nil {
				panic(err.Error())
			}
			year, err := strconv.Atoi(rawCSVData[3])
			if err != nil {
				//panic(err.Error())
			}
			naSales, err := strconv.ParseFloat(rawCSVData[6], 64)
			if err != nil {
				panic(err.Error())
			}
			euSales, err := strconv.ParseFloat(rawCSVData[7], 64)
			if err != nil {
				panic(err.Error())
			}
			jpSales, err := strconv.ParseFloat(rawCSVData[8], 64)
			if err != nil {
				panic(err.Error())
			}
			otherSales, err := strconv.ParseFloat(rawCSVData[9], 64)
			if err != nil {
				panic(err.Error())
			}
			globalSales, err := strconv.ParseFloat(rawCSVData[10], 64)
			if err != nil {
				panic(err.Error())
			}
			returnMap[count] = DAO.History{
				Rank:         rank,
				Name:         rawCSVData[1],
				Platform:     rawCSVData[2],
				Year:         year,
				Genre:        rawCSVData[4],
				Publisher:    rawCSVData[5],
				NA_Sales:     float32(naSales),
				EU_Sales:     float32(euSales),
				JP_Sales:     float32(jpSales),
				Other_Sales:  float32(otherSales),
				Global_Sales: float32(globalSales),
			}
			CreateDB(returnMap[count])
		}

	}
	return returnMap, err
}
