package main

import (
	"fmt"
	"import/services"
)

func main() {
	records := services.ReadCsvFile("import_products_magento.csv")
	response := services.SendProducts(records)
	fmt.Println(response)
}
