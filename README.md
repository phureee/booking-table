## Booking-table
It's a simple application. To reserve a table in a restaurant. Uses RESTful API by Echo framework

## how to run on local
```
go run main.go
```


## API Init table

This API must be always called first and only once at the beginning

```http
POST /table
```

### Request

```javascript
{
    "table_amount" : int,  //require
    "seater_per_table" : int //optional : default 4
}
```

### Response

```javascript
{
    "message" : string,
    "status" : int
}
```

### example

```javascript
Request
{
    "table_amount": 3
}

Response success
{
    "message": "initialized table success",
    "status": 201
}

Response error
{
    "message": "init table only once",
    "status": 400
}
```

## API Get all table

For show information of table 

```http
GET /table
```


### Response

```javascript
{
    "result" : object,
    "status" : int
}
```

### example

```javascript
Response
{
    "result": [
        {
            "table_id": 1,
            "seater_amount": 4,
            "available": false,
            "create_time": "2023-09-17T13:44:53.1716588+07:00"
        },
        {
            "table_id": 2,
            "seater_amount": 4,
            "available": false,
            "create_time": "2023-09-17T13:44:53.1716588+07:00"
        }
    ],
    "status": 200
}
```

## API Booking table

This API must be always called API Init table first.

```http
POST /booking
```

### Request

```javascript
{
    "customer_amount" : int,  //require
}
```

### Response

```javascript
{
    "result" : {
        "booking_id" : int
        "tables_booked" : int
        "tables_id" : []int
    },
    "status" : int
}
```

### example

```javascript
Request
{
    "customer_amount": 3
}

Response success
{
    "result": {
        "booking_id": 1,
        "tables_booked": 2,
        "tables_id": [
            1,
            2
        ]
    },
    "status": 200
}

Response error
{
    "message": "not enough tables for all customers",
    "status": 400
}
```

## API Get All Booking

```http
GET /booking
```

### Response

```javascript
{
    "result" : {
        "booking_id" : int
        "tables_id" : []int
        "customer_amount" : int
        "available" : bool
        "create_time" : string
    },
    "status" : int
}
```

### example

```javascript
Response
{
    "result": [
        {
            "booking_id": 1,
            "table_id": [
                2,
                3
            ],
            "customer_amount": 8,
            "available": true,
            "create_time": "2023-09-17T13:54:19.1250211+07:00"
        }
    ],
    "status": 200
}
```

## API Cancel Booking

This API must be always called API Init table first.

```http
POST /booking/cancel
```

### Request

```javascript
{
    "booking_id" : int,  //require
}
```

### Response

```javascript
{
    "result" : {
        "amount_of_cancel" : int
        "tables_id" : []int
    },
    "status" : int
}
```

### example

```javascript
Request
{
   "booking_id": 1
}

Response success
{
    "result": {
        "amount_of_cancel": 2,
        "tables_id": [
            2,
            3
        ]
    },
    "status": 200
}

Response error
{
    "message": "this booking id have canceled",
    "status": 400
}

Response error
{
    "message": "booking id not found",
    "status": 404
}
```