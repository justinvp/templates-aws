import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const bar = new aws.elastictranscoder.Preset("bar", {
    audio: {
        audioPackingMode: "SingleTrack",
        bitRate: "96",
        channels: "2",
        codec: "AAC",
        sampleRate: "44100",
    },
    audioCodecOptions: {
        profile: "AAC-LC",
    },
    container: "mp4",
    description: "Sample Preset",
    thumbnails: {
        format: "png",
        interval: "120",
        maxHeight: "auto",
        maxWidth: "auto",
        paddingPolicy: "Pad",
        sizingPolicy: "Fit",
    },
    video: {
        bitRate: "1600",
        codec: "H.264",
        displayAspectRatio: "16:9",
        fixedGop: "false",
        frameRate: "auto",
        keyframesMaxDist: "240",
        maxFrameRate: "60",
        maxHeight: "auto",
        maxWidth: "auto",
        paddingPolicy: "Pad",
        sizingPolicy: "Fit",
    },
    videoCodecOptions: {
        ColorSpaceConversionMode: "None",
        InterlacedMode: "Progressive",
        Level: "2.2",
        MaxReferenceFrames: 3,
        Profile: "main",
    },
    videoWatermarks: [{
        horizontalAlign: "Right",
        horizontalOffset: "10px",
        id: "Test",
        maxHeight: "20%",
        maxWidth: "20%",
        opacity: "55.5",
        sizingPolicy: "ShrinkToFit",
        target: "Content",
        verticalAlign: "Bottom",
        verticalOffset: "10px",
    }],
});

