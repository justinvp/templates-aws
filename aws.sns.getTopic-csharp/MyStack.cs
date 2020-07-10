using Pulumi;
using Aws = Pulumi.Aws;

class MyStack : Stack
{
    public MyStack()
    {
        var example = Output.Create(Aws.Sns.GetTopic.InvokeAsync(new Aws.Sns.GetTopicArgs
        {
            Name = "an_example_topic",
        }));
    }

}

