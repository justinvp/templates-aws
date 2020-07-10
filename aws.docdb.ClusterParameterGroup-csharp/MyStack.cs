using Pulumi;
using Aws = Pulumi.Aws;

class MyStack : Stack
{
    public MyStack()
    {
        var example = new Aws.DocDB.ClusterParameterGroup("example", new Aws.DocDB.ClusterParameterGroupArgs
        {
            Description = "docdb cluster parameter group",
            Family = "docdb3.6",
            Parameters = 
            {
                new Aws.DocDB.Inputs.ClusterParameterGroupParameterArgs
                {
                    Name = "tls",
                    Value = "enabled",
                },
            },
        });
    }

}

