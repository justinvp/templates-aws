using Pulumi;
using Aws = Pulumi.Aws;

class MyStack : Stack
{
    public MyStack()
    {
        var fooRecorderStatus = new Aws.Cfg.RecorderStatus("fooRecorderStatus", new Aws.Cfg.RecorderStatusArgs
        {
            IsEnabled = true,
        }, new CustomResourceOptions
        {
            DependsOn = 
            {
                "aws_config_delivery_channel.foo",
            },
        });
        var role = new Aws.Iam.Role("role", new Aws.Iam.RoleArgs
        {
            AssumeRolePolicy = @"{
  ""Version"": ""2012-10-17"",
  ""Statement"": [
    {
      ""Action"": ""sts:AssumeRole"",
      ""Principal"": {
        ""Service"": ""config.amazonaws.com""
      },
      ""Effect"": ""Allow"",
      ""Sid"": """"
    }
  ]
}

",
        });
        var rolePolicyAttachment = new Aws.Iam.RolePolicyAttachment("rolePolicyAttachment", new Aws.Iam.RolePolicyAttachmentArgs
        {
            PolicyArn = "arn:aws:iam::aws:policy/service-role/AWSConfigRole",
            Role = role.Name,
        });
        var bucket = new Aws.S3.Bucket("bucket", new Aws.S3.BucketArgs
        {
        });
        var fooDeliveryChannel = new Aws.Cfg.DeliveryChannel("fooDeliveryChannel", new Aws.Cfg.DeliveryChannelArgs
        {
            S3BucketName = bucket.BucketName,
        });
        var fooRecorder = new Aws.Cfg.Recorder("fooRecorder", new Aws.Cfg.RecorderArgs
        {
            RoleArn = role.Arn,
        });
        var rolePolicy = new Aws.Iam.RolePolicy("rolePolicy", new Aws.Iam.RolePolicyArgs
        {
            Policy = Output.Tuple(bucket.Arn, bucket.Arn).Apply(values =>
            {
                var bucketArn = values.Item1;
                var bucketArn1 = values.Item2;
                return @$"{{
  ""Version"": ""2012-10-17"",
  ""Statement"": [
    {{
      ""Action"": [
        ""s3:*""
      ],
      ""Effect"": ""Allow"",
      ""Resource"": [
        ""{bucketArn}"",
        ""{bucketArn1}/*""
      ]
    }}
  ]
}}

";
            }),
            Role = role.Id,
        });
    }

}

