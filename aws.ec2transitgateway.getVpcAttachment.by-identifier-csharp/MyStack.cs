using Pulumi;
using Aws = Pulumi.Aws;

class MyStack : Stack
{
    public MyStack()
    {
        var example = Output.Create(Aws.Ec2TransitGateway.GetVpcAttachment.InvokeAsync(new Aws.Ec2TransitGateway.GetVpcAttachmentArgs
        {
            Id = "tgw-attach-12345678",
        }));
    }

}

