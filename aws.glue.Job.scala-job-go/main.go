package main

import (
	"fmt"

	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/glue"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := glue.NewJob(ctx, "example", &glue.JobArgs{
			Command: &glue.JobCommandArgs{
				ScriptLocation: pulumi.String(fmt.Sprintf("%v%v%v", "s3://", aws_s3_bucket.Example.Bucket, "/example.scala")),
			},
			DefaultArguments: pulumi.StringMap{
				"--job-language": pulumi.String("scala"),
			},
			RoleArn: pulumi.Any(aws_iam_role.Example.Arn),
		})
		if err != nil {
			return err
		}
		return nil
	})
}

