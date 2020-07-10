using Pulumi;
using Aws = Pulumi.Aws;

class MyStack : Stack
{
    public MyStack()
    {
        var example = new Aws.Ec2.VpcEndpointService("example", new Aws.Ec2.VpcEndpointServiceArgs
        {
            AcceptanceRequired = false,
            NetworkLoadBalancerArns = 
            {
                aws_lb.Example.Arn,
            },
        });
    }

}

