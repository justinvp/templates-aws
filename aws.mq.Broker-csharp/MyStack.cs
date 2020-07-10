using Pulumi;
using Aws = Pulumi.Aws;

class MyStack : Stack
{
    public MyStack()
    {
        var example = new Aws.Mq.Broker("example", new Aws.Mq.BrokerArgs
        {
            BrokerName = "example",
            Configuration = new Aws.Mq.Inputs.BrokerConfigurationArgs
            {
                Id = aws_mq_configuration.Test.Id,
                Revision = aws_mq_configuration.Test.Latest_revision,
            },
            EngineType = "ActiveMQ",
            EngineVersion = "5.15.0",
            HostInstanceType = "mq.t2.micro",
            SecurityGroups = 
            {
                aws_security_group.Test.Id,
            },
            Users = 
            {
                new Aws.Mq.Inputs.BrokerUserArgs
                {
                    Password = "MindTheGap",
                    Username = "ExampleUser",
                },
            },
        });
    }

}

