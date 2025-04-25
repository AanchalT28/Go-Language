package main

import (
	"fmt"
	"sync"
)

type ElectionMsg struct {
	Initiator int
}

type CoordinatorMsg struct {
	Leader int
}

type Process struct {
	ID       int
	Alive    bool
	Channels []chan ElectionMsg
	Coord    chan CoordinatorMsg
	Done     chan bool
}

func (p *Process) Start(wg *sync.WaitGroup) {
	defer wg.Done()

	for {
		select {
		case msg := <-p.Channels[p.ID]:
			if !p.Alive {
				continue
			}

			higherExists := false
			fmt.Printf("Process %d received Election(%d)\n", p.ID, msg.Initiator)

			for i := p.ID + 1; i < len(p.Channels); i++ {
				if processes[i] != nil && processes[i].Alive {
					fmt.Printf("Process %d sends Election(%d) to Process %d\n", p.ID, p.ID, i)
					p.Channels[i] <- ElectionMsg{Initiator: p.ID}
					higherExists = true
				}
			}

			if !higherExists {
				fmt.Printf("\nProcess %d declares itself as the NEW LEADER\n\n", p.ID)
				for i := 1; i < len(p.Channels); i++ {
					if i != p.ID && processes[i].Alive {
						processes[i].Coord <- CoordinatorMsg{Leader: p.ID}
					}
				}
				// Shut down all
				for _, proc := range processes {
					if proc != nil {
						close(proc.Done)
					}
				}
				return
			}

		case msg := <-p.Coord:
			if p.Alive {
				fmt.Printf("Process %d receives Coordinator(%d)\n", p.ID, msg.Leader)
			}
		case <-p.Done:
			return
		}
	}
}

var processes []*Process

func main() {
	var total, failedLeader, initiator, failedExtra int
	fmt.Print("Enter number of processes: ")
	fmt.Scanln(&total)

	fmt.Print("Enter the leader process that fails (e.g., the highest): ")
	fmt.Scanln(&failedLeader)

	fmt.Print("Enter the process that initiates the election: ")
	fmt.Scanln(&initiator)

	fmt.Print("Enter any additional failed process (or 0 for none): ")
	fmt.Scanln(&failedExtra)

	processes = make([]*Process, total+1) // 1-based
	channels := make([]chan ElectionMsg, total+1)

	for i := 1; i <= total; i++ {
		channels[i] = make(chan ElectionMsg, total)
	}

	for i := 1; i <= total; i++ {
		processes[i] = &Process{
			ID:       i,
			Alive:    true,
			Channels: channels,
			Coord:    make(chan CoordinatorMsg),
			Done:     make(chan bool),
		}
	}

	// Fail leader and optional extra process
	processes[failedLeader].Alive = false
	if failedExtra > 0 && failedExtra <= total {
		processes[failedExtra].Alive = false
	}

	fmt.Printf("\nThe process that failed is: %d\n", failedLeader)
	if failedExtra > 0 {
		fmt.Printf("Additional failed process: %d\n", failedExtra)
	}
	fmt.Printf("Process %d starts the election.\n\n", initiator)

	var wg sync.WaitGroup
	for i := 1; i <= total; i++ {
		wg.Add(1)
		go processes[i].Start(&wg)
	}

	// Start the election
	channels[initiator] <- ElectionMsg{Initiator: initiator}

	wg.Wait()
}
