var objDiv = document.getElementById("chat-box");

function sendLogin() {
	var user = {
		name: document.getElementById('user-name').value,
		password: document.getElementById('user-pass').value
	}

	var req = new XMLHttpRequest()
	req.open('POST', "https://localhost:8080/login", true)
	req.setRequestHeader('Content-type', 'application/json')
	req.onreadystatechange = function() {
		if(req.readyState == 4)
		{
			if(req.status == 400)
			{
				alert("User or password incorrect")
			}
			else if(req.status == 200)
			{
				location.href="/index";

				localStorage['username'] = user.name
			}		
		}
	}
	req.send(JSON.stringify(user))
}

function getMessages() {

	var user = {
		name: localStorage['username']
	}

	var req = new XMLHttpRequest()
	req.open('POST', "https://localhost:8080/getMessages", false)
	req.setRequestHeader('Content-type', 'application/json')
	req.send(JSON.stringify(user))

	if(req.status == 400)
	{
		alert("errroooor")
	}
	else if(req.status == 200)
	{
		pack = JSON.parse(req.responseText)
		document.getElementById('chat-box').innerHTML += "<div class='message-row'> <p class='user-name'>" + pack.Name + "</p>" +
		 "<p class='user-message'>" + pack.Body + "</p> </div>"
	}		

	objDiv.scrollTop = objDiv.scrollHeight;	
}

function sendMessage() {
	var jeison = {
		name: localStorage['username'],
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
				pack = JSON.parse(req.responseText)
				document.getElementById('chat-box').innerHTML += "<div class='message-row'> <p class='user-name'>" + pack.Name + "</p>" + 
				 "<p class='user-message'>" + pack.Body + "</p> </div>"
			}		
			objDiv.scrollTop = objDiv.scrollHeight;	
		}
	}
	req.send(JSON.stringify(jeison))
}
