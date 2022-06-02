# Exercise

## Installation

Install the Go programming language and toolset: [golang.org](https://golang.org/doc/install)

You can verify you've installed correctly by running:
```
go version
```

For this exercise I used version `1.17`

## Build

You must build the application prior to running:
```
go build golang/exercise
```

This will result in the executable file `exercise`

## Running

Run the application with the following command:

```
./exercise
```

## Add

Used to add a transaction to the ledger.

URL: `/add`

Method: `POST`

Data constraints

```json
{
  "payer" : "[string - payer name]",
  "points" : "[int - amount of points]",
  "timestamp": "[timestamp - valid timestamp]"
}
```

Data example
```json
{
  "payer" : "ACME",
  "points" : 400,
  "timestamp": "2022-04-07T10:00:00Z"
}
```

Example call

```shell
curl --location --request POST 'http://localhost:80/add' \
--header 'Content-Type: application/json' \
--data-raw '{ "payer": "DANNON", "points": 300, "timestamp": "2020-10-31T10:00:00Z" }'
```

## Spend

Used to spend points.

URL: `/spend`

Method: `POST`

Data constraints

```json
{
  "user_id" : "[string - valid user id]",
  "points" : "[int - amount of points to spend]"
}
```

Data example
```json
{
  "user_id" : "4958309",
  "points" : 400
}
```

Example call

```shell
curl --location --request POST 'http://localhost:80/spend' \
--header 'Content-Type: application/json' \
--data-raw '{ "points": 5000 }'
```

## Balances

Used to spend points.

URL: `/balances`

Method: `GET`

Example call

```shell
curl --location --request GET 'http://localhost:80/balances'
```
