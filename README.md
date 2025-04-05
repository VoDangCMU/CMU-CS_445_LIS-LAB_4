## Quick Start
### On Windows 
 ```shell
  start.cmd
 ```

### On Linux
 ```shell
  start.sh
 ```

### Endpoint:
* Check With Token
```http request
    GET localhost:8080/api/user/check-with-token
```

* Register

```http request
    PUT localhost:8080/api/user/auth/register
```
```json
    {
    "Username": "admin",
    "Password": "1234567",
    "Fullname": "Admin",
    "Email": "vodang@gmail.com",
    "Phone" : "1234566",
    "DateOfBirth": "01/01/2004"
    }
```
* Login

```http request
    POST localhost:8080/api/user/auth/login
```
```json
    {
    "Username": "admin",
    "Password": "1234567",
    "KeepLogin" : "true"
    }
```

* Logout

```http request
    POST localhost:8080/api/user/auth/logout
```









