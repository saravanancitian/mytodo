
var xml = new XMLHttpRequest();

function ajxcall(url, method, value, callback, async){
	xml.open(method, url);
	xml.setRequestHeader("Content-type", "application/json");
	xml.onreadystatechange = function() {
		 if(this.readyState === 4 && this.status === 200) {
			 callback(this.responseText);
		 }
	};
	xml.send(value);
	
}