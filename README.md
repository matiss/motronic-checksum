## Bosch Motronic checksum correction tool

Cross-platform checksum correction tool for Bosch Motronic ECUs.

#### Supported ECUs:

- [x] BMW DME 402
- [x] BMW DME 403
- [x] BMW DME 403
- [] BMW DME 405
- [] BMW DME 413
- [] BMW DME 173
    
#### Requirement:

1. Postgres database
2. Golang
3. GNU Make (Optional)

#### Usage:

1. Build tool
    ```
    make build
    ```

2. Use it to correct checksums
    ```
    motronic_checksum 403_mod.bin
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