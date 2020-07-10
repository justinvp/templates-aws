using Pulumi;
using Aws = Pulumi.Aws;

class MyStack : Stack
{
    public MyStack()
    {
        var @default = new Aws.Rds.ParameterGroup("default", new Aws.Rds.ParameterGroupArgs
        {
            Family = "mysql5.6",
            Parameters = 
            {
                new Aws.Rds.Inputs.ParameterGroupParameterArgs
                {
                    Name = "character_set_server",
                    Value = "utf8",
                },
                new Aws.Rds.Inputs.ParameterGroupParameterArgs
                {
                    Name = "character_set_client",
                    Value = "utf8",
                },
            },
        });
    }

}

