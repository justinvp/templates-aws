using Pulumi;
using Aws = Pulumi.Aws;

class MyStack : Stack
{
    public MyStack()
    {
        var account = new Aws.Cfg.ConfigurationAggregator("account", new Aws.Cfg.ConfigurationAggregatorArgs
        {
            AccountAggregationSource = new Aws.Cfg.Inputs.ConfigurationAggregatorAccountAggregationSourceArgs
            {
                AccountIds = 
                {
                    "123456789012",
                },
                Regions = 
                {
                    "us-west-2",
                },
            },
        });
    }

}

