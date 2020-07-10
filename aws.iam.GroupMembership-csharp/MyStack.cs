using Pulumi;
using Aws = Pulumi.Aws;

class MyStack : Stack
{
    public MyStack()
    {
        var @group = new Aws.Iam.Group("group", new Aws.Iam.GroupArgs
        {
        });
        var userOne = new Aws.Iam.User("userOne", new Aws.Iam.UserArgs
        {
        });
        var userTwo = new Aws.Iam.User("userTwo", new Aws.Iam.UserArgs
        {
        });
        var team = new Aws.Iam.GroupMembership("team", new Aws.Iam.GroupMembershipArgs
        {
            Group = @group.Name,
            Users = 
            {
                userOne.Name,
                userTwo.Name,
            },
        });
    }

}

