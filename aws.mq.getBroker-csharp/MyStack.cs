using Pulumi;
using Aws = Pulumi.Aws;

class MyStack : Stack
{
    public MyStack()
    {
        var config = new Config();
        var brokerId = config.Get("brokerId") ?? "";
        var brokerName = config.Get("brokerName") ?? "";
        var byId = Output.Create(Aws.Mq.GetBroker.InvokeAsync(new Aws.Mq.GetBrokerArgs
        {
            BrokerId = brokerId,
        }));
        var byName = Output.Create(Aws.Mq.GetBroker.InvokeAsync(new Aws.Mq.GetBrokerArgs
        {
            BrokerName = brokerName,
        }));
    }

}

