using Pulumi;
using Aws = Pulumi.Aws;

class MyStack : Stack
{
    public MyStack()
    {
        var main = Output.Create(Aws.CloudTrail.GetServiceAccount.InvokeAsync());
        var bucket = new Aws.S3.Bucket("bucket", new Aws.S3.BucketArgs
        {
            ForceDestroy = true,
            Policy = Output.Tuple(main, main).Apply(values =>
            {
                var main = values.Item1;
                var main1 = values.Item2;
                return @$"{{
  ""Version"": ""2008-10-17"",
  ""Statement"": [
    {{
      ""Sid"": ""Put bucket policy needed for trails"",
      ""Effect"": ""Allow"",
      ""Principal"": {{
        ""AWS"": ""{main.Arn}""
      }},
      ""Action"": ""s3:PutObject"",
      ""Resource"": ""arn:aws:s3:::tf-cloudtrail-logging-test-bucket/*""
    }},
    {{
      ""Sid"": ""Get bucket policy needed for trails"",
      ""Effect"": ""Allow"",
      ""Principal"": {{
        ""AWS"": ""{main1.Arn}""
      }},
      ""Action"": ""s3:GetBucketAcl"",
      ""Resource"": ""arn:aws:s3:::tf-cloudtrail-logging-test-bucket""
    }}
  ]
}}

";
            }),
        });
    }

}

