using Pulumi;
using Aws = Pulumi.Aws;

class MyStack : Stack
{
    public MyStack()
    {
        var example = Output.Create(Aws.Sqs.GetQueue.InvokeAsync(new Aws.Sqs.GetQueueArgs
        {
            Name = "queue",
        }));
    }

}

