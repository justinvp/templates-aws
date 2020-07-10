package main

import (
	"fmt"

	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws"
	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/iam"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		current, err := aws.GetPartition(ctx, nil, nil)
		if err != nil {
			return err
		}
		_, err = iam.GetPolicyDocument(ctx, &iam.GetPolicyDocumentArgs{
			Statements: []iam.GetPolicyDocumentStatement{
				iam.GetPolicyDocumentStatement{
					Actions: []string{
						"s3:ListBucket",
					},
					Resources: []string{
						fmt.Sprintf("%v%v%v", "arn:", current.Partition, ":s3:::my-bucket"),
					},
					Sid: "1",
				},
			},
		}, nil)
		if err != nil {
			return err
		}
		return nil
	})
}

