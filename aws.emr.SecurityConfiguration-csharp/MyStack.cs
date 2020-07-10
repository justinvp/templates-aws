using Pulumi;
using Aws = Pulumi.Aws;

class MyStack : Stack
{
    public MyStack()
    {
        var foo = new Aws.Emr.SecurityConfiguration("foo", new Aws.Emr.SecurityConfigurationArgs
        {
            Configuration = @"{
  ""EncryptionConfiguration"": {
    ""AtRestEncryptionConfiguration"": {
      ""S3EncryptionConfiguration"": {
        ""EncryptionMode"": ""SSE-S3""
      },
      ""LocalDiskEncryptionConfiguration"": {
        ""EncryptionKeyProviderType"": ""AwsKms"",
        ""AwsKmsKey"": ""arn:aws:kms:us-west-2:187416307283:alias/tf_emr_test_key""
      }
    },
    ""EnableInTransitEncryption"": false,
    ""EnableAtRestEncryption"": true
  }
}

",
        });
    }

}

