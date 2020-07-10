using Pulumi;
using Aws = Pulumi.Aws;

class MyStack : Stack
{
    public MyStack()
    {
        var example = new Aws.StorageGateway.SmbFileShare("example", new Aws.StorageGateway.SmbFileShareArgs
        {
            Authentication = "GuestAccess",
            GatewayArn = aws_storagegateway_gateway.Example.Arn,
            LocationArn = aws_s3_bucket.Example.Arn,
            RoleArn = aws_iam_role.Example.Arn,
        });
    }

}

