let elements = document.querySelectorAll('.page');
let last = elements.length;
let pageIndex = getPos();

function previousPage() {
    if (pageIndex > 1) {
	pageIndex--
	window.location.hash = pageIndex
    }
}

function nextPage() {
    if (pageIndex < last) {
	pageIndex++
	window.location.hash = pageIndex
    }
}

document.addEventListener('keydown', function(event) {
    if (previousPageKey(event)) {
	event.preventDefault()		
	previousPage()
	return
    }
    if (nextPageKey(event)) {
	event.preventDefault()
	nextPage()
	return
    }
    var index = getPos()
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

window.addEventListener('resize', handleResize);

function handleResize() {
    window.location.hash = getPos()
}


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

