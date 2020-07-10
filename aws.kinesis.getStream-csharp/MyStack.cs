using Pulumi;
using Aws = Pulumi.Aws;

class MyStack : Stack
{
    public MyStack()
    {
        var stream = Output.Create(Aws.Kinesis.GetStream.InvokeAsync(new Aws.Kinesis.GetStreamArgs
        {
            Name = "stream-name",
        }));
    }

}

