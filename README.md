# chart
The latest CLI tool to view charts for both equities and crypto markets. 

Sometimes you need to see the chart for a symbol urgently, and navigating to Tradingview (or whatever platform you use), finding the ticker, and then picking the timeframe can all be time consuming. `chart` is made with that need in mind: **quick charts**. By typing in as little as 9 characters on your keyboard, you can get access to a state of the art chart for hundreds of tickers. You will also be able to navigate between other timeframes and tickers on your watchlist while the chart is open by just pressing arrow keys.

# Prerequisites
 - MacOS/Linux system
 - Git
     - [Install Guide](https://www.youtube.com/watch?v=B4qsvQ5IqWk) 
     - Verify with `git --version`
 - Make
     - This should come with your Mac/Linux system 
    - Verify with `make --version`
- Go >= 1.23.1
    - [Install Guide](https://www.youtube.com/watch?v=fPjcp48dpPM) 
    - Verify with `go version`

# Install
After verifying all prerequisites, run these commands in your terminal
```
git clone https://github.com/chartcmd/chart
cd chart
make prod-mac
chart
```
You can run this command from anywhere in your terminal after installation. Be sure to increase the size of your terminal window for the best experience :)

## Usage

```
chart btc 15m
```
```
>>
    BTC: $67293.44                                                             19:52:50
    0.23%
                                                                                       |      
                                                                                       | 68000
                                                                                       |      
                                                                                       |      
                                                                                       |      
                                                                                       |      
                                                                                       |      
                                                                                       |      
                                                                         │             | 67500
                                                                         ┃┃            |      
                                                                        ┃┃┃    │       |      
                                                                        ┃│┃    ┃┃┃  ▁▁▁| 67293
                                                                        ┃ ┃   ┃┃┃      |      
                                                                        ┃ ┃┃┃┃┃        |      
                                                                        ┃  ┃┃│         |      
                                                                        ┃  ┃┃          | 67000
 │   │┃                                                                 ┃              |      
 ┃│┃ ┃┃                                                                 ┃              |      
 │┃┃│┃┃                     │  │                                        ┃              |      
    ┃││┃                    ┃┃││┃                                  │    ┃              |      
     │ ┃┃                   ┃┃│┃┃                                  │┃  ┃               |      
        ┃                   ┃┃┃ ┃                        ┃         ┃ ┃┃┃               |      
        ┃│   │┃┃   ┃       ┃ ┃┃ ┃                       ┃│┃      ┃┃   ┃┃               |      
        │┃ │┃┃┃┃┃ ┃┃ │  │┃┃┃ │┃ ┃                   │  ┃ ││┃ │┃┃┃┃                     | 66500
        │┃ ┃    ┃┃ │┃┃│ │┃ │    │┃                 │┃┃│┃ │ ┃ ┃┃┃ ┃                     |      
         ┃ ┃     ┃  │┃│┃┃        ┃                 ┃│┃│┃   │┃┃┃┃                       |      
         ┃┃┃         │┃┃│        ┃│  │             ┃  ┃                                |      
         ┃┃           ┃┃         ┃┃ │┃            │┃                                   |      
                       │          ┃│┃┃            ┃┃                                   |      
                                  │┃┃┃ │         │┃│                                   |      
                                   ┃┃│┃││        ┃┃                                    | 66000
                                   ┃┃ ┃│┃       ┃┃┃                                    |      
                                   │  ┃┃┃ │┃    ┃│┃                                    |      
                                        │┃┃┃│   ┃                                      |      
                                         ┃┃┃│   ┃                                      |      
                                         ┃┃│┃┃ ┃                                       |      
                                         │┃  ┃ ┃                                       |      
                                             ┃┃┃                                       | 65500
                                             │┃                                        |      
                                             ││                                        |      
                                             ││                                        |      
                                              │                                        |      
                                              │                                        |      
                                                                                       |      
                                                                                       | 65000
                                                                                       |      
  --------------------------------------------------------------------------------------      
  Oct 23           04:00           08:00           12:00           16:00                      

                        [1m]    [5m]    [15m]<    [1h]    [6h]    [1d]  

         [BTC]    [ETH]    [SOL]    [DOGE]    [BONK]    [SUI]    [SEI]    [ONDO]    [TIA]  

```

# Backlog
These items are currently being worked on and will be released in upcoming versions:
 - `chart list`
 - Fix equity chart bugs
 - `chart crypto` | `chart equities`
     - This will display 4 charts at a time from the tickers in the watchlist config
 - Fix the Github release workflows
 - Add testing workflows
