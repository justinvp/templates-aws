using Pulumi;
using Aws = Pulumi.Aws;

class MyStack : Stack
{
    public MyStack()
    {
        var examplekms = new Aws.Kms.Key("examplekms", new Aws.Kms.KeyArgs
        {
            DeletionWindowInDays = 7,
            Description = "KMS key 1",
        });
        var examplebucket = new Aws.S3.Bucket("examplebucket", new Aws.S3.BucketArgs
        {
            Acl = "private",
        });
        var examplebucketObject = new Aws.S3.BucketObject("examplebucketObject", new Aws.S3.BucketObjectArgs
        {
            Bucket = examplebucket.Id,
            Key = "someobject",
            KmsKeyId = examplekms.Arn,
            Source = new FileAsset("index.html"),
        });
    }

}

