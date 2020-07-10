using Pulumi;
using Aws = Pulumi.Aws;

class MyStack : Stack
{
    public MyStack()
    {
        var fooLocalGateways = Output.Create(Aws.Ec2.GetLocalGateways.InvokeAsync(new Aws.Ec2.GetLocalGatewaysArgs
        {
            Tags = 
            {
                { "service", "production" },
            },
        }));
        this.Foo = fooLocalGateways.Apply(fooLocalGateways => fooLocalGateways.Ids);
    }

    [Output("foo")]
    public Output<string> Foo { get; set; }
}

