[![English](https://img.shields.io/badge/English-%F0%9F%87%AC%F0%9F%87%A7-blue?style=for-the-badge)](README.md)

# Go биндинги для Permissions Plugify плагина

Этот репозиторий содержит Go биндинги для плагина [plugify-plugin-permissions](https://github.com/untrustedmodders/plugify-plugin-permissions). Биндинги генерируются автоматически и синхронизируются с оригинальным плагином.

## Документация по API

Полную документацию по API можно найти здесь на сайте [API Hub](https://api.plugify.net?file=https://raw.githubusercontent.com/untrustedmodders/plugify-plugin-permissions/refs/heads/main/plugify-plugin-permissions.pplugin.in)

## Установка

```bash
go get github.com/fr0nch/go-plugify-permissions
```

## Обновление

Этот репозиторий использует GitHub Actions для автоматической проверки обновлений в [plugify-plugin-permissions](https://github.com/untrustedmodders/plugify-plugin-permissions) в каждые 2:00 по UTC. При обнаружении новой версии, биндинги обновляются, и создается новый релиз.

## Лицензия

Этот модуль Go для Plugify лицензирован на условиях [MIT License](LICENSE).