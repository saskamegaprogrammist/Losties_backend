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

## Log OUT

"/logout" **DELETE**

### Answers

- 200 - OK
- 400 - Bad Request
- 500 - Internal error

## Authenticate

"/auth" **GET**

### Answers

- 200 - OK
- 401 - Unauthorized
- 500 - Internal error

## Update user info

"/user/{id}" **PUT**

### Answers

- 200 - OK
- 400 - Bad Request
- 500 - Internal error

## Get user info

"/user/{id}" **GET**

### Answers

- 200 - OK
- 400 - Bad Request
- 500 - Internal error

# Ad
## Ad ad to user

"/user/{id}/card" **POST**

### Answers

- 201 - Created ad
- 400 - Bad Request
- 401 - Unauthorized
- 500 - Internal error

## Get user ads

"/user/{id}/cards" **GET**

### Query params

- type: found/lost

### Answers

- 200 - OK
- 400 - Bad Request
- 401 - Unauthorized
- 500 - Internal error

## Get user ads number

"/user/{id}/cards/number" **GET**

### Query params

- type: found/lost

### Answers

- 200 - OK
- 400 - Bad Request
- 500 - Internal error

## Get all ads

"/cards" **GET**

### Query params

- type: found/lost
- sort: date/comments
- search

### Answers

- 200 - OK
- 400 - Bad Request
- 500 - Internal error

## Get ad

"/card/{id}" **GET**

### Answers

- 200 - OK
- 400 - Bad Request
- 500 - Internal error

# Pet
## Ad pet to ad

"/card/{id}/pet" **POST**

### Answers

- 201 - Created pet
- 400 - Bad Request
- 500 - Internal error

## Get ads pet

"/card/{id}/pet" **GET**


### Answers

- 200 - OK
- 400 - Bad Request
- 500 - Internal error

# Pic

## Ad ads picture

"/card/{id}/pic" **POST**

### Answers

- 201 - Created pic
- 400 - Bad Request
- 500 - Internal error

## Get ads picture

"/card/{id}/pic" **GET**


### Answers

- 200 - OK
- 400 - Bad Request
- 500 - Internal error

## Ad users picture

"/user/{id}/pic" **POST**

### Answers

- 201 - Created pic
- 400 - Bad Request
- 401 - Unauthorized
- 500 - Internal error

## Get users picture

"/user/{id}/pic" **GET**


### Answers

- 200 - OK
- 400 - Bad Request
- 500 - Internal error


# Coords

## Ad ads coords

"/card/{id}/coords" **POST**

### Answers

- 201 - Created coords
- 400 - Bad Request
- 500 - Internal error

## Get ads coords

"/card/{id}/coords" **GET**


### Answers

- 200 - OK
- 400 - Bad Request
- 500 - Internal error

## Get all coords

"/coords" **GET**

### Answers

- 200 - OK
- 400 - Bad Request
- 500 - Internal error

# Comments

## Ad comment to ad

"/card/{id}/comment" **POST**

### Answers

- 201 - Created comment
- 400 - Bad Request
- 500 - Internal error

## Get ads comments

"/card/{id}/coomments" **GET**


### Answers

- 200 - OK
- 400 - Bad Request
- 500 - Internal error