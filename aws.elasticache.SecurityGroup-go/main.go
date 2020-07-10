package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/ec2"
	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/elasticache"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		barSecurityGroup, err := ec2.NewSecurityGroup(ctx, "barSecurityGroup", nil)
		if err != nil {
			return err
		}
		_, err = elasticache.NewSecurityGroup(ctx, "barElasticache_securityGroupSecurityGroup", &elasticache.SecurityGroupArgs{
			SecurityGroupNames: pulumi.StringArray{
				barSecurityGroup.Name,
			},
		})
		if err != nil {
			return err
		}
		return nil
	})
}

