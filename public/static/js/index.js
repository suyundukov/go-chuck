var fact = document.getElementById('fact')

function randomFact() {
  var req = new XMLHttpRequest()
  var path = window.location.pathname
  req.open('GET', '/api')
  req.send()
  req.onload = function() {
    resValue = JSON.parse(req.responseText)
    chuckFact = resValue.fact
    fact.innerHTML = chuckFact
    url = resValue.id
    permaLink.setAttribute('href', '/fact/' + url)
  }
  // If user is not on the home page, return him back
  if (path != '/') {
    history.replaceState('', 'Chuck Norris Facts', '/')
  }
}

var btn = document.getElementById('btn')
btn.addEventListener('click', randomFact)
