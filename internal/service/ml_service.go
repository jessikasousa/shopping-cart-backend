package service

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

const API_BASE_URL = "https://api.mercadolibre.com"

type Produto struct {
	ID        string  `json:"id"`
	Title     string  `json:"title"`
	Price     float64 `json:"price"`
	Thumbnail string  `json:"thumbnail"`
	Permalink string  `json:"permalink"`
}

func SearchProducts(query string) ([]Produto, error) {
	url := fmt.Sprintf("%s/sites/MLB/search?q=%s", API_BASE_URL, query)
	resp, err := http.Get(url)
	if err != nil {
		log.Printf("Erro ao buscar produtos: %v", err)
		return nil, err
	}
	defer resp.Body.Close()

	var result map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		log.Printf("Erro ao decodificar resposta: %v", err)
		return nil, err
	}

	items := result["results"].([]interface{})
	var produtos []Produto
	for _, item := range items {
		data := item.(map[string]interface{})
		produtos = append(produtos, Produto{
			ID:        data["id"].(string),
			Title:     data["title"].(string),
			Price:     data["price"].(float64),
			Thumbnail: data["thumbnail"].(string),
			Permalink: data["permalink"].(string),
		})
	}

	return produtos, nil
}
