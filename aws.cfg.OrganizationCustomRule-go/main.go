package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/cfg"
	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/lambda"
	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/organizations"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := lambda.NewPermission(ctx, "examplePermission", &lambda.PermissionArgs{
			Action:    pulumi.String("lambda:InvokeFunction"),
			Function:  pulumi.Any(aws_lambda_function.Example.Arn),
			Principal: pulumi.String("config.amazonaws.com"),
		})
		if err != nil {
			return err
		}
		_, err = organizations.NewOrganization(ctx, "exampleOrganization", &organizations.OrganizationArgs{
			AwsServiceAccessPrincipals: pulumi.StringArray{
				pulumi.String("config-multiaccountsetup.amazonaws.com"),
			},
			FeatureSet: pulumi.String("ALL"),
		})
		if err != nil {
			return err
		}
		_, err = cfg.NewOrganizationCustomRule(ctx, "exampleOrganizationCustomRule", &cfg.OrganizationCustomRuleArgs{
			LambdaFunctionArn: pulumi.Any(aws_lambda_function.Example.Arn),
			TriggerTypes: pulumi.StringArray{
				pulumi.String("ConfigurationItemChangeNotification"),
			},
		}, pulumi.DependsOn([]pulumi.Resource{
			"aws_lambda_permission.example",
			"aws_organizations_organization.example",
		}))
		if err != nil {
			return err
		}
		return nil
	})
}

