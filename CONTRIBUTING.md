# Contributing to GREYE Helm Chart

## Introduction

Thank you for considering contributing to the GREYE Helm Chart! This document provides guidelines and instructions to help you get started.

## Getting Started

### Reporting Bugs

If you find a bug, please report it by opening an issue on our [GitHub repository](https://github.com/greye-monitoring/helm-charts/issues). Include as much detail as possible to help us understand and reproduce the issue.

### Feature Requests

We welcome new feature requests! If you have an idea for a new feature, please open an issue on our [GitHub repository](https://github.com/greye-monitoring/helm-charts/issues) and describe the feature in detail.

### Prerequisites

- Kubernetes 1.16+
- Helm 3.0+
- Go development environment (for testing)

### Setup

1. Fork the repository on GitHub
2. Clone your fork locally


## Development Workflow

**Important: Every development must be associated with an issue. Create an issue before starting any work.**

1. Make your changes

2Test your changes locally:
   ```bash
   helm lint greye/
   helm template greye/
   ```

3Commit your changes with meaningful commit messages, referencing the issue:
   ```bash
   git commit -m "feat: add new monitoring capability (#123)"
   ```

## Pull Request Process

1. Update the README.md or relevant documentation
2. Update the `Chart.yaml` version using semantic versioning
3. Submit a pull request to the main repository
4. Ensure your PR references the associated issue number

## Testing Requirements

- All new features should include appropriate tests
- Run existing tests before submitting

## Style Guidelines

- Follow standard Helm chart best practices
- Use YAML indentation of 2 spaces
- Organize chart resources logically
- Comment complex templates

## Documentation

- Update documentation affected by your changes
- Document new values in `values.yaml`
- Provide examples for significant features

## Communication

- Use GitHub Issues for bug reports and feature requests

Thank you for contributing to GREYE!
