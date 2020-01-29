package main

import "fmt"
import "sort"
import "math"

func coinChangeH (coins []int, amount int, seen map[int] int) int {
  if amount < 0 {
    return -1
  }

  if amount == 0 {
    return 0
  }

  if count, saw := seen[amount]; saw {
    return count
  }

  topBoundIndex := sort.Search(len(coins), func (i int) bool { return coins[i] >= amount
  })

  if len(coins) == topBoundIndex {
    topBoundIndex--
  }

  ret := -1

  for i := topBoundIndex; i >= 0; i-- {
    tmp := coinChangeH(coins, amount-coins[i], seen)

    if tmp != -1 {
      if ret == -1 {
        ret = tmp+1
      } else {
        ret = int(math.Min(float64(ret), float64(tmp+1)))
      }
    }
  }

  seen[amount] = ret

  return ret
}

func coinChange(coins []int, amount int) int {

  coinsCpy := make([]int, len(coins))
  copy(coinsCpy, coins)

  sort.Slice(coinsCpy, func (i, j int) bool {
    return coinsCpy[i] < coinsCpy[j]
  })

  return coinChangeH(coinsCpy, amount, make(map[int]int))
}

func main() {
  fmt.Printf("coins needed %d", coinChange([]int{186, 419, 83, 408}, 6249))
}