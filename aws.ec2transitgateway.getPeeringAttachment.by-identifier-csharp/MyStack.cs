using Pulumi;
using Aws = Pulumi.Aws;

class MyStack : Stack
{
    public MyStack()
    {
        var attachment = Output.Create(Aws.Ec2TransitGateway.GetPeeringAttachment.InvokeAsync(new Aws.Ec2TransitGateway.GetPeeringAttachmentArgs
        {
            Id = "tgw-attach-12345678",
        }));
    }

}

