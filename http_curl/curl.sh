#!/bin/bash

url="http://localhost:8080/v1/swift-codes/BREXPLPWLOD"
url1="http://localhost:8080/v1/swift-codes/PL"

iter=10000

for ((i=0;i<=iter;i++))
do
  echo "Request $i:"
  if (( i % 2 == 0)); then
    curl -X GET "$url"
  else
    curl -X GET "$url1"
  fi

done