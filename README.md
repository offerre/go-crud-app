# Go Simple CRUD App

[![Go Reference](https://pkg.go.dev/badge/golang.org/x/example.svg)](https://pkg.go.dev/golang.org/x/example)

This repository contains a collection of Go programs and libraries that
demonstrate the language, standard libraries, and tools.

## Clone the project

```
$ git clone https://github.com/offerre/go-crud-app.git
$ cd go-crud-app
```

## Run the server

```
$ go run main.go
```

## Result
![](https://i.postimg.cc/HxQrMKrb/Screenshot-2567-08-02-at-17-07-18.png)
![](https://i.postimg.cc/qkq2DvfC/Screenshot-2567-08-02-at-17-07-07.png)


## Services

#### Get all user data
```
curl --location 'localhost:1323/users'
```
#### Get specific user data by Id
```
curl --location 'localhost:1323/users/2' \
--header 'Content-Type: application/json'
```
#### Create user data
```
curl --location 'localhost:1323/users' \
--header 'Content-Type: application/json' \
--data '{
    "name": "Bucky",
    "cards": [
        {
            "id": 1,
            "cardType": "credit",
            "cardNumber": "1234-5678-90",
            "balance": 10000.0
        },
        {
            "id": 2,
            "cardType": "dedit",
            "cardNumber": "1234-5678-90",
            "balance": 20000.0
        }
    ]
}'
```
#### Update user data by Id
```
curl --location --request PUT 'localhost:1323/users/2' \
--header 'Content-Type: application/json' \
--data '{
    "name": "Ducky",
    "cards": [
        {
            "id": 1,
            "cardType": "credit",
            "cardNumber": "1234-5678-90",
            "balance": 12345.6789
        },
        {
            "id": 2,
            "cardType": "dedit",
            "cardNumber": "1234-5678-90",
            "balance": 55555.5555
        }
    ]
}'
```
#### Delete specific card from specific user
```
curl --location --request DELETE 'localhost:1323/users/2?cardID=2' \
--header 'Content-Type: application/json'
```
