package main

import (
	"fmt"

	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws"
	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/cognito"
	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/iam"
	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/pinpoint"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		current, err := aws.GetCallerIdentity(ctx, nil, nil)
		if err != nil {
			return err
		}
		testUserPool, err := cognito.NewUserPool(ctx, "testUserPool", nil)
		if err != nil {
			return err
		}
		testApp, err := pinpoint.NewApp(ctx, "testApp", nil)
		if err != nil {
			return err
		}
		testRole, err := iam.NewRole(ctx, "testRole", &iam.RoleArgs{
			AssumeRolePolicy: pulumi.String(fmt.Sprintf("%v%v%v%v%v%v%v%v%v%v%v%v%v%v", "{\n", "  \"Version\": \"2012-10-17\",\n", "  \"Statement\": [\n", "    {\n", "      \"Action\": \"sts:AssumeRole\",\n", "      \"Principal\": {\n", "        \"Service\": \"cognito-idp.amazonaws.com\"\n", "      },\n", "      \"Effect\": \"Allow\",\n", "      \"Sid\": \"\"\n", "    }\n", "  ]\n", "}\n", "\n")),
		})
		if err != nil {
			return err
		}
		_, err = iam.NewRolePolicy(ctx, "testRolePolicy", &iam.RolePolicyArgs{
			Policy: testApp.ApplicationId.ApplyT(func(applicationId string) (string, error) {
				return fmt.Sprintf("%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v", "{\n", "  \"Version\": \"2012-10-17\",\n", "  \"Statement\": [\n", "    {\n", "      \"Action\": [\n", "        \"mobiletargeting:UpdateEndpoint\",\n", "        \"mobiletargeting:PutItems\"\n", "      ],\n", "      \"Effect\": \"Allow\",\n", "      \"Resource\": \"arn:aws:mobiletargeting:*:", current.AccountId, ":apps/", applicationId, "*\"\n", "    }\n", "  ]\n", "}\n", "\n"), nil
			}).(pulumi.StringOutput),
			Role: testRole.ID(),
		})
		if err != nil {
			return err
		}
		_, err = cognito.NewUserPoolClient(ctx, "testUserPoolClient", &cognito.UserPoolClientArgs{
			AnalyticsConfiguration: &cognito.UserPoolClientAnalyticsConfigurationArgs{
				ApplicationId:  testApp.ApplicationId,
				ExternalId:     pulumi.String("some_id"),
				RoleArn:        testRole.Arn,
				UserDataShared: pulumi.Bool(true),
			},
			UserPoolId: testUserPool.ID(),
		})
		if err != nil {
			return err
		}
		return nil
	})
}

