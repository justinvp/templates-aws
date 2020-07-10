using Pulumi;
using Aws = Pulumi.Aws;

class MyStack : Stack
{
    public MyStack()
    {
        var db = new Aws.OpsWorks.MysqlLayer("db", new Aws.OpsWorks.MysqlLayerArgs
        {
            StackId = aws_opsworks_stack.Main.Id,
        });
    }

}

