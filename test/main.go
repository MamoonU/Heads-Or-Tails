package main

import (
	"fmt"
	"math/rand"
	"net/http"
)

type Coin struct {
	name  string
	heads string
	tails string
}

var coins = []Coin{
	{
		name:  "UK Pound Coin",
		heads: "https://www.royalmint.com/globalassets/the-royal-mint/images/pages/new-pound-coin/large_new_pound.png",
		tails: "https://www.royalmint.com/globalassets/the-royal-mint/images/pages/new-pound-coin/large_new_pound_rev.png",
	},
	{
		name:  "UK 50p Coin",
		heads: "https://www.royalmint.com/globalassets/consumer/_campaigns/2020/dinosaurs/product-photography/the_dinosauria_collection_-_megalosaurus_2020_uk_50p_brilliant_uncirculated_coin_obverse_-_uk20mgbu.jpg",
		tails: "https://i.pinimg.com/originals/ea/72/e3/ea72e37a89868d18ac2176586f0d5722.jpg",
	},
	{
		name:  "US Cent",
		heads: "https://onlinecoin.club/images/coins/United_States/ec377acb-a1a2-4fa5-a637-e1e8c007f7aa.jpg",
		tails: "https://onlinecoin.club/images/coins/United_States/c2462a57-bc3a-4197-91de-7eba4f5fa366.jpg",
	},
	{
		name:  "SA Riyal",
		heads: "https://coin-brothers.com/photos/Saudi_Arabia_Riyal_1/2016_04.11.2020_19.35.jpg",
		tails: "https://coin-brothers.com/photos/Saudi_Arabia_Riyal_1/2016_04.11.2020_19.35_01.jpg",
	},
	{
		name:  "JP Yen",
		heads: "https://s3.amazonaws.com/ngccoin-production/world-coin-price-guide/1442695-4579461-009rf.jpg",
		tails: "https://s3.amazonaws.com/ngccoin-production/world-coin-price-guide/1442695-4579461-009b.jpg",
	},
}

func main() {
	http.HandleFunc("/", fortnite)

	fmt.Println(http.ListenAndServe(":6969", nil))
}

func headsortails(w http.ResponseWriter, req *http.Request) {
	/*
		fmt.Println(req.URL.Query().Get("coin"))

		http.SetCookie(w, &http.Cookie{
			Name:  "coin",
			Value: req.URL.Query().Get("coin"),
		})
	*/
	page := `
	
	<head>
	</head>
	<body>
		<form action="/"	method="get">
		%s
		<form>
	</body>
	`
	buttons := ""
	for _, coin := range coins {
		buttons += fmt.Sprintf(`<button type="submit" name="coin" value="%s"> <img src="%s" style="max-width: 300px;" > </button>`, coin.name, coin.heads)
	}

	page = fmt.Sprintf(page, buttons)
	fmt.Fprintln(w, page)
}

func fortnite(w http.ResponseWriter, req *http.Request) {

	if req.URL.Query().Get("coin") == "" {
		headsortails(w, req)
		return
	}

	coin := Coin{}
	for _, c := range coins {
		if c.name == req.URL.Query().Get("coin") {
			coin = c
		}

	}
	/*
		coinName := ""

		coinCookie, err := req.Cookie("coin")
		if err != nil {
			if err == http.ErrNoCookie {
				coinName = coins[0].name
			} else {
				http.Error(w, "Error getting Cookie", http.StatusInternalServerError)
			}

		}
		coinName = coinCookie.Value

		coin := Coin{}
		for _, c := range coins {
			if c.name == coinName {
				coin = c
			}

		}
	*/
	page := `
	
	<head>
	</head>
	<body>
		<img src="%s" style="max-width: 300px;" ><br>
		<button onclick="location.href='/'" type="button">back</button>
		
	</body>
	
	
	`
	buhh := rand.Intn(2)
	if buhh == 0 {
		fmt.Fprintf(w, page, coin.heads)
	} else {
		fmt.Fprintf(w, page, coin.tails)
	}

}

//https://www.royalmint.com/globalassets/the-royal-mint/images/pages/new-pound-coin/large_new_pound.png
//https://www.royalmint.com/globalassets/the-royal-mint/images/pages/new-pound-coin/large_new_pound_rev.png
