# golang-expression-evaluator

## Program:

```bash
#MAIN
VAR $SOME_VALUE (32)

IF (rand() > 0.5) #MORE #LESS

OUT (" Try Next?")

#MORE
OUT ("you're WIN!")

#LESS
VAR $OUT_RESULT (sum("you're lose ", $SOME_VALUE, " points!"))

OUT ($OUT_RESULT)
```

## AST:

```bash
root
└── #MAIN
    └── VAR
        └── $SOME_VALUE
        └── 32
    └── IF
        └── >
            └── rand
            └── 0.5
        └── #MORE
        └── #LESS
    └── OUT
        └── " Try Next?"
└── #MORE
    └── OUT
        └── "you're WIN!"
└── #LESS
    └── VAR
        └── $OUT_RESULT
        └── sum
            └── "you're lose "
            └── $SOME_VALUE
            └── " points!"
    └── OUT
        └── $OUT_RESULT
```

## IL:

```bash
MARK #MAIN
PUSH 32
VAR $SOME_VALUE
CALL rand 0
PUSH 0.5
CALL > 2
IF #MORE #LESS
PUSH " Try Next?"
CALL OUT 1
MARK #MORE
PUSH "you're WIN!"
CALL OUT 1
MARK #LESS
PUSH "you're lose "
PUSH $SOME_VALUE
PUSH " points!"
CALL sum 3
VAR $OUT_RESULT
PUSH $OUT_RESULT
CALL OUT 1
```

## Output:

```bash
you're lose 32 points! Try Next?
```