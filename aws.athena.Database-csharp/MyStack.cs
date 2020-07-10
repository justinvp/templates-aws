using Pulumi;
using Aws = Pulumi.Aws;

class MyStack : Stack
{
    public MyStack()
    {
        var hogeBucket = new Aws.S3.Bucket("hogeBucket", new Aws.S3.BucketArgs
        {
        });
        var hogeDatabase = new Aws.Athena.Database("hogeDatabase", new Aws.Athena.DatabaseArgs
        {
            Bucket = hogeBucket.BucketName,
            Name = "database_name",
        });
    }

}

