package config

type Configuration struct {
	PubSub PubSubConfig
}

type PubSubConfig struct {
	Project      string
	Topic        string
	Subscription string
}
