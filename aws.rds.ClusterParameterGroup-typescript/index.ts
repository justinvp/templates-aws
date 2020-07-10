import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const defaultClusterParameterGroup = new aws.rds.ClusterParameterGroup("default", {
    description: "RDS default cluster parameter group",
    family: "aurora5.6",
    parameters: [
        {
            name: "character_set_server",
            value: "utf8",
        },
        {
            name: "character_set_client",
            value: "utf8",
        },
    ],
});

