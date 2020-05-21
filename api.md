# User
## Sign UP

"/signup" **POST**

### Answers

- 201 - Created user
- 409 - Already registered
- 500 - Internal error

## Log IN

"/login" **POST**

### Answers

- 200 - OK
- 400 - Bad Request
- 500 - Internal error


"/auth" **GET**

### Answers

- 200 - OK
- 401 - Unauthorized
- 500 - Internal error


"/user/{id}" **PUT**

### Answers

- 200 - OK
- 400 - Bad Request
- 500 - Internal error