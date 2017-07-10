Wrapper for text proofreading API.
Note: Spanish only

Link to docs is <a href="https://www.meaningcloud.com/developer/text-proofreading">here</a>

Example usage:

```go
	package main

	import(
		"fmt"
		"github.com/bas24/meaningcloud-api/proofer"
	)

	const text string = `No vengas hacernos creer lo que todos vemos los mismos pogramas de television.`

	func main(){
		client, err := proofer.New("YOUR_API_KEY")
		if err != nil {
			fmt.Println(err)
		}

		result, err := client.Check(text)
		if err != nil {
			fmt.Println(err)
		}

		fmt.Println(result)
	}
```
