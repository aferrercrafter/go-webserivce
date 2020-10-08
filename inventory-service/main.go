package main

import (
	"encoding/json"
	"log"
	"net/http"
)

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

	}
}

func main() {
	http.HandleFunc("/products", productsHandler)
	http.ListenAndServe(":5000", nil)
}
