using Pulumi;
using Aws = Pulumi.Aws;

class MyStack : Stack
{
    public MyStack()
    {
        var defaultBucket = new Aws.S3.Bucket("defaultBucket", new Aws.S3.BucketArgs
        {
        });
        var defaultBucketObject = new Aws.S3.BucketObject("defaultBucketObject", new Aws.S3.BucketObjectArgs
        {
            Bucket = defaultBucket.Id,
            Key = "beanstalk/go-v1.zip",
            Source = new FileAsset("go-v1.zip"),
        });
        var defaultApplication = new Aws.ElasticBeanstalk.Application("defaultApplication", new Aws.ElasticBeanstalk.ApplicationArgs
        {
            Description = "tf-test-desc",
        });
        var defaultApplicationVersion = new Aws.ElasticBeanstalk.ApplicationVersion("defaultApplicationVersion", new Aws.ElasticBeanstalk.ApplicationVersionArgs
        {
            Application = "tf-test-name",
            Bucket = defaultBucket.Id,
            Description = "application version",
            Key = defaultBucketObject.Id,
        });
    }

}

