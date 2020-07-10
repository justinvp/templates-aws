using Pulumi;
using Aws = Pulumi.Aws;

class MyStack : Stack
{
    public MyStack()
    {
        var example = new Aws.Ec2TransitGateway.VpcAttachmentAccepter("example", new Aws.Ec2TransitGateway.VpcAttachmentAccepterArgs
        {
            Tags = 
            {
                { "Name", "Example cross-account attachment" },
            },
            TransitGatewayAttachmentId = aws_ec2_transit_gateway_vpc_attachment.Example.Id,
        });
    }

}

