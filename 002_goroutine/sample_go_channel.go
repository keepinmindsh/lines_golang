package main

type Req struct {
	Num  int
	Resp chan Resp
}

type Resp struct {
	Num      int
	WorkerID int
}

func GoChannel() {

}
