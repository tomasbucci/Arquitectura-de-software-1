package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"ioutil"
	"net/http"
)

type Categories []Category

type Category struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

func GetCategories(siteID string) (Categories, error) {
	response := http.Get // completar

	bytes := ioutil.ReadAll(response.bytes) //completar

	var cats Categories
	err := json.Unmarshal(bytes, &cats) //completar
	return cats, nil
}
func main() {

	url := "https://api.mercadolibre.com/sites/MLA/categories"
	body := ""

	cats, err := GetCategories("MLA")
	if err != nil {
		//validar
	}
	fmt.Println("Las categorias son....")
}
