## Create User in Library Management System
Description : This API would be used for creating a new user.Access would be given only to admin and superadmin.

### HTTP Request
POST /api/user/createNew

### URL Parameters
N/A

### Query Parameters
N/A


### Request Headers
```
Content-Type: application/x-www-form-urlencoded
```

### Request Body
| Parameter    | Format | Description                                |
|--------------|--------|--------------------------------------------|
| First_Name   | String | First Name of User |
| Last_Name    | String | Last Name of User |
| Email_Address| String | Email Id of user  |
| Password     | String | Set Password of User |
| Contact_No   | longint| Contact number of user |
| Address      | String | Address of User       |
| Role         | String | Role of User(Admin,Superadmin,User) |


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
    "Message": User Created Successfully "
}
```

### Bad Request Response when wrong info entered
```
{
    "Message": "Invalid info entered.Please check again"
}
```

### Bad Request Response when user already exists
```
{
    "Message": "Your username has already been set. Try logging in."
}
```

### Forbidden Response when field is empty
```
{
    "Message": Please enter valid credentials"
}
```
### Forbidden Response when role doesn't match
```
{
    "Message": "Access restricted"
}
```
