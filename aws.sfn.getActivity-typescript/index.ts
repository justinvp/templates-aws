import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const sfnActivity = pulumi.output(aws.sfn.getActivity({
    name: "my-activity",
}, { async: true }));

