using Pulumi;
using Aws = Pulumi.Aws;

class MyStack : Stack
{
    public MyStack()
    {
        var app = new Aws.OpsWorks.NodejsAppLayer("app", new Aws.OpsWorks.NodejsAppLayerArgs
        {
            StackId = aws_opsworks_stack.Main.Id,
        });
    }

}

