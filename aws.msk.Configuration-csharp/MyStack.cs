using Pulumi;
using Aws = Pulumi.Aws;

class MyStack : Stack
{
    public MyStack()
    {
        var example = new Aws.Msk.Configuration("example", new Aws.Msk.ConfigurationArgs
        {
            KafkaVersions = 
            {
                "2.1.0",
            },
            ServerProperties = @"auto.create.topics.enable = true
delete.topic.enable = true

",
        });
    }

}

