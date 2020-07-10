using Pulumi;
using Aws = Pulumi.Aws;

class MyStack : Stack
{
    public MyStack()
    {
        var app = new Aws.OpsWorks.JavaAppLayer("app", new Aws.OpsWorks.JavaAppLayerArgs
        {
            StackId = aws_opsworks_stack.Main.Id,
        });
    }

}

