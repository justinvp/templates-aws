package main

import (
	"fmt"

	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/cfg"
	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/iam"
	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/s3"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		bucket, err := s3.NewBucket(ctx, "bucket", &s3.BucketArgs{
			ForceDestroy: pulumi.Bool(true),
		})
		if err != nil {
			return err
		}
		_, err = cfg.NewDeliveryChannel(ctx, "fooDeliveryChannel", &cfg.DeliveryChannelArgs{
			S3BucketName: bucket.Bucket,
		}, pulumi.DependsOn([]pulumi.Resource{
			"aws_config_configuration_recorder.foo",
		}))
		if err != nil {
			return err
		}
		role, err := iam.NewRole(ctx, "role", &iam.RoleArgs{
			AssumeRolePolicy: pulumi.String(fmt.Sprintf("%v%v%v%v%v%v%v%v%v%v%v%v%v%v", "{\n", "  \"Version\": \"2012-10-17\",\n", "  \"Statement\": [\n", "    {\n", "      \"Action\": \"sts:AssumeRole\",\n", "      \"Principal\": {\n", "        \"Service\": \"config.amazonaws.com\"\n", "      },\n", "      \"Effect\": \"Allow\",\n", "      \"Sid\": \"\"\n", "    }\n", "  ]\n", "}\n", "\n")),
		})
		if err != nil {
			return err
		}
		_, err = cfg.NewRecorder(ctx, "fooRecorder", &cfg.RecorderArgs{
			RoleArn: role.Arn,
		})
		if err != nil {
			return err
		}
		_, err = iam.NewRolePolicy(ctx, "rolePolicy", &iam.RolePolicyArgs{
			Policy: pulumi.All(bucket.Arn, bucket.Arn).ApplyT(func(_args []interface{}) (string, error) {
				bucketArn := _args[0].(string)
				bucketArn1 := _args[1].(string)
				return fmt.Sprintf("%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v", "{\n", "  \"Version\": \"2012-10-17\",\n", "  \"Statement\": [\n", "    {\n", "      \"Action\": [\n", "        \"s3:*\"\n", "      ],\n", "      \"Effect\": \"Allow\",\n", "      \"Resource\": [\n", "        \"", bucketArn, "\",\n", "        \"", bucketArn1, "/*\"\n", "      ]\n", "    }\n", "  ]\n", "}\n", "\n"), nil
			}).(pulumi.StringOutput),
			Role: role.ID(),
		})
		if err != nil {
			return err
		}
		return nil
	})
}

