## Logger

It is a helper package to initiate the slog logger handler. It provides certain options to set the slog configuration, or if not provided, uses default.


### Examples:

```go
// This will initiate a text logger.
l := slogger.New()

// This will initiate a json logger.
l := slogger.New(
  slogger.WithFormat(slogger.JSON),
)

// This will initiate a json logger and disable printing the source of the log.
// Which includes the file name and line number where it was printed. 
l := slogger.New(
  slogger.WithFormat(slogger.JSON),
  slogger.WithOutSource(),
)

// This can use a io.Writer that writes the logs to grafana logger.
l := slogger.New(
  slogger.WithWriter(grafana.Logger),
)

// It can be printed like this.
log.Info(
  "Build service is up",
  slog.String("environment", appEnv),
  slog.Int("port", config.Http.Port),
)

// Or it can be printed like this.
log.Info(
  "Build service is up",
  "environment", appEnv,
  "port", config.Http.Port,
)
```
