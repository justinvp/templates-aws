package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/cloudformation"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := cloudformation.NewStackSetInstance(ctx, "example", &cloudformation.StackSetInstanceArgs{
			AccountId:    pulumi.String("123456789012"),
			Region:       pulumi.String("us-east-1"),
			StackSetName: pulumi.Any(aws_cloudformation_stack_set.Example.Name),
		})
		if err != nil {
			return err
		}
		return nil
	})
}

