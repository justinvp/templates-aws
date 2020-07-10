using Pulumi;
using Aws = Pulumi.Aws;

class MyStack : Stack
{
    public MyStack()
    {
        var strict = new Aws.Iam.AccountPasswordPolicy("strict", new Aws.Iam.AccountPasswordPolicyArgs
        {
            AllowUsersToChangePassword = true,
            MinimumPasswordLength = 8,
            RequireLowercaseCharacters = true,
            RequireNumbers = true,
            RequireSymbols = true,
            RequireUppercaseCharacters = true,
        });
    }

}

