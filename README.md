# Sample project

This is my sample project to try out different things for my POC purposes.


## Services:

There are two services at the moment of writing this README. The first one is sample, and another one is message. Both resides in the [services][1] directory.

There is [pkg][2] directory as well which contains common packages to be used by the services.

Please note that this is not a showcase of go best practices. It just random code for my POC purpose.

### Sample:

I have tried following in the sample service:

- [viper][4] package to read the configuration from yaml structure and to read the environment variables.
- [launchdarky][5] package to try out the launch darky flag configuration.
- [opentel][15] package is a go framework to export the metrics and traces to different services.
- [prometheus][16] Prometheus is a powerful metrics collector and is used to try out the basic metrics collection.
- [grafana][17] Grafana is a dashboarding and alerting tool to view the metrics from different sources by applying the different sets of conditions.
- [jaeger][19] Jaeger is a distributed open source tracing platform with a good user interface.

### Message:

I have tried following in the message service:

- [cobra][6] To try out the command line tool with basic publish and reading from pub/sub.
- [protobuf][7] To try out the serialization with protobuf library.
- [pubsub][8] To try out the pub/sub messaging service using [emulator][9].
- [concurrency][10] To try out the goroutine and channels for subscribing the messages from pub/sub.


## Try the application

There are certain dependencies that we have to install first before we can run the application. Those would be following:

- [golang][11] ofcourse, with minimum version 1.20.
- [pubsub-emulator][9]
- [task][12] to use the application provided commands to use the application.
- [protobuf][13] though, it can be installed using task command.
- [docker][14] The pub/sub and mysql will be running inside the docker container.

### Start the application

#### Message Application:

There are certain task commands to run the applications. Lets go step by step:

- `task message:setup-deps`: To install the `protobuf` dependency.
- `task message:init`: To set the environment files and generate the code from proto definitions.
- `task message:pubsub-up`: To start the pub/sub emulator as a docker container. Please note that this will run the foreground process, so you need to open a new tab for next commands. Keep it running please.
- `task message:read`: This will create the topic and subscription if required and start reading to the pub/sub messages. This will also be a foreground process. Keep it running and open next tab please.
-  `task message:publish`: It will generate a dummy message using [protobuf][13] and publish it to the local pub/sub emulator. The `read` command above will read the message and unmarshal it using `protobuf` library itself.
- `task message:pubsub-down`: To stop the pubsub emulator docker container.

#### Sample Application:

- `task sample:init`: To set the environment files.
- `task sample:db-up`: To start the database container.
- `task sample:run`: To start the application. A http process will start and once done, the application can be accessed via the `localhost:8080` where 8080 is the port number defined in the [config.yaml][3].
- `task prometheus:up`: To start the prometheus server. There is a [prometheus.yml][18] file which contains the configuration to import the metrics from our local server. Please note that if you use a different port than `:8080` for your application, you would need to update the same in the yaml config.
- `task jaeger:up`: To start the jaeger service.
- `task grafana:up`: It is to start the grafana server. Once started, you can login using `admin/admin` and configure your dashboard. Please check the screenshots below to see the configuration.
- `task grafana:down`: To stop the grafana server.
- `task jaeger:down`: To stop the jaeger service.
- `task prometheus:down`: To stop the prometheus server.
- `task sample:db-down`: To stop the database container.

![grafana datasource config](./docs/images/grafana-datasource-config.png)

![grafana dashboard](./docs/images/grafana-dashboard.png)

There are some commented code as well which is basically either because of no secret key and/or there are multiple way of doing things, like reloading the configuration on change without re-running the application.

[1]:./services
[2]:./pkg/
[3]:./services/sample/config/config.yaml
[4]:https://github.com/spf13/viper
[5]:https://github.com/launchdarkly/go-sdk-common
[6]:https://github.com/spf13/cobra
[7]:https://github.com/golang/protobuf
[8]:https://pkg.go.dev/cloud.google.com/go/pubsub
[9]:https://cloud.google.com/pubsub/docs/emulator
[10]:https://medium.com/nerd-for-tech/learning-go-concurrency-goroutines-channels-8836b3c34152
[11]:https://go.dev/doc/install
[12]:https://taskfile.dev/
[13]:https://protobuf.dev/
[14]:https://www.docker.com/get-started/
[15]:https://github.com/open-telemetry/opentelemetry-go
[16]:https://prometheus.io/
[17]:https://grafana.com/
[18]:./build/prometheus/prometheus.yml
[19]:https://www.jaegertracing.io/docs/1.51/
