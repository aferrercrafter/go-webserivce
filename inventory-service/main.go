package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

//Product entity
type Product struct {
	ProductID      int    `json:"productId"`
	Manufacturer   string `json:"manufacturer"`
	Sku            string `json:"ski"`
	Upc            string `json:"upc"`
	PricePerUnit   string `json:"pricePerUnit"`
	QuantityOnHand int    `json:"quantityOnHand"`
	ProductName    string `json:"productName"`
}

var productList []Product

func init() {
	productsJSON := `[
		{
			"productId": 1,
			"manufacturer": "Johns-Jenkins",
			"sku": "p5z343vds",
			"upc": "13123234324",
			"pricePerUnit": "497.45",
			"quantityOnHand": 9703,
			"productName": "sticky note"
			},
		{
			"productId": 2,
			"manufacturer": "Hessel",
			"sku": "p5z343ass",
			"upc": "13122223324",
			"pricePerUnit": "197.45",
			"quantityOnHand": 703,
			"productName": "leg warmeres"
		},
		{
			"productId": 3,
			"manufacturer": "Swaniawski",
			"sku": "p5z3ff34s",
			"upc": "1312221239",
			"pricePerUnit": "457.45",
			"quantityOnHand": 403,
			"productName": "lamp shade"
		}	
	]`
	err := json.Unmarshal([]byte(productsJSON), &productList)
	if err != nil {
		log.Fatal(err)
	}
}

func getNextID() int {
	highestID := 1
	for _, product := range productList {
		if highestID < product.ProductID {
			highestID = product.ProductID
		}
	}
	return highestID + 1
}

func productsHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		productsJSON, err := json.Marshal(productList)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write(productsJSON)
	case http.MethodPost:
		var newProduct Product
		bodyBytes, err := ioutil.ReadAll(r.Body)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		err = json.Unmarshal(bodyBytes, &newProduct)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		if newProduct.ProductID != 0 {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		newProduct.ProductID = getNextID()
		productList = append(productList, newProduct)
		w.WriteHeader(http.StatusCreated)
		return
	default:

	}
}

func main() {
	http.HandleFunc("/products", productsHandler)
	http.ListenAndServe(":5000", nil)
}
