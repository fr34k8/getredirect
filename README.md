# getredirect

Get redirect of URL, programmed in Go. Install with:

```shell
go get github.com/binaryfigments/getredirect
```

# Example usage

```go
package main

import (
	"encoding/json"
	"fmt"

	"github.com/binaryfigments/getredirect"
)

func main() {
	testsite := "https://www.binaryfigments.com"
	header := getredirect.From(testsite)
	json, err := json.MarshalIndent(header, "", "   ")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(json))
}
```

# Example response

```json
{
   "Question": {
      "JobURL": "https://www.binaryfigments.com",
      "JobStatus": "OK",
      "JobMessage": "Job done!",
      "JobTime": "2016-08-21T20:26:55.710800393+02:00"
   },
   "Answer": {
      "HTTPStatus": 301,
      "HTTPStatusText": "301 Moved Permanently",
      "HTTPRedirect": "https://binaryfigments.com",
      "HTTPProto": "HTTP/2.0",
      "IPAddress": [
         "2a01:448:1003::130",
         "213.249.93.130"
      ]
   }
}
```

# Dependencies

 * Go 1.6.x tested https://golang.org
 * govalidator https://github.com/asaskevich/govalidator

```
go get github.com/asaskevich/govalidator
```

# The MIT License (MIT)

Copyright (c) 2016 Sebastian Broekhoven
~~~
Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.
~~~