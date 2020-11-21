currentGameID = ""
currentGame = {numActions: 0}
actionHistory = []
currentlySelectable = []

playerNo = 0
playerToken = ""

pendingActions = []
actionInProgress = {}

autoUpdateStatus = true
allStates = ["waitForJoin", "chooseSpirit", "chooseMove", "chooseTarget", "waitForOpponent", "watching"]
verboseStateText = {
	"waitForJoin": "Waiting for an opponent to join the game...",
	"chooseSpirit": "Please select a spirit",
	"chooseMove": "Please select a move", 
	"chooseTarget": "Please selet a target for the move",
	"waitForOpponent": "Waiting for your opponent to select their actions...",
	"watching": "You are watching the game"
}
currentState = "waitForJoin"
serverUrl = "http://localhost:8080"

var processState = () => {
	console.log("processingState", currentState)
	if (currentState == "waitForJoin" || 
		currentState == "waitForOpponent" || 
		currentState == "watching") {
		res = updateStatus()

		if (currentState == "waitForOpponent" && res.newActions.length > 0) {
			currentState = "chooseSpirit"
		} else if (currentState == "waitForJoin" && currentOpponent() != null) {
			currentState = "chooseSpirit"
		}
	}
	updateSelectable()

	updateStatusText()
}

var processClicked = (elem) => {
	// console.log("clicked", elem)
	obj = getObj(elem)
	if (!currentlySelectable.includes(obj.id)) {
		// console.log("not selectable")
		return
	}
	// console.log("selectable")

	if (currentState == "chooseSpirit") {
		actionInProgress.spirit = elem
		currentState = "chooseMove"
	} else if (currentState == "chooseMove") {
		actionInProgress.move = elem
		currentState = "chooseTarget"
	} else if (currentState == "chooseTarget") {
		// console.log("chose target")
		actionInProgress.target = elem
		pendingActions.push(actionInProgress)
		actionInProgress = {}

		// console.log("updating state asdf")
		if (pendingActions.length == Object.entries(currentPlayer().spirits).length) {
			submitActions()
			currentState = "waitForOpponent"
		} else {
			currentState = "chooseSpirit"
		}
	}
}

var updateSelectable = () => {
	// console.log("updating selectability")

	selectableElements = []

	if (currentState == "chooseSpirit") {
		selectableElements = document.getElementById("player-team").getElementsByClassName("spirit")
	} else if (currentState == "chooseMove") {
		selectableElements = actionInProgress.spirit.nextElementSibling.getElementsByClassName("move")
	} else if (currentState == "chooseTarget") {
		move = getObj(actionInProgress.move)
		if (move.teamTargetable == "self") {
			selectableElements = document.getElementById("player-team").getElementsByClassName("equipment")
		} else {
			selectableElements = document.getElementById("opponent-team").getElementsByClassName("equipment")
		}
	}

	currentlySelectable = []
	for (let i = 0; i < selectableElements.length; i++) {
		// selectableElements[i].classList.add("currentlySelectable")
		currentlySelectable.push(getObj(selectableElements[i]).id)
	}

	highlighted = document.querySelectorAll(".currentlySelectable")
	// console.log(highlighted.length, " elements highlighted")
	// console.log(highlighted.length, "currently selectable")
	for (let i = 0; i < highlighted.length; i++) {
		highlighted[i].classList.remove("currentlySelectable")
	}

	for (let i = 0; i < currentlySelectable.length; i++) {
		document.getElementById(currentlySelectable[i]).classList.add("currentlySelectable")
	}

	// selectableElements.forEach(elem => {
	// 	elem.classList.add("currentlySelectable")
	// 	currentlySelectable.push(getObj(elem).id)
	// });
}

var getObj = (elem) => {
	return JSON.parse(elem.getElementsByTagName("span")[0].innerText)
}

var submitActions = () => {
	// console.log(pendingActions)

	actions = []

	pendingActions.forEach(act => {
		actions.push({
			user: getObj(act.spirit),
			move: getObj(act.move),
			targets: [getObj(act.target)],
			turn: currentGame.numTurns + 1,
			actionText: ""
		})
	});

	payload = {
		token: playerToken,
		actions: actions
	}

	// console.log("sending action payload", payload, JSON.stringify(payload))

	res = post(`${serverUrl}/take-action/${currentGameID}`, payload)
	// console.log(res)

	pendingActions = []
}

var updateStatusText = () => {
	document.getElementById("status-text").innerText = verboseStateText[currentState]
}

var get = (url) => {
	var request = new XMLHttpRequest();
	request.open('GET', url, false);
	request.send(null);

	// console.log(request)

	if (request.status != 200) {
		// alert(request.responseText)
		return null;

	}
	return request
}

var post = (url, body) => {
	var request = new XMLHttpRequest();
	request.open('POST', url, true);
	request.setRequestHeader("Content-type", "application/json");
	request.send(JSON.stringify(body));

	// console.log(request)

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
	req = get(`${serverUrl}/status/${currentGameID}/${currentGame.numActions}`)
	res = {}

	try {
		res = JSON.parse(req.responseText)

		currentGame = res
		actionHistory = actionHistory.concat(res.newActions)
	} catch (e) {}
	
	// console.log(currentGame)
	drawState()
	return res
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
	updateSelectable()
}

var drawHistory = () => {
	document.getElementById("history").innerHTML = ""

	let turnNo = -1
	let currentDiv = document.createElement("h2")
	currentDiv.innerText = "History"
	for (let i = 0; i < actionHistory.length; i++) {
		let a = actionHistory[i]
		if (a.turn > turnNo) {
			turnNo = a.turn
			document.getElementById("history").appendChild(currentDiv)
			el = document.createElement("div")
			el.classList.add("Card")
			el.style.width = "100%"
			el.innerHTML = `<h3 class="card-title">Turn ${turnNo}</h3>`
			currentDiv = el
		}
		action = document.createElement("p")
		action.innerText = a.actionText
		currentDiv.appendChild(action)
	}
	document.getElementById("history").appendChild(currentDiv)
}

var drawPlayer = (elemId, player) => {
	if (player == null) {
		return
	}
	document.getElementById(elemId).innerHTML = ""
	// console.log(player)

	header = document.createElement("div")
	header.classList.add("row")
	if (currentState === "watching") {
		header.innerHTML = `<h2 class='card-title mx-auto'>Team A</h2>`
	} else if (player === currentPlayer()) {
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
		el.innerHTML = `
			<div class="card-body equipment" onclick="processClicked(this)" id="${v.id}">
				<h6 class="card-title">${v.name}</h6>
				<p class="card-text">
					HP: ${v.hp}/${v.maxHP}<br>
					ATK: ${v.atk}<br>
					FLAM: ${v.defenses.FLAM}<br>
					NAME: ${v.defenses.NAME}<br>
					STRN: ${v.defenses.STRN}<br>
					WEAR: ${v.defenses.WEAR}
				</p>
				<span hidden>${JSON.stringify(v)}</span>
			</div>`

		if (v.inhabited) {
			s = player.spirits[v.inhabitedBy]
			el.innerHTML += `
				<div class="card-body spirit" onclick="processClicked(this)" id="${s.id}">
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
					<span hidden>${JSON.stringify(s)}</span>
				</div>`

			moveContainer = `<div class="row move-list">`

			for (const [k, v] of Object.entries(s.moves)) {
				move = Object.assign({}, v, {id: `${s.id}_${v.name}`})
				moveContainer += `
				<div class="col move" onclick="processClicked(this)" id="${move.id}">
					<h6 class="card-title">${move.name}</h6>
					<p class="card-text">
						PWR: ${move.power}<br>
						PRI: ${move.priority}<br>
						TYPE: ${move.type}
					</p>
					<span hidden>${JSON.stringify(move)}</span>
				</div>`
			}

			moveContainer += `</div>`
			el.innerHTML += moveContainer
		}

		equip.appendChild(el)
	}

	document.getElementById(elemId).appendChild(header)
	document.getElementById(elemId).appendChild(equip)
}

var newGame = (operation) => {
	currentGameID = document.getElementById("game-id").value
	res = get(`${serverUrl}/${operation}-game/${currentGameID}`)
	playerToken = res.responseText
	console.log("create/join response", res)
	if (operation == "create") {
		playerNo = 1
	} else {
		playerNo = 2
	}
	updateStatus()
	toggleHidden("modal-container")
	toggleHidden("game-container")
}

var watchGame = () => {
	currentGameID = document.getElementById("game-id").value
	currentState = "watching"
	updateStatus()
	toggleHidden("modal-container")
	toggleHidden("game-container")
}


setInterval(processState, 500)
