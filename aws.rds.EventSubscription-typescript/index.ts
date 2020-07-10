import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const defaultInstance = new aws.rds.Instance("default", {
    allocatedStorage: 10,
    dbSubnetGroupName: "my_database_subnet_group",
    engine: "mysql",
    engineVersion: "5.6.17",
    instanceClass: "db.t2.micro",
    name: "mydb",
    parameterGroupName: "default.mysql5.6",
    password: "bar",
    username: "foo",
});
const defaultTopic = new aws.sns.Topic("default", {});
const defaultEventSubscription = new aws.rds.EventSubscription("default", {
    eventCategories: [
        "availability",
        "deletion",
        "failover",
        "failure",
        "low storage",
        "maintenance",
        "notification",
        "read replica",
        "recovery",
        "restoration",
    ],
    snsTopic: defaultTopic.arn,
    sourceIds: [defaultInstance.id],
    sourceType: "db-instance",
});

