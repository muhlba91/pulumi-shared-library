package mocks

import (
	"fmt"

	"github.com/pulumi/pulumi/sdk/v3/go/common/resource"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/rs/zerolog/log"
)

type Mocks int

func (Mocks) NewResource(args pulumi.MockResourceArgs) (string, resource.PropertyMap, error) {
	log.Info().Msgf("Mocking resource of type %s with name %s", args.TypeToken, args.Name)

	id := args.Name + "_id"

	outs := args.Inputs
	if _, ok := outs["name"]; !ok {
		outs["name"] = resource.NewStringProperty(args.Name)
	}

	if args.TypeToken == "random:index/randomPassword:RandomPassword" {
		outs["result"] = resource.NewStringProperty(
			fmt.Sprintf("mocked-password-%v", args.Inputs["length"].NumberValue()),
		)
	}
	if args.TypeToken == "random:index/randomString:RandomString" {
		outs["result"] = resource.NewStringProperty(
			fmt.Sprintf("mocked-string-%v", args.Inputs["length"].NumberValue()),
		)
	}

	if args.TypeToken == "tls:index/privateKey:PrivateKey" {
		outs["privateKeyPem"] = resource.MakeSecret(
			resource.NewStringProperty(
				fmt.Sprintf(
					"mocked-private-key-%s-%v",
					args.Inputs["algorithm"].StringValue(),
					args.Inputs["rsaBits"].NumberValue(),
				),
			),
		)
		outs["publicKeyPem"] = resource.NewStringProperty(
			fmt.Sprintf(
				"mocked-public-key-%s-%v",
				args.Inputs["algorithm"].StringValue(),
				args.Inputs["rsaBits"].NumberValue(),
			),
		)
	}

	if len(args.TypeToken) >= 4 && args.TypeToken[:4] == "aws:" {
		outs["arn"] = resource.NewStringProperty(fmt.Sprintf("mocked-arn-%s", args.Name))
	}
	if args.TypeToken == "aws:iam/accessKey:AccessKey" {
		outs["secret"] = resource.MakeSecret(
			resource.NewStringProperty(fmt.Sprintf("mocked-secret-%s", args.Name)),
		)
	}

	if len(args.TypeToken) >= 8 && args.TypeToken[:8] == "scaleway:" {
		outs["arn"] = resource.NewStringProperty(fmt.Sprintf("mocked-arn-%s", args.Name))
	}
	if args.TypeToken == "scaleway:iam/apiKey:ApiKey" {
		outs["accessKey"] = resource.MakeSecret(
			resource.NewStringProperty(fmt.Sprintf("mocked-access-key-%s", args.Name)),
		)
		outs["secretKey"] = resource.MakeSecret(
			resource.NewStringProperty(fmt.Sprintf("mocked-secret-key-%s", args.Name)),
		)
	}

	if len(args.TypeToken) >= 7 && args.TypeToken[:7] == "hcloud:" {
		id = "1"
	}
	if args.TypeToken == "hcloud:index/primaryIp:PrimaryIp" {
		outs["ipAddress"] = resource.NewStringProperty(
			fmt.Sprintf(
				"mocked-ip-address-%s-%s",
				args.Inputs["type"].StringValue(),
				args.Inputs["location"].StringValue(),
			),
		)
	}

	if args.TypeToken == "pulumiservice:index:AccessToken" {
		outs["value"] = resource.MakeSecret(
			resource.NewStringProperty(fmt.Sprintf("mocked-pulumi-access-token-%s", args.Name)),
		)
	}

	return id, outs, nil
}

func (Mocks) Call(
	args pulumi.MockCallArgs,
) (resource.PropertyMap, error) {
	log.Info().Msgf("Mocking call of type %s with args: %+v", args.Token, args.Args)

	outs := args.Args
	if _, ok := outs["name"]; !ok {
		outs["name"] = resource.NewStringProperty(args.Args["name"].StringValue())
	}

	if args.Token == "hcloud:index/getNetwork:getNetwork" {
		outs["id"] = resource.NewStringProperty(fmt.Sprintf("mocked-network-id-%s", args.Args["name"].StringValue()))
	}

	return args.Args, nil
}
