package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	done := time.After(10 * time.Second)
	echo := make(chan []byte)

	go readStdin(echo)
	for {
		select {
		case buf := <-echo:
			os.Stdout.Write(buf)
		case <-done:
			fmt.Println("Timed out!")
			os.Exit(0)

		}
	}
	/*
		var wg sync.WaitGroup

		w := newWords()
		for _, f := range os.Args[1:] {
			wg.Add(1)
			go func(file string) {
				if err := tallyWords(file, w); err != nil {
					fmt.Println(err.Error())
				}
				wg.Done()
			}(f)
		}
		wg.Wait()

		fmt.Println("Words that appear more than once:")
		w.Lock()

		for word, count := range w.found {
			if count > 1 {
				fmt.Printf("%s: %d\n", word, count)
			}
		}
		w.Unlock()
	*/
}

func readStdin(out chan<- []byte) {
	for {
		data := make([]byte, 1024)
		l, _ := os.Stdin.Read(data)
		if l > 0 {
			out <- data
		}
	}
}

/*
type words struct {
	sync.Mutex
	found map[string]int
}

func newWords() *words {
	return &words{found: map[string]int{}}
}

func (w *words) add(word string, n int) {
	w.Lock()
	defer w.Unlock()
	count, ok := w.found[word]
	if !ok {
		w.found[word] = n
		return
	}
	w.found[word] = count + n
}

func tallyWords(filename string, dict *words) error {
	file, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanWords)
	for scanner.Scan() {
		word := strings.ToLower(scanner.Text())
		dict.add(word, 1)
	}
	return scanner.Err()
}
*/
