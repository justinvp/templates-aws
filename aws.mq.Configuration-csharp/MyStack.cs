using Pulumi;
using Aws = Pulumi.Aws;

class MyStack : Stack
{
    public MyStack()
    {
        var example = new Aws.Mq.Configuration("example", new Aws.Mq.ConfigurationArgs
        {
            Data = @"<?xml version=""1.0"" encoding=""UTF-8"" standalone=""yes""?>
<broker xmlns=""http://activemq.apache.org/schema/core"">
  <plugins>
    <forcePersistencyModeBrokerPlugin persistenceFlag=""true""/>
    <statisticsBrokerPlugin/>
    <timeStampingBrokerPlugin ttlCeiling=""86400000"" zeroExpirationOverride=""86400000""/>
  </plugins>
</broker>

",
            Description = "Example Configuration",
            EngineType = "ActiveMQ",
            EngineVersion = "5.15.0",
        });
    }

}

