using Pulumi;
using Aws = Pulumi.Aws;

class MyStack : Stack
{
    public MyStack()
    {
        var example = Output.Create(Aws.Ram.GetResourceShare.InvokeAsync(new Aws.Ram.GetResourceShareArgs
        {
            Name = "example",
            ResourceOwner = "SELF",
        }));
    }

}

