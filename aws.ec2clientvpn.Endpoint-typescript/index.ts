import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = new aws.ec2clientvpn.Endpoint("example", {
    authenticationOptions: [{
        rootCertificateChainArn: aws_acm_certificate_root_cert.arn,
        type: "certificate-authentication",
    }],
    clientCidrBlock: "10.0.0.0/16",
    connectionLogOptions: {
        cloudwatchLogGroup: aws_cloudwatch_log_group_lg.name,
        cloudwatchLogStream: aws_cloudwatch_log_stream_ls.name,
        enabled: true,
    },
    description: "clientvpn-example",
    serverCertificateArn: aws_acm_certificate_cert.arn,
});

