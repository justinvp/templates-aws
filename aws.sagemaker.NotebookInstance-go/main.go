package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/sagemaker"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := sagemaker.NewNotebookInstance(ctx, "ni", &sagemaker.NotebookInstanceArgs{
			InstanceType: pulumi.String("ml.t2.medium"),
			RoleArn:      pulumi.Any(aws_iam_role.Role.Arn),
			Tags: pulumi.StringMap{
				"Name": pulumi.String("foo"),
			},
		})
		if err != nil {
			return err
		}
		return nil
	})
}

