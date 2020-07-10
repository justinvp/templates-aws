using Pulumi;
using Aws = Pulumi.Aws;

class MyStack : Stack
{
    public MyStack()
    {
        var ec2 = new Aws.Ec2.VpcEndpoint("ec2", new Aws.Ec2.VpcEndpointArgs
        {
            PrivateDnsEnabled = true,
            SecurityGroupIds = 
            {
                aws_security_group.Sg1.Id,
            },
            ServiceName = "com.amazonaws.us-west-2.ec2",
            VpcEndpointType = "Interface",
            VpcId = aws_vpc.Main.Id,
        });
    }

}

