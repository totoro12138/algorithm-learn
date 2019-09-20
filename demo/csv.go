package main

import (
	"archive/zip"
	"bytes"
	"encoding/csv"
	"fmt"
	"net/http"
	"strconv"
	"time"
)

func main() {
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Println(time.Now().Format("15:04:05"))
		zipBuffer := bytes.NewBuffer(nil)
		zipWriter := zip.NewWriter(zipBuffer)
		for i := 0; i < 10; i++ {
			fmt.Printf("zip: %v.csv\n", i+1)
			csvBuffer := Csv()
			csvFile, err := zipWriter.Create(strconv.Itoa(i+1) + ".csv")
			if err != nil {
				panic(err)
			}

			_, err = csvFile.Write(csvBuffer.Bytes())
			if err != nil {
				panic(err)
			}
		}
		if err := zipWriter.Close(); err != nil {
			panic(err)
		}

		writer.Header().Add("Content-Type", "application/form-data")
		writer.Header().Add("Content-Disposition", fmt.Sprintf(`attachment; filename="%v.zip"`, time.Now().Unix()))
		_, err := zipBuffer.WriteTo(writer)
		if err != nil {
			panic(err)
		}
	})

	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}

func Csv() *bytes.Buffer {
	csvBuffer := bytes.NewBuffer(nil)
	csvWriter := csv.NewWriter(csvBuffer)
	if err := csvWriter.WriteAll(getData()); err != nil {
		panic(err)
	}

	return csvBuffer
}

type SQLValue struct {
	ID       uint
	CreateAt time.Time
	UpdateAt time.Time
	DeleteAt *time.Time

	Username string
	Sex      int // 1.男 , 2.女 , 3.未知
	Age      int // year
	Height   int // cm
}

func getData() [][]string {
	sqlValue := make([]*SQLValue, 1000000)
	for i := range sqlValue {
		sqlValue[i] = &SQLValue{
			ID:       uint(i),
			CreateAt: time.Now(),
			UpdateAt: time.Now(),
			DeleteAt: nil,
			Username: strconv.Itoa(i),
			Sex:      3,
			Age:      20,
			Height:   170,
		}
	}

	data := make([][]string, len(sqlValue))
	data[0] = []string{
		"ID", "CreateAt", "UpdateAt", "DeleteAt", "Username", "Sex", "Age", "Height",
	}

	// 此处可改为数据库查询数据
	for i, v := range sqlValue {
		if i == 0 {
			continue
		}
		id := strconv.Itoa(int(v.ID))
		createAt := v.CreateAt.Format("2006-01-02 15:04:05")
		updateAt := v.UpdateAt.Format("2006-01-02 15:04:05")
		deleteAt := ""
		if v.DeleteAt != nil {
			deleteAt = v.DeleteAt.Format("2006-01-02 15:04:05")
		}
		sex := ""
		switch v.Sex {
		case 1:
			sex = "男"
		case 2:
			sex = "女"
		case 3:
			sex = "未知"
		default:
		}
		age := strconv.Itoa(v.Age)
		height := strconv.Itoa(v.Height)
		data[i] = []string{
			id,
			createAt,
			updateAt,
			deleteAt,
			v.Username,
			sex,
			age,
			height,
		}
	}
	return data
}
