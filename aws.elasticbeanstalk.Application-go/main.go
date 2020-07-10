package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/elasticbeanstalk"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := elasticbeanstalk.NewApplication(ctx, "tftest", &elasticbeanstalk.ApplicationArgs{
			AppversionLifecycle: &elasticbeanstalk.ApplicationAppversionLifecycleArgs{
				DeleteSourceFromS3: pulumi.Bool(true),
				MaxCount:           pulumi.Int(128),
				ServiceRole:        pulumi.Any(aws_iam_role.Beanstalk_service.Arn),
			},
			Description: pulumi.String("tf-test-desc"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}

