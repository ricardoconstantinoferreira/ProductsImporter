package services

import (
	"bytes"
	"encoding/csv"
	"encoding/json"
	"fmt"
	"import/gateway"
	"log"
	"os"

	"strings"
)

const CREDENTIALS = "credentials.txt"

type CategoryLinks struct {
	Position   int    `json:"position"`
	CategoryId string `json:"category_id"`
}

func createCategoryLinks(category string) []string {
	categories := strings.Split(category, ",")
	result := make([]string, len(categories))

	for key := range categories {

		cat := CategoryLinks{key, categories[key]}
		b, err := json.Marshal(cat)

		if err != nil {
			fmt.Println(err)
		}
		result[key] = string(b)
	}

	return result
}

func SendProducts(records [][]string) []string {

	messages := make([]string, len(records))

	for i := 0; i < len(records); i++ {
		if i == 1 {
			category_links := createCategoryLinks(records[i][8])
			categorys := strings.Join(category_links, ",") + ","
			params, _ := json.Marshal(map[string]string{
				"sku":              records[i][0],
				"name":             records[i][1],
				"attribute_set_id": records[i][2],
				"price":            records[i][3],
				"visibility":       records[i][4],
				"type_id":          records[i][5],
				"weight":           records[i][6],
				"extension_attributes": `{
					"stock_item": {
						"qty":         ` + records[i][9] + `,
						"is_in_stock": ` + records[i][10] + `,
					},
					"category_links": [
						` + categorys + `
					],
				}`,
			})

			payload := bytes.NewBuffer(params)

			token := getToken()
			url := ReadCredentials(CREDENTIALS, "product")
			response := gateway.Send(payload, url, token)

			fmt.Println("____________________________________________________________________________________________________________________________________________________________")
			messages[i] = response
		}
	}

	return messages
}

func getToken() string {
	login := ReadCredentials(CREDENTIALS, "login")
	pass := ReadCredentials(CREDENTIALS, "pass")
	url := ReadCredentials(CREDENTIALS, "admin")

	params, _ := json.Marshal(map[string]string{
		"username": login,
		"password": pass,
	})

	payload := bytes.NewBuffer(params)
	return gateway.TokenAdmin(payload, url)
}

func ReadCsvFile(filePath string) [][]string {
	f, err := os.Open(filePath)
	if err != nil {
		log.Fatal("Error to open input file "+filePath, err)
	}
	defer f.Close()

	csvReader := csv.NewReader(f)
	records, err := csvReader.ReadAll()
	if err != nil {
		log.Fatal("Error to read and parse as CSV file "+filePath, err)
	}

	return records
}

func ReadCredentials(filePath string, typeFile string) string {

	var txt string
	result := ReadCsvFile(filePath)

	if typeFile == "login" {
		txt = strings.Split(result[0][0], ": ")[1]
	}

	if typeFile == "pass" {
		txt = strings.Split(result[1][0], ": ")[1]
	}

	if typeFile == "product" {
		txt = strings.Split(result[2][0], ": ")[1]
	}

	if typeFile == "admin" {
		txt = strings.Split(result[3][0], ": ")[1]
	}

	return txt
}
