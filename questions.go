package main

func makeOptions(opts [][2]interface{}) []Option {
	letters := []string{"A", "B", "C", "D"}

	out := make([]Option, len(opts))

	for i, o := range opts {
		out[i] = Option{
			Idx:    i,
			Letter: letters[i],
			Text:   o[0].(string),
			Score:  o[1].(int),
		}
	}

	return out
}

var questions = []Question{

	{
		ID: 1,
		EN: "You witness someone drop their groceries in a busy street. What do you do?",
		TL: "Nakita mo ang isang tao na nahulog ang kanilang mga groceries sa maingay na kalye. Ano ang gagawin mo?",
		Options: makeOptions([][2]interface{}{
			{"Stop and help them pick everything up", 10},
			{"Slow down and offer a hand if they seem to struggle", 7},
			{"Glance over but keep walking — others will help", 3},
			{"Keep walking, you're in a hurry", 0},
		}),
	},

	{
		ID: 2,
		EN: "A friend confides they're going through a tough time. You have plans later. What do you do?",
		TL: "Ang isang kaibigan ay nagkumpisal na sila ay nasa mahirap na panahon. Mayroon kang mga plano mamaya. Ano ang gagawin mo?",
		Options: makeOptions([][2]interface{}{
			{"Cancel your plans and stay with them", 10},
			{"Stay for a while then explain you have to go", 7},
			{"Listen briefly, then remind them you have plans", 3},
			{"Tell them you'll talk later and leave", 0},
		}),
	},

	{
		ID: 3,
		EN: "You find a wallet on the ground with cash and an ID inside. What do you do?",
		TL: "Nakakita ka ng pitaka sa lupa na may pera at ID sa loob. Ano ang gagawin mo?",
		Options: makeOptions([][2]interface{}{
			{"Turn it in to the police or find the owner", 10},
			{"Post about it online", 5},
			{"Keep cash but return wallet", 3},
			{"Keep everything", 0},
		}),
	},

	{
		ID: 4,
		EN: "A classmate is being left out of a group activity. You barely know them.",
		TL: "Ang isang kaklase ay naiiwan sa group activity. Halos hindi mo sila kilala.",
		Options: makeOptions([][2]interface{}{
			{"Invite them to join", 10},
			{"Tell someone else", 7},
			{"Ignore it", 3},
			{"Do nothing", 0},
		}),
	},

	{
		ID: 5,
		EN: "Someone online is being harassed in a comment section.",
		TL: "May taong inaapi sa comment section online.",
		Options: makeOptions([][2]interface{}{
			{"Defend them and report", 10},
			{"Send private support", 7},
			{"Report only", 5},
			{"Ignore it", 0},
		}),
	},

	{
		ID: 6,
		EN: "A homeless person is staring at your food.",
		TL: "May taong walang tirahan na nakatingin sa pagkain mo.",
		Options: makeOptions([][2]interface{}{
			{"Offer food", 10},
			{"Give spare food", 7},
			{"Ignore guilt", 3},
			{"Move away", 0},
		}),
	},

	{
		ID: 7,
		EN: "A child breaks something valuable accidentally.",
		TL: "Isang bata ang nakabasag ng mahalagang bagay.",
		Options: makeOptions([][2]interface{}{
			{"Stay calm and help", 10},
			{"Talk calmly", 7},
			{"Scold but help", 3},
			{"Lose temper", 0},
		}),
	},

	{
		ID: 8,
		EN: "You hear a false rumor about someone you know.",
		TL: "Narinig mo ang maling tsismis tungkol sa kakilala mo.",
		Options: makeOptions([][2]interface{}{
			{"Correct it immediately", 10},
			{"Tell the person", 7},
			{"Stay silent", 3},
			{"Join in", 0},
		}),
	},

	{
		ID: 9,
		EN: "A stray injured animal is on the road.",
		TL: "May sugatang hayop sa daan.",
		Options: makeOptions([][2]interface{}{
			{"Help and contact rescue", 10},
			{"Post online", 6},
			{"Assume others help", 3},
			{"Ignore", 0},
		}),
	},

	{
		ID: 10,
		EN: "You got extra change from a store by mistake.",
		TL: "Nakakuha ka ng sobrang sukli.",
		Options: makeOptions([][2]interface{}{
			{"Return it", 10},
			{"Return if big amount", 6},
			{"Keep it", 3},
			{"Keep it fully", 0},
		}),
	},
}
