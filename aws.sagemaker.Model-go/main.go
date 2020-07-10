package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/iam"
	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/sagemaker"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := sagemaker.NewModel(ctx, "model", &sagemaker.ModelArgs{
			ExecutionRoleArn: pulumi.Any(aws_iam_role.Foo.Arn),
			PrimaryContainer: &sagemaker.ModelPrimaryContainerArgs{
				Image: pulumi.String("174872318107.dkr.ecr.us-west-2.amazonaws.com/kmeans:1"),
			},
		})
		if err != nil {
			return err
		}
		assumeRole, err := iam.GetPolicyDocument(ctx, &iam.GetPolicyDocumentArgs{
			Statements: []iam.GetPolicyDocumentStatement{
				iam.GetPolicyDocumentStatement{
					Actions: []string{
						"sts:AssumeRole",
					},
					Principals: []iam.GetPolicyDocumentStatementPrincipal{
						iam.GetPolicyDocumentStatementPrincipal{
							Identifiers: []string{
								"sagemaker.amazonaws.com",
							},
							Type: "Service",
						},
					},
				},
			},
		}, nil)
		if err != nil {
			return err
		}
		_, err = iam.NewRole(ctx, "role", &iam.RoleArgs{
			AssumeRolePolicy: pulumi.String(assumeRole.Json),
		})
		if err != nil {
			return err
		}
		return nil
	})
}

