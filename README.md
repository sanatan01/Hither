
# Hither

**Hither** is a command-line tool designed to simplify the management of containers and images, leveraging the Containerd runtime. With Hither, you can create, start, stop, and delete containers, as well as pull and delete images with ease.

## Table of Contents
- [Overview](#overview)
- [Installation](#installation)
- [Usage](#usage)
  - [Container Commands](#container-commands)
  - [Image Commands](#image-commands)
  - [Other Commands](#other-commands)
  - [Options](#options)
  - [Examples](#examples)
- [Building and Running Hither](#building-and-running-hither)
- [Troubleshooting](#troubleshooting)
- [Development](#development)
- [License](#license)
- [Acknowledgments](#acknowledgments)

## Overview

Hither offers a simple and intuitive interface for managing containers and images through the Containerd runtime. It provides commands for:
- Creating and managing containers.
- Pulling and deleting images from the local store or a registry.

## Installation

To install Hither, ensure that you have Go installed, then use the following command:

```bash
go get github.com/sanatan01/hither
```

## Usage

After installation, Hither offers various commands for managing containers and images.

### Container Commands

- `hither container run <image-name>`: Create and run a new container from the specified image.
- `hither container start <container-id>`: Start an existing, stopped container.
- `hither container stop <container-id>`: Stop a running container.
- `hither container rm <container-id>`: Delete a stopped container.

### Image Commands

- `hither pull <image-name>`: Pull an image from a registry.
- `hither image rm <image-name>`: Remove an image from the local store.
- `hither image list`: List all images on the system.

### Other Commands

- `hither container list`: List all containers on the system.

### Options

Available option for most commands:
- `-h`, `--help`: Display help information for a command.

### Examples

Here are some common examples of how to use Hither:

```bash
# Pull an image from a registry
hither pull my-image

# Create and run a new container from the pulled image
hither container run my-image

# Start a stopped container
hither container start my-container

# Stop a running container
hither container stop my-container

# Delete a stopped container
hither container rm my-container

# Delete an image from the local store
hither image rm my-image

# List all images on the system
hither image list

# List all containers on the system
hither container list
```

## Building and Running Hither

To build and run Hither, you can use the provided `Makefile`. Ensure you are in the project directory and use the following commands:

```bash
# Clone the repository
git clone git@github.com:sanatan01/Hither.git
cd Hither

# Build the project
make build

# Run tests
make test

# Run Hither
make run
```

## Troubleshooting

If you encounter issues with Hither, consider the following steps:

1. Check the Containerd logs for any errors.
2. Verify that the image or container you are trying to manage exists on your system.
3. Run the command with help option (`-h` or `--help`) to get detailed command information.

## Development

Hither is written in Go and utilizes the Cobra framework for building the command-line interface. If you would like to contribute or build the project from source, follow these steps:

```bash
# Clone the repository
git clone git@github.com:sanatan01/Hither.git
cd Hither

# Build the project
make build

# Run tests
make test

# Run Hither
make run
```

## Acknowledgments

Hither uses the following third-party libraries:
- **Containerd**: A container runtime for building and managing container images.
- **Cobra**: A framework for building command-line interfaces in Go.
- **Moby/Buildkit**: A toolkit for building container images.

We would like to extend our gratitude to the authors and maintainers of these libraries for their contributions to the open-source community.
