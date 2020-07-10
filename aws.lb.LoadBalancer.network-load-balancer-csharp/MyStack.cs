using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

class MyStack : Stack
{
    public MyStack()
    {
        var test = new Aws.LB.LoadBalancer("test", new Aws.LB.LoadBalancerArgs
        {
            EnableDeletionProtection = true,
            Internal = false,
            LoadBalancerType = "network",
            Subnets = 
            {
                aws_subnet.Public.Select(__item => __item.Id).ToList(),
            },
            Tags = 
            {
                { "Environment", "production" },
            },
        });
    }

}

