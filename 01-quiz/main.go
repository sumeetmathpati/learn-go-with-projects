package main

type Question struct {
	Data          string
	Score         int
	OptionA       string
	OptionB       string
	OptionC       string
	OptionD       string
	CorrectOption string
}

func main() {

	// Question DB
	questions := []Question{
		Question{"Which Linux distro uses names from movie Toy story for it's releases?", 2, "Ubuntu", "Debian", "Fedora", "Arch", "b"},
		Question{"Where were The Lord of the Rings movies filmed?", 2, "Ireland", "Iceland", "New Zealand", "Australia", "c"},
		Question{"Which country does Forrest Gump travel to as part of the All-American Ping-Pong Team?", 2, "Vietnam", "China", "Sweden", "France", "b"},
		Question{"What are the words the orcs of Mordor interrogated out of Gollum?", 1, "Shire, Baggins", "Gandalf, Hobbiton", " Ring, Baggins", " Shire, Frodo", "a"},
		Question{"What's the name of the dog guarding the Sorcerer’s Stone?", 2, "Padfoot", "Grim", "Fluffy", "Fang", "c"},
	}

}
