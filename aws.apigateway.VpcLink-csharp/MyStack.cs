using Pulumi;
using Aws = Pulumi.Aws;

class MyStack : Stack
{
    public MyStack()
    {
        var exampleLoadBalancer = new Aws.LB.LoadBalancer("exampleLoadBalancer", new Aws.LB.LoadBalancerArgs
        {
            Internal = true,
            LoadBalancerType = "network",
            SubnetMappings = 
            {
                new Aws.LB.Inputs.LoadBalancerSubnetMappingArgs
                {
                    SubnetId = "12345",
                },
            },
        });
        var exampleVpcLink = new Aws.ApiGateway.VpcLink("exampleVpcLink", new Aws.ApiGateway.VpcLinkArgs
        {
            Description = "example description",
            TargetArn = exampleLoadBalancer.Arn,
        });
    }

}

