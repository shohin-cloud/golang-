package handlers

type Friend struct {
    Name string `json:"name"`
    Age  int    `json:"age"`
}

var friends = []Friend{
    {Name: "John", Age: 30},
    {Name: "Alice", Age: 25},
	{Name: "Shohin", Age: 20},
	{Name: "Dalila", Age: 21},
	{Name: "Shahzod", Age:24}

}
