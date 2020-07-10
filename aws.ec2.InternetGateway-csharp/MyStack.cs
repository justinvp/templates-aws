using Pulumi;
using Aws = Pulumi.Aws;

class MyStack : Stack
{
    public MyStack()
    {
        var gw = new Aws.Ec2.InternetGateway("gw", new Aws.Ec2.InternetGatewayArgs
        {
            Tags = 
            {
                { "Name", "main" },
            },
            VpcId = aws_vpc.Main.Id,
        });
    }

}

