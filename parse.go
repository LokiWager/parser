package parser

type State string

const (
	InitialState   State = "InitialState"
	SpaceState           = "SpaceState"
	WordState            = "WordState"
	DotState             = "DotState"
	NumberState          = "NumberState"
	TimeState            = "TimeState"
	UnderlineState       = "UnderlineState"
)

type expr struct {
	index        int
	state        State
	wordIndex    int
	generateWord bool
}

const (
	InitialIndex = -1
)

func newExpr() *expr {
	return &expr{
		index:     InitialIndex,
		state:     InitialState,
		wordIndex: InitialIndex,
	}
}

func Parse(input []byte) (tokens []string) {
	expr := newExpr()
	cacheWord := ""

	for i, b := range input {
		expr.index = i
		if b >= '0' && b <= '9' {
			expr.ParseNumber()
		} else if b == '.' {
			expr.ParseDot()
		} else if b == ' ' {
			expr.ParseSpace()
		} else if b == ':' {
			expr.ParseColon()
		} else if b == '-' {
			expr.ParseUnderline()
		} else if (b >= 'a' && b <= 'z') || (b >= 'A' && b <= 'Z') {
			cacheWord += string(b)
			expr.ParseWord(cacheWord)
		}

		if expr.generateWord {
			tokens = append(tokens, generateToken(input[expr.wordIndex:expr.
				index]))
			if expr.index != i {
				expr.wordIndex = expr.index + 1
			}
			expr.generateWord = false
			cacheWord = ""
		}
	}

	if expr.state != InitialState {
		tokens = append(tokens, generateToken(input[expr.wordIndex:]))
	}

	return
}

func (e *expr) ParseNumber() {
	switch e.state {
	case SpaceState, InitialState, UnderlineState:
		e.state = NumberState
		e.wordIndex = e.index
	case DotState:
		e.state = NumberState
	}
}

func (e *expr) ParseDot() {
	switch e.state {
	case WordState, TimeState:
		e.generateWord = true
		e.state = InitialState
	case SpaceState:
		e.state = InitialState
	case NumberState, InitialState:
		e.state = DotState
	}
}

func (e *expr) ParseSpace() {
	switch e.state {
	case WordState, DotState:
		e.generateWord = true
		e.state = InitialState
	}
}

func (e *expr) ParseColon() {
	switch e.state {
	case NumberState:
		e.state = TimeState
	}
}

func (e *expr) ParseUnderline() {
	switch e.state {
	case NumberState:
		e.generateWord = true
		e.state = InitialState
	}
}

func (e *expr) ParseWord(preWord string) {
	switch e.state {
	case InitialState, SpaceState, DotState:
		e.wordIndex = e.index
		e.state = WordState
	case TimeState, NumberState:
		if !isTime(preWord) {
			e.state = WordState
			e.generateWord = true
			e.index -= len(preWord)
		}
	}
}

func isTime(preWord string) bool {
	return preWord == "a" || preWord == "p" || preWord == "am" || preWord == "pm"
}

func generateToken(input []byte) string {
	index := 0
	token := make([]byte, 0, len(input))
	for index < len(input) {
		if input[index] == '-' && index+1 < len(
			input) && input[index+1] == '\n' {
			index += 2
			continue
		}
		if index == len(input)-1 && input[index] == '.' {
			index += 1
			continue
		}
		token = append(token, input[index])
		index += 1
	}
	return string(token)
}
