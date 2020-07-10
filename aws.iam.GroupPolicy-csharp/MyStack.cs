using Pulumi;
using Aws = Pulumi.Aws;

class MyStack : Stack
{
    public MyStack()
    {
        var myDevelopers = new Aws.Iam.Group("myDevelopers", new Aws.Iam.GroupArgs
        {
            Path = "/users/",
        });
        var myDeveloperPolicy = new Aws.Iam.GroupPolicy("myDeveloperPolicy", new Aws.Iam.GroupPolicyArgs
        {
            Group = myDevelopers.Id,
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

