package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"

	. "../helpers"
	. "../models"
)

var productStore = make(map[string]Product)

var id int = 0

//HTTP post -/api/prouduct
func PostProductHandler(w http.ResponseWriter, r *http.Request) {
	var product Product
	err := json.NewDecoder(r.Body).Decode(&product)
	fmt.Printf("Decoder aşaması", err.Error())
	CheckError(err)
	product.CreatedOn = time.Now()
	id++
	product.ID = id
	key := strconv.Itoa(id)
	productStore[key] = product

	data, err := json.Marshal(product)
	fmt.Println("json aşaması")
	CheckError(err)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

	w.Write(data)

}

//HTTP get's -/api/prouduct
func GetProductsHandler(w http.ResponseWriter, r *http.Request) {
	var products []Product

	for _, product := range productStore {
		products = append(products, product)
	}

	data, err := json.Marshal(products)
	CheckError(err)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	w.Write(data)
}

//HTTP get -/api/prouduct/{id}
func GetProductHandler(w http.ResponseWriter, r *http.Request) {
	var product Product
	vars := mux.Vars(r) //request içindeki variable 'ları alır..
	key, _ := strconv.Atoi(vars["id"])
	for _, prd := range productStore {
		if prd.ID == key {
			product = prd
		}
	}

	data, err := json.Marshal(product)
	CheckError(err)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	w.Write(data)
}

//HTTP put -/api/prouduct/{id}
func PutProductHandler(w http.ResponseWriter, r *http.Request) {

	var err error
	vars := mux.Vars(r)
	key := vars["id"]

	var prdUpd Product
	err = json.NewDecoder(r.Body).Decode(&prdUpd)
	CheckError(err)

	if _, ok := productStore[key]; ok {
		prdUpd.ID, _ = strconv.Atoi(key)
		prdUpd.ChangedOn = time.Now()
		delete(productStore, key)
		productStore[key] = prdUpd
	} else {

		log.Printf("Değer bulunamadı : %s", key)

	}

	w.WriteHeader(http.StatusOK)

}

//HTTP delete -/api/prouduct/{id}
func DeleteProductHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["id"]

	if _, ok := productStore[key]; ok {
		delete(productStore, key)
	} else {
		log.Printf("Değer bulunamadı : %s", key)
	}
	w.WriteHeader(http.StatusOK)
}
