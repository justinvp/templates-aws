package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/elasticbeanstalk"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		example, err := elasticbeanstalk.LookupApplication(ctx, &elasticbeanstalk.LookupApplicationArgs{
			Name: "example",
		}, nil)
		if err != nil {
			return err
		}
		ctx.Export("arn", example.Arn)
		ctx.Export("description", example.Description)
		return nil
	})
}

