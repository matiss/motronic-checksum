## Bosch Motronic checksum correction tool

Cross-platform checksum correction tool for Bosch Motronic ECUs. Supported firmware auto recognition built-in.

#### Supported ECUs:

- [x] BMW DME 0 261 200 402
- [x] BMW DME 0 261 200 403
- [X] BMW DME 0 261 200 405
- [ ] BMW DME 0 261 200 413
- [ ] BMW DME 0 261 200 173

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

Before running unit tests make sure you have necessary original firmware version inside firmwares directory
```
make test
```

Test coverage
```
make coverage
```