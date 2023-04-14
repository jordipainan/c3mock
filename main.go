package main

import (
	"encoding/json"
	"net/http"
	"strings"
)

type Token struct {
	Id            string `json:"id"`
	Name          string `json:"name"`
	Type          string `json:"type"`
	CreationBlock int    `json:"creation_block"`
}

type TokensResponse struct {
	Tokens []Token `json:"tokens"`
}

type SupportedTokensResponse struct {
	SupportedTokens []string `json:"supported_tokens"`
}

type TokenDetailsResponse struct {
	Id            string          `json:"id"`
	Type          string          `json:"type"`
	Decimals      int             `json:"decimals"`
	Symbol        string          `json:"symbol"`
	CreationBlock int             `json:"creation_block"`
	Name          string          `json:"name"`
	TotalSupply   string          `json:"total_supply"`
	Status        TokenStatusInfo `json:"status"`
}

type TokenStatusInfo struct {
	Synced   bool `json:"synced"`
	AtBlock  int  `json:"at_block"`
	Progress int  `json:"progress"`
}

type Strategy struct {
	Id        int    `json:"id"`
	Predicate string `json:"predicate"`
}

type StrategiesResponse struct {
	Strategies []Strategy `json:"strategies"`
}

type TokenInfo struct {
	Id         string `json:"id"`
	Name       string `json:"name"`
	MinBalance string `json:"minimum_balance"`
	Method     string `json:"method"`
}

type StrategyDetailsResponse struct {
	Tokens    []TokenInfo `json:"tokens"`
	Predicate string      `json:"predicate"`
}

type CensusRequest struct {
	StrategyId  string `json:"strategy_id"`
	BlockNumber int    `json:"block_number"`
}

type CensusResponse struct {
	CensusId string `json:"census_id"`
}

type CensusDetailsResponse struct {
	Root        string `json:"root"`
	BlockNumber int    `json:"block_number"`
	Uri         string `json:"uri"`
}

func tokenDetailsHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	id := strings.TrimPrefix(r.URL.Path, "/api/tokens/")

	response := TokenDetailsResponse{
		Id:            id,
		Type:          "wANT",
		Decimals:      18,
		Symbol:        "wANT",
		CreationBlock: 123456,
		Name:          "Aragon Wrapped ANT",
		TotalSupply:   "123456",
		Status: TokenStatusInfo{
			Synced:   false,
			AtBlock:  12345,
			Progress: 87,
		},
	}

	json.NewEncoder(w).Encode(response)
}

func tokensHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	response := TokensResponse{
		Tokens: []Token{
			{
				Id:            "0x5713c2d9e9d4381bff966b1cdbf52cf4e8addc2c",
				Name:          "Aragon Wrapped ANT",
				Type:          "wANT",
				CreationBlock: 123456,
			},
		},
	}

	json.NewEncoder(w).Encode(response)
}

func tokensTypesHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	response := SupportedTokensResponse{
		SupportedTokens: []string{
			"erc20", "erc721", "erc777",
			"erc1155", "nation3", "wANT",
		},
	}

	json.NewEncoder(w).Encode(response)
}

func strategiesHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	response := StrategiesResponse{
		Strategies: []Strategy{
			{
				Id:        0,
				Predicate: "wANT AND USDC",
			},
			{
				Id:        1,
				Predicate: "wANT AND USDC OR DAI",
			},
		},
	}

	json.NewEncoder(w).Encode(response)
}

func strategyDetailsHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	response := StrategyDetailsResponse{
		Tokens: []TokenInfo{
			{
				Id:         "0x5713c2d9e9d4381bff966b1cdbf52cf4e8addc2c",
				Name:       "Wrapped Aragon Network Token",
				MinBalance: "10000",
				Method:     "balanceOfAt",
			},
			{
				Id:         "0xa0b86991c6218b36c1d19d4a2e9eb0ce3606eb48",
				Name:       "Circle USD",
				MinBalance: "20000",
				Method:     "balanceOf",
			},
			{
				Id:         "0xa117000000f279D81A1D3cc75430fAA017FA5A2e",
				Name:       "Aragon Network Token",
				MinBalance: "1",
				Method:     "balanceOf",
			},
		},
		Predicate: "(wANT OR ANT) AND USDC",
	}

	json.NewEncoder(w).Encode(response)
}

func strategiesByTokenHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	response := StrategiesResponse{
		Strategies: []Strategy{
			{
				Id:        0,
				Predicate: "wANT AND USDC",
			},
			{
				Id:        1,
				Predicate: "wANT AND USDC OR DAI",
			},
		},
	}

	json.NewEncoder(w).Encode(response)
}

func censusHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	var censusRequest CensusRequest
	err := json.NewDecoder(r.Body).Decode(&censusRequest)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	response := CensusResponse{
		CensusId: "fb2aa2aadaecbd68395f20f06f6a59980e6c42a1eb4db4ee1ba7210e0a672608",
	}

	json.NewEncoder(w).Encode(response)
}

func censusDetailsHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	response := CensusDetailsResponse{
		Root:        "fb2aa2aadaecbd68395f20f06f6a59980e6c42a1eb4db4ee1ba7210e0a672608",
		BlockNumber: 12345,
		Uri:         "ipfs://bagaaieraiycmycdcbbfglbz3ga4ou5yt2eeh2gpzzmsu5wku3fquuyefnxiq",
	}

	json.NewEncoder(w).Encode(response)
}

func main() {
	http.HandleFunc("/api/tokens", tokensHandler)
	http.HandleFunc("/api/tokens/types", tokensTypesHandler)
	http.HandleFunc("/api/tokens/0x5713c2d9e9d4381bff966b1cdbf52cf4e8addc2c", tokenDetailsHandler)
	http.HandleFunc("/api/strategies/page/0", strategiesHandler)
	http.HandleFunc("/api/strategies/0", strategyDetailsHandler)
	http.HandleFunc("/api/strategies/token/0x5713c2d9e9d4381bff966b1cdbf52cf4e8addc2c", strategiesByTokenHandler)
	http.HandleFunc("/api/census", censusHandler)
	http.HandleFunc("/api/census/0", censusDetailsHandler)

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}
}
