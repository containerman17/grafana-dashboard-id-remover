# Grafana Dashboard ID Remover

A simple utility to clean Grafana dashboard JSON files by removing UIDs, IDs, and datasource references. This is particularly useful when moving dashboards between Grafana instances or version control systems.

## Installation

```bash
go install github.com/containerman17/grafana-dashboard-id-remover@latest
```

## Usage

Run with default path (`/etc/grafana/provisioning/dashboards`):
```bash
go run github.com/containerman17/grafana-dashboard-id-remover
```

Or specify a custom path:
```bash
go run github.com/containerman17/grafana-dashboard-id-remover /path/to/dashboards
```

## What it does

The tool:
- Removes dashboard `uid` and `id` fields
- Removes `datasource` references from panel targets
- Processes all `.json` files in the specified directory
- Preserves the original file structure and formatting

## Development

To run tests:
```bash
go test ./...
```

## License

Unlicense
