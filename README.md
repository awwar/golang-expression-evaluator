# golang-expression-evaluator

```bash
"result = " + (1 + 2 * sum(3 + 4) / 5 + 6)
[
	{value: "result = ", type: string}
	{value: "+", type: operation}
	{value: "(", type: bracket}
	{value: "1", type: number}
	{value: "+", type: operation}
	{value: "2", type: number}
	{value: "*", type: operation}
	{value: "sum", type: word}
	{value: "(", type: bracket}
	{value: "3", type: number}
	{value: "+", type: operation}
	{value: "4", type: number}
	{value: ")", type: bracket}
	{value: "/", type: operation}
	{value: "5", type: number}
	{value: "+", type: operation}
	{value: "6", type: number}
	{value: ")", type: bracket}
]
+
└── #0 "result = "
└── #1 +
      └── #0 +
            └── #0 1
            └── #1 /
                  └── #0 *
                        └── #0 2
                        └── #1 sum
                              └── #0 +
                                    └── #0 3
                                    └── #1 4
                  └── #1 5
      └── #1 6

"result = 9.8"
```