using Pulumi;
using Aws = Pulumi.Aws;

class MyStack : Stack
{
    public MyStack()
    {
        var example = new Aws.Ec2TransitGateway.VpcAttachment("example", new Aws.Ec2TransitGateway.VpcAttachmentArgs
        {
            SubnetIds = 
            {
                aws_subnet.Example.Id,
            },
            TransitGatewayId = aws_ec2_transit_gateway.Example.Id,
            VpcId = aws_vpc.Example.Id,
        });
    }

}

