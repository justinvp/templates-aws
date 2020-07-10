package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/sns"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := sns.LookupTopic(ctx, &sns.LookupTopicArgs{
			Name: "an_example_topic",
		}, nil)
		if err != nil {
			return err
		}
		return nil
	})
}

