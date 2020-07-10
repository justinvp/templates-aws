using Pulumi;
using Aws = Pulumi.Aws;

class MyStack : Stack
{
    public MyStack()
    {
        var test_queue = Output.Create(Aws.Batch.GetJobQueue.InvokeAsync(new Aws.Batch.GetJobQueueArgs
        {
            Name = "tf-test-batch-job-queue",
        }));
    }

}

