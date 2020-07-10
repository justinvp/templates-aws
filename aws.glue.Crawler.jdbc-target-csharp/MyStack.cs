using Pulumi;
using Aws = Pulumi.Aws;

class MyStack : Stack
{
    public MyStack()
    {
        var example = new Aws.Glue.Crawler("example", new Aws.Glue.CrawlerArgs
        {
            DatabaseName = aws_glue_catalog_database.Example.Name,
            JdbcTargets = 
            {
                new Aws.Glue.Inputs.CrawlerJdbcTargetArgs
                {
                    ConnectionName = aws_glue_connection.Example.Name,
                    Path = "database-name/%",
                },
            },
            Role = aws_iam_role.Example.Arn,
        });
    }

}

