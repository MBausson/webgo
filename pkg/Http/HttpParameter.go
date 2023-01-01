package Http

import "strings"

type HttpParameter struct {
	Name  string
	Value string
}

type ParameterContainer struct {
	elements []HttpParameter
}

func (pc ParameterContainer) Get(name string) (param HttpParameter, found bool) {
	for _, e := range pc.elements {
		if e.Name == name {
			return e, true
		}
	}

	return HttpParameter{}, false
}

func NewParameterContainer(params []HttpParameter) ParameterContainer {
	return ParameterContainer{
		elements: params,
	}
}

func ParameterFromString(content string) HttpParameter {
	splitted := strings.Split(content, "=")

	return HttpParameter{
		Name:  splitted[0],
		Value: strings.Join(splitted[1:], "="),
	}
}

func UrlToParameters(url string) ParameterContainer {
	if len(url) == 0 {
		return NewParameterContainer([]HttpParameter{})
	}

	// If the url starts with a &, remove it
	if url[0] == '?' {
		url = url[1:]
	}

	params := []HttpParameter{}
	splitted := strings.Split(url, "&")

	for _, e := range splitted {
		params = append(params, ParameterFromString(e))
	}

	return NewParameterContainer(params)
}
