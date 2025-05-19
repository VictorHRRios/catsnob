let currentReviewId = null;

function deleteReview(id) {
  fetch('/app/deleteAlbumReview', {
    method: 'DELETE',
    headers: {
      'Content-Type': 'application/json'
    },
    body: JSON.stringify({ reviewId: id })
  })
  .then(response => {
    if (response.ok) {
      alert("Review deleted");
	window.location.reload();
    } else {
      alert("Failed to delete review");
    }
  });
}

function submitUpdate(id) {
  const updatedData = {
    title: document.getElementById('title').value,
    content: document.getElementById('content').value,
    rating: document.getElementById('rating').value
  };

  fetch('/app/updateAlbumReview', {
    method: 'PUT',
    headers: {
      'Content-Type': 'application/json'
    },
    body: JSON.stringify({
      id: currentReviewId,
      ...updatedData
    })
  })
  .then(response => {
    if (response.ok) {
      alert("Review updated!");
      window.location.reload(); // or update DOM if you prefer
    } else {
      alert("Failed to update review");
    }
    closeDialog();
  });
}

function showDialog(id) {
  currentReviewId = id;
  document.getElementById("updateDialog").style.display = "block";

  // Optional: preload existing values via DOM or fetch
  // Example: if already in the DOM, you can fetch values like:
  // document.getElementById("dialog-title").value = document.getElementById(`title-${id}`).textContent;
}

function closeDialog() {
  document.getElementById("updateDialog").style.display = "none";
  currentReviewId = null;
}
