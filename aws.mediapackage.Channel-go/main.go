package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/mediapackage"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := mediapackage.NewChannel(ctx, "kittens", &mediapackage.ChannelArgs{
			ChannelId:   pulumi.String("kitten-channel"),
			Description: pulumi.String("A channel dedicated to amusing videos of kittens."),
		})
		if err != nil {
			return err
		}
		return nil
	})
}

