using Pulumi;
using Aws = Pulumi.Aws;

class MyStack : Stack
{
    public MyStack()
    {
        var example = new Aws.Xray.SamplingRule("example", new Aws.Xray.SamplingRuleArgs
        {
            Attributes = 
            {
                { "Hello", "Tris" },
            },
            FixedRate = 0.05,
            Host = "*",
            HttpMethod = "*",
            Priority = 10000,
            ReservoirSize = 1,
            ResourceArn = "*",
            RuleName = "example",
            ServiceName = "*",
            ServiceType = "*",
            UrlPath = "*",
            Version = 1,
        });
    }

}

