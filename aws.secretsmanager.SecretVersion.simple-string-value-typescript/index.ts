import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = new aws.secretsmanager.SecretVersion("example", {
    secretId: aws_secretsmanager_secret_example.id,
    secretString: "example-string-to-protect",
});

