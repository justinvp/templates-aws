import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const myDetector = new aws.guardduty.Detector("MyDetector", {
    enable: true,
});

