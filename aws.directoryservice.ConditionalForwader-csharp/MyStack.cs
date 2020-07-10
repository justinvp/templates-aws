using Pulumi;
using Aws = Pulumi.Aws;

class MyStack : Stack
{
    public MyStack()
    {
        var example = new Aws.DirectoryService.ConditionalForwader("example", new Aws.DirectoryService.ConditionalForwaderArgs
        {
            DirectoryId = aws_directory_service_directory.Ad.Id,
            DnsIps = 
            {
                "8.8.8.8",
                "8.8.4.4",
            },
            RemoteDomainName = "example.com",
        });
    }

}

