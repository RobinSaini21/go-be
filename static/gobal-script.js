document.getElementById("product-button").addEventListener("click", () => {
    const url = window.location.origin + "/product-gallery";
    window.history.pushState({
        refetch: false
    }, "", url);
});

window.addEventListener('popstate', function(event) {
    htmx.ajax('get',window.location.pathname , null, {
        headers: {
          'X-Requested-With': 'XMLHttpRequest'
        },
        onLoad: function(responseText, xhr, request) {
          console.log('HTMX request completed:', responseText);
          // Append the received HTML to the target element
         // document.getElementById('targetElement').innerHTML = responseText;
        },
        onError: function(xhr, request) {
          console.error('HTMX request failed:', xhr.status);
          // Handle errors here
        }
      });
});
