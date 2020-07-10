using Pulumi;
using Aws = Pulumi.Aws;

class MyStack : Stack
{
    public MyStack()
    {
        var testQueue = new Aws.Batch.JobQueue("testQueue", new Aws.Batch.JobQueueArgs
        {
            ComputeEnvironments = 
            {
                aws_batch_compute_environment.Test_environment_1.Arn,
                aws_batch_compute_environment.Test_environment_2.Arn,
            },
            Priority = 1,
            State = "ENABLED",
        });
    }

}

