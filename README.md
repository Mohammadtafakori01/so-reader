# SO Reader

This is a simple tool written in Go by Kaveh Tafakori (tafakori.dev@gmail.com) to read shared object files (`.so`) in Linux environments.

## Prerequisites
- Linux environment

## Develop and Build
After making changes to the `soreader.go` file, you can run the following command to build and execute the tool:
```
go run soreader.go
```

## How to Use

### 1. Single File:
- With Path:
  ```
  ./soreader file.so /path
  ```
- Without Path:
  ```
  ./soreader file.so
  ```

### 2. Recursive Use:
- With Path:
  ```
  find "$(pwd)" -type f -name "*.so" | xargs -I {file} sh -c '/home/tafakori.k/Downloads/soreader/soreader "{file}" /path/jsons;'
  ```
- Without Path:
  ```
  find "$(pwd)" -type f -name "*.so" | xargs -I {file} sh -c '/home/tafakori.k/Downloads/soreader/soreader "{file}";'
  ```

For any inquiries or feedback, feel free to reach out to Kaveh Tafakori at tafakori.dev@gmail.com.
