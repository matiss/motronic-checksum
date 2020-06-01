## Bosch Motronic checksum correction tool

Cross-platform checksum correction tool for Bosch Motronic ECUs. Supported firmware auto recognition built-in.

#### Supported ECUs:

- [x] BMW DME 0 261 200 173 (705, 794)
- [x] BMW DME 0 261 200 402 (098, 599)
- [x] BMW DME 0 261 200 403 (547, 950)
- [x] BMW DME 0 261 200 404 (689)
- [x] BMW DME 0 261 200 405 (951)
- [x] BMW DME 0 261 200 413 (609, 623, 715)
- [x] BMW DME 0 261 203 484 (582)
- [x] BMW DME 0 261 203 590 (597)

Some of software versions might not be supported by this tool even though hardware codes are same. If you want to add support for some specific ROM, you can make a pull request or create a new issue ticket with hardware and software (chip) codes and link to original firmware/ROM. The tool will not try to calculate checksum if it can't recognise firmware/ROM.

#### Usage:

1. Build tool
    ```
    make build
    make build-windows
    make build-linux
    ```

2. Use it to correct checksums. It will ask for input file. File can be dragged into command window. Tool will overwrite input file.
    ```
    motronic_checksum
    ```

#### Downloads:

Latest platform specific pre-compiled binaries can be downloaded under [releases](https://github.com/matiss/motronic-checksum/releases) section


#### Test:

Before running unit tests make sure you have necessary original firmware version inside firmwares directory. ROMs can be downloaded [here](https://www.dropbox.com/sh/7waxylurxvu9qo1/AADMF_GXHVlr8CWAhL7JcERna?dl=0)
```
make test
```

Test coverage
```
make coverage
```

### Coffee Support

You can support me by buying me a [coffee](https://www.buymeacoffee.com/matiss)

### License

Copyright (c) 2020-present [matiss](https://github.com/matiss). Motronic Checksum tool is free and open-source software licensed under the MIT License.