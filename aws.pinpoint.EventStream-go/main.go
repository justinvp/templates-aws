package main

import (
	"fmt"

	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/iam"
	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/kinesis"
	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/pinpoint"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		app, err := pinpoint.NewApp(ctx, "app", nil)
		if err != nil {
			return err
		}
		testStream, err := kinesis.NewStream(ctx, "testStream", &kinesis.StreamArgs{
			ShardCount: pulumi.Int(1),
		})
		if err != nil {
			return err
		}
		testRole, err := iam.NewRole(ctx, "testRole", &iam.RoleArgs{
			AssumeRolePolicy: pulumi.String(fmt.Sprintf("%v%v%v%v%v%v%v%v%v%v%v%v%v%v", "{\n", "  \"Version\": \"2012-10-17\",\n", "  \"Statement\": [\n", "    {\n", "      \"Action\": \"sts:AssumeRole\",\n", "      \"Principal\": {\n", "        \"Service\": \"pinpoint.us-east-1.amazonaws.com\"\n", "      },\n", "      \"Effect\": \"Allow\",\n", "      \"Sid\": \"\"\n", "    }\n", "  ]\n", "}\n", "\n")),
		})
		if err != nil {
			return err
		}
		_, err = pinpoint.NewEventStream(ctx, "stream", &pinpoint.EventStreamArgs{
			ApplicationId:        app.ApplicationId,
			DestinationStreamArn: testStream.Arn,
			RoleArn:              testRole.Arn,
		})
		if err != nil {
			return err
		}
		_, err = iam.NewRolePolicy(ctx, "testRolePolicy", &iam.RolePolicyArgs{
			Policy: pulumi.String(fmt.Sprintf("%v%v%v%v%v%v%v%v%v%v%v%v%v%v", "{\n", "  \"Version\": \"2012-10-17\",\n", "  \"Statement\": {\n", "    \"Action\": [\n", "      \"kinesis:PutRecords\",\n", "      \"kinesis:DescribeStream\"\n", "    ],\n", "    \"Effect\": \"Allow\",\n", "    \"Resource\": [\n", "      \"arn:aws:kinesis:us-east-1:*:*/*\"\n", "    ]\n", "  }\n", "}\n", "\n")),
			Role:   testRole.ID(),
		})
		if err != nil {
			return err
		}
		return nil
	})
}

