using Pulumi;
using Aws = Pulumi.Aws;

class MyStack : Stack
{
    public MyStack()
    {
        var example = new Aws.Ec2.TransitGatewayPeeringAttachmentAccepter("example", new Aws.Ec2.TransitGatewayPeeringAttachmentAccepterArgs
        {
            Tags = 
            {
                { "Name", "Example cross-account attachment" },
            },
            TransitGatewayAttachmentId = aws_ec2_transit_gateway_peering_attachment.Example.Id,
        });
    }

}

