package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/ec2"
	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/s3"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		exampleBucket, err := s3.NewBucket(ctx, "exampleBucket", nil)
		if err != nil {
			return err
		}
		_, err = ec2.NewFlowLog(ctx, "exampleFlowLog", &ec2.FlowLogArgs{
			LogDestination:     exampleBucket.Arn,
			LogDestinationType: pulumi.String("s3"),
			TrafficType:        pulumi.String("ALL"),
			VpcId:              pulumi.Any(aws_vpc.Example.Id),
		})
		if err != nil {
			return err
		}
		return nil
	})
}

