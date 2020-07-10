using Pulumi;
using Aws = Pulumi.Aws;

class MyStack : Stack
{
    public MyStack()
    {
        var foo = new Aws.Swf.Domain("foo", new Aws.Swf.DomainArgs
        {
            Description = "SWF Domain",
            WorkflowExecutionRetentionPeriodInDays = "30",
        });
    }

}

