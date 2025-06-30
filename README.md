## Program:

```bash
#MAIN
call #LOTTERY () $LOTTERY_RESULT

#LOTTERY() void
var (32) $SOME_VALUE

call #FLIP_COIN (0.5) $FLIP_COIN_RESULT

if ($FLIP_COIN_RESULT) (#MORE, #LESS)

print ("Try Next?")

return (void(0))

#MORE
print ("you're WIN!")

#LESS
var (sum("you're lose ", $SOME_VALUE, " points!")) $OUT_RESULT

print ($OUT_RESULT)

#FLIP_COIN(float($FLIP_COIN_ARG0)) bool

return (bool(rand() > $FLIP_COIN_ARG0))
```

## AST:

```bash
root
└── #MAIN
    └── call
        └── #LOTTERY
        └── $LOTTERY_RESULT
└── #LOTTERY
    └── void
    └── var
        └── $SOME_VALUE
        └── 32
    └── call
        └── #FLIP_COIN
        └── $FLIP_COIN_RESULT
        └── 0.500000
    └── if
        └── $FLIP_COIN_RESULT
        └── #MORE
        └── #LESS
    └── print
        └── Try Next?
    └── return
        └── void
            └── 0
└── #MORE
    └── print
        └── you're WIN!
└── #LESS
    └── var
        └── $OUT_RESULT
        └── sum
            └── you're lose 
            └── $SOME_VALUE
            └──  points!
    └── print
        └── $OUT_RESULT
└── #FLIP_COIN
    └── float
        └── $FLIP_COIN_ARG0
    └── bool
    └── return
        └── bool
            └── >
                └── rand
                └── $FLIP_COIN_ARG0
```

## IL:

```bash
0: MARK #MAIN 
1: JMP #LOTTERY 
2: VAR $LOTTERY_RESULT 
3: MARK #LOTTERY 
4: PUSH 32 
5: VAR $SOME_VALUE 
6: PUSH 0.500000 
7: JMP #FLIP_COIN 
8: VAR $FLIP_COIN_RESULT 
9: PUSH $FLIP_COIN_RESULT 
10: CSKIP 2 
11: JMP #LESS 
12: SKIP 1 
13: JMP #MORE 
14: PUSH Try Next? 
15: EXEC print 1 
16: PUSH 0 
17: EXEC void 1 
18: MARK #MORE 
19: PUSH you're WIN! 
20: EXEC print 1 
21: MARK #LESS 
22: PUSH you're lose  
23: PUSH $SOME_VALUE 
24: PUSH  points! 
25: EXEC sum 3 
26: VAR $OUT_RESULT 
27: PUSH $OUT_RESULT 
28: EXEC print 1 
29: MARK #FLIP_COIN 
30: EXEC float 1 
31: VAR $FLIP_COIN_ARG0 
32: EXEC rand 0 
33: PUSH $FLIP_COIN_ARG0 
34: EXEC > 2 
35: EXEC bool 1 
```

## Output:

```bash
you're WIN!
Try Next?
```
