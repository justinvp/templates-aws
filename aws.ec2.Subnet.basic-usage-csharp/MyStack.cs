using Pulumi;
using Aws = Pulumi.Aws;

class MyStack : Stack
{
    public MyStack()
    {
        var main = new Aws.Ec2.Subnet("main", new Aws.Ec2.SubnetArgs
        {
            CidrBlock = "10.0.1.0/24",
            Tags = 
            {
                { "Name", "Main" },
            },
            VpcId = aws_vpc.Main.Id,
        });
    }

}

