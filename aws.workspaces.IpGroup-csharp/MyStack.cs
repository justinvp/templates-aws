using Pulumi;
using Aws = Pulumi.Aws;

class MyStack : Stack
{
    public MyStack()
    {
        var contractors = new Aws.Workspaces.IpGroup("contractors", new Aws.Workspaces.IpGroupArgs
        {
            Description = "Contractors IP access control group",
        });
    }

}

