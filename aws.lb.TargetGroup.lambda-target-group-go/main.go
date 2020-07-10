package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/lb"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := lb.NewTargetGroup(ctx, "lambda_example", &lb.TargetGroupArgs{
			TargetType: pulumi.String("lambda"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}

