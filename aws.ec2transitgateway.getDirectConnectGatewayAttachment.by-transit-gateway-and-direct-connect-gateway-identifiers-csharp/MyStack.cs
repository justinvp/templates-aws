using Pulumi;
using Aws = Pulumi.Aws;

class MyStack : Stack
{
    public MyStack()
    {
        var example = Output.Create(Aws.Ec2TransitGateway.GetDirectConnectGatewayAttachment.InvokeAsync(new Aws.Ec2TransitGateway.GetDirectConnectGatewayAttachmentArgs
        {
            DxGatewayId = aws_dx_gateway.Example.Id,
            TransitGatewayId = aws_ec2_transit_gateway.Example.Id,
        }));
    }

}

