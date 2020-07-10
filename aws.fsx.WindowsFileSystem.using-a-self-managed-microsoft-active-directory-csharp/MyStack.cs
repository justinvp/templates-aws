using Pulumi;
using Aws = Pulumi.Aws;

class MyStack : Stack
{
    public MyStack()
    {
        var example = new Aws.Fsx.WindowsFileSystem("example", new Aws.Fsx.WindowsFileSystemArgs
        {
            KmsKeyId = aws_kms_key.Example.Arn,
            SelfManagedActiveDirectory = new Aws.Fsx.Inputs.WindowsFileSystemSelfManagedActiveDirectoryArgs
            {
                DnsIps = 
                {
                    "10.0.0.111",
                    "10.0.0.222",
                },
                DomainName = "corp.example.com",
                Password = "avoid-plaintext-passwords",
                Username = "Admin",
            },
            StorageCapacity = 300,
            SubnetIds = aws_subnet.Example.Id,
            ThroughputCapacity = 1024,
        });
    }

}

