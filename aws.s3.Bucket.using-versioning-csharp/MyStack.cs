using Pulumi;
using Aws = Pulumi.Aws;

class MyStack : Stack
{
    public MyStack()
    {
        var bucket = new Aws.S3.Bucket("bucket", new Aws.S3.BucketArgs
        {
            Acl = "private",
            Versioning = new Aws.S3.Inputs.BucketVersioningArgs
            {
                Enabled = true,
            },
        });
    }

}

