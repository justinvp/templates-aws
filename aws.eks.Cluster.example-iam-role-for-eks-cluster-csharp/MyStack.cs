using Pulumi;
using Aws = Pulumi.Aws;

class MyStack : Stack
{
    public MyStack()
    {
        var example = new Aws.Iam.Role("example", new Aws.Iam.RoleArgs
        {
            AssumeRolePolicy = @"{
  ""Version"": ""2012-10-17"",
  ""Statement"": [
    {
      ""Effect"": ""Allow"",
      ""Principal"": {
        ""Service"": ""eks.amazonaws.com""
      },
      ""Action"": ""sts:AssumeRole""
    }
  ]
}

",
        });
        var example_AmazonEKSClusterPolicy = new Aws.Iam.RolePolicyAttachment("example-AmazonEKSClusterPolicy", new Aws.Iam.RolePolicyAttachmentArgs
        {
            PolicyArn = "arn:aws:iam::aws:policy/AmazonEKSClusterPolicy",
            Role = example.Name,
        });
        var example_AmazonEKSServicePolicy = new Aws.Iam.RolePolicyAttachment("example-AmazonEKSServicePolicy", new Aws.Iam.RolePolicyAttachmentArgs
        {
            PolicyArn = "arn:aws:iam::aws:policy/AmazonEKSServicePolicy",
            Role = example.Name,
        });
    }

}

