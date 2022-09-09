## Get All Books details of Library Management System
Description : This API would be used for printing/accessing all the books details available in the library.

### HTTP Request
GET /api/book/all

### URL Parameters
N/A

### Query Parameters
N/A


### Request Headers
```
Content-Type: application/x-www-form-urlencoded
```

### Request Body
N/A


### Status codes and errors
| Value | Description           |
|-------|-----------------------|
| 200   | OK                    |
| 400   | Bad Request           |
| 403   | Forbidden             |
| 410   | Gone                  |
| 500   | Internal Server Error |

### Response Headers
N/A

### Success Response Body
```
{
    "Message": "Printed details of all the books in the library."
}
```

### Bad Request Response when wrong request given.
```
{
    "Message": "Book details couldn't be accessed due to bad request."
}
```
