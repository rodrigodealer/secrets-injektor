![Docker](https://github.com/rodrigodealer/secrets-injector/workflows/Docker/badge.svg)
![Rust](https://github.com/rodrigodealer/secrets-injector/workflows/Rust/badge.svg)
# secrets-injector

Application to deal with secrets from various providers and inject them into your application using environment variables.

## Building

You need rust [installed](https://rustup.rs/).

`$ cargo build --release`

## Running tests

`$ cargo test`

## Building your docker image

`$ docker build . -t youruser/secrets-injector:latest`

## Providers supported at the moment

- [x] Vault
- [x] AWS Secrets Manager
- [ ] AWS Parameter Store
- [ ] GCP Secrets Manager

## Contributing

[Contribution guidelines for this project](docs/CONTRIBUTING.md)
