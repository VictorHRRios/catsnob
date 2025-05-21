let currentShoutId = null;
function deleteShout(id) {
  fetch('/app/deleteShout', {
    method: 'DELETE',
    headers: {
      'Content-Type': 'application/json'
    },
    body: JSON.stringify({ shoutId: id })
  })
  .then(response => {
    if (response.ok) {
      alert("Shout deleted");
      window.location.reload();
    } else {
      alert("Failed to delete shout");
    }
  });
}
function submitUpdate(id) {
    const updatedData = {
        title: document.getElementById('title').value,
        shout_text: document.getElementById('shout_text').value
    };
    
    fetch('/app/updateShout', {
        method: 'PUT',
        headers: {
        'Content-Type': 'application/json'
        },
        body: JSON.stringify({
        id: currentShoutId,
        ...updatedData
        })
    })
    .then(response => {
        if (response.ok) {
        alert("Shout updated!");
        window.location.reload(); // or update DOM if you prefer
        } else {
        alert("Failed to update shout");
        }
        closeDialog();
    });
    }

function showDialog(id) {   
    currentShoutId = id;
    document.getElementById("updateDialog").style.display = "block";
    
    // Optional: preload existing values via DOM or fetch
    // Example: if already in the DOM, you can fetch values like:
    // document.getElementById("dialog-title").value = document.getElementById(`title-${id}`).textContent;
    }
function closeDialog() {
    document.getElementById("updateDialog").style.display = "none";
    currentShoutId = null;
}




