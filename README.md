### PARSE WORD

Extract words from a string.

Features:

1. Treat `in-\nstreresting` as `insteresting`
2. Treat `11:00 am` as one word
3. Treat `11.11` as a number

Examples:

```golang
    word := ` Hello World 100-20. in-
teresting 11.2 at 11:20 am`
    tokens := Parse([]byte(word))
    assert.Equal(t, []string{"Hello", "World", "100", "20", "interesting", 
"11.2", "at", "11:20 am"}, tokens)
    assert.Equal(t, 8, len(tokens))
```

API
```golang
    func Parse(input []byte) (tokens []string)
```
