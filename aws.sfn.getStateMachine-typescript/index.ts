import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = pulumi.output(aws.sfn.getStateMachine({
    name: "an_example_sfn_name",
}, { async: true }));

