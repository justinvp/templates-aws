import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

// Create a new Lightsail Key Pair
const lgKeyPair = new aws.lightsail.KeyPair("lg_key_pair", {});

