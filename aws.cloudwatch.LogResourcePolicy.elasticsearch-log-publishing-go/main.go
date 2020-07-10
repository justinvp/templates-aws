package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/cloudwatch"
	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/iam"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		elasticsearch_log_publishing_policyPolicyDocument, err := iam.GetPolicyDocument(ctx, &iam.GetPolicyDocumentArgs{
			Statements: []iam.GetPolicyDocumentStatement{
				iam.GetPolicyDocumentStatement{
					Actions: []string{
						"logs:CreateLogStream",
						"logs:PutLogEvents",
						"logs:PutLogEventsBatch",
					},
					Principals: []iam.GetPolicyDocumentStatementPrincipal{
						iam.GetPolicyDocumentStatementPrincipal{
							Identifiers: []string{
								"es.amazonaws.com",
							},
							Type: "Service",
						},
					},
					Resources: []string{
						"arn:aws:logs:*",
					},
				},
			},
		}, nil)
		if err != nil {
			return err
		}
		_, err = cloudwatch.NewLogResourcePolicy(ctx, "elasticsearch_log_publishing_policyLogResourcePolicy", &cloudwatch.LogResourcePolicyArgs{
			PolicyDocument: pulumi.String(elasticsearch_log_publishing_policyPolicyDocument.Json),
			PolicyName:     pulumi.String("elasticsearch-log-publishing-policy"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}

