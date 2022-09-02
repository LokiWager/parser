### PARSE WORD

Count the number of English words and digits in the text

**Note:**
1. Only Support ascii codes
2. Treat '\0' as EOF


Features:

1. Treat `in-\nstreresting` as one word
2. Treat `11:00` as one word
3. Treat `11.11` as a number

API
```golang
    func WordCount(input []byte, precision bool) (count int, err error)
```
* When precision is false, if input contains non-ascii characters, 
  processing continues. However, return IllegalCharacter error
* See more examples in test file

![state machine](./docs/state%20machine.svg)
