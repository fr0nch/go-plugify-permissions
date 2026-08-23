[![Русский](https://img.shields.io/badge/Русский-%F0%9F%87%B7%F0%9F%87%BA-green?style=for-the-badge)](README_ru.md)

# Go bindings for the Permissions Plugify plugin

This repository contains Go bindings for the [plugify-plugin-permissions](https://github.com/untrustedmodders/plugify-plugin-permissions) plugin. The bindings are automatically generated and kept in sync with the original plugin.

## API Documentation

Full API documentation can be found here on the [API Hub](https://api.plugify.net?file=https://raw.githubusercontent.com/untrustedmodders/plugify-plugin-permissions/refs/heads/main/plugify-plugin-permissions.pplugin.in) website.

## Installation

```bash
go get github.com/fr0nch/go-plugify-permissions
```

## Updates

This repository uses GitHub Actions to automatically check for updates in [plugify-plugin-permissions](https://github.com/untrustedmodders/plugify-plugin-permissions) every day at 2:00 UTC. When a new version is detected, the bindings are updated, and a new release is created.

## License

This Go module for Plugify is licensed under the [MIT License](LICENSE).