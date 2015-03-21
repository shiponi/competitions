# No.1 Reading a file

## Basic Information

**No.:** 1

**Name:** Reading a file

**Start Data & Time:** Mar. 22, 2015, 12:00 a.m. UTC

**End Data & Time:** Apr. 12, 2015, 11:00 p.m. UTC

## Question

Implement the `Count(path string) uint64` function in the `bench` package, which does the following things:

1. Read the text file which contains only `0 (U+0030)`, `1 (U+0031)` and `LF (U+000A)` on the `path` specified as a parameter
2. Count the number of `1 (U+0031)` in the file and return it

## Sample Program

The [sample](sample) directory contains a sample program and a testing code. You can start with this sample program.
