## Get All Books details of Library Management System
Deescrription : This API would be used for printing/accessing details of a particular book available in the library.

### HTTP Request
GET /api/book/{book_id}

### URL Parameters
{book_id}

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
    "Message": "Printed details of the given book id."
}
```

### Bad Request Response when wrong book_id given.
```
{
    "Message": "Book details for given book id are not available."
}
```
