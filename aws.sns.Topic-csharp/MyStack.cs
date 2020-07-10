using Pulumi;
using Aws = Pulumi.Aws;

class MyStack : Stack
{
    public MyStack()
    {
        var userUpdates = new Aws.Sns.Topic("userUpdates", new Aws.Sns.TopicArgs
        {
        });
    }

}

