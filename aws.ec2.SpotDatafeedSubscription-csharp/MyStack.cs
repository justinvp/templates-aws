using Pulumi;
using Aws = Pulumi.Aws;

class MyStack : Stack
{
    public MyStack()
    {
        var defaultBucket = new Aws.S3.Bucket("defaultBucket", new Aws.S3.BucketArgs
        {
        });
        var defaultSpotDatafeedSubscription = new Aws.Ec2.SpotDatafeedSubscription("defaultSpotDatafeedSubscription", new Aws.Ec2.SpotDatafeedSubscriptionArgs
        {
            Bucket = defaultBucket.BucketName,
            Prefix = "my_subdirectory",
        });
    }

}

