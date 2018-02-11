// PCBMngmnt
package main

import (
	. "fmt"
)

var temp string //用于接收输入流换行符的容器

type PCB struct {
	name      string //进程名
	state     string //进程当前状态
	priority  int    //进程优先级
	timeNd    int    //进程运行所需时间
	timeToRun int    //进程剩余运行时间
	ifPriorer bool   //用于Queue函数，比较两个进程之间是否更优先
	num       int    //进程进入的次序
} //声明结构体用于表示进程

//键入进程块的函数
func (p *PCB) SetPCB() {
	Scanln(&temp)
	_ = temp
	Printf("\n请输入进程名：")
	Scanf("%s", &p.name)
	Scanln(&temp)
	_ = temp
	Printf("请输入优先级数：")
	Scanf("%d", &p.priority)
	Scanln(&temp)
	_ = temp
	Printf("请输入运行程序需要的时间：")
	Scanf("%d", &p.timeNd)
	Scanln(&temp)
	_ = temp
	p.timeToRun = p.timeNd
	p.state = "Waiting"
	p.ifPriorer = false
}

//显示单个进程信息的函数
func (p *PCB) ShowProcess() {
	Printf("%s\t%s\t%d\t%d\t%d\t\n", p.name, p.state, p.priority, p.timeNd, p.timeToRun)
}

//比较两个进程先后的函数
func Queue(pA *PCB, pB *PCB) {
	if pA.priority > pB.priority {
		pA.ifPriorer = true
		pB.ifPriorer = false
	} else if pA.priority < pB.priority {
		pA.ifPriorer = false
		pB.ifPriorer = true
	} else {
		if pA.num < pB.num {
			pA.ifPriorer = true
			pB.ifPriorer = false
		} else {
			pA.ifPriorer = false
			pB.ifPriorer = true
		}
	}
}

func main() {
	var amount int
	Printf("请输入进程数量：")
	Scanf("%d", &amount)
	Scanln(&temp)
	_ = temp

	var pS [100]PCB //创建一个结构体数组来储存所有进程块

	for i := 0; i < amount; i++ {
		Printf("\n这是第%d个进入的进程。\n*******键入回车继续*******", i+1)
		pS[i].SetPCB()
		pS[i].num = i
	}

	Println("所有进程如下：")
	Println("进程名 \t 状态 \t 优先级 \t 需要时间 \t 剩余时间 ")
	for i := 0; i < amount; i++ {
		pS[i].ShowProcess()
	} //展示一次全部进程的最初状态

	var tempStruct PCB  //临时结构体变量用于排序
	for k := 0; ; k++ { //无限循环以在不定的、适合的时间跳出循环
		for i := 0; i < amount-1; i++ {
			for j := 0; j < amount-1-i; j++ {
				Queue(&pS[i], &pS[i+1])
				if pS[i].ifPriorer == true && pS[i+1].ifPriorer == false {
					continue
				} else if pS[i].ifPriorer == false && pS[i+1].ifPriorer == true {
					tempStruct = pS[i]
					pS[i] = pS[i+1]
					pS[i+1] = tempStruct
				}
			}
		} //排序进程

		for i := 0; i < amount; i++ {
			if pS[i].timeToRun == 0 {
				pS[i].state = "Finish"
				pS[i].priority = -1
			} else {
				pS[i].state = "Waiting"
			}
		}
		if pS[0].state == "Waiting" {
			pS[0].state = "Running"
		}
		Println("\n运行结果如下：（优先级若为-1则说明进程已结束）")
		Println("进程名 \t 状态 \t 优先级 \t 需要时间 \t 剩余时间 ")
		for i := 0; i < amount; i++ {
			pS[i].ShowProcess()
		}

		Println("*******按回车键继续*******")
		Scanln(&temp)
		_ = temp

		count := 0
		for i := 0; i < amount; i++ {
			if pS[i].timeToRun == 0 {
				count++
			}
		}
		if count == amount {
			goto finish
		} //判定全部进程已完成并跳出循环

		for i := 0; i < amount; i++ {
			if pS[i].state != "Finish" {
				pS[i].state = "Waiting"
			}
		} //重置状态

		if pS[0].timeToRun > 0 {
			pS[0].timeToRun--
		} //调整剩余时间

		if pS[0].priority > 0 {
			pS[0].priority--
		} //调整优先级

	}

finish:
}
