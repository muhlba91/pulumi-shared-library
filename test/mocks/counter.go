package mocks

import (
	"fmt"
	"sync"

	"github.com/pulumi/pulumi/sdk/v3/go/common/resource"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/rs/zerolog/log"
)

// Counter is a thread-safe struct to keep track of resource counts during testing.
type Counter struct {
	Resources map[string][]string
	mu        sync.Mutex
}

// NewCounter initializes and returns a new Counter instance.
func NewCounter() *Counter {
	return &Counter{
		Resources: make(map[string][]string),
	}
}

// NewResource mocks resource creation for Pulumi.
// args: MockResourceArgs containing information about the resource being mocked.
func (c *Counter) NewResource(args pulumi.MockResourceArgs) (string, resource.PropertyMap, error) {
	log.Info().Msgf("Mocking resource of type %s with name %s", args.TypeToken, args.Name)

	c.mu.Lock()
	c.Resources[args.TypeToken] = append(c.Resources[args.TypeToken], args.Name)
	c.mu.Unlock()

	id := args.Name + "_id"

	outs := args.Inputs
	if _, ok := outs["name"]; !ok {
		outs["name"] = resource.NewStringProperty(args.Name)
	}

	switch args.TypeToken {
	case "random:index/randomPassword:RandomPassword":
		outs["result"] = resource.NewStringProperty(
			fmt.Sprintf("mocked-password-%v", args.Inputs["length"].NumberValue()),
		)
	case "random:index/randomId:RandomId":
		length := args.Inputs["byteLength"].NumberValue()
		outs["hex"] = resource.NewStringProperty(fmt.Sprintf("mocked-bytes-hex-%v", length))
		outs["b64Std"] = resource.NewStringProperty(fmt.Sprintf("mocked-bytes-b64std-%v", length))
		outs["b64Url"] = resource.NewStringProperty(fmt.Sprintf("mocked-bytes-b64url-%v", length))
	case "random:index/randomString:RandomString":
		outs["result"] = resource.NewStringProperty(
			fmt.Sprintf("mocked-string-%v", args.Inputs["length"].NumberValue()),
		)
	case "tls:index/privateKey:PrivateKey":
		key := privateKeyMockSuffix(args.Inputs)
		outs["privateKeyPem"] = resource.MakeSecret(
			resource.NewStringProperty(fmt.Sprintf("mocked-private-key-%s", key)),
		)
		outs["publicKeyPem"] = resource.NewStringProperty(fmt.Sprintf("mocked-public-key-%s", key))
		outs["privateKeyOpenssh"] = resource.MakeSecret(
			resource.NewStringProperty(fmt.Sprintf("mocked-private-key-openssh-%s", key)),
		)
		outs["publicKeyOpenssh"] = resource.NewStringProperty(fmt.Sprintf("mocked-public-key-openssh-%s", key))
	case "aws:iam/accessKey:AccessKey":
		outs["secret"] = resource.MakeSecret(
			resource.NewStringProperty(fmt.Sprintf("mocked-secret-%s", args.Name)),
		)
	case "scaleway:iam/apiKey:ApiKey":
		outs["accessKey"] = resource.MakeSecret(
			resource.NewStringProperty(fmt.Sprintf("mocked-access-key-%s", args.Name)),
		)
		outs["secretKey"] = resource.MakeSecret(
			resource.NewStringProperty(fmt.Sprintf("mocked-secret-key-%s", args.Name)),
		)
	case "hcloud:index/primaryIp:PrimaryIp":
		outs["ipAddress"] = resource.NewStringProperty(
			fmt.Sprintf(
				"mocked-ip-address-%s-%s",
				args.Inputs["type"].StringValue(),
				args.Inputs["location"].StringValue(),
			),
		)
	case "netbird:index/setupKey:SetupKey":
		outs["key"] = resource.MakeSecret(
			resource.NewStringProperty(fmt.Sprintf("mocked-netbird-setup-key-%s", args.Name)),
		)
	case "pulumiservice:index:AccessToken":
		outs["value"] = resource.MakeSecret(
			resource.NewStringProperty(fmt.Sprintf("mocked-pulumi-access-token-%s", args.Name)),
		)
	}

	if len(args.TypeToken) >= 4 && args.TypeToken[:4] == "aws:" {
		outs["arn"] = resource.NewStringProperty(fmt.Sprintf("mocked-arn-%s", args.Name))
	}

	if len(args.TypeToken) >= 7 && args.TypeToken[:7] == "hcloud:" {
		id = "1"
	}

	return id, outs, nil
}

// privateKeyMockSuffix returns the algorithm and its parameter of a mocked private key, e.g. "RSA-4096".
// inputs: The inputs of the private key resource.
func privateKeyMockSuffix(inputs resource.PropertyMap) string {
	algorithm := inputs["algorithm"].StringValue()

	switch {
	case inputs["rsaBits"].IsNumber():
		return fmt.Sprintf("%s-%v", algorithm, inputs["rsaBits"].NumberValue())
	case inputs["ecdsaCurve"].IsString():
		return fmt.Sprintf("%s-%s", algorithm, inputs["ecdsaCurve"].StringValue())
	default:
		return algorithm
	}
}

// Call mocks function calls for Pulumi.
// args: MockCallArgs containing information about the call being mocked.
func (c *Counter) Call(
	args pulumi.MockCallArgs,
) (resource.PropertyMap, error) {
	log.Info().Msgf("Mocking call of type %s with args: %+v", args.Token, args.Args)

	outs := args.Args

	switch args.Token {
	case "hcloud:index/getNetwork:getNetwork":
		name := ""
		if v, ok := args.Args["name"]; ok {
			name = v.StringValue()
		}
		outs["name"] = resource.NewStringProperty(name)
		outs["id"] = resource.NewStringProperty("1")
	case "netbird:index/getPeer:getPeer":
		name := ""
		if v, ok := args.Args["name"]; ok {
			name = v.StringValue()
		}
		outs["name"] = resource.NewStringProperty(name)
		outs["id"] = resource.NewStringProperty("1")
	case "gitlab:index/getCurrentUser:getCurrentUser":
		outs["id"] = resource.NewStringProperty("1")
	}

	return outs, nil
}
