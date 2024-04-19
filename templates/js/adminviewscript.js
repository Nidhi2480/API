document.addEventListener("DOMContentLoaded", function() {
    fetchMobiles();
});

function fetchMobiles() {
    fetch('http://localhost:8080/mobiles')
        .then(response => {
            if (!response.ok) {
                throw new Error('Network response was not ok');
            }
            return response.json();
        })
        .then(data => {
            displayMobiles(data);
        })
        .catch(error => {
            console.error('There was a problem with the fetch operation:', error);
        });
}

function displayMobiles(mobiles) {
    
    const mobileList = document.getElementById('mobileList');
    mobileList.innerHTML = '';
    var counter=0
    mobiles.forEach(mobile => {
        const row = document.createElement('tr');
        row.innerHTML = `
            <th scope="row">${++counter}</th>
            <td>
                <div class="img-container">
                    <a href="viewmobile.html?id=${mobile.id}">
                        <img class="img-fluid rounded " src="${mobile.image}" alt="Mobile Phone">
                    </a>
                </div>
            </td>
            <td class="movie-details">
                <p class="my_title4">
                    <b>${mobile.name}</b><br>
                    Mobile specs:<b>${mobile.specs}<br></b>
                    <p class="text-right">Rs.<b>${mobile.price}<br></b></p>
                    </p>
                    <hr>
                    <a class="btn btn-warning" href="update.html?id=${mobile.id}"> update</a>
                    <button onclick="confirmDelete('${mobile.id}')" class="btn btn-danger">Delete Mobile</button>
                    
            </td>
        `;
        mobileList.appendChild(row);
    });
}
function confirmDelete(studentId) {
    
if (confirm("Are you sure you want to delete this mobile?")) {
    const sessionID = sessionStorage.getItem('sessionID');
    const stmt='http://localhost:8080/delmobile/'+studentId
    console.log(stmt)
    fetch(stmt, {
    method: 'DELETE',
    headers: {
        "Authorization": sessionID
    }
    })
    .then(response => {
        if (!response.ok) {
            throw new Error('Network response was not ok');
        }
        return response.json();
    })
    .then(data => {
        window.alert(data.message)
        fetchMobiles();
    })
    .catch(error => {
        console.error('There was a problem with the fetch operation:', error);
    });
}
}
document.getElementById("searchForm").addEventListener("submit", function(event) {
event.preventDefault();
const query = document.getElementById("searchInput").value.trim();
console.log(query)
fetch(`http://localhost:8080/search?query=${query}`)
.then(response => {
    if (!response.ok) {
        throw new Error('Network response was not ok');
    }
    return response.json();
})
.then(data => {
    if (data!=null){
    header2.innerHTML = '';
    document.getElementById('header2').innerHTML = `You've searched for "${query}"`
    console.log(data);
    displayMobiles(data);}
    else{
        console.log("no results found")
        document.getElementById('header2').innerHTML = `no results found for "${query}"`
        mobileList.innerHTML = '';
    }
})
.catch(error => {
    console.error('There was a problem with the fetch operation:', error);
});
});
