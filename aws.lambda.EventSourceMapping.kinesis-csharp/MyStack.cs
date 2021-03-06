using Pulumi;
using Aws = Pulumi.Aws;

class MyStack : Stack
{
    public MyStack()
    {
        var example = new Aws.Lambda.EventSourceMapping("example", new Aws.Lambda.EventSourceMappingArgs
        {
            EventSourceArn = aws_kinesis_stream.Example.Arn,
            FunctionName = aws_lambda_function.Example.Arn,
            StartingPosition = "LATEST",
        });
    }

}

