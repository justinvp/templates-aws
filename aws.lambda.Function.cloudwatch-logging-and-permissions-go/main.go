package main

import (
	"fmt"

	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/cloudwatch"
	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/iam"
	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/lambda"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := lambda.NewFunction(ctx, "testLambda", nil, pulumi.DependsOn([]pulumi.Resource{
			"aws_cloudwatch_log_group.example",
			"aws_iam_role_policy_attachment.lambda_logs",
		}))
		if err != nil {
			return err
		}
		_, err = cloudwatch.NewLogGroup(ctx, "example", &cloudwatch.LogGroupArgs{
			RetentionInDays: pulumi.Int(14),
		})
		if err != nil {
			return err
		}
		lambdaLogging, err := iam.NewPolicy(ctx, "lambdaLogging", &iam.PolicyArgs{
			Description: pulumi.String("IAM policy for logging from a lambda"),
			Path:        pulumi.String("/"),
			Policy:      pulumi.String(fmt.Sprintf("%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v", "{\n", "  \"Version\": \"2012-10-17\",\n", "  \"Statement\": [\n", "    {\n", "      \"Action\": [\n", "        \"logs:CreateLogGroup\",\n", "        \"logs:CreateLogStream\",\n", "        \"logs:PutLogEvents\"\n", "      ],\n", "      \"Resource\": \"arn:aws:logs:*:*:*\",\n", "      \"Effect\": \"Allow\"\n", "    }\n", "  ]\n", "}\n", "\n")),
		})
		if err != nil {
			return err
		}
		_, err = iam.NewRolePolicyAttachment(ctx, "lambdaLogs", &iam.RolePolicyAttachmentArgs{
			PolicyArn: lambdaLogging.Arn,
			Role:      pulumi.Any(aws_iam_role.Iam_for_lambda.Name),
		})
		if err != nil {
			return err
		}
		return nil
	})
}

