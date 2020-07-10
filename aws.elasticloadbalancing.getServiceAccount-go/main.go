package main

import (
	"fmt"

	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/elb"
	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/s3"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		main, err := elb.GetServiceAccount(ctx, nil, nil)
		if err != nil {
			return err
		}
		elbLogs, err := s3.NewBucket(ctx, "elbLogs", &s3.BucketArgs{
			Acl:    pulumi.String("private"),
			Policy: pulumi.String(fmt.Sprintf("%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v", "{\n", "  \"Id\": \"Policy\",\n", "  \"Version\": \"2012-10-17\",\n", "  \"Statement\": [\n", "    {\n", "      \"Action\": [\n", "        \"s3:PutObject\"\n", "      ],\n", "      \"Effect\": \"Allow\",\n", "      \"Resource\": \"arn:aws:s3:::my-elb-tf-test-bucket/AWSLogs/*\",\n", "      \"Principal\": {\n", "        \"AWS\": [\n", "          \"", main.Arn, "\"\n", "        ]\n", "      }\n", "    }\n", "  ]\n", "}\n", "\n")),
		})
		if err != nil {
			return err
		}
		_, err = elb.NewLoadBalancer(ctx, "bar", &elb.LoadBalancerArgs{
			AccessLogs: &elb.LoadBalancerAccessLogsArgs{
				Bucket:   elbLogs.Bucket,
				Interval: pulumi.Int(5),
			},
			AvailabilityZones: pulumi.StringArray{
				pulumi.String("us-west-2a"),
			},
			Listeners: elb.LoadBalancerListenerArray{
				&elb.LoadBalancerListenerArgs{
					InstancePort:     pulumi.Int(8000),
					InstanceProtocol: pulumi.String("http"),
					LbPort:           pulumi.Int(80),
					LbProtocol:       pulumi.String("http"),
				},
			},
		})
		if err != nil {
			return err
		}
		return nil
	})
}

