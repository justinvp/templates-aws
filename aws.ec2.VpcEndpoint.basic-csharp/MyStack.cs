using Pulumi;
using Aws = Pulumi.Aws;

class MyStack : Stack
{
    public MyStack()
    {
        var s3 = new Aws.Ec2.VpcEndpoint("s3", new Aws.Ec2.VpcEndpointArgs
        {
            ServiceName = "com.amazonaws.us-west-2.s3",
            VpcId = aws_vpc.Main.Id,
        });
    }

}

