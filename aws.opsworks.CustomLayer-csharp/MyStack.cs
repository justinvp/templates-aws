using Pulumi;
using Aws = Pulumi.Aws;

class MyStack : Stack
{
    public MyStack()
    {
        var custlayer = new Aws.OpsWorks.CustomLayer("custlayer", new Aws.OpsWorks.CustomLayerArgs
        {
            ShortName = "awesome",
            StackId = aws_opsworks_stack.Main.Id,
        });
    }

}

