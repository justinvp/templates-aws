using Pulumi;
using Aws = Pulumi.Aws;

class MyStack : Stack
{
    public MyStack()
    {
        var example = Output.Create(Aws.Ec2.GetVpcDhcpOptions.InvokeAsync(new Aws.Ec2.GetVpcDhcpOptionsArgs
        {
            DhcpOptionsId = "dopts-12345678",
        }));
    }

}

