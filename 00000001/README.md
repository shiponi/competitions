# No.1 Find words

## Basic Information

<table>
  <tbody>
    <tr>
      <td><strong>No.</strong></td>
      <td>1</td>
    </tr>
    <tr>
      <td><strong>Name</strong></td>
      <td>Find words</td>
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

Implement the `Find(path, s string) (string, error)` function in the `bench` package, which does the following things:

1. Return some kind of `error` if `s` is empty
2. Read the text file on the `path`
3. Find the `s` words on the file
4. Return the row numbers and indexs in the form of `r:c,r:c,...,r:c`, at which the `s` words exist

## Example

If the text file on the `./data.txt` contains the following data,

```
aabbaa
xxxxxx
aaabbb
```

`Find("./data.txt", "aa")` should return `"1:0,1:4,3:0,3:1", nil`.

## Sample Program

The [sample](sample) directory contains a sample program and a testing code. You can start with this sample program.
