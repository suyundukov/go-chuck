(function() {
  var btn = document.querySelector('.boxy.green');
  var checkbox = document.querySelector('#switch');
  var fact = document.querySelector('.fact');
  var link = document.querySelector('.link');

  function getRandomFact() {
    var req = new XMLHttpRequest();
    var path = window.location.pathname;

    req.open('GET', '/api');
    req.onload = function() {
      res = JSON.parse(req.responseText);
      fact.innerHTML = res.value.fact;
      link.setAttribute('href', '/fact/' + res.value.id);
    }
    req.send();

    // If user is not on the home page, return him back
    if (path != '/') {
      history.replaceState('', 'Chuck Norris Facts', '/');
    }
  }

  function darkModeSwitch() {
    if (checkbox.checked) {
      document.body.classList.add('dark');
    } else {
      document.body.classList.remove('dark');
    }
    localStorage.setItem('switch', checkbox.checked);
  }

  btn.addEventListener('click', getRandomFact);
  checkbox.addEventListener('click', darkModeSwitch);

  checkbox.checked = localStorage.getItem('switch') === 'true' ? true : false;
  darkModeSwitch();
})();
