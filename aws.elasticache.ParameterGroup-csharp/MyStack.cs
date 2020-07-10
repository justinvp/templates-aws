using Pulumi;
using Aws = Pulumi.Aws;

class MyStack : Stack
{
    public MyStack()
    {
        var @default = new Aws.ElastiCache.ParameterGroup("default", new Aws.ElastiCache.ParameterGroupArgs
        {
            Family = "redis2.8",
            Parameters = 
            {
                new Aws.ElastiCache.Inputs.ParameterGroupParameterArgs
                {
                    Name = "activerehashing",
                    Value = "yes",
                },
                new Aws.ElastiCache.Inputs.ParameterGroupParameterArgs
                {
                    Name = "min-slaves-to-write",
                    Value = "2",
                },
            },
        });
    }

}

