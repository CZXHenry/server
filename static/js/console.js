$(function() {
  $(".nav-link[href='"+window.location.pathname+"']").addClass("active")
  $("a.back").click(function() {
    window.history.back();
  });
  $("a.edit").click(function() {
    $(".need-enable").prop("disabled", false);
    $(".need-disappear").addClass("display-none");
    $(".need-appear").removeClass("display-none");
  });
  $("a.cancel").click(function() {
    window.location.reload();
  });
});
