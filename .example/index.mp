:META name  LOTTERY
:META input float

#LOTTERY
VAR (32) $SOME_VALUE

CALL #RAND $RAND_RESULT

IF ($RAND_RESULT > 0.5) (#MORE, #LESS)

PRINT ("Try Next?")

#MORE
PRINT ("you're WIN!")

#LESS
VAR (sum("you're lose ", $SOME_VALUE, " points!")) $OUT_RESULT

PRINT ($OUT_RESULT)

#RAND
RETURN (rand())