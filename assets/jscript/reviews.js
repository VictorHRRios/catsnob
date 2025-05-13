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

function updateReview(id) {
  fetch('/app/updateAlbumReview', {
    method: 'UPDATE',
    headers: {
      'Content-Type': 'application/json'
    },
    body: JSON.stringify({
			reviewId: id
		})
  })
  .then(response => {
    if (response.ok) {
      alert("Review updated");
	window.location.reload();
    } else {
      alert("Failed to update review");
    }
  });
}
