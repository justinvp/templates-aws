package main

import (
	"fmt"

	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/apigateway"
	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/iam"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		cloudwatchRole, err := iam.NewRole(ctx, "cloudwatchRole", &iam.RoleArgs{
			AssumeRolePolicy: pulumi.String(fmt.Sprintf("%v%v%v%v%v%v%v%v%v%v%v%v%v%v", "{\n", "  \"Version\": \"2012-10-17\",\n", "  \"Statement\": [\n", "    {\n", "      \"Sid\": \"\",\n", "      \"Effect\": \"Allow\",\n", "      \"Principal\": {\n", "        \"Service\": \"apigateway.amazonaws.com\"\n", "      },\n", "      \"Action\": \"sts:AssumeRole\"\n", "    }\n", "  ]\n", "}\n", "\n")),
		})
		if err != nil {
			return err
		}
		_, err = apigateway.NewAccount(ctx, "demo", &apigateway.AccountArgs{
			CloudwatchRoleArn: cloudwatchRole.Arn,
		})
		if err != nil {
			return err
		}
		_, err = iam.NewRolePolicy(ctx, "cloudwatchRolePolicy", &iam.RolePolicyArgs{
			Policy: pulumi.String(fmt.Sprintf("%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v", "{\n", "    \"Version\": \"2012-10-17\",\n", "    \"Statement\": [\n", "        {\n", "            \"Effect\": \"Allow\",\n", "            \"Action\": [\n", "                \"logs:CreateLogGroup\",\n", "                \"logs:CreateLogStream\",\n", "                \"logs:DescribeLogGroups\",\n", "                \"logs:DescribeLogStreams\",\n", "                \"logs:PutLogEvents\",\n", "                \"logs:GetLogEvents\",\n", "                \"logs:FilterLogEvents\"\n", "            ],\n", "            \"Resource\": \"*\"\n", "        }\n", "    ]\n", "}\n", "\n")),
			Role:   cloudwatchRole.ID(),
		})
		if err != nil {
			return err
		}
		return nil
	})
}

