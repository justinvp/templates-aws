import pulumi
import pulumi_aws as aws

bar = aws.elastictranscoder.Preset("bar",
    audio={
        "audioPackingMode": "SingleTrack",
        "bitRate": 96,
        "channels": 2,
        "codec": "AAC",
        "sampleRate": 44100,
    },
    audio_codec_options={
        "profile": "AAC-LC",
    },
    container="mp4",
    description="Sample Preset",
    thumbnails={
        "format": "png",
        "interval": 120,
        "maxHeight": "auto",
        "maxWidth": "auto",
        "paddingPolicy": "Pad",
        "sizingPolicy": "Fit",
    },
    video={
        "bitRate": "1600",
        "codec": "H.264",
        "displayAspectRatio": "16:9",
        "fixedGop": "false",
        "frameRate": "auto",
        "keyframesMaxDist": 240,
        "maxFrameRate": "60",
        "maxHeight": "auto",
        "maxWidth": "auto",
        "paddingPolicy": "Pad",
        "sizingPolicy": "Fit",
    },
    video_codec_options={
        "ColorSpaceConversionMode": "None",
        "InterlacedMode": "Progressive",
        "Level": "2.2",
        "MaxReferenceFrames": 3,
        "Profile": "main",
    },
    video_watermarks=[{
        "horizontalAlign": "Right",
        "horizontalOffset": "10px",
        "id": "Test",
        "maxHeight": "20%",
        "maxWidth": "20%",
        "opacity": "55.5",
        "sizingPolicy": "ShrinkToFit",
        "target": "Content",
        "verticalAlign": "Bottom",
        "verticalOffset": "10px",
    }])

