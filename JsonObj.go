/*
Copyright (c) 2017 Mohammad Mohebi

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.
*/
package mmj

import (
	"strings"
	"encoding/json"
)

type JsonObj map[string]interface{}
type JsonArray []interface{}

//Create a new JSON objet
func NewObj() JsonObj{
	return make(JsonObj)
}

//Create a new JSON Array
func NewArray() JsonArray{
	return make(JsonArray, 0)
}

// Set an objet to a specific path by using a concatenated path
// if path already exist, it will overrid the value
// if the path doesnt exist, it will create it
//	ex: path.to.the.object
//
//		{
//			"path": {
//				"to": {
//					"the": {
//						"object":"value"
//					}
//				}
//			}
//		}
//
func (j *JsonObj) SetP(obj interface{}, path string) {
	p := strings.Split(path, ".")
	j.setObj(p, obj, 0)
}

// Do the same thing as SetP, but the path is given as different parameters
func (j *JsonObj) Set(d interface{}, path... string) {
	j.setObj(path, d, 0)
}

//Append an objet to a JSON array
func (j *JsonArray) Append(o interface{}){
	*j = append(*j, o)
}

//Return JSON output as array of bytes
func (j *JsonArray) Bytes() ([]byte, error){
	b, e := json.Marshal(j)
	return b, e
}

//Return JSON output as string
func (j *JsonArray) String() (string, error){
	b, e := json.Marshal(j)
	s := string(b[:])
	return s, e
}

//Function the is intern to the package and will create the
//JSON structure inside JSON object
func (j *JsonObj) setObj(path []string, d interface{}, iteration int) {
	if path != nil && len(path) > 0 && iteration < len(path){
		if iteration < len(path)-1 {
			key := path[iteration]
			iteration++
			v, OK := (*j)[key]
			if OK {
				val, OK := v.(JsonObj)
				if OK {
					val.setObj(path, d, iteration)
					(*j)[key] = val
				} else {
					return
				}

			} else {
				newObj := make(JsonObj)
				newObj.setObj(path, d, iteration)
				(*j)[key] = newObj
			}
		} else {
			(*j)[path[iteration]] = d
		}
	}
}
