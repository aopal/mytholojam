currentGameID = ""
playerNo = 0
currentGame = {}
autoUpdateStatus = true
serverUrl = "http://localhost:8080"

var get = (url) => {
	var request = new XMLHttpRequest();
	request.open('GET', url, false);
	request.send(null);

	console.log(request)

	if (request.status != 200) {
		// alert(request.responseText)
		return null;

	}
	return request
}

var toggleHidden = (elemId) => {
  var x = document.getElementById(elemId);
  if (x.hidden) {
		x.hidden = false;
		x.style.zIndex = 1;
  } else {
		x.hidden = true;
		x.style.zIndex = -1;
  }
}

var updateStatus = () => {
	req = get(`${serverUrl}/status/${currentGameID}/0`)
	res = JSON.parse(req.responseText)
	currentGame = res
	console.log(currentGame)
	drawState()
}

var currentPlayer = () => {
	if (playerNo == 1) {
		return currentGame.player1
	} else {
		return currentGame.player2
	}
}

var currentOpponent = () => {
	if (playerNo == 1) {
		return currentGame.player2
	} else {
		return currentGame.player1
	}
}

var drawState = () => {
	drawHistory()
	drawPlayer("opponent-team", currentOpponent())
	drawPlayer("player-team", currentPlayer())
}

var drawHistory = () => {
	document.getElementById("history").innerHTML = ""

	let turnNo = -1
	let currentDiv = document.createElement("h2")
	currentDiv.innerText = "Game Start"
	for (let i = 0; i < currentGame.newActions.length; i++) {
		let a = currentGame.newActions[i]
		if (a.turn > turnNo) {
			turnNo = a.turn
			document.getElementById("history").appendChild(currentDiv)
			el = document.createElement("div")
			el.classList.add("Card")
			el.style.width = "100%"
			el.innerHTML = `<h3 class="card-title">Turn ${turnNo}</h3>`
			currentDiv = el
			// document.getElementById("history").appendChild(el)
		}
		action = document.createElement("p")
		action.innerText = a.actionText
		currentDiv.appendChild(action)
	}
	document.getElementById("history").appendChild(currentDiv)
}

var drawPlayer = (elemId, player) => {
	document.getElementById(elemId).innerHTML = ""
	console.log(player)

	header = document.createElement("div")
	header.classList.add("row")
	if (player === currentPlayer()) {
		header.innerHTML = "<h2 class='card-title mx-auto'>Your team</h2>"
	} else {
		header.innerHTML = "<h2 class='card-title mx-auto'>Opponent's team</h2>"
	}

	equip = document.createElement("div")
	equip.classList.add("row")

	for (const [k, v] of Object.entries(player.equipment)) {
		el = document.createElement("div")
		el.classList.add("card")
		el.classList.add("col")
		el.style.overflow = "auto"
		el.innerHTML =`
			<div class="card-body">
				<h6 class="card-title">${v.name}</h6>
				<p class="card-text">
					HP: ${v.hp}/${v.maxHP}<br>
					ATK: ${v.atk}<br>
					FLAM: ${v.defenses.FLAM}<br>
					NAME: ${v.defenses.NAME}<br>
					STRN: ${v.defenses.STRN}<br>
					WEAR: ${v.defenses.WEAR}
				</p>
			</div>`

		if (v.inhabited) {
			s = player.spirits[v.inhabitedBy]
			el.innerHTML += `
				<div class="card-body">
					<h6 class="card-title">${s.name}</h6>
					<p class="card-text">
						HP: ${s.hp}/${s.maxHP}<br>
						SPD: ${s.speed}<br>
						ATK: ${s.atk}<br>
						FLAM: ${s.defenses.FLAM}<br>
						NAME: ${s.defenses.NAME}<br>
						STRN: ${s.defenses.STRN}<br>
						WEAR: ${s.defenses.WEAR}
					</p>
				</div`
		}

		equip.appendChild(el)
	}

	document.getElementById(elemId).appendChild(header)
	document.getElementById(elemId).appendChild(equip)
}

var newGame = (operation) => {
	currentGameID = document.getElementById("game-id").value
	get(`${serverUrl}/${operation}-game/${currentGameID}`)
	if (operation == "create") {
		playerNo = 1
	} else {
		playerNo = 2
	}
	updateStatus()
	toggleHidden("modal-container")
	toggleHidden("game-container")
}
