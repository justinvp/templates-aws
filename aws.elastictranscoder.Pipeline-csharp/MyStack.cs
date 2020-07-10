using Pulumi;
using Aws = Pulumi.Aws;

class MyStack : Stack
{
    public MyStack()
    {
        var bar = new Aws.ElasticTranscoder.Pipeline("bar", new Aws.ElasticTranscoder.PipelineArgs
        {
            ContentConfig = new Aws.ElasticTranscoder.Inputs.PipelineContentConfigArgs
            {
                Bucket = aws_s3_bucket.Content_bucket.Bucket,
                StorageClass = "Standard",
            },
            InputBucket = aws_s3_bucket.Input_bucket.Bucket,
            Role = aws_iam_role.Test_role.Arn,
            ThumbnailConfig = new Aws.ElasticTranscoder.Inputs.PipelineThumbnailConfigArgs
            {
                Bucket = aws_s3_bucket.Thumb_bucket.Bucket,
                StorageClass = "Standard",
            },
        });
    }

}

