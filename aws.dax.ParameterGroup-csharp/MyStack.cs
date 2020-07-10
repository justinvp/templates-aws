using Pulumi;
using Aws = Pulumi.Aws;

class MyStack : Stack
{
    public MyStack()
    {
        var example = new Aws.Dax.ParameterGroup("example", new Aws.Dax.ParameterGroupArgs
        {
            Parameters = 
            {
                new Aws.Dax.Inputs.ParameterGroupParameterArgs
                {
                    Name = "query-ttl-millis",
                    Value = "100000",
                },
                new Aws.Dax.Inputs.ParameterGroupParameterArgs
                {
                    Name = "record-ttl-millis",
                    Value = "100000",
                },
            },
        });
    }

}

