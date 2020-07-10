package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/swf"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := swf.NewDomain(ctx, "foo", &swf.DomainArgs{
			Description:                            pulumi.String("SWF Domain"),
			WorkflowExecutionRetentionPeriodInDays: pulumi.String("30"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}

