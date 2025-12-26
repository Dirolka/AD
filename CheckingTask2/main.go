package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/Dirolka/CheckingTask2/Agents"
	"github.com/Dirolka/CheckingTask2/Tickets"
)

func main() {
	agents := make(map[string]Agents.Agent)
	store := Tickets.NewTicketStore()

	in := bufio.NewScanner(os.Stdin)

	for {
		printMenu()
		choice := readLine(in, "Choose: ")

		switch choice {
		case "1":
			createTicket(in, store)
		case "2":
			addAgent(in, agents)
		case "3":
			assignTicket(in, store, agents)
		case "4":
			resolveTicket(in, store)
		case "5":
			printTickets(store.ListAll())
		case "6":
			printTickets(store.ListByStatus("OPEN"))
		case "7":
			printTickets(store.ListByStatus("DONE"))
		case "8":
			printTickets(store.ListUnassigned())
		case "9":
			return
		default:
			fmt.Println("Invalid option")
		}

		fmt.Println()
	}
}

func printMenu() {
	fmt.Println("1. Create Ticket")
	fmt.Println("2. Add Agent (Human or Bot)")
	fmt.Println("3. Assign Ticket to Agent")
	fmt.Println("4. Resolve Ticket")
	fmt.Println("5. List All Tickets")
	fmt.Println("6. List OPEN Tickets")
	fmt.Println("7. List DONE Tickets")
	fmt.Println("8. List Unassigned Tickets")
	fmt.Println("9. Exit")
}

func readLine(in *bufio.Scanner, prompt string) string {
	fmt.Print(prompt)
	if !in.Scan() {
		return ""
	}
	return strings.TrimSpace(in.Text())
}

func readInt(in *bufio.Scanner, prompt string) (int, error) {
	s := readLine(in, prompt)
	return strconv.Atoi(s)
}

func createTicket(in *bufio.Scanner, store *Tickets.TicketStore) {
	id := readLine(in, "Ticket ID: ")
	title := readLine(in, "Title: ")
	desc := readLine(in, "Description: ")
	priority, err := readInt(in, "Priority (1..3): ")
	if err != nil {
		fmt.Println("Invalid priority")
		return
	}

	t := Tickets.Ticket{
		ID:          id,
		Title:       title,
		Description: desc,
		Priority:    priority,
		AssigneeID:  "",
		Status:      "OPEN",
	}

	if err := store.Create(t); err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("Ticket created")
}

func addAgent(in *bufio.Scanner, agents map[string]Agents.Agent) {
	typeChoice := strings.ToUpper(readLine(in, "Agent type (H/B): "))
	id := readLine(in, "Agent ID: ")
	name := readLine(in, "Name: ")

	if id == "" {
		fmt.Println("Error: agent id must be non-empty")
		return
	}
	if name == "" {
		fmt.Println("Error: agent name must be non-empty")
		return
	}
	if _, exists := agents[id]; exists {
		fmt.Println("Error: agent id already exists")
		return
	}

	switch typeChoice {
	case "H":
		a := Agents.HumanAgent{ID: id, Name: name}
		agents[id] = a
		fmt.Println("Added:", Agents.FormatAgent(a))
	case "B":
		version := readLine(in, "Bot Version: ")
		if version == "" {
			fmt.Println("Error: bot version must be non-empty")
			return
		}
		a := Agents.BotAgent{ID: id, Name: name, Version: version}
		agents[id] = a
		fmt.Println("Added:", Agents.FormatAgent(a))
	default:
		fmt.Println("Invalid agent type")
	}
}

func assignTicket(in *bufio.Scanner, store *Tickets.TicketStore, agents map[string]Agents.Agent) {
	ticketID := readLine(in, "Ticket ID: ")
	agentID := readLine(in, "Agent ID: ")

	if _, ok := agents[agentID]; !ok {
		fmt.Println("Error: agent does not exist")
		return
	}

	if err := store.Assign(ticketID, agentID); err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("Ticket assigned")
}

func resolveTicket(in *bufio.Scanner, store *Tickets.TicketStore) {
	ticketID := readLine(in, "Ticket ID: ")
	if err := store.Resolve(ticketID); err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("Ticket resolved")
}

func printTickets(tickets []Tickets.Ticket) {
	if len(tickets) == 0 {
		fmt.Println("No tickets")
		return
	}

	for _, t := range tickets {
		assignee := t.AssigneeID
		if assignee == "" {
			assignee = "(unassigned)"
		}
		fmt.Printf("%s | %s | priority:%d | %s | %s\n", t.ID, t.Title, t.Priority, assignee, t.Status)
	}
}
