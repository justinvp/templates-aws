using Pulumi;
using Aws = Pulumi.Aws;

class MyStack : Stack
{
    public MyStack()
    {
        var test = Output.Create(Aws.Ec2.GetVpcEndpointService.InvokeAsync(new Aws.Ec2.GetVpcEndpointServiceArgs
        {
            Filters = 
            {
                new Aws.Ec2.Inputs.GetVpcEndpointServiceFilterArgs
                {
                    Name = "service-name",
                    Values = 
                    {
                        "some-service",
                    },
                },
            },
        }));
    }

}

