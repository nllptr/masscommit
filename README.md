# masscommit
A simple program that generates a lot of git commits to a specified file. This could be used for performance tests on git blame.

## Usage

```masscommit -n 10 -f "File"```

```-n```	This is the number of commits to perform.
```-f```	This is the file to edit and commit. At the moment, it must exist before running masscommit.

The edits to the specified file are just appended to the end and have the format "Edit number n", where n is just a counter.
