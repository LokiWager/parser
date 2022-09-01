package parser

import "log"

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
	cache [1]Event
}

const (
	InitialCount = 0
)

var (
	initialState = &InitialState{}
	wordState    = &WordState{}
	digitState   = &DigitState{}
	floatState   = &FloatState{}
	timeState    = &TimeState{}
	errorState   = &ErrorState{}
)

// InitialState which will be entered when recounting
type InitialState struct{}

func (state *InitialState) Transition(ctx *context, event Event) State {
	if (event >= 'a' && event <= 'z') || (event >= 'A' && event <= 'Z') {
		return wordState
	}
	if event >= '0' && event <= '9' {
		return digitState
	}
	if event == ' ' || event == '-' || event == '.' ||
		event == '\n' || event == ':' {
		return state
	}
	return errorState
}

// WordState which will enter this state when a letter is received
type WordState struct{}

func (state *WordState) Transition(ctx *context, event Event) State {
	if (event >= 'a' && event <= 'z') || (event >= 'A' && event <= 'Z') {
		return state
	}

	if event == '-' {
		ctx.cache[0] = event
		return state
	}
	// may be one word in the '-\n' case, or two word in the break line
	if event == '\n' {
		if ctx.cache[0] == '-' {
			ctx.cache = [1]Event{}
			return state
		}
		ctx.count += 1
		return initialState
	}
	if event == ' ' || event == '.' {
		ctx.count += 1
		return initialState
	}
	return errorState
}

// DigitState which means currently traversed word may be a digit.
type DigitState struct{}

func (state *DigitState) Transition(ctx *context, event Event) State {
	if event >= '0' && event <= '9' {
		return state
	}
	if event == ' ' {
		ctx.count += 1
		return initialState
	}
	if event == '.' {
		return floatState
	}
	if event == ':' {
		return timeState
	}
	// distinguish the case of letters, numbers are considered as two words.
	if event == '-' {
		ctx.count += 1
		return initialState
	}
	return errorState
}

// TimeState which means currently traversed word may be a time.
type TimeState struct{}

func (state *TimeState) Transition(ctx *context, event Event) State {
	if event >= '0' && event <= '9' {
		return state
	}
	if event == ' ' {
		ctx.count += 1
		return initialState
	}
	return errorState
}

// FloatState which means currently traversed word may be a float number.
type FloatState struct{}

func (state *FloatState) Transition(ctx *context, event Event) State {
	if event >= '0' && event <= '9' {
		return digitState
	}
	if event == ' ' || event == '.' {
		ctx.count += 1
		return initialState
	}

	return errorState
}

// ErrorState which means current event triggered an illegal path.
type ErrorState struct{}

func (state *ErrorState) Transition(ctx *context, event Event) State {
	log.Panicf("Illegal Event")
	return state
}

func newContext() *context {
	return &context{
		count: InitialCount,
	}
}

// Parse: Word count function to slice and dice with spaces and dot as key
// characters.
// " Hello Word." which will be as 2 word
// " 11.1 " which will be as 1 word
// "in-teresting" which will be as 1 word
// see more examples in test file
func Parse(input []byte) (count int) {
	ctx := newContext()

	var state State
	state = &InitialState{}
	for _, b := range input {
		state = state.Transition(ctx, Event(b))
	}

	if state == errorState {
		log.Panicf("Illegal Event")
	}

	if state != initialState {
		ctx.count += 1
	}
	count = ctx.count

	return
}
