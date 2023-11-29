## Logger

The package is a wrapper to the logging libraries that can be used in the project.

It returns a handler instance that implements the `Logger` interface. It uses the options pattern to set the configurations.

### Example code:

```go
log, err := logger.New(
  logger.WithSlogger(),
  logger.WithJSONFormat(),
)
if err != nil {
  // Do something with error
}

// Log a basic information with couple of attributes.
log.Info(
  "Logger connected!",
  "attr-key", "attr-value",
  "attr2-key", "attr2-value",
)
```
