using Pulumi;
using Aws = Pulumi.Aws;

class MyStack : Stack
{
    public MyStack()
    {
        var example = Output.Create(Aws.Workspaces.GetBundle.InvokeAsync(new Aws.Workspaces.GetBundleArgs
        {
            Name = "Value with Windows 10 and Office 2016",
            Owner = "AMAZON",
        }));
    }

}

