using Pulumi;
using Aws = Pulumi.Aws;

class MyStack : Stack
{
    public MyStack()
    {
        var example = new Aws.Glue.Crawler("example", new Aws.Glue.CrawlerArgs
        {
            DatabaseName = aws_glue_catalog_database.Example.Name,
            DynamodbTargets = 
            {
                new Aws.Glue.Inputs.CrawlerDynamodbTargetArgs
                {
                    Path = "table-name",
                },
            },
            Role = aws_iam_role.Example.Arn,
        });
    }

}

