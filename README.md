# before
```go
type (
	exampleType struct {
		String string
		Int    int
		Float  float64
		Struct exampleStruct

		Array   []int
		Map     map[string]string
		Pointer *int
	}

	exampleStruct struct {
		Name string
	}
)
```
# after
```json
{
  "String": "73a31271-5d19-4a0a-bf27-b1d9c47b627f",
  "Int": 6,
  "Float": 7.373282048929731,
  "Struct": {
    "Name": "79b5d696-c893-4734-ad9b-f475e3af4fef"
  },
  "Array": [
    9
  ],
  "Map": {
    "738d7604-2f1a-4608-ae76-c8883560d303": "52cc52c1-0aee-4b3c-9e08-7d634ae90656"
  },
  "Pointer": 3
}
```