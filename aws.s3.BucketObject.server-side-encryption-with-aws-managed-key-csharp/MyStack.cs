using Pulumi;
using Aws = Pulumi.Aws;

class MyStack : Stack
{
    public MyStack()
    {
        var examplebucket = new Aws.S3.Bucket("examplebucket", new Aws.S3.BucketArgs
        {
            Acl = "private",
        });
        var examplebucketObject = new Aws.S3.BucketObject("examplebucketObject", new Aws.S3.BucketObjectArgs
        {
            Bucket = examplebucket.Id,
            Key = "someobject",
            ServerSideEncryption = "AES256",
            Source = new FileAsset("index.html"),
        });
    }

}

