# DPS (Docker Presentation System)

DPS is a command-line utility written in Golang that provides a convenient way to display lists of Docker containers, images, networks, and volumes in a tabular format. With DPS, you can quickly inspect and manage your Docker resources using a user-friendly table view.

## Features

- List Docker containers with detailed information.
- Display Docker images with repository tags and creation timestamps.
- View Docker networks and their driver information.
- Inspect Docker volumes and their usage statistics.
- Easily customizable columns and sorting options.

# Installation

To install DPS, follow these steps:

1. Clone the DPS repository to your local machine:

```bash
git clone https://github.com/8thgencore/dps.git
```

2. Build the DPS binary:

```bash
cd dps
go build
```

3. Move the generated binary to a directory included in your system's PATH.

```bash
sudo mv dps /usr/local/bin/
```

4. Verify the installation by running:

```bash
dps --version
```

## Usage

DPS provides simple and intuitive commands to display Docker resources. Here are some common usage examples:

### List Docker Containers

```bash
dps container
```

### List Docker Images

```bash
dps image
```

### List Docker Networks

```bash
dps network
```

### List Docker Volumes

```bash
dps volume
```

### Customizing Output

You can customize the output by specifying columns and sorting options. For example, to list containers with only specific columns and sort them by name:

```bash
dps containers --columns "Container ID,Name,Status" --sort-by "Name"
```

###Help and Options
For more options and help, you can use the --help flag:

```bash
dps --help
```

### Contributing

We welcome contributions from the community. If you'd like to contribute to DPS, please fork the repository and create a pull request with your changes. Make sure to follow our contribution guidelines.

### License

DPS is open-source software licensed under the MIT License.

### Acknowledgments

DPS is built on top of the Docker SDK for Golang. We would like to thank the Docker development team for their fantastic work.
