# Unofficial Chapa SDK for Golang

[![Chapa Logo](image.png)](https://chapa.co/)

Unofficial SDK for Chapa Payment Gateway in Ethiopia.

## Overview
This SDK provides a convenient way to interact with Chapa's API for payments in Ethiopia.

## Installation
```bash
go get github.com/alazarbeyenenew2/chapasdk
```
## Example Code
``` go
package main

import (
    "context"
    "fmt"
    "log"

    "github.com/alazarbeyenenew2/chapasdk"
    "github.com/alazarbeyenenew2/chapasdk/pkg/model"
    "github.com/shopspring/decimal"
)

func main() {
    // Initialize the Chapa client with your API keys
    chapa := chapasdk.NewClient("CHAPUBK_PUBLIC_KEY", "CHASECK_SECRET_KEY", "ENCRYPTION_KEY", false)

    // Initialize a transaction
    resp, err := chapa.InitTransaction(context.Background(), model.InitTransactionReq{
        Amount:                   decimal.NewFromInt(30),
        Currency:                 "ETB",
        Email:                    "test@gmail.com",
        FirstName:                "Alazar",
        LastName:                 "Beyene",
        PhoneNumber:              "0975146165",
        TxRef:                    "chewatatest-6669",
        CallbackURL:              "http://localhost:3000",
        ReturnURL:                "http://localhost:8000",
        CustomizationTitle:       "Test Payment",
        CustomizationDescription: "Test transaction",
        Meta:                     true,
    })

    if err != nil {
        log.Fatal("Error initializing transaction:", err)
    }

    fmt.Println("Transaction Response:", resp)
}
```

