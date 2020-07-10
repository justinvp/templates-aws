using Pulumi;
using Aws = Pulumi.Aws;

class MyStack : Stack
{
    public MyStack()
    {
        var example = new Aws.MediaStore.Container("example", new Aws.MediaStore.ContainerArgs
        {
        });
    }

}

