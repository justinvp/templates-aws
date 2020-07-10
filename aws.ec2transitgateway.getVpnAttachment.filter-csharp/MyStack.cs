using Pulumi;
using Aws = Pulumi.Aws;

class MyStack : Stack
{
    public MyStack()
    {
        var test = Output.Create(Aws.Ec2TransitGateway.GetVpnAttachment.InvokeAsync(new Aws.Ec2TransitGateway.GetVpnAttachmentArgs
        {
            Filters = 
            {
                new Aws.Ec2TransitGateway.Inputs.GetVpnAttachmentFilterArgs
                {
                    Name = "resource-id",
                    Values = 
                    {
                        "some-resource",
                    },
                },
            },
        }));
    }

}

