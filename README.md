# Gate User Sync

The Gate User Sync service compares active personnel on KolayIK with users in the Alternatif SuperApp customer group. If there are users who should be added, the service sends a request to the Alternatif SuperApp API to include them. In addition, the service fetches information about inactive employees from KolayIK to ensure they are removed from the Alternatif SuperApp customer group.

## Obtaining Environment Information

Visit https://apidocs.kolayik.com to obtain the KolayIK API token. This token should have the `person:list` and `person:bulk-view` permissions.

To acquire the Alternatif SuperApp API token, contact the responsible personnel at Alternatif SuperApp.

## Installation

The project is written in Go 1.20. After cloning the repository onto your machine, fill out the values in the .env file. Then, execute the following commands to run the project:

```bash
cp .env.example .env
go mod download
go run main.go
```

### Running with Docker

The project includes a Makefile to provide shortcuts for Docker commands. Use the following command to run the project with Docker:

```bash
make up
```
