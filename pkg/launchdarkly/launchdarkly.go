package launchdarkly

import (
	"fmt"
	"os"
	"time"

	"github.com/launchdarkly/go-sdk-common/v3/ldcontext"
	ld "github.com/launchdarkly/go-server-sdk/v6"
)

type LaunchDarklyClient struct {
	ldClient *ld.LDClient
}

func NewClient(sdkKey string) *LaunchDarklyClient {
	ldClient, _ := ld.MakeClient(sdkKey, 5*time.Second)

	if ldClient.Initialized() {
		fmt.Println("SDK successfully initialized!")
	} else {
		fmt.Println("SDK failed to initialize" + sdkKey)
		os.Exit(1)
	}

	return &LaunchDarklyClient{
		ldClient: ldClient,
	}
}

func (c *LaunchDarklyClient) ReadFlag(featureFlagKey string) bool {
	// Set up the evaluation context. This context should appear on your LaunchDarkly contexts dashboard
	// soon after you run the demo.
	context := ldcontext.NewBuilder("example-user-key").
		Name("Sandy").
		Build()

	flagValue, err := c.ldClient.BoolVariation(featureFlagKey, context, false)
	if err != nil {
		fmt.Println("error: " + err.Error())
	}

	return flagValue
}
