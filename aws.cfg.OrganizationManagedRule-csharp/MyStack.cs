using Pulumi;
using Aws = Pulumi.Aws;

class MyStack : Stack
{
    public MyStack()
    {
        var exampleOrganization = new Aws.Organizations.Organization("exampleOrganization", new Aws.Organizations.OrganizationArgs
        {
            AwsServiceAccessPrincipals = 
            {
                "config-multiaccountsetup.amazonaws.com",
            },
            FeatureSet = "ALL",
        });
        var exampleOrganizationManagedRule = new Aws.Cfg.OrganizationManagedRule("exampleOrganizationManagedRule", new Aws.Cfg.OrganizationManagedRuleArgs
        {
            RuleIdentifier = "IAM_PASSWORD_POLICY",
        }, new CustomResourceOptions
        {
            DependsOn = 
            {
                "aws_organizations_organization.example",
            },
        });
    }

}

