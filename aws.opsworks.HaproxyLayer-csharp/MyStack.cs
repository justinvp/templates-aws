using Pulumi;
using Aws = Pulumi.Aws;

class MyStack : Stack
{
    public MyStack()
    {
        var lb = new Aws.OpsWorks.HaproxyLayer("lb", new Aws.OpsWorks.HaproxyLayerArgs
        {
            StackId = aws_opsworks_stack.Main.Id,
            StatsPassword = "foobarbaz",
        });
    }

}

