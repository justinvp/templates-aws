package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/cloudwatch"
	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/iam"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		route53_query_logging_policyPolicyDocument, err := iam.GetPolicyDocument(ctx, &iam.GetPolicyDocumentArgs{
			Statements: []iam.GetPolicyDocumentStatement{
				iam.GetPolicyDocumentStatement{
					Actions: []string{
						"logs:CreateLogStream",
						"logs:PutLogEvents",
					},
					Principals: []iam.GetPolicyDocumentStatementPrincipal{
						iam.GetPolicyDocumentStatementPrincipal{
							Identifiers: []string{
								"route53.amazonaws.com",
							},
							Type: "Service",
						},
					},
					Resources: []string{
						"arn:aws:logs:*:*:log-group:/aws/route53/*",
					},
				},
			},
		}, nil)
		if err != nil {
			return err
		}
		_, err = cloudwatch.NewLogResourcePolicy(ctx, "route53_query_logging_policyLogResourcePolicy", &cloudwatch.LogResourcePolicyArgs{
			PolicyDocument: pulumi.String(route53_query_logging_policyPolicyDocument.Json),
			PolicyName:     pulumi.String("route53-query-logging-policy"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}

