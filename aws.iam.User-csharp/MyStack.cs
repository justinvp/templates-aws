using Pulumi;
using Aws = Pulumi.Aws;

class MyStack : Stack
{
    public MyStack()
    {
        var lbUser = new Aws.Iam.User("lbUser", new Aws.Iam.UserArgs
        {
            Path = "/system/",
            Tags = 
            {
                { "tag-key", "tag-value" },
            },
        });
        var lbAccessKey = new Aws.Iam.AccessKey("lbAccessKey", new Aws.Iam.AccessKeyArgs
        {
            User = lbUser.Name,
        });
        var lbRo = new Aws.Iam.UserPolicy("lbRo", new Aws.Iam.UserPolicyArgs
        {
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
            User = lbUser.Name,
        });
    }

}

