# before
```go
type (
	exampleType struct {
		String string
		Int    int
		Struct exampleStruct

		Array []int
		Map   map[string]string
		// Pointer *int
	}

	exampleStruct struct {
		Name string
	}
)
```
# after
```json
{
  "String": "d7d442ac-e440-4296-8f4b-6750b3c307e5",
  "Int": 40,
  "Struct": {
    "Name": "83e95a53-3967-4c96-9ad7-45b77d17f1f8"
  },
  "Array": [
    11
  ],
  "Map": {
    "34901ae8-5ba3-4e9d-b1e0-7e80ab1c5159": "b5fd580a-7632-47b9-a9da-6e03de550cb5"
  }
}
```