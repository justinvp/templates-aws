using Pulumi;
using Aws = Pulumi.Aws;

class MyStack : Stack
{
    public MyStack()
    {
        var gw = new Aws.Ec2.NatGateway("gw", new Aws.Ec2.NatGatewayArgs
        {
            AllocationId = aws_eip.Nat.Id,
            SubnetId = aws_subnet.Example.Id,
        });
    }

}

