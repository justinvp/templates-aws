using Pulumi;
using Aws = Pulumi.Aws;

class MyStack : Stack
{
    public MyStack()
    {
        // Create a new endpoint
        var test = new Aws.Dms.Endpoint("test", new Aws.Dms.EndpointArgs
        {
            CertificateArn = "arn:aws:acm:us-east-1:123456789012:certificate/12345678-1234-1234-1234-123456789012",
            DatabaseName = "test",
            EndpointId = "test-dms-endpoint-tf",
            EndpointType = "source",
            EngineName = "aurora",
            ExtraConnectionAttributes = "",
            KmsKeyArn = "arn:aws:kms:us-east-1:123456789012:key/12345678-1234-1234-1234-123456789012",
            Password = "test",
            Port = 3306,
            ServerName = "test",
            SslMode = "none",
            Tags = 
            {
                { "Name", "test" },
            },
            Username = "test",
        });
    }

}

