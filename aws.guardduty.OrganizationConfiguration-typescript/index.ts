import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const exampleDetector = new aws.guardduty.Detector("exampleDetector", {enable: true});
const exampleOrganizationConfiguration = new aws.guardduty.OrganizationConfiguration("exampleOrganizationConfiguration", {
    autoEnable: true,
    detectorId: exampleDetector.id,
});

