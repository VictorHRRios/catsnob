let currentReviewId = null;

function deleteReview(id) {
  fetch('/app/deleteReview', {
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

function submitCreate(id) {
  const createdData = {
    title: document.getElementById('new_title').value,
    review: document.getElementById('new_review').value,
    rating: document.getElementById('new_rating').value
  };

  fetch('/app/createReviewLong', {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json'
    },
    body: JSON.stringify({
	id: id,
      ...createdData
    })
  })
  .then(response => {
    if (response.ok) {
      alert("Review created!");
      window.location.reload();
    } else {
      alert("Failed to create review");
    }
    closeDialog();
  });
}

function submitUpdate(id) {
  const updatedData = {
    title: document.getElementById('title').value,
    review: document.getElementById('review').value,
    rating: document.getElementById('rating').value
  };

  fetch('/app/updateReview', {
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
      window.location.reload();
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

function showDialogCreate() {
  document.getElementById("createDialog").style.display = "block";
}

function closeDialog() {
  const dialog = document.querySelectorAll(".dialog");
	for (let value of dialog) {
		value.style.display = "none";
	}
  currentReviewId = null;
}
