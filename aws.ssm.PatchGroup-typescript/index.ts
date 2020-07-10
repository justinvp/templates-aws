import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const production = new aws.ssm.PatchBaseline("production", {
    approvedPatches: ["KB123456"],
});
const patchgroup = new aws.ssm.PatchGroup("patchgroup", {
    baselineId: production.id,
    patchGroup: "patch-group-name",
});

