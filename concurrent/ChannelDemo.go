package main

type Task struct {
}

func main() {
	//带缓冲的channel
	ch := make(chan Task, 3)

	//启动固定数量的worker
	for i := 0; i < 10; i++ {
		go worker(ch)
	}

	//发送任务给worker
	hellaTasks := getTaks()

	for _, task := range hellaTasks {
		ch <- task
	}
}

func getTaks() []Task {
	return nil
}

func worker(ch chan Task) {
	for {
		//接受任务
		task := <-ch
		process(task)
	}
}

func process(task Task) {

}
