# tradeRevenueResultsService
RevenueCalculatorService is developed in Go. The service calculates the profit/loss based on the business logic by sorting and processing the stock trading data and show the results grouped based on the financial year. Both the input and output are in JSON.

# Project Structure
## cmd/userservice/userservice.go
Configuration and router intialization.Then the main function continues in controller

## controller
Server handler layer. Server routes are written

## service
Business logic implemented
Datas are sorted and put into appropriate bucket to make a accurate calculation and the results are grouped based on the financial year.
API accepts input and shows results in JSON. 

# Installation

```go
git clone https://github.com/t-revathi/tradeRevenueResultsService.git

```

# Sample Input
```

```
# Sample Output
```
[
    {
        "Year": "2018-2019",
        "Items": [
            {
                "Date": "2019-02-14T12:42:31Z",
                "Market": "NEXTDC Ltd",
                "Quantity": 499,
                "PandL": 340.8061,
                "SellUnitPrice": 6.948978
            },
            {
                "Date": "2019-04-26T12:42:31Z",
                "Market": "NEXTDC Ltd",
                "Quantity": 322,
                "PandL": 59.60878,
                "SellUnitPrice": 6.390155
            }
        ]
    },
    {
        "Year": "2019-2020",
        "Items": [
            {
                "Date": "2019-12-11T12:42:31Z",
                "Market": "NEXTDC Ltd",
                "Quantity": 200,
                "PandL": 54.235153,
                "SellUnitPrice": 6.71
            }
        ]
    }
]
```



