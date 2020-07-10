using Pulumi;
using Aws = Pulumi.Aws;

class MyStack : Stack
{
    public MyStack()
    {
        var test = Output.Create(Aws.CloudFront.GetDistribution.InvokeAsync(new Aws.CloudFront.GetDistributionArgs
        {
            Id = "EDFDVBD632BHDS5",
        }));
    }

}

