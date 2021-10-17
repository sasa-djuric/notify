# Notify

A CLI tool for firing OS native notifications.

## Motivation

Main motivation for building this tool is to get notification when some process is finished.
For example running:

```sh
$ apt-get update; notify -s Update
```

## Example

```sh
$ notify --title "My notification" --description "Example description"
```

## Options

| Option               | Alias | Description                                                                                                 |
| -------------------- | ----- | ----------------------------------------------------------------------------------------------------------- |
| --title              | -t    | Notification title                                                                                          |
| <nobr/>--description | -d    | Notification description                                                                                    |
| --subject            | -s    | Subject of notification. If title option is not provided title of notification will be "{subject} finished" |
| --version            | -v    | Tool version                                                                                                |
| --help               | -h    | Help                                                                                                        |

## License

[MIT](https://github.com/sasa-djuric/notify/blob/master/LICENSE)
