package main

func init() {
	ottomap = map[string]interface{}{"date": Date,
		"otto help":    Help,
		"otto random":  Random,
		"otto say":     Say,
		"otto roll":    Roll,
		"otto mock":    Mock,
		"otto flip":    Flip,
		"otto magic":   Magic,
		"otto will":    Magic,
		"otto tod":     ToD,
		"otto ToD":     ToD,
		"otto Tod":     ToD,
		"otto weather": Weather,
		"otto calc":    Calc,
		"egg":          Egg, //eightball easter egg
		"otto hello":   "hello there!",
		"otto version": "I am Version 1.3.0",
		"otto what":    "I am a imessage virtual assistant that runs when Peter's computer is on. Type 'otto help' to see all the commands I can do.",
		"hi otto ":     "hi there!",
		"otto time":    Time,
		"otto thanks":  "you're welcome",
		"otto google":  Google,   //gets first span
		"otto wiki":    Wiki,     //link
		"otto info":    Wikitext, //intro paragraph
	}
}

type WeatherSettings struct {
	Default string `json:"default"`
	Apikey  string `json:"apikey"`
}

type EightballSettings struct {
	Phrases    []string            `json:"phrases"`
	Eastereggs map[string][]string `json:"eastereggs"`
}
type ChatSettings struct {
	Lastperson     string `json:"lastperson"`
	Lastamount     int    `json:"lastamount"`
	Lasttext       string `json:"lasttext"`
	Lasttextperson string `json:"lasttextperson"`
}
type Results struct {
	Weather      WeatherSettings     `json:"weather"`
	Chat         ChatSettings        `json:"chat"`
	Errormessage string              `json:"errormessage"`
	Maxmessage   string              `json:"maxmessage"`
	Eightball    EightballSettings   `json:"eightball"`
	TruthOrDare  map[string][]string `json:"truthordare"`
}
