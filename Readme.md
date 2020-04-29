# Monte Carlo Exploration

Playground to explore Monte Carlo generators properties.

Inspired by _Fooled by Randomness_.

## Build
```sh 
$ go build .
```

## Coin toss

Simulate multiple rounds of coin tossing. Heads we get $1, tails we pay $1. We have $10 starting capital, but drawing
it down to $0 means _ruin_, and we can't continue the game.

```sh
$ ./monte-carlo-exploration
# Run simulation examples/cointoss.toml
## Simulating 1000000 executions of 100 round coin toss with starting capital of $10
Ruined: 31.931900% (319319 / 1000000)
Less capital: 48.775700% (487757 / 1000000)
More capital: 44.336700% (443367 / 1000000)
p01 $0
p05 $0
p10 $0
p50 $10
p90 $22
p95 $26
p99 $34
```

I wound't play that -- 30%+ chance of ruin. We end up with more capital only <45% of the time, and with less >48%.

Let's ignore ruin for now. $100 and 100 rounds. We can go to ruin only in the last round.

```sh
$ ./monte-carlo-exploration -conf examples/cointoss-no-ruin.toml
# Run simulation examples/cointoss-no-ruin.toml
## Simulating 1000000 executions of 100 round coin toss with starting capital of $100
Ruined: 0.000000% (0 / 1000000)
Less capital: 46.037200% (460372 / 1000000)
More capital: 46.019000% (460190 / 1000000)
p01 $76
p05 $84
p10 $88
p50 $100
p90 $112
p95 $116
p99 $124
```

Good sport, seems a fair game! Only 1% of the time we lose more than $24, and only 1% we earn more than $24.


## More ideas to explore

- Pi calculation.
- Stock price behavior. (what's the distribution? how to get it's params?)
- Options pricing.