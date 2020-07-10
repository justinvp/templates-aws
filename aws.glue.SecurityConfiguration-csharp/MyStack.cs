using Pulumi;
using Aws = Pulumi.Aws;

class MyStack : Stack
{
    public MyStack()
    {
        var example = new Aws.Glue.SecurityConfiguration("example", new Aws.Glue.SecurityConfigurationArgs
        {
            EncryptionConfiguration = new Aws.Glue.Inputs.SecurityConfigurationEncryptionConfigurationArgs
            {
                CloudwatchEncryption = new Aws.Glue.Inputs.SecurityConfigurationEncryptionConfigurationCloudwatchEncryptionArgs
                {
                    CloudwatchEncryptionMode = "DISABLED",
                },
                JobBookmarksEncryption = new Aws.Glue.Inputs.SecurityConfigurationEncryptionConfigurationJobBookmarksEncryptionArgs
                {
                    JobBookmarksEncryptionMode = "DISABLED",
                },
                S3Encryption = new Aws.Glue.Inputs.SecurityConfigurationEncryptionConfigurationS3EncryptionArgs
                {
                    KmsKeyArn = data.Aws_kms_key.Example.Arn,
                    S3EncryptionMode = "SSE-KMS",
                },
            },
        });
    }

}

