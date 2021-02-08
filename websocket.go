package trading212

// wss://live.trading212.com/streaming/events/?app=WC4&appVersion=5.119.0&EIO=3&transport=websocket

// first msg:
// 0{"sid":"aaaaaa-","upgrades":[],"pingInterval":7000,"pingTimeout":5000}

// sends this a lot, i guess frontends just checking for changes on it
// 42["acc",{$accobject}

// every 7 seconds it sends the message "2" and gets the response "3"

// startup
//M"]	30
//19:27:23.384
//> 42["working-schedule-sync",[{.....
//19:27:23.385
//> 42["deposit-declaration-sync",{"accountId":3943424,"status":"FILLED"}]	70
//19:27:23.385
//> 42["platform-message-sync",[]]	30
//19:27:23.385
//42["subscribe","/ACCOUNT"]	26
//19:27:23.399
//42["subscribe","/BUYERSSELLERS"]	32
//19:27:23.399
//42["subscribe","/WORKING-SCHEDULES"]	36
//19:27:23.399
//42["subscribe","/DEPOSIT-DECLARATIONS"]	39
//19:27:23.399
//42["subscribe","/BUCKETS"]	26
//19:27:23.399
//42["subscribe","/TN",["BNGO_US_EQ"]]	36
//19:27:23.399
//42["s-qbulk",["BNGO_US_EQ","GME_US_EQ","BB_US_EQ","AAME_US_EQ","THCB_US_EQ","TSLA_US_EQ","AAPL_US_EQ","AMZN_US_EQ","NFLX_US_EQ"]]	129
//19:27:23.399
//42["acc"]	9
//19:27:23.399
//> 42["buckets-sync",[]]	21
//19:27:23.421
//42["unsubscribe","/BUCKETS",[]]
