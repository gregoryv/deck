document.addEventListener('keydown', function(event) {
    var elements = document.querySelectorAll('.page');
    var last = elements.length;

    index = getPos()
    if ( index > 1 && previousPageKey(event)) {
	index--
	event.preventDefault()	
    }
    if (index < last && nextPageKey(event)) {
	index++
	event.preventDefault()	
    }
    if (event.key === 'Home') {
	index = 1
	event.preventDefault()	
    }
    if (event.key === 'End') {
	index = last
	event.preventDefault()	
    }

    window.location.hash = index
    return
})

function previousPageKey(event) {
    return event.key === 'ArrowLeft' ||
	    event.key === 'ArrowUp' ||
	    event.key === 'PageUp'
}
function nextPageKey(event) {
    return event.key === 'ArrowRight' ||
	    event.key === 'ArrowDown' ||
	    event.key === 'PageDown'
}
function getPos() {
    let currentView = window.location.hash
    if (currentView === '') {
	currentView = "#1"
    }
    currentView =parseInt(currentView.split('#').join(''))
    return currentView
}

