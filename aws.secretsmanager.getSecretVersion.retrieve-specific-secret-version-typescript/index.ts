import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const by_version_stage = aws_secretsmanager_secret_example.id.apply(id => aws.secretsmanager.getSecretVersion({
    secretId: id,
    versionStage: "example",
}, { async: true }));

