function clearOutput() {
    	output = document.getElementById("result").innerHTML = "";
	document.getElementById("input").value = "";
}

function makeRequest() {
	font = document.getElementById('font').value;
	input = document.getElementById('input').value;
	let text = encodeURIComponent(input);
	if (text == "") {
		document.getElementById("result").innerHTML = "";
		return
	}
	let width = window.innerWidth / 6.50;
	width = parseInt(width)
	req = 'font=' + font + '&width=' + width + '&input=' + text;
	httpRequest = new XMLHttpRequest();
	httpRequest.onreadystatechange = function() {
		if (httpRequest.status == 400) {
			document.getElementById("result").innerHTML = "OOPS! Bad Request!";
			return 
		}
		if (httpRequest.status == 500) {
			document.getElementById("result").innerHTML = "OOPS! Internal ServerError!";
			return 
		}
		if (httpRequest.responseText) {
			document.getElementById("result").innerHTML = httpRequest.responseText;
		}
	}
	httpRequest.open("POST", '/api');
	httpRequest.setRequestHeader("Content-Type", "application/x-www-form-urlencoded");
	httpRequest.send(req);
}

function copyText() {
	var copyText = document.getElementById("result").textContent;
	navigator.clipboard.writeText(copyText); 
}

