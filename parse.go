package main

import (
	"encoding/json"
	"log"
)

func parse(str string) (*node, error) {
	root := &node{}
	
	// save the name current node
	var nodeName string
	// basically used to store json string
	var jsonStr string
	var flag bool

	// special characters used by str.
	openBrace := "["
	endBrace := "]"
	separator := ","

	nameTemplate := `{"name": "`
	childrenTemplate := `", "children": [`
	nodeEndingTemplate := `]}`

	for i, v := range str {
		// to get the character from byte.
		ch := string(v)

		// when character is starting bracket.
		if ch == openBrace {
			// to append the nodeName between name & children template.
			jsonStr += nameTemplate + nodeName + childrenTemplate

			// to clear the nodeName once we append it to string.
			nodeName = ""

			// to skip to the next loop cycle.
			continue
		}

		// when character is ending bracket.
		if ch == endBrace {
			// to append the nodeName between name & children template.
			if flag || nodeName != "" {
			
				jsonStr += nameTemplate + nodeName + childrenTemplate
				flag = false
			}

			// to clear the nodeName once we append it to string.
			nodeName = ""
			jsonStr += nodeEndingTemplate

			// to append the outer closing bracket.
			if i+1 == len(str) {
				jsonStr += nodeEndingTemplate
			}

			// to skip to the next loop cycle.
			continue
		}

		// when character is separator.
		if ch == separator {
			if nodeName == "" {
				// end the previous node
				jsonStr += nodeEndingTemplate + separator
			} else {
				// to append the node with no children
				jsonStr += nameTemplate + nodeName + childrenTemplate + nodeEndingTemplate + separator
				flag = true
			}

			// to clear the nodeName once we append it to string.
			nodeName = ""

			// to skip to the next loop cycle.
			continue
		}

		// to append the character to nodeName
		nodeName += ch
	}
	// if the there is no child 
	if nodeName != "" {
		jsonStr += nameTemplate + nodeName + childrenTemplate + nodeEndingTemplate
	}
  	
	// unmarshal the json string into node struct
	if err := json.Unmarshal([]byte(jsonStr), &root); err != nil {
		return nil, err
	}
	return root, nil
}

type node struct {
	Name     string  `json:"name"`
	Children []*node `json:"children,omitempty"`
}

var examples = []string{
	"[a,b,c]",
	"[a[aa[aaa],ab,ac],b,c[ca,cb,cc[cca]]]",
}

func main() {
	for i, example := range examples {
		result, err := parse(example)
		if err != nil {
			panic(err)
		}
		j, err := json.MarshalIndent(result, " ", " ")
		if err != nil {
			panic(err)
		}
		log.Printf("Example %d: %s - %s", i, example, string(j))
	}
}
