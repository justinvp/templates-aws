import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = aws_secretsmanager_secret_example.id.apply(id => aws.secretsmanager.getSecretVersion({
    secretId: id,
}, { async: true }));

