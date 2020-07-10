using Pulumi;
using Aws = Pulumi.Aws;

class MyStack : Stack
{
    public MyStack()
    {
        var user1 = new Aws.Iam.User("user1", new Aws.Iam.UserArgs
        {
        });
        var group1 = new Aws.Iam.Group("group1", new Aws.Iam.GroupArgs
        {
        });
        var group2 = new Aws.Iam.Group("group2", new Aws.Iam.GroupArgs
        {
        });
        var example1 = new Aws.Iam.UserGroupMembership("example1", new Aws.Iam.UserGroupMembershipArgs
        {
            Groups = 
            {
                group1.Name,
                group2.Name,
            },
            User = user1.Name,
        });
        var group3 = new Aws.Iam.Group("group3", new Aws.Iam.GroupArgs
        {
        });
        var example2 = new Aws.Iam.UserGroupMembership("example2", new Aws.Iam.UserGroupMembershipArgs
        {
            Groups = 
            {
                group3.Name,
            },
            User = user1.Name,
        });
    }

}

