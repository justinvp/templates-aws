using Pulumi;
using Aws = Pulumi.Aws;

class MyStack : Stack
{
    public MyStack()
    {
        var exampleUser = new Aws.Iam.User("exampleUser", new Aws.Iam.UserArgs
        {
            ForceDestroy = true,
            Path = "/",
        });
        var exampleUserLoginProfile = new Aws.Iam.UserLoginProfile("exampleUserLoginProfile", new Aws.Iam.UserLoginProfileArgs
        {
            PgpKey = "keybase:some_person_that_exists",
            User = exampleUser.Name,
        });
        this.Password = exampleUserLoginProfile.EncryptedPassword;
    }

    [Output("password")]
    public Output<string> Password { get; set; }
}

