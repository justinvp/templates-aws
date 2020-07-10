using Pulumi;
using Aws = Pulumi.Aws;

class MyStack : Stack
{
    public MyStack()
    {
        var example = new Aws.LB.LoadBalancer("example", new Aws.LB.LoadBalancerArgs
        {
            LoadBalancerType = "network",
            SubnetMappings = 
            {
                new Aws.LB.Inputs.LoadBalancerSubnetMappingArgs
                {
                    AllocationId = aws_eip.Example1.Id,
                    SubnetId = aws_subnet.Example1.Id,
                },
                new Aws.LB.Inputs.LoadBalancerSubnetMappingArgs
                {
                    AllocationId = aws_eip.Example2.Id,
                    SubnetId = aws_subnet.Example2.Id,
                },
            },
        });
    }

}

