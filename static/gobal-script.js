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
        },
        onError: function(xhr, request) {
        }
      });
});
