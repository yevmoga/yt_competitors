package main

import (
	"bytes"
	"encoding/json"
	"math"
	"net/http"
)

// Структура для передачі слова
type WordRequest struct {
	Word string `json:"word"`
}

// Структура для отримання вектору
type VectorResponse struct {
	Vector []float64 `json:"vector"`
}

func getWordVector(word string) ([]float64, error) {
	requestBody, err := json.Marshal(WordRequest{Word: word})
	if err != nil {
		return nil, err
	}

	resp, err := http.Post("http://localhost:5000/vectorize", "application/json", bytes.NewBuffer(requestBody))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var response VectorResponse
	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		return nil, err
	}
	return response.Vector, nil
}

// averageVector обчислює середнє значення векторів
func averageVector(vectors [][]float64) []float64 {
	if len(vectors) == 0 {
		return nil
	}
	dimension := len(vectors[0])
	avg := make([]float64, dimension)
	for _, vec := range vectors {
		for i := range vec {
			avg[i] += vec[i]
		}
	}
	for i := range avg {
		avg[i] /= float64(len(vectors))
	}
	return avg
}

// cosineSimilarity обчислює косинусну схожість між двома векторами
func cosineSimilarity(vecA, vecB []float64) float64 {
	var dotProduct, normA, normB float64
	for i := range vecA {
		dotProduct += vecA[i] * vecB[i]
		normA += vecA[i] * vecA[i]
		normB += vecB[i] * vecB[i]
	}
	if normA == 0 || normB == 0 {
		return 0.0
	}
	return dotProduct / (math.Sqrt(normA) * math.Sqrt(normB))
}

// vectorizeText обчислює середній вектор для тексту (назви або тегів)
func vectorizeText(words []string) ([]float64, error) {
	var vectors [][]float64
	for _, word := range words {
		vec, err := getWordVector(word)
		if err != nil {
			return nil, err
		}
		vectors = append(vectors, vec)
	}
	return averageVector(vectors), nil
}

// calculateSimilarity порівнює назви та теги двох відео
func calculateSimilarity(titleA, titleB, tagsA, tagsB []string) (float64, float64, error) {
	titleVecA, err := vectorizeText(titleA)
	if err != nil {
		return 0.0, 0.0, err
	}
	titleVecB, err := vectorizeText(titleB)
	if err != nil {
		return 0.0, 0.0, err
	}
	tagsVecA, err := vectorizeText(tagsA)
	if err != nil {
		return 0.0, 0.0, err
	}
	tagsVecB, err := vectorizeText(tagsB)
	if err != nil {
		return 0.0, 0.0, err
	}

	titleSimilarity := cosineSimilarity(titleVecA, titleVecB)
	tagsSimilarity := cosineSimilarity(tagsVecA, tagsVecB)

	return titleSimilarity, tagsSimilarity, nil
}
