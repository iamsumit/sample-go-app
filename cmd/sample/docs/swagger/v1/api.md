


# Sample API.
the purpose of this application is to provide basic routes
to play with.

This should demonstrate all the possible comment annotations
that are available to turn go code into a fully compliant swagger 2.0 spec
  

## Informations

### Version

0.0.1

### Terms Of Service

there are no TOS at this moment, use at your own risk we take no responsibility

## Content negotiation

### URI Schemes
  * http
  * https

### Consumes
  * application/json

### Produces
  * application/json

## All endpoints

###  users

| Method  | URI     | Name   | Summary |
|---------|---------|--------|---------|
| POST | /v1/user | [create user](#create-user) | Create a new user by given information. |
| GET | /v1/user/{id} | [get user](#get-user) | ByID returns the user for the given id. |
  


## Paths

### <span id="create-user"></span> Create a new user by given information. (*createUser*)

```
POST /v1/user
```

This will help you create a new user by given information.
It will validate the information and create a new user.
The uniqueness validation will be done if email is provided.

#### URI Schemes
  * http
  * https

#### Consumes
  * application/json

#### Produces
  * application/json

#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| body | `body` | [NewUser](#new-user) | `models.NewUser` | | ✓ | | The body to create a new user. |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [201](#create-user-201) | Created |  |  | [schema](#create-user-201-schema) |
| [400](#create-user-400) | Bad Request |  |  | [schema](#create-user-400-schema) |
| [409](#create-user-409) | Conflict |  |  | [schema](#create-user-409-schema) |

#### Responses


##### <span id="create-user-201"></span> 201
Status: Created

###### <span id="create-user-201-schema"></span> Schema
   
  

[CreateUserCreatedBody](#create-user-created-body)

##### <span id="create-user-400"></span> 400
Status: Bad Request

###### <span id="create-user-400-schema"></span> Schema
   
  

[CreateUserBadRequestBody](#create-user-bad-request-body)

##### <span id="create-user-409"></span> 409
Status: Conflict

###### <span id="create-user-409-schema"></span> Schema
   
  

[CreateUserConflictBody](#create-user-conflict-body)

###### Inlined models

**<span id="create-user-bad-request-body"></span> CreateUserBadRequestBody**


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| Success | boolean| `bool` |  | | Success | `false` |
| Timestamp | int64 (formatted integer)| `int64` |  | | Timestamp | `1639237536` |
| errors | [ErrorResponse](#error-response)| `models.ErrorResponse` |  | |  |  |



**<span id="create-user-conflict-body"></span> CreateUserConflictBody**


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| Errors | [interface{}](#interface)| `interface{}` |  | | Data
in: body | `{"error":"user already exists"}` |
| Success | boolean| `bool` |  | | Success | `false` |
| Timestamp | int64 (formatted integer)| `int64` |  | | Timestamp | `1639237536` |



**<span id="create-user-created-body"></span> CreateUserCreatedBody**


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| Data | [][User](#user)| `[]*models.User` |  | | Data
in: body |  |
| Success | boolean| `bool` |  | | Success | `false` |
| Timestamp | int64 (formatted integer)| `int64` |  | | Timestamp | `1639237536` |



### <span id="get-user"></span> ByID returns the user for the given id. (*getUser*)

```
GET /v1/user/{id}
```

This will help you get a user information by given id.

#### URI Schemes
  * http
  * https

#### Consumes
  * application/json

#### Produces
  * application/json

#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| id | `path` | int64 (formatted integer) | `int64` |  | ✓ |  | The id to get a new user. |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [200](#get-user-200) | OK |  |  | [schema](#get-user-200-schema) |
| [400](#get-user-400) | Bad Request |  |  | [schema](#get-user-400-schema) |
| [404](#get-user-404) | Not Found |  |  | [schema](#get-user-404-schema) |

#### Responses


##### <span id="get-user-200"></span> 200
Status: OK

###### <span id="get-user-200-schema"></span> Schema
   
  

[GetUserOKBody](#get-user-o-k-body)

##### <span id="get-user-400"></span> 400
Status: Bad Request

###### <span id="get-user-400-schema"></span> Schema
   
  

[GetUserBadRequestBody](#get-user-bad-request-body)

##### <span id="get-user-404"></span> 404
Status: Not Found

###### <span id="get-user-404-schema"></span> Schema
   
  

[GetUserNotFoundBody](#get-user-not-found-body)

###### Inlined models

**<span id="get-user-bad-request-body"></span> GetUserBadRequestBody**


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| Success | boolean| `bool` |  | | Success | `false` |
| Timestamp | int64 (formatted integer)| `int64` |  | | Timestamp | `1639237536` |
| errors | [ErrorResponse](#error-response)| `models.ErrorResponse` |  | |  |  |



**<span id="get-user-not-found-body"></span> GetUserNotFoundBody**


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| Errors | [interface{}](#interface)| `interface{}` |  | | Data
in: body | `{"error":"user not found"}` |
| Success | boolean| `bool` |  | | Success | `false` |
| Timestamp | int64 (formatted integer)| `int64` |  | | Timestamp | `1639237536` |



**<span id="get-user-o-k-body"></span> GetUserOKBody**


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| Data | [][User](#user)| `[]*models.User` |  | | Data
in: body |  |
| Success | boolean| `bool` |  | | Success | `false` |
| Timestamp | int64 (formatted integer)| `int64` |  | | Timestamp | `1639237536` |



## Models

### <span id="error-response"></span> ErrorResponse


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| Data | [interface{}](#interface)| `interface{}` |  | | in:body | `{"field":"error message for this specific field"}` |
| Error | string| `string` |  | | in:body | `data is not in proper format` |



### <span id="new-user"></span> NewUser


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| Biography | string| `string` |  | | Bio of the user

in: body
type: string | `I am a developer by profession.` |
| DateOfBirth | string| `string` |  | | Date of birth of the user

in: body
type: string | `1990-01-15` |
| Email | string| `string` |  | | the email address for this user

in: body
type: string | `user@provider.net` |
| Name | string| `string` | ✓ | | Name of the user

in: body
type: string | `Sumit Kumar` |



### <span id="settings"></span> Settings


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| Biography | string| `string` |  | | Bio of the user

type: string | `I am a developer by profession.` |
| DateOfBirth | string| `string` |  | | Date of birth of the user

type: string | `1990-01-15` |
| Email | string| `string` |  | | Email of the user

type: string | `user@provider.net` |
| IsActive | boolean| `bool` |  | | IsActive represents the status of the user.

type: bool | `true` |
| IsSubscribed | boolean| `bool` |  | | IsSubscribed represents the subscription status of the user.

type: bool | `true` |



### <span id="user"></span> User


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| CreatedAt | date-time (formatted string)| `strfmt.DateTime` |  | | CreatedAt represents the time when the user was created.

type: string | `2020-01-01T00:00:00Z` |
| ID | int64 (formatted integer)| `int64` | ✓ | | ID of the user

type: int | `1` |
| Name | string| `string` |  | | Name of the user

type: string | `Sumit Kumar` |
| UpdatedAt | date-time (formatted string)| `strfmt.DateTime` |  | | UpdatedAt represents the time when the user was updated.

type: string | `2020-01-01T00:00:00Z` |
| settings | [Settings](#settings)| `Settings` |  | |  |  |


