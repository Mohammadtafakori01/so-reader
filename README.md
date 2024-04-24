# Develop and build
# after change soreader.go file
go run soreader.go

# How to use
# 1- single file:
# with path
./soreader file.so /path
# without /path
./soreader file.so

# 2- recursive use:
# with /path
find "$(pwd)" -type f -name "*.so"  | xargs -I {file} sh -c '/home/tafakori.k/Downloads/soreader/soreader "{file}" /path/jsons;'
# without /path
find "$(pwd)" -type f -name "*.so"  | xargs -I {file} sh -c '/home/tafakori.k/Downloads/soreader/soreader "{file}";'
