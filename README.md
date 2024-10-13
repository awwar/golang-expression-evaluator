# golang-expression-evaluator


```bash
1 + 2 * (3 + 4) / 5 + 6
[
	{value: "1", type: number}
	{value: "+", type: operation}
	{value: "2", type: number}
	{value: "*", type: operation}
	{value: "(", type: bracket}
	{value: "3", type: number}
	{value: "+", type: operation}
	{value: "4", type: number}
	{value: ")", type: bracket}
	{value: "/", type: operation}
	{value: "5", type: number}
	{value: "+", type: operation}
	{value: "6", type: number}
]
+
└── L +
      └── L 1
      └── R /
            └── L *
                  └── L 2
                  └── R +
                        └── L 3
                        └── R 4
            └── R 5
└── R 6

result = 9.800000
```
