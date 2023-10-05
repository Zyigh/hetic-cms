package internal

func Map[Tin, Tout any](arr []Tin, callable func(Tin) Tout) []Tout {
	out := make([]Tout, len(arr))

	for i, el := range arr {
		out[i] = callable(el)
	}

	return out
}
