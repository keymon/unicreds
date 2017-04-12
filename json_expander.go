package unicreds

import (
	"encoding/json"
	"errors"
	"fmt"
)

// http://stackoverflow.com/questions/30341588/how-to-parse-a-complicated-json-with-go-lang-unmarshal

func JsonExpand(
	rawJson []byte,
	prefix string,
	secretRetriever func(key string) (string, error),
) (expandedRawJson []byte, err error) {

	var walkAndExpandJson func(treeNode interface{}) (interface{}, error)

	walkAndExpandJson = func(treeNode interface{}) (interface{}, error) {
		switch v := treeNode.(type) {
		case string:
			fmt.Println("is string", v)
			return v, nil
		case int:
			return v, nil
		case map[string]interface{}:
			fmt.Println("is an hash:")
			a := map[string]interface{}{}
			for k, element := range v {
				fmt.Println("Processing ", k)
				r, err := walkAndExpandJson(element)
				if err != nil {
					return nil, err
				}
				a[k] = r
			}
			return a, nil
		case []interface{}:
			fmt.Println("is an array:")
			a := []interface{}{}
			for i, element := range v {
				fmt.Println("Processing ", i)
				r, err := walkAndExpandJson(element)
				if err != nil {
					return nil, err
				}
				a = append(a, r)
			}
			return a, nil
		default:
			return nil, errors.New("is of a type I don't know how to handle")
		}
	}

	var baseJson interface{}
	err = json.Unmarshal(rawJson, &baseJson)
	if err != nil {
		return nil, err
	}

	expandedJson, err := walkAndExpandJson(baseJson)
	if err != nil {
		return nil, err
	}

	expandedRawJson, err = json.Marshal(expandedJson)
	if err != nil {
		return nil, err
	}
	return rawJson, nil
}
