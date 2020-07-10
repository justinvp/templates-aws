package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/cloudwatch"
	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/eks"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := eks.NewCluster(ctx, "exampleCluster", &eks.ClusterArgs{
			EnabledClusterLogTypes: pulumi.StringArray{
				pulumi.String("api"),
				pulumi.String("audit"),
			},
		}, pulumi.DependsOn([]pulumi.Resource{
			"aws_cloudwatch_log_group.example",
		}))
		if err != nil {
			return err
		}
		_, err = cloudwatch.NewLogGroup(ctx, "exampleLogGroup", &cloudwatch.LogGroupArgs{
			RetentionInDays: pulumi.Int(7),
		})
		if err != nil {
			return err
		}
		return nil
	})
}

