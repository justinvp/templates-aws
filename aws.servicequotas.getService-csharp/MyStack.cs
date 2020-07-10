using Pulumi;
using Aws = Pulumi.Aws;

class MyStack : Stack
{
    public MyStack()
    {
        var example = Output.Create(Aws.ServiceQuotas.GetService.InvokeAsync(new Aws.ServiceQuotas.GetServiceArgs
        {
            ServiceName = "Amazon Virtual Private Cloud (Amazon VPC)",
        }));
    }

}

