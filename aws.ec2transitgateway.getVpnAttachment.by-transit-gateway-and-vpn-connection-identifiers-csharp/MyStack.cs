using Pulumi;
using Aws = Pulumi.Aws;

class MyStack : Stack
{
    public MyStack()
    {
        var example = Output.Create(Aws.Ec2TransitGateway.GetVpnAttachment.InvokeAsync(new Aws.Ec2TransitGateway.GetVpnAttachmentArgs
        {
            TransitGatewayId = aws_ec2_transit_gateway.Example.Id,
            VpnConnectionId = aws_vpn_connection.Example.Id,
        }));
    }

}

