using Pulumi;
using Aws = Pulumi.Aws;

class MyStack : Stack
{
    public MyStack()
    {
        var main = new Aws.OpsWorks.Stack("main", new Aws.OpsWorks.StackArgs
        {
            CustomJson = @"{
 ""foobar"": {
    ""version"": ""1.0.0""
  }
}

",
            DefaultInstanceProfileArn = aws_iam_instance_profile.Opsworks.Arn,
            Region = "us-west-1",
            ServiceRoleArn = aws_iam_role.Opsworks.Arn,
            Tags = 
            {
                { "Name", "foobar-stack" },
            },
        });
    }

}

