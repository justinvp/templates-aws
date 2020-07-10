using Pulumi;
using Aws = Pulumi.Aws;

class MyStack : Stack
{
    public MyStack()
    {
        var exampleCurReportDefinition = new Aws.Cur.ReportDefinition("exampleCurReportDefinition", new Aws.Cur.ReportDefinitionArgs
        {
            AdditionalArtifacts = 
            {
                "REDSHIFT",
                "QUICKSIGHT",
            },
            AdditionalSchemaElements = 
            {
                "RESOURCES",
            },
            Compression = "GZIP",
            Format = "textORcsv",
            ReportName = "example-cur-report-definition",
            S3Bucket = "example-bucket-name",
            S3Region = "us-east-1",
            TimeUnit = "HOURLY",
        });
    }

}

