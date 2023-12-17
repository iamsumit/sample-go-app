


# Activity API.
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

###  activity

| Method  | URI     | Name   | Summary |
|---------|---------|--------|---------|
| POST | /v1/activity | [create activity](#create-activity) | Create returns the endpoint for the activity handler. |
  


## Paths

### <span id="create-activity"></span> Create returns the endpoint for the activity handler. (*createActivity*)

```
POST /v1/activity
```

This will help you create a new activity by given information.
It will validate the information and create a new activity.

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
| body | `body` | [NewActivity](#new-activity) | `models.NewActivity` | | ✓ | | The body to create a new activity. |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [201](#create-activity-201) | Created |  |  | [schema](#create-activity-201-schema) |
| [400](#create-activity-400) | Bad Request |  |  | [schema](#create-activity-400-schema) |
| [500](#create-activity-500) | Internal Server Error |  |  | [schema](#create-activity-500-schema) |

#### Responses


##### <span id="create-activity-201"></span> 201
Status: Created

###### <span id="create-activity-201-schema"></span> Schema
   
  

[CreateActivityCreatedBody](#create-activity-created-body)

##### <span id="create-activity-400"></span> 400
Status: Bad Request

###### <span id="create-activity-400-schema"></span> Schema
   
  

[CreateActivityBadRequestBody](#create-activity-bad-request-body)

##### <span id="create-activity-500"></span> 500
Status: Internal Server Error

###### <span id="create-activity-500-schema"></span> Schema
   
  

[CreateActivityInternalServerErrorBody](#create-activity-internal-server-error-body)

###### Inlined models

**<span id="create-activity-bad-request-body"></span> CreateActivityBadRequestBody**


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| Success | boolean| `bool` |  | | Success | `false` |
| Timestamp | int64 (formatted integer)| `int64` |  | | Timestamp | `1639237536` |
| errors | [ErrorResponse](#error-response)| `models.ErrorResponse` |  | |  |  |



**<span id="create-activity-created-body"></span> CreateActivityCreatedBody**


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| Data | [][ActivityRes](#activity-res)| `[]*models.ActivityRes` |  | | Data
in: body |  |
| Success | boolean| `bool` |  | | Success | `false` |
| Timestamp | int64 (formatted integer)| `int64` |  | | Timestamp | `1639237536` |



**<span id="create-activity-internal-server-error-body"></span> CreateActivityInternalServerErrorBody**


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| Errors | [interface{}](#interface)| `interface{}` |  | | Data
in: body | `{"error":"some internal error occured"}` |
| Success | boolean| `bool` |  | | Success | `false` |
| Timestamp | int64 (formatted integer)| `int64` |  | | Timestamp | `1639237536` |



## Models

### <span id="activity-res"></span> ActivityRes


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| CreatedAt | date-time (formatted string)| `strfmt.DateTime` |  | | CreatedAt represents the time when the user was created.

type: string | `2020-01-01T00:00:00Z` |
| Entity | string| `string` | ✓ | | Entity of the activity

type: string | `user` |
| ID | int64 (formatted integer)| `int64` | ✓ | | ID of the activity

type: int | `1` |
| Operation | string| `string` | ✓ | | Operation of the activity

type: string | `created` |
| UpdatedAt | date-time (formatted string)| `strfmt.DateTime` |  | | UpdatedAt represents the time when the user was updated.

type: string | `2020-01-01T00:00:00Z` |
| app | [App](#app)| `App` |  | |  |  |



### <span id="app"></span> App


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| Alias | string| `string` | ✓ | | Alias of the app

type: string | `sample` |
| CreatedAt | date-time (formatted string)| `strfmt.DateTime` |  | | CreatedAt represents the time when the app was created.

type: string | `2020-01-01T00:00:00Z` |
| ID | int64 (formatted integer)| `int64` | ✓ | | ID of the app

type: int | `1` |
| Name | string| `string` | ✓ | | Name of the app

type: string | `sample` |
| UpdatedAt | date-time (formatted string)| `strfmt.DateTime` |  | | UpdatedAt represents the time when the app was updated.

type: string | `2020-01-01T00:00:00Z` |



### <span id="error-response"></span> ErrorResponse


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| Data | [interface{}](#interface)| `interface{}` |  | | in:body | `{"field":"error message for this specific field"}` |
| Error | string| `string` |  | | in:body | `data is not in proper format` |



### <span id="new-activity"></span> NewActivity


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| AppName | string| `string` | ✓ | | Name of the application to which this activity belongs

in: body
type: string | `sample` |
| Entity | string| `string` | ✓ | | Name of the entity to which this activity belongs

in: body
type: string | `user` |
| Operation | string| `string` | ✓ | | Operation performed on the entity

in: body
type: string | `created` |



### <span id="pagination"></span> Pagination


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| Direction | string| `string` |  | | The direction of the sort

in: query |  |
| Page | int64 (formatted integer)| `int64` |  | | The current page

in: query
type: integer |  |
| PerPage | int64 (formatted integer)| `int64` |  | | The per page limit

in: query
type: integer |  |
| Sort | string| `string` |  | | The column to sort on

in: query |  |


