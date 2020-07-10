using Pulumi;
using Aws = Pulumi.Aws;

class MyStack : Stack
{
    public MyStack()
    {
        var app = new Aws.OpsWorks.PhpAppLayer("app", new Aws.OpsWorks.PhpAppLayerArgs
        {
            StackId = aws_opsworks_stack.Main.Id,
        });
    }

}

