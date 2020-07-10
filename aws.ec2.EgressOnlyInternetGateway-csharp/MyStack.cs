using Pulumi;
using Aws = Pulumi.Aws;

class MyStack : Stack
{
    public MyStack()
    {
        var exampleVpc = new Aws.Ec2.Vpc("exampleVpc", new Aws.Ec2.VpcArgs
        {
            AssignGeneratedIpv6CidrBlock = true,
            CidrBlock = "10.1.0.0/16",
        });
        var exampleEgressOnlyInternetGateway = new Aws.Ec2.EgressOnlyInternetGateway("exampleEgressOnlyInternetGateway", new Aws.Ec2.EgressOnlyInternetGatewayArgs
        {
            Tags = 
            {
                { "Name", "main" },
            },
            VpcId = exampleVpc.Id,
        });
    }

}

