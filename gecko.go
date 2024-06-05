package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"time"
)

// ======================================Flynn-pack-scraping-GO===============//

const (
	pt = "\033[0;37m"
	lb = "\033[1;36m"
	r  = "\033[1;31m"
	gr = "\033[1;32m"
	y  = "\033[1;33m"
	mg = "\033[35m"
)

func own(url string, ua []string) string {
	for {
		req, err := http.NewRequest("GET", url, nil)
		if err != nil {
			fmt.Println("\x1b[1;33mCheck Your Connection!\n")
			time.Sleep(2 * time.Second)
			continue
		}

		for _, header := range ua {
			parts := strings.SplitN(header, ":", 2)
			req.Header.Set(strings.TrimSpace(parts[0]), strings.TrimSpace(parts[1]))
		}

		resp, err := http.DefaultClient.Do(req)
		if err != nil {
			fmt.Println("\x1b[1;33mCheck Your Connection!\n")
			time.Sleep(2 * time.Second)
			continue
		}
		defer resp.Body.Close()

		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			fmt.Println("\x1b[1;33mCheck Your Connection!\n")
			time.Sleep(2 * time.Second)
			continue
		}

		return string(body)
	}
}

func x(x1, x2, xdata string) string {
	x1Index := strings.Index(xdata, x1)
	if x1Index == -1 {
		return ""
	}
	x1Index += len(x1)
	x2Index := strings.Index(xdata[x1Index:], x2)
	if x2Index == -1 {
		return ""
	}
	return xdata[x1Index : x1Index+x2Index]
}

func timer(clk int) {
	ti := time.Now().Add(time.Duration(clk) * time.Second).Unix()
	for {
		fmt.Print("\r                        \r")
		res := ti - time.Now().Unix()
		if res < 1 {
			break
		}
		fmt.Print(time.Unix(res, 0).Format("15:04:05"))
		time.Sleep(1 * time.Second)
	}
}

// ======================================Flynn-pack-scraping-GO===============//

func main() {
	for {
		//fmt.Print("\033[H\033[2J") // Clear the screen
		defaultCurrency := "usd"
		/*
			change the compare currency "defaultCurrency" with yours!
			available "USD,EUR,RUB,BRL,GBP,INR,AUD,CAD,PLN,TRY
			BTC,ETH"
			nb : {only use "lowercase" usd not USD} // conflict
		*/
		isin := "backend/data/bubbles1000." + defaultCurrency + ".json"
		dmn := "https://cryptobubbles.net/" + isin

		flynn := own(dmn, []string{
			"Host: cryptobubbles.net",
			"User-Agent: Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/122.0.0.0 Safari/537.36",
			"Accept: */*",
			"Referer: https://cryptobubbles.net/",
			"Accept-Language: id-ID,id;q=0.9,en-US;q=0.8,en;q=0.7",
		})
		reader := bufio.NewReader(os.Stdin)
		fmt.Print(pt, "INPUT COIN", y, " => ", lb)
		code, _ := reader.ReadString('\n')
		code = strings.TrimSpace(code)
		upperCode := strings.ToUpper(code)
		/* or use this
		fmt.Print(pt, "INPUT COIN", y, " => ", lb)
		var code string
		fmt.Scanln(&code)
		upperCode := strings.ToUpper(code)
		*/
		a := `"`
		as := a + upperCode + a

		coin := x(as, "{\"id\"", flynn)
		prc := x("\"price\":", ",\"", coin)
		cs := x("\"circulating_supply\":", ",\"", coin)
		dom := x("\"dominance\":", ",\"", coin)
		mrc := x("\"marketcap\":", ",\"", coin)
		perf := x("\"performance\":", "},", coin)
		trd := x("\"symbols\":", ",\"i", coin)
		fmt.Printf("%s<===================================>\n", lb)
		fmt.Printf("%s Ticker    %s:%s %s\n", pt, r, y, upperCode)
		fmt.Printf("%s Price     %s:%s %s USD\n", pt, r, gr, prc)
		fmt.Printf("%s Perfomance%s:%s %s\n", pt, r, pt, perf)
		fmt.Printf("%s c.Supplay %s:%s %s\n", pt, r, pt, cs)
		fmt.Printf("%s MarketCap %s:%s %s\n", pt, r, mg, mrc)
		fmt.Printf("%s Dominance %s:%s %s\n", pt, r, pt, dom)
		fmt.Printf("%s Trade On  %s:%s %s\n", pt, r, pt, trd)
		fmt.Printf("%s<===================================>\n", lb)

		timer(7)

	}
}
