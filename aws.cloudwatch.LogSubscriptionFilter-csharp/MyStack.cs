using Pulumi;
using Aws = Pulumi.Aws;

class MyStack : Stack
{
    public MyStack()
    {
        var testLambdafunctionLogfilter = new Aws.CloudWatch.LogSubscriptionFilter("testLambdafunctionLogfilter", new Aws.CloudWatch.LogSubscriptionFilterArgs
        {
            DestinationArn = aws_kinesis_stream.Test_logstream.Arn,
            Distribution = "Random",
            FilterPattern = "logtype test",
            LogGroup = "/aws/lambda/example_lambda_name",
            RoleArn = aws_iam_role.Iam_for_lambda.Arn,
        });
    }

}

