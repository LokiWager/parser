package parser

import "errors"

// State interface in the State Machine
type State interface {
	// Transition: peek an Event, and return next State
	Transition(ctx *context, event Event) State
}

// Event type in the State Machine
type Event byte

// Information that needs to passed in context
type context struct {
	count int
	cache bool
}

const (
	InitialCount = 0
	EOF          = '\000'
	Hyphen       = '-'
)

var (
	initialState = &InitialState{}
	wordState    = &WordState{}
	digitState   = &DigitState{}
	finialState  = &FinalState{}
)

var IllegalCharacter = errors.New("character must be ascii code")

// InitialState which will be entered when recounting
type InitialState struct{}

func (state *InitialState) Transition(_ *context, event Event) State {
	if isLetter(event) {
		return wordState
	}
	if isDigit(event) {
		return digitState
	}
	return state
}

// WordState which will enter this state when a letter is received
type WordState struct{}

func (state *WordState) Transition(ctx *context, event Event) (result State) {
	if isLetter(event) {
		return state
	}

	if isHyphen(event) {
		ctx.cache = true
		return state
	}

	if isNewLine(event) && ctx.cache {
		return state
	}

	result = initialState

	if isDigit(event) {
		result = digitState
	}

	if event == EOF {
		result = finialState
	}

	ctx.count += 1
	ctx.cache = false
	return
}

// DigitState which means currently traversed word may be a digit.
type DigitState struct{}

func (state *DigitState) Transition(ctx *context, event Event) (result State) {
	if isDigit(event) || isDelimiter(event) {
		return state
	}

	result = initialState

	if isLetter(event) {
		result = wordState
	}

	if event == EOF {
		result = finialState
	}

	ctx.count += 1
	return
}

// FinalState which means current event is EOF. Stop state machine.
type FinalState struct{}

func (state *FinalState) Transition(_ *context, _ Event) State {
	return state
}

func newContext() *context {
	return &context{
		count: InitialCount,
	}
}

// WordCount Count the number of English words and digits in the text.
// precision: when precision is false, if input contains non-ascii characters,
// processing continues. However, return IllegalCharacter error
// see more examples in test file
func WordCount(input []byte, precision bool) (count int, err error) {
	ctx := newContext()

	var state State
	state = &InitialState{}
	for _, b := range input {
		if !isValidEvent(Event(b)) && precision {
			err = IllegalCharacter
			return
		}
		state = state.Transition(ctx, Event(b))

		if state == finialState {
			count = ctx.count
			return
		}
	}
	state = state.Transition(ctx, EOF)

	count = ctx.count

	return
}

// event must be a legitimate ascii code
func isValidEvent(event Event) bool {
	return event <= 127
}

func isLetter(event Event) bool {
	return (event >= 'a' && event <= 'z') || (event >= 'A' && event <= 'Z')
}

func isDigit(event Event) bool {
	return event >= '0' && event <= '9'
}

func isNewLine(event Event) bool {
	return event == '\n' || event == '\r'
}

func isDelimiter(event Event) bool {
	return event == '.' || event == ':' || event == Hyphen
}

func isHyphen(event Event) bool {
	return event == Hyphen
}
