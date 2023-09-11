package driver


type Actions struct {
	None   string
	Right  string
	Left   string
	Pickup string
	Jump   string
	Brake  string
}

var actions = Actions{
	None:   "none",
	Right:  "right",
	Left:   "left",
	Pickup: "pickup",
	Jump:   "jump",
	Brake:  "brake",
}

type Obstacles struct {
	None    string
	Crack   string
	Trash   string
	Penguin string
	Bike    string
	Water   string
	Barrier string
}

var obstacles = Obstacles{
	None:    "",
	Crack:   "crack",
	Trash:   "trash",
	Penguin: "penguin",
	Bike:    "bike",
	Water:   "water",
	Barrier: "barrier",
}
