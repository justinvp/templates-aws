using Pulumi;
using Aws = Pulumi.Aws;

class MyStack : Stack
{
    public MyStack()
    {
        var myApiGatewayVpcLink = Output.Create(Aws.ApiGateway.GetVpcLink.InvokeAsync(new Aws.ApiGateway.GetVpcLinkArgs
        {
            Name = "my-vpc-link",
        }));
    }

}

