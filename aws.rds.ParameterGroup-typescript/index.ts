import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const defaultParameterGroup = new aws.rds.ParameterGroup("default", {
    family: "mysql5.6",
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

