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

	if resp.StatusCode != http.StatusOK {
		log.Printf("API retornou um status inesperado: %d", resp.StatusCode)
		return nil, fmt.Errorf("falha ao buscar produtos, status: %d", resp.StatusCode)
	}

	var result map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		log.Printf("Erro ao decodificar resposta: %v", err)
		return nil, err
	}

	items, ok := result["results"].([]interface{})
	if !ok {
		log.Printf("Formato inesperado no campo 'results'")
		return nil, fmt.Errorf("formato inesperado no campo 'results'")
	}

	var produtos []Produto
	for _, item := range items {
		data, ok := item.(map[string]interface{})
		if !ok {
			log.Printf("Formato inesperado no item do produto")
			continue
		}

		id, _ := data["id"].(string)
		title, _ := data["title"].(string)
		price, _ := data["price"].(float64)
		thumbnail, _ := data["thumbnail"].(string)
		permalink, _ := data["permalink"].(string)

		produtos = append(produtos, Produto{
			ID:        id,
			Title:     title,
			Price:     price,
			Thumbnail: thumbnail,
			Permalink: permalink,
		})
	}

	return produtos, nil
}
