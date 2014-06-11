function getMessages() {

	var req = new XMLHttpRequest()
	req.open('GET', "https://localhost:8080/messages", false)
	req.send(null)
	if(req.status == 200)
	{
		//var response = console.log(req.responseText)
		document.getElementById('chat-box').innerHTML += "<p>" + req.responseText + "</p>"
	}
	else
	{
		alert('no')
	}
}

function sendMessage() {
	var jeison = {
		name: "Paco",
		body: document.getElementById('message-area').value
	}

	var req = new XMLHttpRequest()
	req.open('POST', "https://localhost:8080/messages", true)
	req.setRequestHeader('Content-type', 'application/json')
	req.onreadystatechange = function() {
		if(req.readyState == 4)
		{
			if(req.status == 400)
			{
				alert("errroooor")
			}
			else if(req.status == 200)
			{
				console.log(JSON.parse(req.responseText))
				pack = JSON.parse(req.responseText)
				console.log(pack)
				document.getElementById('chat-box').innerHTML += "<p>" + pack.Name + "</p>"
				document.getElementById('chat-box').innerHTML += "<p>" + pack.Body + "</p>"
			}		
			console.log(JSON.parse(req.responseText))			
		}
	}
	req.send(JSON.stringify(jeison))
	console.log(JSON.stringify(jeison))
}
