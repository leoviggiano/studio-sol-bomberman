```
0 -> Create Grid
1 -> Pass
2 -> Add Bombs
3 -> Explode 1st bombs
4 -> Add Bombs
5 -> Explode 3rd Bombs
6 -> Add Bombs
7 -> Explode 5rd Bombs
8 -> Add Bombs
9 -> Explode 7rd Bombs
```

First pattern of explosions start at 3


* 3 / 2 = 1.5 % 2 = 1 -> First Explosion/Pattern
* 5 / 2 = 2.5 % 2 = 0 -> Second Explosion/Pattern

### What means 
* if (Seconds / 2) % 2 != 0 == First Explosion
* or (Seconds / 2) % 2 == 0 == Second Explosion
