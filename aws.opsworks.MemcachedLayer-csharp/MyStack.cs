using Pulumi;
using Aws = Pulumi.Aws;

class MyStack : Stack
{
    public MyStack()
    {
        var cache = new Aws.OpsWorks.MemcachedLayer("cache", new Aws.OpsWorks.MemcachedLayerArgs
        {
            StackId = aws_opsworks_stack.Main.Id,
        });
    }

}

