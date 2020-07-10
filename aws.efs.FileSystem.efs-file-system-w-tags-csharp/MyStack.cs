using Pulumi;
using Aws = Pulumi.Aws;

class MyStack : Stack
{
    public MyStack()
    {
        var foo = new Aws.Efs.FileSystem("foo", new Aws.Efs.FileSystemArgs
        {
            Tags = 
            {
                { "Name", "MyProduct" },
            },
        });
    }

}

