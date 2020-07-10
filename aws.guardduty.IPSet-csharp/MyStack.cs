using Pulumi;
using Aws = Pulumi.Aws;

class MyStack : Stack
{
    public MyStack()
    {
        var master = new Aws.GuardDuty.Detector("master", new Aws.GuardDuty.DetectorArgs
        {
            Enable = true,
        });
        var bucket = new Aws.S3.Bucket("bucket", new Aws.S3.BucketArgs
        {
            Acl = "private",
        });
        var myIPSetBucketObject = new Aws.S3.BucketObject("myIPSetBucketObject", new Aws.S3.BucketObjectArgs
        {
            Acl = "public-read",
            Bucket = bucket.Id,
            Content = @"10.0.0.0/8

",
            Key = "MyIPSet",
        });
        var myIPSetIPSet = new Aws.GuardDuty.IPSet("myIPSetIPSet", new Aws.GuardDuty.IPSetArgs
        {
            Activate = true,
            DetectorId = master.Id,
            Format = "TXT",
            Location = Output.Tuple(myIPSetBucketObject.Bucket, myIPSetBucketObject.Key).Apply(values =>
            {
                var bucket = values.Item1;
                var key = values.Item2;
                return $"https://s3.amazonaws.com/{bucket}/{key}";
            }),
        });
    }

}

