using Pulumi;
using Aws = Pulumi.Aws;

class MyStack : Stack
{
    public MyStack()
    {
        var monitor = new Aws.OpsWorks.GangliaLayer("monitor", new Aws.OpsWorks.GangliaLayerArgs
        {
            Password = "foobarbaz",
            StackId = aws_opsworks_stack.Main.Id,
        });
    }

}

