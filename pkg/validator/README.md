## Validator

Validator is a wrapper on the [go-playground/validator](github.com/go-playground/validator) to parse and return a satisfying error for internal packages.

### Example

```go

type User struct {
	Name        string  `validate:"required"`
	Email       *string `validate:"email"`
	Biography   *string 
	DateOfBirth *time.Time
}

user := User{
  Name: "something",
  Email: "some"
}

err := validator.Validate(user)
if err != nil {
  return nil, err
}

// err should return a validator.FieldErrors type error.
```
