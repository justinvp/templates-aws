using Pulumi;
using Aws = Pulumi.Aws;

class MyStack : Stack
{
    public MyStack()
    {
        var examplebucket = new Aws.S3.Bucket("examplebucket", new Aws.S3.BucketArgs
        {
            Acl = "private",
            ObjectLockConfiguration = new Aws.S3.Inputs.BucketObjectLockConfigurationArgs
            {
                ObjectLockEnabled = "Enabled",
            },
            Versioning = new Aws.S3.Inputs.BucketVersioningArgs
            {
                Enabled = true,
            },
        });
        var examplebucketObject = new Aws.S3.BucketObject("examplebucketObject", new Aws.S3.BucketObjectArgs
        {
            Bucket = examplebucket.Id,
            ForceDestroy = true,
            Key = "someobject",
            ObjectLockLegalHoldStatus = "ON",
            ObjectLockMode = "GOVERNANCE",
            ObjectLockRetainUntilDate = "2021-12-31T23:59:60Z",
            Source = new FileAsset("important.txt"),
        });
    }

}

