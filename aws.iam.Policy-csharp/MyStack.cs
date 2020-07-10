using Pulumi;
using Aws = Pulumi.Aws;

class MyStack : Stack
{
    public MyStack()
    {
        var policy = new Aws.Iam.Policy("policy", new Aws.Iam.PolicyArgs
        {
            Description = "My test policy",
            Path = "/",
            Policy = @"{
  ""Version"": ""2012-10-17"",
  ""Statement"": [
    {
      ""Action"": [
        ""ec2:Describe*""
      ],
      ""Effect"": ""Allow"",
      ""Resource"": ""*""
    }
  ]
}

",
        });
    }

}

