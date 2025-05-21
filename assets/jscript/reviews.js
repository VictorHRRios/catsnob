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

function getReview(score) {
	const fullStar = `<svg xmlns="http://www.w3.org/2000/svg" height="24px" viewBox="0 -960 960 960" width="24px" fill="#B89230"><path d="m384-334 96-74 96 74-36-122 90-64H518l-38-124-38 124H330l90 64-36 122ZM233-120l93-304L80-600h304l96-320 96 320h304L634-424l93 304-247-188-247 188Zm247-369Z"/></svg>`
	const halfStar = `<svg xmlns="http://www.w3.org/2000/svg" height="24px" viewBox="0 -960 960 960" width="24px" fill="#B89230"><path d="M480-644v236l96 74-36-122 90-64H518l-38-124ZM233-120l93-304L80-600h304l96-320 96 320h304L634-424l93 304-247-188-247 188Z"/></svg>`
    
    const emptyStar = `<svg xmlns="http://www.w3.org/2000/svg" height="24px" viewBox="0 -960 960 960" width="24px" fill="#000000"><path d="m354-287 126-76 126 77-33-144 111-96-146-13-58-136-58 135-146 13 111 97-33 143ZM233-120l65-281L80-590l288-25 112-265 112 265 288 25-218 189 65 281-247-149-247 149Zm247-350Z"/></svg>`;
    score = Number(score)
    const fullStars = Math.floor(score);
    const hasHalfStar = score % 1 >= 0.5;
    
    let starsHTML = '';
    
    for (let i = 0; i < fullStars; i++) {
        starsHTML += fullStar;
    }

    if (hasHalfStar) {
        starsHTML += halfStar;
    }
    
    const emptyStars = 5 - fullStars - (hasHalfStar ? 1 : 0);
    for (let i = 0; i < emptyStars; i++) {
        starsHTML += emptyStar;
    }
    
    document.getElementById("score").innerHTML = starsHTML;
}
