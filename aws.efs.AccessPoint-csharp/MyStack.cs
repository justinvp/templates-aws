using Pulumi;
using Aws = Pulumi.Aws;

class MyStack : Stack
{
    public MyStack()
    {
        var test = new Aws.Efs.AccessPoint("test", new Aws.Efs.AccessPointArgs
        {
            FileSystemId = aws_efs_file_system.Foo.Id,
        });
    }

}

