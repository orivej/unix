A collection of small UNIX tools.

## Every

`every` runs a command periodically. It can snap the period boundary to a rounded grid, so that e.g. `every -r 6m cmd` runs cmd at 1:00:00, 1:06:00, 1:12:00 etc.

## Inplace

`inplace` helps filter files with commands.

`inplace file grep pattern` is like `grep pattern <file >file~ && mv file~ file || rm file~ && exit 1`.

## Regrep

`regrep` combines `grep` and `sed` and supports regular expressions that match multiple lines.

`regrep gs '\(([^(]*?)\)' '$1'` will print everything between `(` and `)`.

## TS

`ts` filters input and timestamps each line with durations since the previous line and since program startup.

`{ sleep 1; echo test; } | ts` may print `  0.998   0.998   test`.
