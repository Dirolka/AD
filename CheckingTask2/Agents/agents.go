package Agents

import "fmt"

type Agent interface {
	GetID() string
	GetName() string
}

type HumanAgent struct {
	ID   string
	Name string
}

func (a HumanAgent) GetID() string {
	return a.ID
}

func (a HumanAgent) GetName() string {
	return a.Name
}

type BotAgent struct {
	ID      string
	Name    string
	Version string
}

func (a BotAgent) GetID() string {
	return a.ID
}

func (a BotAgent) GetName() string {
	return a.Name
}

func FormatAgent(a Agent) string {
	switch v := a.(type) {
	case HumanAgent:
		return fmt.Sprintf("%s | %s", v.ID, v.Name)
	case *HumanAgent:
		return fmt.Sprintf("%s | %s", v.ID, v.Name)
	case BotAgent:
		return fmt.Sprintf("%s | %s | bot:%s", v.ID, v.Name, v.Version)
	case *BotAgent:
		return fmt.Sprintf("%s | %s | bot:%s", v.ID, v.Name, v.Version)
	default:
		return ""
	}
}
