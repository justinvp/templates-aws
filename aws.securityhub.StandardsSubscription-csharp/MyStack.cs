using Pulumi;
using Aws = Pulumi.Aws;

class MyStack : Stack
{
    public MyStack()
    {
        var example = new Aws.SecurityHub.Account("example", new Aws.SecurityHub.AccountArgs
        {
        });
        var cis = new Aws.SecurityHub.StandardsSubscription("cis", new Aws.SecurityHub.StandardsSubscriptionArgs
        {
            StandardsArn = "arn:aws:securityhub:::ruleset/cis-aws-foundations-benchmark/v/1.2.0",
        }, new CustomResourceOptions
        {
            DependsOn = 
            {
                "aws_securityhub_account.example",
            },
        });
        var pci321 = new Aws.SecurityHub.StandardsSubscription("pci321", new Aws.SecurityHub.StandardsSubscriptionArgs
        {
            StandardsArn = "arn:aws:securityhub:us-east-1::standards/pci-dss/v/3.2.1",
        }, new CustomResourceOptions
        {
            DependsOn = 
            {
                "aws_securityhub_account.example",
            },
        });
    }

}

