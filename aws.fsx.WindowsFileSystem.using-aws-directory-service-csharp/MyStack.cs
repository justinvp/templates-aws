using Pulumi;
using Aws = Pulumi.Aws;

class MyStack : Stack
{
    public MyStack()
    {
        var example = new Aws.Fsx.WindowsFileSystem("example", new Aws.Fsx.WindowsFileSystemArgs
        {
            ActiveDirectoryId = aws_directory_service_directory.Example.Id,
            KmsKeyId = aws_kms_key.Example.Arn,
            StorageCapacity = 300,
            SubnetIds = aws_subnet.Example.Id,
            ThroughputCapacity = 1024,
        });
    }

}

