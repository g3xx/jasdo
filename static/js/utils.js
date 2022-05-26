// generate array pagination
var pages = function(currentPage, totalcount) {
    let pages = [];
    totalcount = Math.ceil(totalcount / 10)
    if (totalcount <= 6) {
        for (i = 1; i <= totalcount; i++) {
           pages.push(i)
        }
    } else {
        pages.push(1);
        if (currentPage > 3) {
            pages.push("...");
        }
        if (currentPage == totalcount) {
            pages.push(currentPage - 2);
        }
        if (currentPage > 2) {
            pages.push(currentPage - 1);
        }
        if (currentPage != 1 && currentPage != totalcount) {
            pages.push(currentPage);
        }
        if (currentPage < totalcount - 1) {
            pages.push(currentPage + 1);
        }
        if (currentPage == 1) {
            pages.push(currentPage + 2);
        }
        if (currentPage < totalcount - 2) {
            pages.push("...");
        }
        if (totalcount > 1) {
            pages.push(totalcount);
        }
    }
    return pages;
}

// wrap function post with csrf
// dry csrf token 
var doPost = function(url, m, data){
    return fetch(url, {
        method: m,
        headers: { 
            'X-CSRFToken' : Alpine.store('csrf').val,
            'Content-Type': 'application/json' 
        },
        body: JSON.stringify(data)
    }).then(res => res.json())
    .then(result => {
        return result 
    }).catch(err => {
         throw new Error(err.message);
    });
}

// valdiate Email
var validateEmail = (email) => {
    return email.match(
      /^(([^<>()[\]\\.,;:\s@\"]+(\.[^<>()[\]\\.,;:\s@\"]+)*)|(\".+\"))@((\[[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\])|(([a-zA-Z\-0-9]+\.)+[a-zA-Z]{2,}))$/
    );
}
  