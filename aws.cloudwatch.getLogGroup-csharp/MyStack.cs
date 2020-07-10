using Pulumi;
using Aws = Pulumi.Aws;

class MyStack : Stack
{
    public MyStack()
    {
        var example = Output.Create(Aws.CloudWatch.GetLogGroup.InvokeAsync(new Aws.CloudWatch.GetLogGroupArgs
        {
            Name = "MyImportantLogs",
        }));
    }

}

