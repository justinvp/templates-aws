using Pulumi;
using Aws = Pulumi.Aws;

class MyStack : Stack
{
    public MyStack()
    {
        var example = new Aws.Glue.Crawler("example", new Aws.Glue.CrawlerArgs
        {
            CatalogTargets = 
            {
                new Aws.Glue.Inputs.CrawlerCatalogTargetArgs
                {
                    DatabaseName = aws_glue_catalog_database.Example.Name,
                    Tables = 
                    {
                        aws_glue_catalog_table.Example.Name,
                    },
                },
            },
            Configuration = @"{
  ""Version"":1.0,
  ""Grouping"": {
    ""TableGroupingPolicy"": ""CombineCompatibleSchemas""
  }
}

",
            DatabaseName = aws_glue_catalog_database.Example.Name,
            Role = aws_iam_role.Example.Arn,
            SchemaChangePolicy = new Aws.Glue.Inputs.CrawlerSchemaChangePolicyArgs
            {
                DeleteBehavior = "LOG",
            },
        });
    }

}

