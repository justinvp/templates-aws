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
                "guardduty.amazonaws.com",
            },
            FeatureSet = "ALL",
        });
        var exampleDetector = new Aws.GuardDuty.Detector("exampleDetector", new Aws.GuardDuty.DetectorArgs
        {
        });
        var exampleOrganizationAdminAccount = new Aws.GuardDuty.OrganizationAdminAccount("exampleOrganizationAdminAccount", new Aws.GuardDuty.OrganizationAdminAccountArgs
        {
            AdminAccountId = "123456789012",
        }, new CustomResourceOptions
        {
            DependsOn = 
            {
                exampleOrganization,
            },
        });
    }

}

