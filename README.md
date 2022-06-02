# Exercise

## Installation

### Checkout code
```
git clone https://github.com/jmiller-shipt/exercise.git
```

### Install Go

Install the Go programming language and toolset: [golang.org](https://golang.org/doc/install)

You can verify you've installed correctly by running:
```
go version
```

For this exercise I used version `1.17`

## Build
Make sure you are in the `exercise` directory

```
cd exercise
```

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
  "user_id" : "[string - valid user id]",
  "payer" : "[string - payer name]",
  "points" : "[int - amount of points]",
  "timestamp": "[timestamp - valid timestamp]"
}
```

Data example
```json
{
  "user_id": "4958309",
  "payer" : "DANNON",
  "points" : 300,
  "timestamp": "2020-10-31T10:00:00Z"
}
```

Example call

```shell
curl --location --request POST 'http://localhost:80/add' \
--header 'Content-Type: application/json' \
--data-raw '{
    "user_id": "4958309",
    "payer": "DANNON",
    "points": 300,
    "timestamp": "2020-10-31T10:00:00Z"
}'
```

Example success response
```json
{"success":true,"message":"Successfully added transaction","data":{"user_id":"4958309","payer":"DANNON"}}
```

Example failure response
```json
{"success":false,"message":"Unprocessable entity","data":null}
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
  "points" : 5000
}
```

Example call

```shell
curl --location --request POST 'http://localhost:80/spend' \
--header 'Content-Type: application/json' \
--data-raw '{ "user_id": "4958309", "points": 5000 }'
```

Example success response
```json
{"success":true,"message":"Points spent by payer for user","data":{"user_id":"4958309","spend_records":[{"payer":"DANNON","points":-100},{"payer":"UNILEVER","points":-200},{"payer":"MILLER COORS","points":-4700}]}}
```

Example failure response
```json
{"success":false,"message":"Unprocessable entity","data":null}
```

## Balances

Used to spend points.

URL: `/balances/{userId}`

Method: `GET`

Example call

```shell
curl --location --request GET 'http://localhost:80/balances/4958309'
```

Example success response
```json
{"success":true,"message":"Current payer balances","data":[{"payer":"DANNON","points":900,"user_id":"4958309"}]}
```

Example failure response
```json
{"success":false,"message":"Unprocessable entity","data":null}
```
