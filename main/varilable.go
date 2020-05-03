package main

import (
    "fmt"
    "time"

    "strconv"
)

// 两个协程替代锁方案，变量高级用法: 大包包小包，小包可以用大包的变量

type Person struct {
    Name string

    salary float64

    chF chan func()
}

func NewPerson(name string, salary float64) *Person {

    p := &Person{name, salary, make(chan func())}

    go p.backend()

    return p

}

func (p *Person) backend() {

    for f := range p.chF {

        f()

    }

}

// 设置 salary.

func (p *Person) SetSalary(sal float64) {

    p.chF <- func() { p.salary = sal }

}

// 取回 salary.

func (p *Person) Salary() float64 {

    // fChan := make(chan float64)
    //	//	//
    //	//	// p.chF <- func() { fChan <- p.salary } // ?
    //	//	//
    //	//	// return <-fChan // ?

    //fChan := make(chan float64)
    var fChan float64
    p.chF <- func() {
        fChan = p.salary
    } // ?

    time.Sleep(time.Second * 1)
    return fChan // ?
}

func (p *Person) String() string {

    return "Person - name is: " + p.Name + " - salary is: " +
        strconv.FormatFloat(p.Salary(), 'f', 2, 64)

}

func main() {

    bs := NewPerson("Smith Bill", 2500.5)

    fmt.Println(bs)

    bs.SetSalary(4000.25)

    fmt.Println("Salary changed:")

    fmt.Println(bs)

}
