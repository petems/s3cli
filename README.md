# s3cli

A simple s3 client for moving files to and from a fixed s3 bucket.

## But why?

I needed a simple way to periodically copy backup files from a pod in a
kubernetes cluster to s3 without having to run a sidecar container.

## Adding

### mount this repo as a Filesystem in the pod

### add execution values to the environment

### execute to upload

## Building

Some versions were built using the [build.sh] script and are ready to go
so that this repo can be mounted directly as described above.
