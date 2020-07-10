using Pulumi;
using Aws = Pulumi.Aws;

class MyStack : Stack
{
    public MyStack()
    {
        var examplePermission = new Aws.Lambda.Permission("examplePermission", new Aws.Lambda.PermissionArgs
        {
            Action = "lambda:InvokeFunction",
            Function = aws_lambda_function.Example.Arn,
            Principal = "config.amazonaws.com",
        });
        var exampleOrganization = new Aws.Organizations.Organization("exampleOrganization", new Aws.Organizations.OrganizationArgs
        {
            AwsServiceAccessPrincipals = 
            {
                "config-multiaccountsetup.amazonaws.com",
            },
            FeatureSet = "ALL",
        });
        var exampleOrganizationCustomRule = new Aws.Cfg.OrganizationCustomRule("exampleOrganizationCustomRule", new Aws.Cfg.OrganizationCustomRuleArgs
        {
            LambdaFunctionArn = aws_lambda_function.Example.Arn,
            TriggerTypes = 
            {
                "ConfigurationItemChangeNotification",
            },
        }, new CustomResourceOptions
        {
            DependsOn = 
            {
                "aws_lambda_permission.example",
                "aws_organizations_organization.example",
            },
        });
    }

}

