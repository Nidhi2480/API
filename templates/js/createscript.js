document.getElementById("mobileDetailForm").addEventListener("submit", function(event) {
    event.preventDefault();
    const formData = new FormData(this);
    const sessionID = sessionStorage.getItem('sessionID');
    fetch("http://localhost:8080/addmobile", {
        method: "POST",
        body: formData,
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
        createdName=data.name;
        window.alert(createdName+ "details added successfully!");
        window.location.assign('listview.html');
    })
    .catch(error => {
        console.error('There was a problem with the fetch operation:', error);
        alert("Failed to add mobile detail. Please try again.");
    });
});