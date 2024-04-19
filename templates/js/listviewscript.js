document.addEventListener("DOMContentLoaded", function() {
    const sessionID = sessionStorage.getItem('sessionID');
    console.log(sessionID)
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

function displayMobiles(mobile) {
   
    const mobileList = document.getElementById('mobileList');
    mobileList.innerHTML = '';
    mobile.forEach(mobile => {
        const mobileDiv = document.createElement('div');
        mobileDiv.classList.add('col-md-4','mt-4');
        mobileDiv.innerHTML = `
        <div class="card  text-center" style="width: 18rem;">
            <a href="viewmobile.html?id=${mobile.id}"><img class="img-fluid rounded card-image-top myimage" src="${mobile.image}" alt="Card image cap"></a>
        <div class="card-body">
        <div class="card-title"><h5>${mobile.name}</h5>
        <p class="card-text"><b>Rs.</b>${mobile.price}/-</p>
        </div>
        `;
        mobileList.appendChild(mobileDiv);
    });
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
// const sessionID = sessionStorage.getItem('sessionID');
// if(sessionID){
//     console.log(sessionID)
//     const button1 = document.getElementById("loginButon");
//     if (button1) {
//         button1.style.display = "none";
//     }
//     const button = document.getElementById("logoutButton");
//     if (button) {
//     button.style.display = "block";}
// }else{
//     const button1 = document.getElementById("loginButon");
//     if (button1) {
//         button1.style.display = "block";
//     }
//     const button = document.getElementById("logoutButton");
//     if (button) {
//     button.style.display = "none";
// }
// }

const sessionROLE = sessionStorage.getItem('sessionROLE');
if (sessionROLE === "admin") {
document.getElementById("adminNavItem").style.display = "block";
document.getElementById("adminNavItem1").style.display = "block";
const button1 = document.getElementById("loginButon");
if (button1) {
    button1.style.display = "none";
}

} else if (sessionROLE === "user"){
document.getElementById("adminNavItem").style.display = "none";
document.getElementById("adminNavItem1").style.display = "none";
const button = document.getElementById("loginButon");
if (button) {
    button.style.display = "none";
}
}else{
document.getElementById("adminNavItem").style.display = "none";
document.getElementById("adminNavItem1").style.display = "none";
const button1 = document.getElementById("loginButon");
if (button1) {
    button1.style.display = "block";
}
const button = document.getElementById("logoutButton");
if (button) {
    button.style.display = "none";
}
}

document.getElementById("logoutButton").addEventListener("click", function() {
    logout();
});
function logout() {
    sessionStorage.clear();
    window.location.assign('login.html');
}
document.getElementById("logoutButton").addEventListener("click", function() {
    logout(); 
});