using Pulumi;
using Aws = Pulumi.Aws;

class MyStack : Stack
{
    public MyStack()
    {
        var exampleBucket = new Aws.S3.Bucket("exampleBucket", new Aws.S3.BucketArgs
        {
        });
        var exampleFlowLog = new Aws.Ec2.FlowLog("exampleFlowLog", new Aws.Ec2.FlowLogArgs
        {
            LogDestination = exampleBucket.Arn,
            LogDestinationType = "s3",
            TrafficType = "ALL",
            VpcId = aws_vpc.Example.Id,
        });
    }

}

