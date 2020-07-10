using Pulumi;
using Aws = Pulumi.Aws;

class MyStack : Stack
{
    public MyStack()
    {
        var mykey = new Aws.Kms.Key("mykey", new Aws.Kms.KeyArgs
        {
            DeletionWindowInDays = 10,
            Description = "This key is used to encrypt bucket objects",
        });
        var mybucket = new Aws.S3.Bucket("mybucket", new Aws.S3.BucketArgs
        {
            ServerSideEncryptionConfiguration = new Aws.S3.Inputs.BucketServerSideEncryptionConfigurationArgs
            {
                Rule = new Aws.S3.Inputs.BucketServerSideEncryptionConfigurationRuleArgs
                {
                    ApplyServerSideEncryptionByDefault = new Aws.S3.Inputs.BucketServerSideEncryptionConfigurationRuleApplyServerSideEncryptionByDefaultArgs
                    {
                        KmsMasterKeyId = mykey.Arn,
                        SseAlgorithm = "aws:kms",
                    },
                },
            },
        });
    }

}

