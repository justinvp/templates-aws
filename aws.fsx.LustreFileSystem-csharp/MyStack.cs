using Pulumi;
using Aws = Pulumi.Aws;

class MyStack : Stack
{
    public MyStack()
    {
        var example = new Aws.Fsx.LustreFileSystem("example", new Aws.Fsx.LustreFileSystemArgs
        {
            ImportPath = $"s3://{aws_s3_bucket.Example.Bucket}",
            StorageCapacity = 1200,
            SubnetIds = aws_subnet.Example.Id,
        });
    }

}

