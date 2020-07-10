using Pulumi;
using Aws = Pulumi.Aws;

class MyStack : Stack
{
    public MyStack()
    {
        var web = new Aws.OpsWorks.StaticWebLayer("web", new Aws.OpsWorks.StaticWebLayerArgs
        {
            StackId = aws_opsworks_stack.Main.Id,
        });
    }

}

