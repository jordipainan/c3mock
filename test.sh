#!/bin/bash

API_BASE_URL="http://localhost:8080/api"

echo "Testing /tokens"
curl -s -X GET "${API_BASE_URL}/tokens" | jq
echo -e "\n"

echo "Testing /tokens/types"
curl -s -X GET "${API_BASE_URL}/tokens/types" | jq
echo -e "\n"

echo "Testing /tokens/0x5713c2d9e9d4381bff966b1cdbf52cf4e8addc2c"
curl -s -X GET "${API_BASE_URL}/tokens/0x5713c2d9e9d4381bff966b1cdbf52cf4e8addc2c" | jq
echo -e "\n"

echo "Testing /strategies/page/0"
curl -s -X GET "${API_BASE_URL}/strategies/page/0" | jq
echo -e "\n"

echo "Testing /strategies/0"
curl -s -X GET "${API_BASE_URL}/strategies/0" | jq
echo -e "\n"

echo "Testing /strategies/token/0x5713c2d9e9d4381bff966b1cdbf52cf4e8addc2c"
curl -s -X GET "${API_BASE_URL}/strategies/token/0x5713c2d9e9d4381bff966b1cdbf52cf4e8addc2c" | jq
echo -e "\n"

echo "Testing /census (POST)"
curl -s -X POST "${API_BASE_URL}/census" -H "Content-Type: application/json" -d '{"strategyId": "0", "blockNumber": 123445}' | jq
echo -e "\n"

echo "Testing /census/0"
curl -s -X GET "${API_BASE_URL}/census/0" | jq
echo -e "\n"
