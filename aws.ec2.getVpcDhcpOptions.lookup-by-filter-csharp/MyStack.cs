using Pulumi;
using Aws = Pulumi.Aws;

class MyStack : Stack
{
    public MyStack()
    {
        var example = Output.Create(Aws.Ec2.GetVpcDhcpOptions.InvokeAsync(new Aws.Ec2.GetVpcDhcpOptionsArgs
        {
            Filters = 
            {
                new Aws.Ec2.Inputs.GetVpcDhcpOptionsFilterArgs
                {
                    Name = "key",
                    Values = 
                    {
                        "domain-name",
                    },
                },
                new Aws.Ec2.Inputs.GetVpcDhcpOptionsFilterArgs
                {
                    Name = "value",
                    Values = 
                    {
                        "example.com",
                    },
                },
            },
        }));
    }

}

