using Pulumi;
using Aws = Pulumi.Aws;

class MyStack : Stack
{
    public MyStack()
    {
        var org = new Aws.Organizations.Organization("org", new Aws.Organizations.OrganizationArgs
        {
            AwsServiceAccessPrincipals = 
            {
                "cloudtrail.amazonaws.com",
                "config.amazonaws.com",
            },
            FeatureSet = "ALL",
        });
    }

}

