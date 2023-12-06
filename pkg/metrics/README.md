## Metrics

The package is a wrapper to the different metrics provider that can be used

It returns a metrics provider [common.Provider](./common/common.go#L9) that can be used to create a new counter and can used to show the metrics to a path.

## Example

```go
  // To create the metrics provider with prometheus exporter.
  mProvider, err := metrics.New(&metrics.Config{
    Name:     "sample",
    Type:     metrics.Otel,
    Exporter: metrics.Prometheus,
  })
  if err != nil {
    return err
  }

  // The request counter to store the metrics for number of incoming requests.
  requestCounter := mProvider.NewCounter("sample_request", "Number of requests", "path", "method")
  // To record the metrics.
  requestCounter.Record(r.Context(), 1, r.URL.Path, r.Method)
```

For handler, it provides it in the func type that [api](../api/) wrapper for routing provides. Here is an example using the `api` package as routing:

```go
  shutdown := make(chan os.Signal, 1)
  signal.Notify(shutdown, syscall.SIGINT, syscall.SIGTERM)

  // The api package is a wrapper to the 
  api := api.New(shutdown)
  api.Handle(http.MethodGet, "/metrics", mProvider.Handler)
```
