using Pulumi;
using Aws = Pulumi.Aws;

class MyStack : Stack
{
    public MyStack()
    {
        var yada = new Aws.CloudWatch.LogGroup("yada", new Aws.CloudWatch.LogGroupArgs
        {
        });
        var foo = new Aws.CloudWatch.LogStream("foo", new Aws.CloudWatch.LogStreamArgs
        {
            LogGroupName = yada.Name,
        });
    }

}

