using Pulumi;
using Aws = Pulumi.Aws;

class MyStack : Stack
{
    public MyStack()
    {
        var central = new Aws.Provider("central", new Aws.ProviderArgs
        {
            Region = "eu-central-1",
        });
        var replicationRole = new Aws.Iam.Role("replicationRole", new Aws.Iam.RoleArgs
        {
            AssumeRolePolicy = @"{
  ""Version"": ""2012-10-17"",
  ""Statement"": [
    {
      ""Action"": ""sts:AssumeRole"",
      ""Principal"": {
        ""Service"": ""s3.amazonaws.com""
      },
      ""Effect"": ""Allow"",
      ""Sid"": """"
    }
  ]
}

",
        });
        var destination = new Aws.S3.Bucket("destination", new Aws.S3.BucketArgs
        {
            Region = "eu-west-1",
            Versioning = new Aws.S3.Inputs.BucketVersioningArgs
            {
                Enabled = true,
            },
        });
        var bucket = new Aws.S3.Bucket("bucket", new Aws.S3.BucketArgs
        {
            Acl = "private",
            Region = "eu-central-1",
            ReplicationConfiguration = new Aws.S3.Inputs.BucketReplicationConfigurationArgs
            {
                Role = replicationRole.Arn,
                Rules = 
                {
                    new Aws.S3.Inputs.BucketReplicationConfigurationRuleArgs
                    {
                        Destination = new Aws.S3.Inputs.BucketReplicationConfigurationRuleDestinationArgs
                        {
                            Bucket = destination.Arn,
                            StorageClass = "STANDARD",
                        },
                        Id = "foobar",
                        Prefix = "foo",
                        Status = "Enabled",
                    },
                },
            },
            Versioning = new Aws.S3.Inputs.BucketVersioningArgs
            {
                Enabled = true,
            },
        }, new CustomResourceOptions
        {
            Provider = "aws.central",
        });
        var replicationPolicy = new Aws.Iam.Policy("replicationPolicy", new Aws.Iam.PolicyArgs
        {
            Policy = Output.Tuple(bucket.Arn, bucket.Arn, destination.Arn).Apply(values =>
            {
                var bucketArn = values.Item1;
                var bucketArn1 = values.Item2;
                var destinationArn = values.Item3;
                return @$"{{
  ""Version"": ""2012-10-17"",
  ""Statement"": [
    {{
      ""Action"": [
        ""s3:GetReplicationConfiguration"",
        ""s3:ListBucket""
      ],
      ""Effect"": ""Allow"",
      ""Resource"": [
        ""{bucketArn}""
      ]
    }},
    {{
      ""Action"": [
        ""s3:GetObjectVersion"",
        ""s3:GetObjectVersionAcl""
      ],
      ""Effect"": ""Allow"",
      ""Resource"": [
        ""{bucketArn1}/*""
      ]
    }},
    {{
      ""Action"": [
        ""s3:ReplicateObject"",
        ""s3:ReplicateDelete""
      ],
      ""Effect"": ""Allow"",
      ""Resource"": ""{destinationArn}/*""
    }}
  ]
}}

";
            }),
        });
        var replicationRolePolicyAttachment = new Aws.Iam.RolePolicyAttachment("replicationRolePolicyAttachment", new Aws.Iam.RolePolicyAttachmentArgs
        {
            PolicyArn = replicationPolicy.Arn,
            Role = replicationRole.Name,
        });
    }

}

