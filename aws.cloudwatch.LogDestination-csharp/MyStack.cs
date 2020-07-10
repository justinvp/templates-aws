using Pulumi;
using Aws = Pulumi.Aws;

class MyStack : Stack
{
    public MyStack()
    {
        var testDestination = new Aws.CloudWatch.LogDestination("testDestination", new Aws.CloudWatch.LogDestinationArgs
        {
            RoleArn = aws_iam_role.Iam_for_cloudwatch.Arn,
            TargetArn = aws_kinesis_stream.Kinesis_for_cloudwatch.Arn,
        });
    }

}

