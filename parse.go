package parser

import "log"

type State interface {
	Transition(ctx *context, event Event) State
}
type Event byte

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

type WordState struct{}

func (state *WordState) Transition(ctx *context, event Event) State {
	if (event >= 'a' && event <= 'z') || (event >= 'A' && event <= 'Z') {
		return state
	}
	if event == '-' {
		ctx.cache[0] = event
		return state
	}
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
	if event == '-' {
		ctx.count += 1
		return initialState
	}
	return errorState
}

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

func Parse(input []byte) (count int) {
	ctx := newContext()

	var state State
	state = &InitialState{}
	for _, b := range input {
		state = state.Transition(ctx, Event(b))
	}

	if state != initialState {
		ctx.count += 1
	}
	count = ctx.count

	return
}
