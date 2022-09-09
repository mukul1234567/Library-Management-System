## Login for Library Management System
Deescrription : This API would be for user /admin/superadmin login. It would be accepting username(email address) and password.

### HTTP Request
 POST /api/login
### URL Parameters
N/A

### Query Parameters
N/A


### Request Headers
```
Content-Type: application/x-www-form-urlencoded
```
	
### Request Body
| Parameter  | Format | Description                                |
|------------|--------|--------------------------------------------|
| Email      | String | Email Id of person requesting authentication|
| Password   | String | Password of the person       |


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
    "Message": "Successfully logged in"
}
```

### Bad Request Response when Invalid Username/Password received
```
{
    "Message": "Invalid credentials"
}
```
