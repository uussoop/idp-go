# Identity Provider Service

This service provides user identity and authentication functions.

# Endpoints

## Register User

```
POST /register
```

Creates a new user record.

### Request

```json
{
  "address": "0x...",
  "username": "john" 
}
```

### Response 

201 Created on success

500 Error on failure

## Get Nonce

```
POST /nonce
```

Gets a nonce for a user to sign.

### Request

```json
{
  "address": "0x..."
}
```

### Response

```json
{
  "nonce": "12938191293"
}
```

400 Error if address is invalid

403 Forbidden if user not found

## Login

```
POST /login
```

Authenticates a user and returns a JWT token.

### Request

```json
{
  "address": "0x...",
  "nonce": "12938191293", 
  "sig": "0x..." 
}
```

### Response

```json
{
  "token": "eyJ..." 
}
```

400 Error if invalid request

500 Error if JWT token generation fails


Here is the markdown documentation for the provided Go code:

## Pull

```
GET /pull
```

Pulls the public key for JWT verification.

#### Request Headers

- **Authorization**: JWT token 

#### Response 

- 200 OK with public key in body
- 401 Unauthorized if missing or invalid JWT token
- 500 Internal Server Error if error generating public key

#### Example

```json
{
  "public_key": "048a9d...341fe2" 
}
```