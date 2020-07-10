using Pulumi;
using Aws = Pulumi.Aws;

class MyStack : Stack
{
    public MyStack()
    {
        var example = new Aws.Neptune.ClusterParameterGroup("example", new Aws.Neptune.ClusterParameterGroupArgs
        {
            Description = "neptune cluster parameter group",
            Family = "neptune1",
            Parameters = 
            {
                new Aws.Neptune.Inputs.ClusterParameterGroupParameterArgs
                {
                    Name = "neptune_enable_audit_log",
                    Value = "1",
                },
            },
        });
    }

}

