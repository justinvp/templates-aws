using Pulumi;
using Aws = Pulumi.Aws;

class MyStack : Stack
{
    public MyStack()
    {
        var main = Output.Create(Aws.GetBillingServiceAccount.InvokeAsync());
        var billingLogs = new Aws.S3.Bucket("billingLogs", new Aws.S3.BucketArgs
        {
            Acl = "private",
            Policy = Output.Tuple(main, main).Apply(values =>
            {
                var main = values.Item1;
                var main1 = values.Item2;
                return @$"{{
  ""Id"": ""Policy"",
  ""Version"": ""2012-10-17"",
  ""Statement"": [
    {{
      ""Action"": [
        ""s3:GetBucketAcl"", ""s3:GetBucketPolicy""
      ],
      ""Effect"": ""Allow"",
      ""Resource"": ""arn:aws:s3:::my-billing-tf-test-bucket"",
      ""Principal"": {{
        ""AWS"": [
          ""{main.Arn}""
        ]
      }}
    }},
    {{
      ""Action"": [
        ""s3:PutObject""
      ],
      ""Effect"": ""Allow"",
      ""Resource"": ""arn:aws:s3:::my-billing-tf-test-bucket/*"",
      ""Principal"": {{
        ""AWS"": [
          ""{main1.Arn}""
        ]
      }}
    }}
  ]
}}

";
            }),
        });
    }

}

