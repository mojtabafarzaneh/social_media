# Social Media Replica Project

This repository contains a social media replica project developed in Go. 

## Table of Contents

- [Installation](#installation)
- [Documentation](#documentation)
- [Contributing](#contributing)
- [License](#license)

## Installation

### Prerequisites

Before installing the project, ensure you have the following prerequisites installed:

- Go 1.18 or higher
- Git
- Docker

### Steps

1. **Clone the repository**:

    ```bash
    git clone https://github.com/mojtabafarzaneh/social_media.git
    cd social_media
    ```

2. **Install dependencies**:

    The Go module system will automatically install the necessary dependencies when you build the project. However, you can download the dependencies explicitly using:

    ```bash
    go mod tidy
    ```

3. **Build the project**:

    You can build the project by running:

    ```bash
    make build
    docker-compose up -d
    ```

4. **Run the project**:

    After building, you can run the project with:

    ```bash
    make serve
    ```

