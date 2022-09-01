### PARSE WORD

Extract words from a string.

Use spaces or periods as separators, and also support line feeds, see the State Machine for detailed processing logic


Legal character set：
* a-z
* A-Z
* space
* . : - \n

Features:

1. Treat `in-\nstreresting` as one word
2. Treat `11:00` as one word
3. Treat `11.11` as a number

Examples:

```golang
    word := ` Hello World 100-20. in-
teresting 11.2 at 11:20 am`
    count := Parse([]byte(word))
    assert.Equal(t, 9, count)
```

API
```golang
    func Parse(input []byte) (count int)
```

![state machine](./docs/state%20machine.svg)
