package setup

import (
	"bufio"
	"io/fs"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/hashfunc/project-milky-way/internal"
	"github.com/hashfunc/project-milky-way/internal/model"
)

const (
	DefaultDataPath = ".data/"
)

const (
	ColumnIndexCode        = 0
	ColumnIndexName        = 1
	ColumnIndexBCode       = 17
	ColumnIndexAddress     = 24
	ColumnIndexRoadAddress = 31
)

func GetCsvFiles() ([]fs.FileInfo, error) {
	files, err := ioutil.ReadDir(DefaultDataPath)

	if err != nil {
		return nil, internal.NewError("CSV 파일 조회 오류", err)
	}

	var csvFiles []fs.FileInfo

	for _, file := range files {
		if strings.HasSuffix(file.Name(), ".csv") {
			csvFiles = append(csvFiles, file)
		}
	}

	return csvFiles, nil
}

func GetStarsFromFiles(files []fs.FileInfo) ([]*model.Star, error) {
	var ret []*model.Star

	for _, file := range files {
		stars, err := parseStarsFromFile(DefaultDataPath + file.Name())

		if err != nil {
			return nil, err
		}

		ret = append(ret, stars...)
	}

	return ret, nil
}

func parseStarsFromFile(filePath string) ([]*model.Star, error) {
	file, err := os.Open(filePath)

	if err != nil {
		return nil, internal.NewError("CSV 파일 열기 오류", err)
	}

	defer internal.CloseOrPanic(file)

	var data []*model.Star

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		row := scanner.Text()
		if isCafe(row) && isStar(row) {
			data = append(data, parseStarFromRow(row))
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, internal.NewError("CSV 파일 읽기 오류", err)
	}

	return data, nil
}

func parseStarFromRow(row string) *model.Star {
	column := strings.Split(row, ",")

	longitude, err := strconv.ParseFloat(trimQuote(column[len(column)-2]), 64)

	if err != nil {
		log.Panic(err)
	}

	latitude, err := strconv.ParseFloat(trimQuote(column[len(column)-1]), 64)

	if err != nil {
		println(row)
		log.Panic(err)
	}

	return &model.Star{
		Code:        trimQuote(column[ColumnIndexCode]),
		Name:        trimQuote(column[ColumnIndexName]),
		BCode:       trimQuote(column[ColumnIndexBCode]),
		Address:     trimQuote(column[ColumnIndexAddress]),
		RoadAddress: trimQuote(column[ColumnIndexRoadAddress]),
		Longitude:   longitude,
		Latitude:    latitude,
		Point: &model.GeoJSON{
			Type:        "Point",
			Coordinates: []interface{}{longitude, latitude},
		},
	}
}

func trimQuote(target string) string {
	return strings.Trim(target, "\"")
}

func isStar(row string) bool {
	return strings.Contains(row, "스타벅스")
}

func isCafe(row string) bool {
	return strings.Contains(row, "커피점/카페")
}
