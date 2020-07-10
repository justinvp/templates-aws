using Pulumi;
using Aws = Pulumi.Aws;

class MyStack : Stack
{
    public MyStack()
    {
        var config = new Config();
        var fileSystemId = config.Get("fileSystemId") ?? "";
        var byId = Output.Create(Aws.Efs.GetFileSystem.InvokeAsync(new Aws.Efs.GetFileSystemArgs
        {
            FileSystemId = fileSystemId,
        }));
    }

}

