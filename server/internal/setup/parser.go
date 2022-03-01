package setup

import (
	"bufio"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/fs"
	"io/ioutil"
	"os"
	"strings"

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
	ColumnIndexLongitude   = 37
	ColumnIndexLatitude    = 38
)

func GetCsvFiles() ([]fs.FileInfo, error) {
	files, err := ioutil.ReadDir(DefaultDataPath)

	if err != nil {
		return nil, newError("CSV 파일 조회 오류", err)
	}

	var csvFiles []fs.FileInfo

	for _, file := range files {
		if strings.HasSuffix(file.Name(), ".csv") {
			csvFiles = append(csvFiles, file)
		}
	}

	return csvFiles, nil
}

func SaveData(stars []*model.Star) error {
	file, err := os.Create(DefaultDataPath + "data.json")

	if err != nil {
		return newError("파일 생성 오류", err)
	}

	defer closeOrPanic(file)

	data, err := json.Marshal(stars)

	if err != nil {
		return newError("데이터 변환 오류", err)
	}

	if _, err := file.Write(data); err != nil {
		return newError("데이터 저장 오류", err)
	}

	return nil
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
		return nil, newError("CSV 파일 열기 오류", err)
	}

	defer closeOrPanic(file)

	var data []*model.Star

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		row := scanner.Text()
		if isCafe(row) && isStar(row) {
			data = append(data, parseStarFromRow(row))
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, newError("CSV 파일 읽기 오류", err)
	}

	return data, nil
}

func parseStarFromRow(row string) *model.Star {
	column := strings.Split(row, ",")
	return &model.Star{
		Code:        trimQuote(column[ColumnIndexCode]),
		Name:        trimQuote(column[ColumnIndexName]),
		BCode:       trimQuote(column[ColumnIndexBCode]),
		Address:     trimQuote(column[ColumnIndexAddress]),
		RoadAddress: trimQuote(column[ColumnIndexRoadAddress]),
		Longitude:   trimQuote(column[ColumnIndexLongitude]),
		Latitude:    trimQuote(column[ColumnIndexLatitude]),
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

func newError(message string, err error) error {
	text := fmt.Sprintf("%v: %v", message, err)
	return errors.New(text)
}

func closeOrPanic(closer io.Closer) func() {
	return func() {
		if err := closer.Close(); err != nil {
			text := fmt.Sprintf("Close error: %v", err)
			panic(text)
		}
	}
}
