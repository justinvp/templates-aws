using Pulumi;
using Aws = Pulumi.Aws;

class MyStack : Stack
{
    public MyStack()
    {
        var role = new Aws.Iam.Role("role", new Aws.Iam.RoleArgs
        {
            AssumeRolePolicy = @"{
    ""Version"": ""2012-10-17"",
    ""Statement"": [
        {
            ""Action"": ""sts:AssumeRole"",
            ""Principal"": {
               ""Service"": ""ec2.amazonaws.com""
            },
            ""Effect"": ""Allow"",
            ""Sid"": """"
        }
    ]
}

",
            Path = "/",
        });
        var testProfile = new Aws.Iam.InstanceProfile("testProfile", new Aws.Iam.InstanceProfileArgs
        {
            Role = role.Name,
        });
    }

}

