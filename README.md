# How to Deploy a Go Web Server to Container Apps at Lightning Speed

[日本語のREADMEへ](README_ja.md)

This guide introduces how to deploy a Go web server to Container Apps at lightning speed.

## Azure Configuration

The resources to be created are as follows:

- User-assigned Managed Identity
- Azure Container Registry
- Azure Container App Environment
- Azure Container Apps

The user-assigned managed identity is necessary for Azure Container Apps to access the Azure Container Registry, and the ACR Pull role is assigned.

The Azure Container App Environment is the runtime environment for Container Apps, under which Container Apps are created.

## Go Web Server

A simple web server that returns `Hello, World!` at the `/` endpoint.

## Deployment Steps

First, appropriately modify the `SUFFIX` in the `Makefile` to ensure that resource names do not conflict globally. Then proceed with the following steps:

```sh
make rg    # Create resource group
make env   # Create Container Apps environment and ACR

export ACR_NAME=<acr-name>  # Set the name of the ACR

make acr-login              # Log in to the ACR
make build-image            # Build & push the image
make deploy                 # Deploy to Container Apps
```