# GREYE Helm Chart

## Introduction

GREYE is a Kubernetes application that provides cluster monitoring and notification capabilities. This helm chart facilitates easy deployment of GREYE in your Kubernetes environment.

## Prerequisites

- Kubernetes 1.16+
- Helm 3.0+

## Installation

Add the GREYE repository:

```shell
helm repo add greye https://ftrigari.github.io/helm-charts
```

Install the chart:

```shell
helm install greye greye/greye --namespace greye --create-namespace
```

## Configuration

The following table lists the configurable parameters for the GREYE helm chart:

| Parameter | Description | Default |
|-----------|-------------|---------|
| `notification.telegram.destination` | Telegram destination ID | `""` |
| `notification.telegram.token` | Telegram bot token | `""` |
| `notification.email.destination` | Email destination | `""` |
| `notification.email.token` | Email service token | `""` |
| `monitoring.enabled` | Enable monitoring functionality | `false` |
| `cluster.intervalSeconds` | Interval between cluster checks in seconds | `60` |
| `cluster.timeoutSeconds` | Timeout for requests in seconds | `5` |
| `cluster.maxFailedRequests` | Maximum number of failed requests before action | `3` |
| `cluster.myIp` | IP address of this node | `""` |
| `cluster.ip` | List of cluster node IPs | `[]` |

### Custom Values

To customize the installation, create a `values.yaml` file and install the chart using:

```shell
helm install greye greye/greye --namespace greye --create-namespace -f values.yaml
```

## Uninstallation

```shell
helm uninstall greye -n greye
```

Note: The chart includes a pre-delete hook that cleans up resources by making a DELETE request to the GREYE API.

## Version Information

- Chart Version: 0.1.0
- App Version: 1.16.0

## Links

- [GitHub Repository](https://github.com/greye-monitoring/helm-charts)