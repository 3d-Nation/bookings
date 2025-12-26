package forms

type errors map[string][]string

// adds error message for a given form field
func (e errors) Add(field, message string) { //should reciever ba a pointer??  Isn't in video (yet at least)
	e[field] = append(e[field], message)
}

// returns the first error message
func (e errors) Get(field string) string { //same comment as above about receiver
	es := e[field]
	if len(es) == 0 {
		return ""
	}

	return es[0]
}
