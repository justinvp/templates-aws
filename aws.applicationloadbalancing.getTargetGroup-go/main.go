package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/lb"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		opt0 := lbTgArn
		opt1 := lbTgName
		_, err := lb.LookupTargetGroup(ctx, &lb.LookupTargetGroupArgs{
			Arn:  &opt0,
			Name: &opt1,
		}, nil)
		if err != nil {
			return err
		}
		return nil
	})
}

