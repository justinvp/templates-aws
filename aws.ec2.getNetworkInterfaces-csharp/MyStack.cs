using Pulumi;
using Aws = Pulumi.Aws;

class MyStack : Stack
{
    public MyStack()
    {
        var exampleNetworkInterfaces = Output.Create(Aws.Ec2.GetNetworkInterfaces.InvokeAsync());
        this.Example = exampleNetworkInterfaces.Apply(exampleNetworkInterfaces => exampleNetworkInterfaces.Ids);
    }

    [Output("example")]
    public Output<string> Example { get; set; }
}

