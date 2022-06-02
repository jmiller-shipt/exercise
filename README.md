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

## Calling

Once the application is running, you can call it with `curl`:

### Example Add
```shell
curl --location --request POST 'http://localhost:80/add' \
--header 'Content-Type: application/json' \
--data-raw '{ "payer": "DANNON", "points": 300, "timestamp": "2020-10-31T10:00:00Z" }'
```

### Example Spend
```shell
curl --location --request POST 'http://localhost:80/spend' \
--header 'Content-Type: application/json' \
--data-raw '{ "points": 5000 }'
```

### Example Balances
```shell
curl --location --request GET 'http://localhost:80/balances'
```
