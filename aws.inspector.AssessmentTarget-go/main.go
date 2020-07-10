package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/inspector"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		bar, err := inspector.NewResourceGroup(ctx, "bar", &inspector.ResourceGroupArgs{
			Tags: pulumi.StringMap{
				"Env":  pulumi.String("bar"),
				"Name": pulumi.String("foo"),
			},
		})
		if err != nil {
			return err
		}
		_, err = inspector.NewAssessmentTarget(ctx, "foo", &inspector.AssessmentTargetArgs{
			ResourceGroupArn: bar.Arn,
		})
		if err != nil {
			return err
		}
		return nil
	})
}

