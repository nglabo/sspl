# Spell

Bypass [SSL Pinning](https://medium.com/@appmattus/android-security-ssl-pinning-1db8acb6621e) on Android

## Installation

### Arch Linux

Install [sspl](https://aur.archlinux.org/packages/sspl/) <sup>AUR</sup> :

```sh
$ yay -S sspl
```

## Usage

- **sspl** _FILE_...

### Example

```sh
$ sspl com.example.apk
[success] /tmp/sspl-696463700/build/com.example.apk
sspl com.example.apk  46.69s user 2.01s system 208% cpu 23.377 total
```

## License

This project is distributed under the [MIT license](LICENSE)

##### Thanks to [levyitay/AddSecurityExceptionAndroid](https://github.com/levyitay/AddSecurityExceptionAndroid)
