Betfair Golang Library
===
Betfair is the leading company in online betting exchange. They are currently developing the next generation API for developers which want to automate their betting systems. For more info, please refer to [Betfair API-NG home](https://api.developer.betfair.com/services/webapps/docs/display/1smk3cen4v3lu3yomq5qye0ni/API-NG+Overview).

I will write some articles [here](aded.it/tag/comp/betfair-golang-library).

Please note that the only login method implemented is the [non-interactive (bot) login](https://api.developer.betfair.com/services/webapps/docs/display/1smk3cen4v3lu3yomq5qye0ni/Non-Interactive+%28bot%29+login).

TODO
---

* Betting API: order-related methods
* Betting API: listTimeRanges, listVenues
* Betting API: add all params to methods
* Betting API: add all params to market filter
* Betting API: improve support for horse racing

Quick usage
---
See examples/*.

Account info:
<pre>
go run examples/account_info.go -conf=betfair_test.conf.json
</pre>
Output:
<pre>
Alessandro De Donno (it)
	Available: EUR x.xx
	Exposure : EUR 0.00
</pre>

Football events in GB:
<pre>
go run examples/football_events.go -conf=betfair_test.conf.json
</pre>
Output:
<pre>
(...)
2014-01-04 15:00:00 +0000 UTC: [GB] Barnsley v Coventry (40 markets)
2014-01-04 15:00:00 +0000 UTC: [GB] Bolton v Blackpool (40 markets)
2014-01-04 15:00:00 +0000 UTC: [GB] Stoke v Leicester (78 markets)
2014-01-04 15:00:00 +0000 UTC: [GB] West Brom v C Palace (78 markets)
2014-01-04 15:00:00 +0000 UTC: [GB] Wigan v MK Dons (40 markets)
2014-02-18 19:45:00 +0000 UTC: [GB] Man City v Barcelona (3 markets)
2014-01-04 17:15:00 +0000 UTC: [GB] Arsenal v Tottenham (78 markets)
2014-01-04 17:15:00 +0000 UTC: [GB] Fixtures 04 January   (1 markets)
(...)
</pre>

Tests
---
To run tests, you have to provide your credentials in the file
<pre>
betfair_test.conf.json
</pre>
You can find a sample configuration file in the main directory:
<pre>
betfair_test.conf.json.sample
</pre>

Please note that only username, password, certfile and keyfile are mandatory.

License
---
The library is dual licensed. For free software/open source projects, please refer to GPLv3. For commercial projects, please contact the author.

Note
---
I apologize for my poor English :-).