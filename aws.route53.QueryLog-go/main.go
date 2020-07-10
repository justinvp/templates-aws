package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/cloudwatch"
	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/iam"
	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/providers"
	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/route53"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := providers.Newaws(ctx, "us_east_1", &providers.awsArgs{
			Region: pulumi.String("us-east-1"),
		})
		if err != nil {
			return err
		}
		awsRoute53ExampleCom, err := cloudwatch.NewLogGroup(ctx, "awsRoute53ExampleCom", &cloudwatch.LogGroupArgs{
			RetentionInDays: pulumi.Int(30),
		}, pulumi.Provider("aws.us-east-1"))
		if err != nil {
			return err
		}
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
		}, pulumi.Provider("aws.us-east-1"))
		if err != nil {
			return err
		}
		exampleComZone, err := route53.NewZone(ctx, "exampleComZone", nil)
		if err != nil {
			return err
		}
		_, err = route53.NewQueryLog(ctx, "exampleComQueryLog", &route53.QueryLogArgs{
			CloudwatchLogGroupArn: awsRoute53ExampleCom.Arn,
			ZoneId:                exampleComZone.ZoneId,
		}, pulumi.DependsOn([]pulumi.Resource{
			"aws_cloudwatch_log_resource_policy.route53-query-logging-policy",
		}))
		if err != nil {
			return err
		}
		return nil
	})
}

