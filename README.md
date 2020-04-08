## Bosch Motronic checksum correction tool

Cross-platform checksum correction tool for Bosch Motronic ECUs. Supported firmware auto recognition built-in.

#### Supported ECUs:

- [x] BMW DME 402
- [x] BMW DME 403
- [ ] BMW DME 405
- [ ] BMW DME 413
- [ ] BMW DME 173

#### Usage:

1. Build tool
    ```
    make build
    ```

2. Use it to correct checksums. It will ask for input file. File can be dragged into command window. Tool will overwrite input file.
    ```
    motronic_checksum
    ```

#### Test:

Before running unit tests make sure you have necessary original firmware version inside firmwares directory
```
make test
```

Test coverage
```
make coverage
```