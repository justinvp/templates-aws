using Pulumi;
using Aws = Pulumi.Aws;

class MyStack : Stack
{
    public MyStack()
    {
        var @default = new Aws.Rds.ClusterParameterGroup("default", new Aws.Rds.ClusterParameterGroupArgs
        {
            Description = "RDS default cluster parameter group",
            Family = "aurora5.6",
            Parameters = 
            {
                new Aws.Rds.Inputs.ClusterParameterGroupParameterArgs
                {
                    Name = "character_set_server",
                    Value = "utf8",
                },
                new Aws.Rds.Inputs.ClusterParameterGroupParameterArgs
                {
                    Name = "character_set_client",
                    Value = "utf8",
                },
            },
        });
    }

}

