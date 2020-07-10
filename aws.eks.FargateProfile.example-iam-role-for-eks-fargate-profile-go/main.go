package main

import (
	"encoding/json"

	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/iam"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		tmpJSON0, err := json.Marshal(map[string]interface{}{
			"Statement": []map[string]interface{}{
				map[string]interface{}{
					"Action": "sts:AssumeRole",
					"Effect": "Allow",
					"Principal": map[string]interface{}{
						"Service": "eks-fargate-pods.amazonaws.com",
					},
				},
			},
			"Version": "2012-10-17",
		})
		if err != nil {
			return err
		}
		json0 := string(tmpJSON0)
		example, err := iam.NewRole(ctx, "example", &iam.RoleArgs{
			AssumeRolePolicy: pulumi.String(json0),
		})
		if err != nil {
			return err
		}
		_, err = iam.NewRolePolicyAttachment(ctx, "example_AmazonEKSFargatePodExecutionRolePolicy", &iam.RolePolicyAttachmentArgs{
			PolicyArn: pulumi.String("arn:aws:iam::aws:policy/AmazonEKSFargatePodExecutionRolePolicy"),
			Role:      example.Name,
		})
		if err != nil {
			return err
		}
		return nil
	})
}

