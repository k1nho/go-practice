package main

var assertReady bool

func main() {
	go prepareQuery()

}

func prepareQuery() {
	sleep()
	assertReady = true
}

func sleep() {
	rand.Seed(time.Now().UnixNano())
	waitTime := time.Duration(1+rand.Intn(5)) * time.Second
	time.Sleep(someTime)
}
