fetch("api/hello")
.then(r => r.json())
.then(data => document.getElementById('label').innerText = data.greets)
