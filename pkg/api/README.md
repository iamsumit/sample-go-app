## API

The API package is a wrapper to the mux router with additional features such as common middlewares to use, with a `Handler` method to have a common control over the incoming requests and returned response.

It also provides helper methods to read the arguments or contexts to pass between middlewares.

More features will be added to it.

### Examples:

```go
  shutdown := make(chan os.Signal, 1)
  signal.Notify(shutdown, syscall.SIGINT, syscall.SIGTERM)

  // Middlewares
  mw := make([]api.Middleware, 0)
  mw = append(mw, middleware.Logger(cfg.Log))
  mw = append(mw, middleware.Errors(cfg.Log))

  // API handler
  api := api.New(shutdown, mw...)

  // Registering a route to get user by id.
  api.Handle(http.MethodGet, "/v1/user/{id}", userV1.GetByID)
  
  server := &http.Server{
    Addr:    ":8080", // Set your desired port
    Handler: api,
  }

  // Start the server
  server.ListenAndServe()

```

```go
// Get returns the user for the given id.
func (h *Handler) GetByID(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
  // Read the id parameter value.
  id := api.Param(r, "id")

  user := User{ID: id}
  // Return a simple user object in response.
  api.Respond(ctx, w, user, http.StatusOK)

  return nil
}
```

The output will look like following:
```json
{
  "success": true,
  "timestamp": 1701852258,
  "data": {
    "id": "1"
  }
}
```
