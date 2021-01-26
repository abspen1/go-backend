package home

import (
	"fmt"
	"net/http"
)

// Page func that displays the endpoints on my API
func Page(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, `<body style="text-align:center;">
	<h1>Austin's back-end API</h1>
	<h3>Random</h3>
	<p><a href="/austinapi/projects">Projects</a></p>
	<p><a href="/austinapi/rmprojects">Remove Projects</a></p>
	<p><a href="/austinapi/rps/login">Rock Paper Scissors Login</a></p>
	<p><a href="/austinapi/rps/">Rock Paper Scissors Save</a></p>
	<p><a href="/austinapi/bdayemail">Bday Emailer</a></p>
	<p><a href="/austinapi/email">Contact Page Emailer</a></p>
	<p><a href="/austinapi/todos">Todos</a></p>
	<p><a href="/austinapi/tendie-intern">Twitter Data</a></p>
	<p><a href="/austinapi/go-tweet">Go Twitter Bot</a></p>
	<h3>Battle of the States Flag Football League</h3>
	<p><a href="/austinapi/botsffl">Standings</a></p>
	<p><a href="/austinapi/botsffl/teams/midwest">Midwest Teams</a></p>
	<p><a href="/austinapi/botsffl/teams/midwest/roster">Midwest Teams' Rosters</a></p>
	<p><a href="/austinapi/botsffl/teams/west">West Teams</a></p>
	<p><a href="/austinapi/botsffl/teams/west/roster">West Teams' Rosters</a></p>
	<p><a href="/austinapi/botsffl/teams/northeast">Northeast Teams</a></p>
	<p><a href="/austinapi/botsffl/teams/northeast/roster">Northeast Teams' Rosters</a></p>
	<p><a href="/austinapi/botsffl/teams/southeast">Southeast Teams</a></p>
	<p><a href="/austinapi/botsffl/teams/southeast/roster">Southeast Teams' Rosters</a></p>
	<p><a href="/austinapi/botsffl/trending/daily/add">sleeper's trending added players(24hrs)</a></p>
	<p><a href="/austinapi/botsffl/trending/daily/drop">sleeper's trending dropped players(24hrs)</a></p>
	<p><a href="/austinapi/botsffl/trending/weekly/add">sleeper's trending added players(5days)</a></p>
	<p><a href="/austinapi/botsffl/trending/weekly/drop">sleeper's trending added players(5days)</a></p>
	</body>`)
}
