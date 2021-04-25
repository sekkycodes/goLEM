package provide

import "io/ioutil"

// JsonFileGolem is a very simple provider reading in JSON files
type JsonFile struct {
	Source string
}

func (jsg *JsonFile) Provide() (json []byte, err error) {
	content, err := ioutil.ReadFile(jsg.Source)
	if err != nil {
		return nil, err
	}

	return content, nil
}
