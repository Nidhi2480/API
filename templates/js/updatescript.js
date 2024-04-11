window.onload = function() {
    const urlParams = new URLSearchParams(window.location.search);
    const productId = urlParams.get('id');
    var secondAPIURL = "http://localhost:8080/getmobile/" + productId;
    fetch(secondAPIURL)
    .then(response => {
        if (!response.ok) {
            throw new Error('Network response was not ok');
        }
        return response.json();
    })
    .then(data => {
        document.getElementById('name').value = data.name;
        document.getElementById('specs').value = data.specs;
        document.getElementById('price').value = data.price;
    })
    .catch(error => {
        console.error('There was a problem with the fetch operation:', error);
    });

}
document.getElementById("mobileDetailForm").addEventListener("submit", function(event) {
event.preventDefault(); 
const formData = new FormData(this);
const urlParams = new URLSearchParams(window.location.search);
const productId = urlParams.get('id');
fetch("http://localhost:8080/update/"+productId, {
method: "PUT",
body: formData
})
.then(response => {
if (!response.ok) {
    throw new Error('Network response was not ok');
}
return response.text();
})
.then(data => {
alert("Mobile detail updated successfully!");
console.log("Server response:", data);
window.location.assign('listview.html');
})
.catch(error => {
console.error('There was a problem with the fetch operation:', error);
alert("Failed to add mobile detail. Please try again.");
});

});
