# No.1 Reading a file

## Basic Information

<table>
  <tbody>
    <tr>
      <td><strong>No.</strong></td>
      <td>1</td>
    </tr>
    <tr>
      <td><strong>Name</strong></td>
      <td>Reading a file</td>
    </tr>
    <tr>
      <td><strong>Start Data & Time</strong></td>
      <td>2015-03-22 12:00 a.m. UTC</td>
    </tr>
    <tr>
      <td><strong>End Data & Time</strong></td>
      <td>2015-04-12 11:00 p.m. UTC</td>
    </tr>
  </tbody>
</table>

## Question

Implement the `Count(path string) uint64` function in the `bench` package, which does the following things:

1. Read the text file which contains only `0 (U+0030)`, `1 (U+0031)` and `LF (U+000A)` on the `path` specified as a parameter
2. Count the number of `1 (U+0031)` in the file and return it

## Sample Program

The [sample](sample) directory contains a sample program and a testing code. You can start with this sample program.
