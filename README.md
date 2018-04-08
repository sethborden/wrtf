# wrtf
Write from stdin to a file...yeah `>` does this, but`xargs` and `>` don't get along.

## Usage
Pipe somthing to `wrtf` via stdin like so:

```bash
date | wrtf -o date.txt
```
