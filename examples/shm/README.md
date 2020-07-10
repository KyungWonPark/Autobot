# shm/

Shared memory example

# How to
```bash
$ go run producer/main.go
$ gcc consumer/main.c
$ ./a.out <Shared_Memory_Region_ID>
```

# Caution
- Producer must be running while consumer is reading
- "int" in Golang is 64 bit by default, in AMD64 while int is 32 bit in C. So use int32 in Golang
