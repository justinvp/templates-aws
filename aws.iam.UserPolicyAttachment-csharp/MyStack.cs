using Pulumi;
using Aws = Pulumi.Aws;

class MyStack : Stack
{
    public MyStack()
    {
        var user = new Aws.Iam.User("user", new Aws.Iam.UserArgs
        {
        });
        var policy = new Aws.Iam.Policy("policy", new Aws.Iam.PolicyArgs
        {
            Description = "A test policy",
            Policy = "",
        });
        // insert policy here
        var test_attach = new Aws.Iam.UserPolicyAttachment("test-attach", new Aws.Iam.UserPolicyAttachmentArgs
        {
            PolicyArn = policy.Arn,
            User = user.Name,
        });
    }

}

