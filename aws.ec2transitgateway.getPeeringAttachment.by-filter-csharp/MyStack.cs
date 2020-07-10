using Pulumi;
using Aws = Pulumi.Aws;

class MyStack : Stack
{
    public MyStack()
    {
        var example = Output.Create(Aws.Ec2TransitGateway.GetPeeringAttachment.InvokeAsync(new Aws.Ec2TransitGateway.GetPeeringAttachmentArgs
        {
            Filters = 
            {
                new Aws.Ec2TransitGateway.Inputs.GetPeeringAttachmentFilterArgs
                {
                    Name = "transit-gateway-attachment-id",
                    Values = 
                    {
                        "tgw-attach-12345678",
                    },
                },
            },
        }));
    }

}

