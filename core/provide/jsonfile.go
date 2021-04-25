package provide

import "io/ioutil"

type JsonFileGolem struct {
}

type Source interface {
}

func (jsg *JsonFileGolem) Provide(source string) (json []byte, err error) {
	content, err := ioutil.ReadFile(source)
	if err != nil {
		return nil, err
	}

	return content, nil
}
