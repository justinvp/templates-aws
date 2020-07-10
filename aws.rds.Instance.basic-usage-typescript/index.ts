import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const defaultInstance = new aws.rds.Instance("default", {
    allocatedStorage: 20,
    engine: "mysql",
    engineVersion: "5.7",
    instanceClass: "db.t2.micro",
    name: "mydb",
    parameterGroupName: "default.mysql5.7",
    password: "foobarbaz",
    storageType: "gp2",
    username: "foo",
});

