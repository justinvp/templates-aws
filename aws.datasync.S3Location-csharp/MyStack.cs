using Pulumi;
using Aws = Pulumi.Aws;

class MyStack : Stack
{
    public MyStack()
    {
        var example = new Aws.DataSync.S3Location("example", new Aws.DataSync.S3LocationArgs
        {
            S3BucketArn = aws_s3_bucket.Example.Arn,
            S3Config = new Aws.DataSync.Inputs.S3LocationS3ConfigArgs
            {
                BucketAccessRoleArn = aws_iam_role.Example.Arn,
            },
            Subdirectory = "/example/prefix",
        });
    }

}

