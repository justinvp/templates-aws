using Pulumi;
using Aws = Pulumi.Aws;

class MyStack : Stack
{
    public MyStack()
    {
        var batch_mongo = Output.Create(Aws.Batch.GetComputeEnvironment.InvokeAsync(new Aws.Batch.GetComputeEnvironmentArgs
        {
            ComputeEnvironmentName = "batch-mongo-production",
        }));
    }

}

