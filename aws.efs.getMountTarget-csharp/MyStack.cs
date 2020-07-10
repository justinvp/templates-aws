using Pulumi;
using Aws = Pulumi.Aws;

class MyStack : Stack
{
    public MyStack()
    {
        var config = new Config();
        var mountTargetId = config.Get("mountTargetId") ?? "";
        var byId = Output.Create(Aws.Efs.GetMountTarget.InvokeAsync(new Aws.Efs.GetMountTargetArgs
        {
            MountTargetId = mountTargetId,
        }));
    }

}

