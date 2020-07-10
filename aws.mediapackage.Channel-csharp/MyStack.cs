using Pulumi;
using Aws = Pulumi.Aws;

class MyStack : Stack
{
    public MyStack()
    {
        var kittens = new Aws.MediaPackage.Channel("kittens", new Aws.MediaPackage.ChannelArgs
        {
            ChannelId = "kitten-channel",
            Description = "A channel dedicated to amusing videos of kittens.",
        });
    }

}

