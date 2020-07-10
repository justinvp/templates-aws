using Pulumi;
using Aws = Pulumi.Aws;

class MyStack : Stack
{
    public MyStack()
    {
        var test = new Aws.MediaConvert.Queue("test", new Aws.MediaConvert.QueueArgs
        {
        });
    }

}

