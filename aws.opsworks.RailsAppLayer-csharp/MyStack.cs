using Pulumi;
using Aws = Pulumi.Aws;

class MyStack : Stack
{
    public MyStack()
    {
        var app = new Aws.OpsWorks.RailsAppLayer("app", new Aws.OpsWorks.RailsAppLayerArgs
        {
            StackId = aws_opsworks_stack.Main.Id,
        });
    }

}

