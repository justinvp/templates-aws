using Pulumi;
using Aws = Pulumi.Aws;

class MyStack : Stack
{
    public MyStack()
    {
        var bar = new Aws.RedShift.ParameterGroup("bar", new Aws.RedShift.ParameterGroupArgs
        {
            Family = "redshift-1.0",
            Parameters = 
            {
                new Aws.RedShift.Inputs.ParameterGroupParameterArgs
                {
                    Name = "require_ssl",
                    Value = "true",
                },
                new Aws.RedShift.Inputs.ParameterGroupParameterArgs
                {
                    Name = "query_group",
                    Value = "example",
                },
                new Aws.RedShift.Inputs.ParameterGroupParameterArgs
                {
                    Name = "enable_user_activity_logging",
                    Value = "true",
                },
            },
        });
    }

}

