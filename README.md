### PartialCopy
A command line tool for copying a portion of file, Like linux "dd" command, But much faster than it.

### How to build it
There's only one golang source file, Install [golang sdk](https://golang.org), Then run a command: `go build PartialCopy.go`

### How to use
Output from `PartialCopy.exe -h`

    Usage of PartialCopy.exe:
        -append
            append if file exists
        -count int
            copy max count bytes.
        -input string
            input filename. (position parameter is also ok)
        -output string
            output filename. (position parameter is also ok)
        -seek int
            seek bytes of output file. (negative number is supported)
        -skip int
            seek bytes of input file. (negative number is supported)

    Example:
        1. PartialCopy.exe src.txt dst.txt
           #copy src.txt to dst.txt, use position parameter

        2. PartialCopy.exe -output dst.txt src.txt
           #copy src.txt to dst.txt, use position and key parameter
           #unspecified parammeter will use position parameter

        3. PartialCopy.exe -skip 100 src.txt dst.txt
           #seek 100 bytes from beginning of src.txt before copy

        4. PartialCopy.exe -skip -100 src.txt dst.txt
           #seek 100 bytes from ending of src.txt before copy

        5. PartialCopy.exe -seek 100 src.txt dst.txt
           #seek 100 bytes from beginning of dst.txt before copy

        6. PartialCopy.exe -seek -100 src.txt dst.txt
           #seek 100 bytes from ending of dst.txt before copy

        7. PartialCopy.exe -count 50 src.txt dst.txt
           #copy 50 bytes from src.txt to dst.txt

        8. PartialCopy.exe -append src.txt dst.txt
           #use append mode instead of truncate if file exists

### Why need this tool
One day, I got a large of mega bytes file, But I don't want to use some first 100 bytes, Firstly, I use linux dd command like this: `dd if=xx of=xx ibs=1 skip=100`, I must specify `ibs=1` because I want to skip exact 100 bytes. But it will very very slow because it will only read 1 byte every time, Too much system call if the file is large. if not specify `ibs=1`, It will skip 512(default bs/ibs value) * 100 bytes, That's not what I want.

I'd like someone tell me if the linux release already have similiar command to do the same thing. mail: wly65535@163.com
