let elements = document.querySelectorAll('.page');
const last = elements.length;

function previousPage() {
    let pageIndex = getPos()
    if (pageIndex > 1) {
	pageIndex--
	window.location.hash = pageIndex
    }
}

function nextPage() {
    let pageIndex = getPos()
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
    if (event.key === 'Home') {
	pageIndex = 1
	event.preventDefault()
	window.location.hash = pageIndex	
    }
    if (event.key === 'End') {
	pageIndex = last
	event.preventDefault()
	window.location.hash = pageIndex	
    }
    if (event.key === 't') {
	pageIndex = 2 // toc
	event.preventDefault()
	window.location.hash = pageIndex	
    }
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

