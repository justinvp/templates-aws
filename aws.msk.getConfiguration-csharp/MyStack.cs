using Pulumi;
using Aws = Pulumi.Aws;

class MyStack : Stack
{
    public MyStack()
    {
        var example = Output.Create(Aws.Msk.GetConfiguration.InvokeAsync(new Aws.Msk.GetConfigurationArgs
        {
            Name = "example",
        }));
    }

}

