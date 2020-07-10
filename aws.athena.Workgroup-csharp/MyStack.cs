using Pulumi;
using Aws = Pulumi.Aws;

class MyStack : Stack
{
    public MyStack()
    {
        var example = new Aws.Athena.Workgroup("example", new Aws.Athena.WorkgroupArgs
        {
            Configuration = new Aws.Athena.Inputs.WorkgroupConfigurationArgs
            {
                EnforceWorkgroupConfiguration = true,
                PublishCloudwatchMetricsEnabled = true,
                ResultConfiguration = new Aws.Athena.Inputs.WorkgroupConfigurationResultConfigurationArgs
                {
                    EncryptionConfiguration = new Aws.Athena.Inputs.WorkgroupConfigurationResultConfigurationEncryptionConfigurationArgs
                    {
                        EncryptionOption = "SSE_KMS",
                        KmsKeyArn = aws_kms_key.Example.Arn,
                    },
                    OutputLocation = "s3://{aws_s3_bucket.example.bucket}/output/",
                },
            },
        });
    }

}

