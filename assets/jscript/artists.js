function deleteArtist(id) {
  fetch('/admin/deleteArtist', {
    method: 'DELETE',
    headers: {
      'Content-Type': 'application/json'
    },
    body: JSON.stringify({ albumId: id })
  })
  .then(response => {
    if (response.ok) {
      alert("Artist deleted");
	window.location.reload();
    } else {
      alert("Failed to delete artist");
    }
  });
}

