using Pulumi;
using Aws = Pulumi.Aws;

class MyStack : Stack
{
    public MyStack()
    {
        var @default = new Aws.Rds.Instance("default", new Aws.Rds.InstanceArgs
        {
            AllocatedStorage = 20,
            Engine = "mysql",
            EngineVersion = "5.7",
            InstanceClass = "db.t2.micro",
            Name = "mydb",
            ParameterGroupName = "default.mysql5.7",
            Password = "foobarbaz",
            StorageType = "gp2",
            Username = "foo",
        });
    }

}

