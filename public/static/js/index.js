var btn = document.getElementById('btn')
var checkbox = document.getElementById('switch')
var fact = document.getElementById('fact')

function randomFact() {
  var req = new XMLHttpRequest()
  var path = window.location.pathname
  req.open('GET', '/api')
  req.send()
  req.onload = function() {
    res = JSON.parse(req.responseText)
    resValue = res.value
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

function darkModeSwitch() {
  if (checkbox.checked) {
    document.body.classList.add('dark')
  } else {
    document.body.classList.remove('dark')
  }
  localStorage.setItem('switch', checkbox.checked)
}

btn.addEventListener('click', randomFact)
checkbox.addEventListener('click', darkModeSwitch)

checkbox.checked = localStorage.getItem('switch') === 'true' ? true : false
darkModeSwitch()
