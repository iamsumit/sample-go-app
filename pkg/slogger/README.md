## Logger

It is a helper package to initiate the logger handler. It provides support for different kinds of logger that can be used for your site.

It uses slog as its logger for most of the parts. It also returns a `ServerLogger` that converts the slog to `log.Logger` so that the `http.Server` can use the same kind of logger.

It has support for following loggers:

- `GCloudLogger`: To be used in the google cloud for google cloud specific logging structure.
- `TextLogger`: It can be used for text formatted logs, probably on local.
- `JSONLogger`: It can be used for json formatted logs, probably on cloud.
- `TintLogger`: It can be used for colored text formmatted logs, probably on local.
