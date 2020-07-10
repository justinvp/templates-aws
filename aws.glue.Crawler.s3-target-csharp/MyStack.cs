using Pulumi;
using Aws = Pulumi.Aws;

class MyStack : Stack
{
    public MyStack()
    {
        var example = new Aws.Glue.Crawler("example", new Aws.Glue.CrawlerArgs
        {
            DatabaseName = aws_glue_catalog_database.Example.Name,
            Role = aws_iam_role.Example.Arn,
            S3Targets = 
            {
                new Aws.Glue.Inputs.CrawlerS3TargetArgs
                {
                    Path = $"s3://{aws_s3_bucket.Example.Bucket}",
                },
            },
        });
    }

}

