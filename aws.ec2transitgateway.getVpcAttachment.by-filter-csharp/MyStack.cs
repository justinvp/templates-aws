using Pulumi;
using Aws = Pulumi.Aws;

class MyStack : Stack
{
    public MyStack()
    {
        var example = Output.Create(Aws.Ec2TransitGateway.GetVpcAttachment.InvokeAsync(new Aws.Ec2TransitGateway.GetVpcAttachmentArgs
        {
            Filters = 
            {
                new Aws.Ec2TransitGateway.Inputs.GetVpcAttachmentFilterArgs
                {
                    Name = "vpc-id",
                    Values = 
                    {
                        "vpc-12345678",
                    },
                },
            },
        }));
    }

}

