import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const config = new pulumi.Config();
const mountTargetId = config.get("mountTargetId") || "";

const byId = pulumi.output(aws.efs.getMountTarget({
    mountTargetId: mountTargetId,
}, { async: true }));

