using Pulumi;
using Aws = Pulumi.Aws;

class MyStack : Stack
{
    public MyStack()
    {
        var exampleAccount = new Aws.SecurityHub.Account("exampleAccount", new Aws.SecurityHub.AccountArgs
        {
        });
        var current = Output.Create(Aws.GetRegion.InvokeAsync());
        var exampleProductSubscription = new Aws.SecurityHub.ProductSubscription("exampleProductSubscription", new Aws.SecurityHub.ProductSubscriptionArgs
        {
            ProductArn = current.Apply(current => $"arn:aws:securityhub:{current.Name}:733251395267:product/alertlogic/althreatmanagement"),
        }, new CustomResourceOptions
        {
            DependsOn = 
            {
                "aws_securityhub_account.example",
            },
        });
    }

}

