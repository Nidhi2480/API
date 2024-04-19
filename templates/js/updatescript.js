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
        document.getElementById('imagePreview').src = data.image;
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
const imageFile = formData.get('image');
const sessionID = sessionStorage.getItem('sessionID');
console.log(imageFile)
    if (imageFile.name == "") {

        const imageUrl = document.getElementById('imagePreview').src;
        const imageName = imageUrl.substring(imageUrl.lastIndexOf('/') + 1);
        formData.append('image1', imageName);
        console.log(imageName)
    }
fetch("http://localhost:8080/update/"+productId, {
method: "PUT",
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
    const updatedName = data.name;
    window.alert(updatedName + " details updated successfully!");
    window.location.assign('listview.html');
   
})
.catch(error => {
console.error('There was a problem with the fetch operation:', error);
alert("Failed to add mobile detail. Please try again.");
});

});
