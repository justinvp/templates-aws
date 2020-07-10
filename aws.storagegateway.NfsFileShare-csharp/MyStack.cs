using Pulumi;
using Aws = Pulumi.Aws;

class MyStack : Stack
{
    public MyStack()
    {
        var example = new Aws.StorageGateway.NfsFileShare("example", new Aws.StorageGateway.NfsFileShareArgs
        {
            ClientLists = 
            {
                "0.0.0.0/0",
            },
            GatewayArn = aws_storagegateway_gateway.Example.Arn,
            LocationArn = aws_s3_bucket.Example.Arn,
            RoleArn = aws_iam_role.Example.Arn,
        });
    }

}

