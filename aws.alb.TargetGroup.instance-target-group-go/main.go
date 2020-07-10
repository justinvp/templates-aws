package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/ec2"
	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/lb"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		main, err := ec2.NewVpc(ctx, "main", &ec2.VpcArgs{
			CidrBlock: pulumi.String("10.0.0.0/16"),
		})
		if err != nil {
			return err
		}
		_, err = lb.NewTargetGroup(ctx, "test", &lb.TargetGroupArgs{
			Port:     pulumi.Int(80),
			Protocol: pulumi.String("HTTP"),
			VpcId:    main.ID(),
		})
		if err != nil {
			return err
		}
		return nil
	})
}

