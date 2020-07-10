using Pulumi;
using Aws = Pulumi.Aws;

class MyStack : Stack
{
    public MyStack()
    {
        var example = new Aws.Neptune.ParameterGroup("example", new Aws.Neptune.ParameterGroupArgs
        {
            Family = "neptune1",
            Parameters = 
            {
                new Aws.Neptune.Inputs.ParameterGroupParameterArgs
                {
                    Name = "neptune_query_timeout",
                    Value = "25",
                },
            },
        });
    }

}

