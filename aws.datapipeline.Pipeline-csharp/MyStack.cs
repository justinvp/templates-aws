using Pulumi;
using Aws = Pulumi.Aws;

class MyStack : Stack
{
    public MyStack()
    {
        var @default = new Aws.DataPipeline.Pipeline("default", new Aws.DataPipeline.PipelineArgs
        {
        });
    }

}

