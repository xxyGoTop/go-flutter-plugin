# runbash

This Go package implements the host-side of the Flutter [runbash](https://github.com/xxyGOTop/go-flutter-plugin) plugin.

## Usage

Import as:

```go
import runbash "https://github.com/xxyGOTop/go-flutter-plugin/runbash"
```

Then add the following option to your go-flutter [application options](https://github.com/go-flutter-desktop/go-flutter/wiki/Plugin-info):

```go
flutter.AddPlugin(&runbash.OpenFilePlugin{}),
```
