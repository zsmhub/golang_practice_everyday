package main

import (
    "errors"
    "fmt"
    "math/rand"
    "sort"
    "time"
)

func main() {
    amount := 100000
    count := 50

    r, _ := DivideRedPack(amount, count, false)
    fmt.Println(r)

    total := 0
    for k := 0; k < count; k++ {
        total += r[k]
    }
    fmt.Printf("总金额：%d，金额差：%d\n", total, amount-total)
}

func init() {
    rand.Seed(time.Now().UnixNano())
}

/**
 * 分割红包金额算法，返回按数量分割好的红包slice
 * @param int amount 红包金额，单位分
 * @param int count 红包数量
 * @param bool equally 是否均分
 * @return []int, error
 */
func DivideRedPack(amount, count int, equally bool) ([]int, error) {
    if amount == 0 || count == 0 {
        return []int{}, errors.New("红包金额或数量不能为0")
    }

    if amount < count {
        return []int{}, errors.New("单个红包金额不能小于 1 分钱")
    }

    // 获取红包平均值，均分时为平均数，随机时为 1 分钱
    var avg int
    if equally {
        avg = amount / count // tip: 整数除整数，结果还是整数，舍去法取整
    } else {
        avg = 1
    }

    if avg == 0 {
        avg = 1
    }

    // 红包 slice 初始化
    ret := make([]int, count)
    for k := range ret {
        ret[k] = avg
    }

    // 剩余的钱
    leftAmount := amount - count*avg

    if leftAmount == 0 {
        return ret, nil
    }

    // 剩余的钱，按照不同的方式分
    if equally {
        // 均分
        for i := 0; i < leftAmount; i++ {
            ret[i] += 1
        }
    } else {
        // 随机分配
        var luck []int
        var leftCount = count
        if leftAmount < count {
            leftCount = leftAmount
        }

        luck = make([]int, leftCount)

        // 方案一：随机范围不断缩小，不可用，因为红包数越多，越容易出现一堆 1 分钱的情况
        /*for k := range luck {
            // 最后一个红包直接承包最后剩余的金额，确保总金额正确
            if k == leftCount - 1 {
                luck[k] = leftAmount
                break
            }

            // 这里随机数最大可取到 leftAmount，则可能出现很多人只有 1 分钱的情况
            luck[k] = rand.Intn(leftAmount + 1)
            leftAmount -= luck[k]

            if leftAmount == 0 {
                break
            }
        }*/

        // 方案二：线段法，1. 在 0 ~ leftAmount 的线段中，插入 n 个节点；2. 根据节点，将线段裁成 n 个小线段
        /*for k := range luck {
              luck[k] = rand.Intn(leftAmount + 1)
          }
          sort.Ints(luck)
          luck[leftCount-1] = leftAmount
          for k := leftCount - 1; k > 0; k-- {
              luck[k] = luck[k] - luck[k-1]
          }*/

        // 方案三：按权重分配法
        randAmountSum := 0                          // 随机金额总和
        randAmountMax := leftAmount / leftCount * 2 // 最大可领金额 = 剩余金额的平均值x2 = (剩余金额 / 剩余数量) * 2
        for k := range luck {
            luck[k] = rand.Intn(randAmountMax)
            randAmountSum += luck[k]
        }

        weightAmountSum := 0 // 按权重分配的金额总和
        for k := range luck {
            luck[k] = int(float32(luck[k]) / float32(randAmountSum) * float32(leftAmount))
            weightAmountSum += luck[k]
        }

        // 剩余金额继续分配
        if weightAmountSum < leftAmount {
            randAmountSum := leftAmount - weightAmountSum

            randAmountMax = randAmountSum / leftCount * 2
            if randAmountMax == 0 {
                randAmountMax = randAmountSum
            }

            var temp int
            for k := range luck {
                if k == (leftCount - 1) {
                    temp = randAmountSum
                } else {
                    temp = rand.Intn(randAmountMax)
                }

                luck[k] += temp
                randAmountSum -= temp
                if randAmountSum == 0 {
                    break
                }
            }
        }

        // 再次调整分配金额，使得更平均，否则部分人拿到很高金额，部分人拿到很低金额，有些人心境会崩溃，导致影响上班心情的哈
        sort.Ints(luck)
        for i := 0; i < leftCount>>2; i++ {
            delta := luck[leftCount-1-i]>>1 + luck[leftCount-1-i]>>4
            if delta > 0 && delta < luck[leftCount-1-i] {
                luck[i] += delta
                luck[leftCount-1-i] -= delta
            }
        }

        // 叠加到初始分配上
        for k := range luck {
            ret[k] += luck[k]
        }
    }

    if equally == false || leftAmount != count {
        // 打乱顺序
        rand.Shuffle(len(ret), func(i, j int) {
            ret[i], ret[j] = ret[j], ret[i]
        })
    }

    return ret, nil
}
